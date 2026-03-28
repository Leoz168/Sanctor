package search

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// SimpleTokenizer is a baseline tokenizer for vertical search.
type SimpleTokenizer struct {
	cfg TokenizerConfig
}

// NewSimpleTokenizer constructs a SimpleTokenizer with safe defaults.
func NewSimpleTokenizer(cfg TokenizerConfig) *SimpleTokenizer {
	if cfg.MinLen <= 0 {
		cfg.MinLen = 2
	}
	if cfg.Stopwords == nil {
		cfg.Stopwords = DefaultStopwords()
	}

	return &SimpleTokenizer{cfg: cfg}
}

// Tokenize normalizes text, splits into terms, and filters low-value tokens.
func (t *SimpleTokenizer) Tokenize(input string) []Token {
	normalized := strings.ToLower(input)

	tokens := make([]Token, 0, 16)
	var b strings.Builder
	position := 0

	flush := func() {
		if b.Len() == 0 {
			return
		}

		term := b.String()
		b.Reset()

		if utf8.RuneCountInString(term) < t.cfg.MinLen {
			return
		}
		if _, ok := t.cfg.Stopwords[term]; ok {
			return // skip stopwords
		}

		tokens = append(tokens, Token{Term: term, Pos: position})
		position++
	}

	for _, r := range normalized {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
		case r == '#' && t.cfg.KeepHashtags && b.Len() == 0:
			b.WriteRune(r)
		case r == '@' && t.cfg.KeepMentions && b.Len() == 0:
			b.WriteRune(r)
		default:
			flush()
		}
	}

	flush()
	return tokens
}
