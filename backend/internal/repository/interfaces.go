// Package repository defines data access interfaces for the Executo backend.
// Implementations live in sub-packages (e.g. repository/postgres).
// Services depend on these interfaces, not on concrete implementations.
package repository

import (
	"context"

	"github.com/executo/backend/internal/models"
)

// ── Problem Repository ────────────────────────

// ProblemRepository defines data access methods for problems.
type ProblemRepository interface {
	// GetByID returns a single problem by its database ID.
	GetByID(ctx context.Context, id int64) (*models.Problem, error)

	// GetBySlug returns a single problem by its URL slug.
	GetBySlug(ctx context.Context, slug string) (*models.Problem, error)

	// List returns a paginated, filterable list of problem summaries.
	List(ctx context.Context, filter ProblemFilter) ([]models.ProblemSummary, int, error)

	// Create inserts a new problem and returns its ID.
	Create(ctx context.Context, req *models.CreateProblemRequest) (int64, error)

	// UpdateStats increments submission/accepted counters.
	UpdateStats(ctx context.Context, problemID int64, accepted bool) error

	// GetPlatformStats returns aggregate platform statistics.
	GetPlatformStats(ctx context.Context) (*PlatformStats, error)
}

// ProblemFilter holds filtering/pagination options for listing problems.
type ProblemFilter struct {
	Difficulty string // "easy", "medium", "hard", or "" for all
	Search     string // title search (ILIKE)
	Page       int
	PageSize   int
}

// PlatformStats holds aggregate platform statistics.
type PlatformStats struct {
	TotalProblems    int64   `json:"total_problems"`
	TotalSubmissions int64   `json:"total_submissions"`
	TotalAccepted    int64   `json:"total_accepted"`
	AcceptanceRate   float64 `json:"acceptance_rate"`
}

// ── Submission Repository ─────────────────────

// SubmissionRepository defines data access methods for submissions.
type SubmissionRepository interface {
	// Create inserts a new submission and returns its ID.
	Create(ctx context.Context, req CreateSubmissionParams) (int64, error)

	// GetByID returns a single submission by ID.
	GetByID(ctx context.Context, id int64) (*models.Submission, error)

	// ListByProblem returns recent submissions for a problem.
	ListByProblem(ctx context.Context, problemID int64, limit int) ([]models.Submission, error)

	// UpdateResult updates a submission with execution results.
	UpdateResult(ctx context.Context, id int64, result SubmissionResult) error

	// UpdateStatus updates just the status field.
	UpdateStatus(ctx context.Context, id int64, status models.SubmissionStatus) error
}

// CreateSubmissionParams holds the data needed to create a submission.
type CreateSubmissionParams struct {
	ProblemID      int64
	Language       models.Language
	SourceCode     string
	TestCasesTotal int
}

// SubmissionResult holds the execution result to save.
type SubmissionResult struct {
	Status          models.SubmissionStatus
	Verdict         string
	Stdout          string
	Stderr          string
	CompileOutput   string
	RuntimeMS       *int64
	MemoryKB        *int64
	TestCasesPassed int
}

// ── TestCase Repository ───────────────────────

// TestCaseRepository defines data access for user-suggested test cases.
type TestCaseRepository interface {
	// Suggest creates a new pending test case suggestion.
	Suggest(ctx context.Context, problemID int64, input, expectedOutput, note, submittedBy string) (int64, error)

	// GetApproved returns approved test cases for a problem.
	GetApproved(ctx context.Context, problemSlug string) ([]SuggestedTestCase, error)

	// ListPending returns all pending suggestions (admin).
	ListPending(ctx context.Context, status string) ([]AdminTestCase, error)

	// Patch updates the status of a suggestion (approve/reject).
	Patch(ctx context.Context, id int64, status, adminNote string) error
}

// SuggestedTestCase is a public-facing approved test case.
type SuggestedTestCase struct {
	ID             int64  `json:"id"`
	ProblemID      int64  `json:"problem_id"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
	Note           string `json:"note,omitempty"`
}

// AdminTestCase is a test case with admin-visible fields.
type AdminTestCase struct {
	ID             int64  `json:"id"`
	ProblemID      int64  `json:"problem_id"`
	ProblemTitle   string `json:"problem_title"`
	ProblemSlug    string `json:"problem_slug"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
	Note           string `json:"note,omitempty"`
	Status         string `json:"status"`
	SubmittedBy    string `json:"submitted_by,omitempty"`
	AdminNote      string `json:"admin_note,omitempty"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
