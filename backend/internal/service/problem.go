// Package service contains business logic for the Executo backend.
// Services orchestrate between repositories, external APIs, and domain rules.
package service

import (
	"context"

	"github.com/executo/backend/internal/models"
	"github.com/executo/backend/internal/repository"
)

// ProblemService handles business logic for problems.
type ProblemService struct {
	repo repository.ProblemRepository
}

// NewProblemService creates a new ProblemService.
func NewProblemService(repo repository.ProblemRepository) *ProblemService {
	return &ProblemService{repo: repo}
}

// List returns a paginated list of problems.
func (s *ProblemService) List(ctx context.Context, filter repository.ProblemFilter) ([]models.ProblemSummary, int, error) {
	// Enforce page size limits
	if filter.PageSize < 1 || filter.PageSize > 100 {
		filter.PageSize = 20
	}
	if filter.Page < 1 {
		filter.Page = 1
	}
	return s.repo.List(ctx, filter)
}

// GetByID returns a single problem (hides test cases from response).
func (s *ProblemService) GetByID(ctx context.Context, id int64) (*models.Problem, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// Never expose test cases to the client
	p.TestCases = nil
	return p, nil
}

// GetBySlug returns a single problem by slug.
func (s *ProblemService) GetBySlug(ctx context.Context, slug string) (*models.Problem, error) {
	p, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	p.TestCases = nil
	return p, nil
}

// GetStats returns platform-wide statistics.
func (s *ProblemService) GetStats(ctx context.Context) (*repository.PlatformStats, error) {
	return s.repo.GetPlatformStats(ctx)
}
