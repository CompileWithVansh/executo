package service

import (
	"fmt"

	"github.com/executo/backend/internal/domain"
	"github.com/executo/backend/internal/executor"
)

// PlaygroundService handles direct code execution (no problem context).
type PlaygroundService struct {
	judge0 *executor.Judge0Client
}

// NewPlaygroundService creates a new PlaygroundService.
func NewPlaygroundService(judge0 *executor.Judge0Client) *PlaygroundService {
	return &PlaygroundService{judge0: judge0}
}

// RunRequest is the input for running code in the playground.
type RunRequest struct {
	SourceCodeB64 string // base64-encoded source code
	LanguageID    int
	StdinB64      string // base64-encoded stdin (optional)
}

// Validate checks the run request.
func (r *RunRequest) Validate() error {
	if r.SourceCodeB64 == "" {
		return &domain.ErrValidation{Field: "source_code", Message: "is required"}
	}
	if r.LanguageID == 0 {
		return &domain.ErrValidation{Field: "language_id", Message: "is required"}
	}
	return nil
}

// Submit sends code to Judge0 and returns a polling token.
func (s *PlaygroundService) Submit(req *RunRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	token, err := s.judge0.SubmitRaw(req.SourceCodeB64, req.LanguageID, req.StdinB64)
	if err != nil {
		return "", fmt.Errorf("judge0 submit: %w", err)
	}
	return token, nil
}

// GetResult polls Judge0 for the result of a run.
func (s *PlaygroundService) GetResult(token string) (*executor.Judge0Result, error) {
	if token == "" {
		return nil, &domain.ErrValidation{Field: "token", Message: "is required"}
	}
	return s.judge0.GetResult(token)
}
