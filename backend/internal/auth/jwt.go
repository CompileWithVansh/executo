// Package auth provides JWT token generation and validation for Executo.
package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Claims represents the JWT payload.
type Claims struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IssuedAt int64  `json:"iat"`
	ExpiresAt int64 `json:"exp"`
}

// IsExpired returns true if the token has expired.
func (c *Claims) IsExpired() bool {
	return time.Now().Unix() > c.ExpiresAt
}

// GenerateToken creates a new JWT token for a user.
// Token expires in 7 days by default.
func GenerateToken(userID int64, email, role string) (string, error) {
	secret := getSecret()
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET is not configured")
	}

	now := time.Now()
	claims := Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(7 * 24 * time.Hour).Unix(), // 7 days
	}

	// Header
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, _ := json.Marshal(header)
	headerB64 := base64URLEncode(headerJSON)

	// Payload
	payloadJSON, _ := json.Marshal(claims)
	payloadB64 := base64URLEncode(payloadJSON)

	// Signature
	signingInput := headerB64 + "." + payloadB64
	signature := sign(signingInput, secret)

	return signingInput + "." + signature, nil
}

// ValidateToken parses and validates a JWT token.
// Returns the claims if valid, or an error if invalid/expired.
func ValidateToken(tokenString string) (*Claims, error) {
	secret := getSecret()
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not configured")
	}

	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	// Verify signature
	signingInput := parts[0] + "." + parts[1]
	expectedSig := sign(signingInput, secret)
	if !hmac.Equal([]byte(parts[2]), []byte(expectedSig)) {
		return nil, fmt.Errorf("invalid token signature")
	}

	// Decode payload
	payloadJSON, err := base64URLDecode(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid token payload: %w", err)
	}

	var claims Claims
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return nil, fmt.Errorf("invalid token claims: %w", err)
	}

	// Check expiration
	if claims.IsExpired() {
		return nil, fmt.Errorf("token expired")
	}

	return &claims, nil
}

// ── Helpers ───────────────────────────────────

func getSecret() string {
	return os.Getenv("JWT_SECRET")
}

func sign(input, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(input))
	return base64URLEncode(mac.Sum(nil))
}

func base64URLEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}

func base64URLDecode(s string) ([]byte, error) {
	// Add padding back
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}
