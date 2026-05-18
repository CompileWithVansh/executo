package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/executo/backend/internal/domain"
	"github.com/executo/backend/internal/models"
	"github.com/executo/backend/internal/queue"
	"github.com/executo/backend/internal/repository"
)

// SubmissionService handles business logic for code submissions.
type SubmissionService struct {
	submissionRepo repository.SubmissionRepository
	problemRepo    repository.ProblemRepository
	queue          *queue.Client
}

// NewSubmissionService creates a new SubmissionService.
func NewSubmissionService(
	submissionRepo repository.SubmissionRepository,
	problemRepo repository.ProblemRepository,
	queue *queue.Client,
) *SubmissionService {
	return &SubmissionService{
		submissionRepo: submissionRepo,
		problemRepo:    problemRepo,
		queue:          queue,
	}
}

// Create validates and creates a new submission, then enqueues it for execution.
func (s *SubmissionService) Create(ctx context.Context, req *models.CreateSubmissionRequest) (*models.CreateSubmissionResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, &domain.ErrValidation{Message: err.Error()}
	}

	// Verify problem exists and count test cases
	problem, err := s.problemRepo.GetByID(ctx, req.ProblemID)
	if err != nil {
		return nil, err
	}

	testCasesTotal := 0
	if problem.TestCases != nil {
		var testCases []models.TestCase
		// TestCases is a JSONArray, marshal/unmarshal to count
		data, _ := json.Marshal(problem.TestCases)
		if json.Unmarshal(data, &testCases) == nil {
			testCasesTotal = len(testCases)
		}
	}

	// Create submission record
	submissionID, err := s.submissionRepo.Create(ctx, repository.CreateSubmissionParams{
		ProblemID:      req.ProblemID,
		Language:       req.Language,
		SourceCode:     req.SourceCode,
		TestCasesTotal: testCasesTotal,
	})
	if err != nil {
		return nil, fmt.Errorf("creating submission: %w", err)
	}

	// Enqueue for async execution (non-blocking — submission is saved even if queue fails)
	if err := s.queue.EnqueueSubmission(submissionID); err != nil {
		// Log but don't fail — a retry mechanism can pick this up
		fmt.Printf("Warning: failed to enqueue submission #%d: %v\n", submissionID, err)
	}

	return &models.CreateSubmissionResponse{
		ID:      submissionID,
		Status:  models.StatusPending,
		Message: "Submission received and queued for execution",
	}, nil
}

// GetByID returns a submission by ID (used for polling).
func (s *SubmissionService) GetByID(ctx context.Context, id int64) (*models.Submission, error) {
	return s.submissionRepo.GetByID(ctx, id)
}

// ListByProblem returns recent submissions for a problem.
func (s *SubmissionService) ListByProblem(ctx context.Context, problemID int64) ([]models.Submission, error) {
	return s.submissionRepo.ListByProblem(ctx, problemID, 50)
}
