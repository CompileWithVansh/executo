package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const (
	// RequestIDHeader is the HTTP header for request tracing.
	RequestIDHeader = "X-Request-ID"
	// RequestIDKey is the context key for the request ID.
	RequestIDKey contextKey = "request_id"
)

// RequestID generates a unique request ID for each request.
// If the client sends X-Request-ID, it's reused (for distributed tracing).
// Otherwise, a new random ID is generated.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use existing request ID from header, or generate a new one
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = generateID()
		}

		// Set in response header for client-side tracing
		w.Header().Set(RequestIDHeader, requestID)

		// Inject into context for use in handlers/logs
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID returns the request ID from context.
func GetRequestID(ctx context.Context) string {
	id, _ := ctx.Value(RequestIDKey).(string)
	return id
}

// generateID creates a random 8-byte hex string (16 chars).
func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}
