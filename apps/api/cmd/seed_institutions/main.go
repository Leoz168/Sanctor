package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type upstreamInstitution struct {
	Name          string  `json:"name"`
	Country       string  `json:"country"`
	StateProvince *string `json:"state-province"`
}

type createInstitutionRequest struct {
	Name    string `json:"name"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
}

func main() {
	apiBaseURL := flag.String("api-url", "http://localhost:8080", "Base URL for Sanctor API")
	country := flag.String("country", "Canada", "Country name to import from public API")
	flag.Parse()

	publicURL := fmt.Sprintf("http://universities.hipolabs.com/search?country=%s", url.QueryEscape(*country))

	client := &http.Client{Timeout: 20 * time.Second}

	upstreamResp, err := client.Get(publicURL)
	if err != nil {
		log.Fatalf("failed to fetch public institutions: %v", err)
	}
	defer upstreamResp.Body.Close()

	if upstreamResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(upstreamResp.Body)
		log.Fatalf("public API returned %d: %s", upstreamResp.StatusCode, strings.TrimSpace(string(body)))
	}

	var upstream []upstreamInstitution
	if err := json.NewDecoder(upstreamResp.Body).Decode(&upstream); err != nil {
		log.Fatalf("failed to decode public API response: %v", err)
	}

	if len(upstream) == 0 {
		log.Println("no institutions found from public API")
		return
	}

	// Deduplicate by institution name so repeated records do not spam inserts.
	seen := make(map[string]struct{})
	payloads := make([]createInstitutionRequest, 0, len(upstream))
	for _, inst := range upstream {
		name := strings.TrimSpace(inst.Name)
		if name == "" {
			continue
		}
		if _, exists := seen[name]; exists {
			continue
		}
		seen[name] = struct{}{}

		region := ""
		if inst.StateProvince != nil {
			region = strings.TrimSpace(*inst.StateProvince)
		}

		payloads = append(payloads, createInstitutionRequest{
			Name:    name,
			Country: strings.TrimSpace(inst.Country),
			Region:  region,
		})
	}

	sort.Slice(payloads, func(i, j int) bool {
		return payloads[i].Name < payloads[j].Name
	})

	createURL := strings.TrimRight(*apiBaseURL, "/") + "/api/institutions/create"

	created := 0
	skipped := 0
	failed := 0

	for _, payload := range payloads {
		body, err := json.Marshal(payload)
		if err != nil {
			failed++
			log.Printf("marshal failed for %q: %v", payload.Name, err)
			continue
		}

		req, err := http.NewRequest(http.MethodPost, createURL, bytes.NewReader(body))
		if err != nil {
			failed++
			log.Printf("request build failed for %q: %v", payload.Name, err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			failed++
			log.Printf("request failed for %q: %v", payload.Name, err)
			continue
		}

		respBody, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		respText := strings.TrimSpace(string(respBody))

		switch resp.StatusCode {
		case http.StatusCreated:
			created++
		case http.StatusBadRequest:
			if strings.Contains(strings.ToLower(respText), "already exists") {
				skipped++
			} else {
				failed++
				log.Printf("bad request for %q: %s", payload.Name, respText)
			}
		default:
			failed++
			log.Printf("unexpected status %d for %q: %s", resp.StatusCode, payload.Name, respText)
		}
	}

	log.Printf("seed complete: created=%d skipped=%d failed=%d total=%d", created, skipped, failed, len(payloads))
}
