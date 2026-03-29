# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Ensure go.sum is up-to-date
RUN go mod tidy

# Build binaries for API server and migrations
RUN go build -o main ./cmd/api
RUN go build -o migrate ./cmd/migrate

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Set the entrypoint to the compiled binary
ENTRYPOINT ["/root/main"]
