package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strings"
)

// HashPassword creates a salted SHA-256 hash of the password.
// Format: "salt:hash" where salt is 16 random bytes (hex-encoded).
//
// Note: For production at scale, use bcrypt or argon2. SHA-256 with salt
// is used here to avoid CGO dependencies (bcrypt in Go uses CGO on some platforms).
// The salt prevents rainbow table attacks.
func HashPassword(password string) (string, error) {
	// Generate 16-byte random salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generating salt: %w", err)
	}

	saltHex := hex.EncodeToString(salt)
	hash := hashWithSalt(password, saltHex)

	return saltHex + ":" + hash, nil
}

// CheckPassword verifies a password against a stored hash.
// Returns true if the password matches.
func CheckPassword(password, storedHash string) bool {
	parts := strings.SplitN(storedHash, ":", 2)
	if len(parts) != 2 {
		return false
	}

	salt := parts[0]
	expectedHash := parts[1]
	actualHash := hashWithSalt(password, salt)

	// Constant-time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare([]byte(expectedHash), []byte(actualHash)) == 1
}

// hashWithSalt performs SHA-256(salt + password) with multiple iterations.
func hashWithSalt(password, salt string) string {
	// Use 10000 iterations for key stretching
	data := []byte(salt + password)
	for i := 0; i < 10000; i++ {
		h := sha256.Sum256(data)
		data = h[:]
	}
	return hex.EncodeToString(data)
}
