# ─────────────────────────────────────────────
#  Executo Backend — Production Dockerfile
#  Multi-stage build: builder → minimal runtime
# ─────────────────────────────────────────────

# ── Stage 1: Build ────────────────────────────
FROM golang:1.22-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copy go.mod and go.sum first for layer caching
COPY go.mod go.sum ./
RUN go mod download || go mod tidy

# Copy source code
COPY . .

# Build the binary
# CGO_ENABLED=0 for a fully static binary
# -ldflags="-s -w" strips debug info to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w -X main.version=$(git describe --tags --always --dirty 2>/dev/null || echo dev)" \
    -o /app/bin/executo \
    ./cmd/server

# ── Stage 2: Runtime ──────────────────────────
FROM scratch

# Copy timezone data and CA certificates from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /app/bin/executo /executo

# Copy migration files (embedded at runtime)
COPY --from=builder /app/internal/db/migrations /migrations

# Expose the API port
EXPOSE 8080

# Run as non-root user
USER 65534:65534

ENTRYPOINT ["/executo"]
