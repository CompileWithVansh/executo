package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/executo/backend/internal/auth"
	"github.com/executo/backend/internal/api/middleware"
	"github.com/executo/backend/internal/db"
	"github.com/executo/backend/internal/models"
)

// AuthHandler handles authentication endpoints.
type AuthHandler struct {
	db *db.DB
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(database *db.DB) *AuthHandler {
	return &AuthHandler{db: database}
}

// ── POST /auth/register ───────────────────────

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Check if email already exists
	var exists bool
	err := h.db.QueryRowContext(r.Context(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.Email,
	).Scan(&exists)
	if err != nil {
		log.Printf("Error checking email existence: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to check email")
		return
	}
	if exists {
		writeError(w, http.StatusConflict, "email already registered")
		return
	}

	// Check if username already exists
	err = h.db.QueryRowContext(r.Context(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", req.Username,
	).Scan(&exists)
	if err != nil {
		log.Printf("Error checking username existence: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to check username")
		return
	}
	if exists {
		writeError(w, http.StatusConflict, "username already taken")
		return
	}

	// Hash password
	passwordHash, err := auth.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create account")
		return
	}

	// Generate avatar URL
	avatarURL := fmt.Sprintf(
		"https://api.dicebear.com/7.x/bottts/svg?seed=%s&backgroundColor=b6e3f4,c0aede,d1d4f9",
		req.Username,
	)

	// Determine role (first user or matching admin email gets admin)
	role := "user"
	var userCount int
	h.db.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM users").Scan(&userCount)
	if userCount == 0 {
		role = "admin" // First user is always admin
	}

	// Insert user
	var userID int64
	err = h.db.QueryRowContext(r.Context(), `
		INSERT INTO users (username, email, password_hash, role, avatar_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, req.Username, req.Email, passwordHash, role, avatarURL).Scan(&userID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create account")
		return
	}

	// Generate JWT
	token, err := auth.GenerateToken(userID, req.Email, role)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	writeJSON(w, http.StatusCreated, models.AuthResponse{
		Success: true,
		Token:   token,
		User: &models.PublicUser{
			ID:        userID,
			Username:  req.Username,
			Email:     req.Email,
			Role:      models.Role(role),
			AvatarURL: avatarURL,
		},
	})
}

// ── POST /auth/login ──────────────────────────

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Find user by email
	var user models.User
	err := h.db.QueryRowContext(r.Context(), `
		SELECT id, username, email, password_hash, role, 
		       COALESCE(avatar_url, ''), COALESCE(bio, ''),
		       problems_solved, total_submissions, is_active
		FROM users WHERE email = $1
	`, req.Email).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role,
		&user.AvatarURL, &user.Bio,
		&user.ProblemsSolved, &user.TotalSubmissions, &user.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusUnauthorized, "invalid credentials")
			return
		}
		log.Printf("Error finding user: %v", err)
		writeError(w, http.StatusInternalServerError, "login failed")
		return
	}

	// Check if account is active
	if !user.IsActive {
		writeError(w, http.StatusForbidden, "account is deactivated")
		return
	}

	// Verify password
	if !auth.CheckPassword(req.Password, user.PasswordHash) {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// Generate JWT
	token, err := auth.GenerateToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		log.Printf("Error generating token: %v", err)
		writeError(w, http.StatusInternalServerError, "login failed")
		return
	}

	writeJSON(w, http.StatusOK, models.AuthResponse{
		Success: true,
		Token:   token,
		User:    user.ToPublic(),
	})
}

// ── GET /auth/me ──────────────────────────────

func (h *AuthHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if userID == 0 {
		writeError(w, http.StatusUnauthorized, "not authenticated")
		return
	}

	var user models.User
	err := h.db.QueryRowContext(r.Context(), `
		SELECT id, username, email, role,
		       COALESCE(avatar_url, ''), COALESCE(bio, ''),
		       problems_solved, total_submissions
		FROM users WHERE id = $1
	`, userID).Scan(
		&user.ID, &user.Username, &user.Email, &user.Role,
		&user.AvatarURL, &user.Bio,
		&user.ProblemsSolved, &user.TotalSubmissions,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "user not found")
			return
		}
		log.Printf("Error fetching user: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to fetch user")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"user":    user.ToPublic(),
	})
}

// ── PUT /auth/profile ─────────────────────────

func (h *AuthHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if userID == 0 {
		writeError(w, http.StatusUnauthorized, "not authenticated")
		return
	}

	var req struct {
		Bio       string `json:"bio"`
		AvatarURL string `json:"avatar_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Limit bio length
	if len(req.Bio) > 160 {
		req.Bio = req.Bio[:160]
	}

	// Build update query dynamically
	// Always update bio (allows clearing it to empty string)
	updates := []string{}
	args := []interface{}{}
	argIdx := 1

	updates = append(updates, fmt.Sprintf("bio = $%d", argIdx))
	args = append(args, req.Bio)
	argIdx++

	if req.AvatarURL != "" {
		updates = append(updates, fmt.Sprintf("avatar_url = $%d", argIdx))
		args = append(args, req.AvatarURL)
		argIdx++
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d",
		strings.Join(updates, ", "), argIdx)
	args = append(args, userID)

	_, err := h.db.ExecContext(r.Context(), query, args...)
	if err != nil {
		log.Printf("Error updating profile: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to update profile")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "profile updated",
	})
}
