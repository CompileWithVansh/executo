package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/executo/backend/internal/auth"
)

// contextKey is a custom type for context keys to avoid collisions.
type contextKey string

const (
	// UserIDKey is the context key for the authenticated user's ID.
	UserIDKey contextKey = "user_id"
	// UserEmailKey is the context key for the authenticated user's email.
	UserEmailKey contextKey = "user_email"
	// UserRoleKey is the context key for the authenticated user's role.
	UserRoleKey contextKey = "user_role"
)

// ── RequireAuth ───────────────────────────────
// Blocks the request if no valid JWT token is present.
// Injects user_id, user_email, user_role into request context.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := extractAndValidateToken(r)
		if err != nil {
			writeAuthError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Inject user info into context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ── OptionalAuth ──────────────────────────────
// Attaches user info to context IF a valid token is present.
// Does NOT block the request if no token or invalid token.
func OptionalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := extractAndValidateToken(r)
		if err == nil && claims != nil {
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
			ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

// ── RequireAdmin ──────────────────────────────
// Must be used AFTER RequireAuth. Blocks if user is not an admin.
func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(UserRoleKey).(string)
		if role != "admin" {
			writeAuthError(w, http.StatusForbidden, "admin access required")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ── Context Helpers ───────────────────────────
// Use these in handlers to get the authenticated user's info.

// GetUserID returns the authenticated user's ID from context, or 0 if not authenticated.
func GetUserID(ctx context.Context) int64 {
	id, _ := ctx.Value(UserIDKey).(int64)
	return id
}

// GetUserEmail returns the authenticated user's email from context.
func GetUserEmail(ctx context.Context) string {
	email, _ := ctx.Value(UserEmailKey).(string)
	return email
}

// GetUserRole returns the authenticated user's role from context.
func GetUserRole(ctx context.Context) string {
	role, _ := ctx.Value(UserRoleKey).(string)
	return role
}

// IsAuthenticated returns true if the request has a valid user in context.
func IsAuthenticated(ctx context.Context) bool {
	return GetUserID(ctx) > 0
}

// IsAdmin returns true if the authenticated user is an admin.
func IsAdmin(ctx context.Context) bool {
	return GetUserRole(ctx) == "admin"
}

// ── Internal Helpers ──────────────────────────

// extractAndValidateToken extracts the Bearer token from the Authorization header
// and validates it.
func extractAndValidateToken(r *http.Request) (*auth.Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, &authError{"no authorization header"}
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, &authError{"invalid authorization format (expected: Bearer <token>)"}
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return nil, &authError{"empty token"}
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		return nil, &authError{err.Error()}
	}

	return claims, nil
}

type authError struct {
	message string
}

func (e *authError) Error() string {
	return e.message
}

func writeAuthError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   message,
	})
}
