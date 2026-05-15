package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Prometheus metrics for HTTP requests
var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	submissionsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "executo_submissions_total",
			Help: "Total number of code submissions by status",
		},
		[]string{"status", "language"},
	)
)

// responseWriter wraps http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Metrics returns a middleware that records Prometheus metrics for each request.
func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(rw.statusCode)

		// Normalize path to avoid high cardinality (e.g. /submissions/123 → /submissions/:id)
		path := normalizePath(r.URL.Path)

		httpRequestsTotal.WithLabelValues(r.Method, path, status).Inc()
		httpRequestDuration.WithLabelValues(r.Method, path).Observe(duration)
	})
}

// RecordSubmission records a submission result in Prometheus.
func RecordSubmission(status, language string) {
	submissionsTotal.WithLabelValues(status, language).Inc()
}

// normalizePath replaces numeric path segments with ":id" to reduce cardinality.
func normalizePath(path string) string {
	// Simple normalization: replace /submissions/123 with /submissions/:id
	// A proper implementation would use a router-aware approach
	result := make([]byte, 0, len(path))
	i := 0
	for i < len(path) {
		if path[i] == '/' {
			result = append(result, '/')
			i++
			// Check if next segment is all digits
			j := i
			for j < len(path) && path[j] != '/' {
				j++
			}
			segment := path[i:j]
			if isNumeric(segment) {
				result = append(result, []byte(":id")...)
			} else {
				result = append(result, []byte(segment)...)
			}
			i = j
		} else {
			i++
		}
	}
	if len(result) == 0 {
		return "/"
	}
	return string(result)
}

func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
