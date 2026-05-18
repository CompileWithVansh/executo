package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Role represents a user's permission level.
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

// User represents a registered user.
type User struct {
	ID               int64     `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"-"` // never expose in JSON
	Role             Role      `json:"role"`
	AvatarURL        string    `json:"avatar_url,omitempty"`
	Bio              string    `json:"bio,omitempty"`
	ProblemsSolved   int       `json:"problems_solved"`
	TotalSubmissions int       `json:"total_submissions"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// PublicUser is the safe-to-expose version of User (no sensitive fields).
type PublicUser struct {
	ID               int64  `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Role             Role   `json:"role"`
	AvatarURL        string `json:"avatar_url,omitempty"`
	Bio              string `json:"bio,omitempty"`
	ProblemsSolved   int    `json:"problems_solved"`
	TotalSubmissions int    `json:"total_submissions"`
}

// ToPublic converts a User to its public representation.
func (u *User) ToPublic() *PublicUser {
	return &PublicUser{
		ID:               u.ID,
		Username:         u.Username,
		Email:            u.Email,
		Role:             u.Role,
		AvatarURL:        u.AvatarURL,
		Bio:              u.Bio,
		ProblemsSolved:   u.ProblemsSolved,
		TotalSubmissions: u.TotalSubmissions,
	}
}

// IsAdmin returns true if the user has admin role.
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// ── Registration Request ──────────────────────

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,30}$`)
)

func (r *RegisterRequest) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))

	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	if !usernameRegex.MatchString(r.Username) {
		return fmt.Errorf("username must be 3-30 characters (letters, numbers, _ or -)")
	}
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !emailRegex.MatchString(r.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(r.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}
	if len(r.Password) > 128 {
		return fmt.Errorf("password must be at most 128 characters")
	}
	return nil
}

// ── Login Request ─────────────────────────────

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	if r.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

// ── Auth Response ─────────────────────────────

type AuthResponse struct {
	Success bool        `json:"success"`
	Token   string      `json:"token"`
	User    *PublicUser `json:"user"`
}
