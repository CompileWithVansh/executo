// Package queue provides the Asynq worker that processes code execution jobs.
package queue

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hibiken/asynq"

	"github.com/executo/backend/internal/db"
	"github.com/executo/backend/internal/executor"
	"github.com/executo/backend/internal/models"
)

// Worker processes submission execution tasks from the Redis queue.
type Worker struct {
	db     *db.DB
	judge0 *executor.Judge0Client
	server *asynq.Server
}

// NewWorker creates a new Asynq worker connected to Redis.
func NewWorker(database *db.DB, judge0Client *executor.Judge0Client) *Worker {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}

	// Parse Redis URL
	redisOpt, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			// Process up to 10 submissions concurrently
			Concurrency: 10,
			// Queue priorities: critical > default > low
			Queues: map[string]int{
				QueueCritical: 6,
				QueueDefault:  3,
				QueueLow:      1,
			},
			// Log errors
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Printf("Task %s failed: %v", task.Type(), err)
			}),
		},
	)

	return &Worker{
		db:     database,
		judge0: judge0Client,
		server: server,
	}
}

// Start registers task handlers and begins processing jobs.
// This is a blocking call — run it in a goroutine.
func (w *Worker) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeExecuteSubmission, w.handleExecuteSubmission)

	log.Println("✓ Asynq worker started, listening for jobs...")
	return w.server.Run(mux)
}

// Shutdown gracefully stops the worker.
func (w *Worker) Shutdown() {
	w.server.Shutdown()
}

// ── Task Handler ──────────────────────────────

// handleExecuteSubmission processes a single code submission:
// 1. Fetches submission + problem from DB
// 2. Runs each test case through Judge0
// 3. Updates submission with results
func (w *Worker) handleExecuteSubmission(ctx context.Context, task *asynq.Task) error {
	payload, err := ParseExecuteSubmissionPayload(task)
	if err != nil {
		return fmt.Errorf("parsing payload: %w", err)
	}

	submissionID := payload.SubmissionID
	log.Printf("Processing submission #%d", submissionID)

	// Mark as processing
	if err := w.updateStatus(ctx, submissionID, models.StatusProcessing); err != nil {
		return fmt.Errorf("updating status to processing: %w", err)
	}

	// Fetch submission from DB
	submission, err := w.getSubmission(ctx, submissionID)
	if err != nil {
		return w.failSubmission(ctx, submissionID, fmt.Sprintf("fetching submission: %v", err))
	}

	// Fetch problem (for test cases)
	problem, err := w.getProblem(ctx, submission.ProblemID)
	if err != nil {
		return w.failSubmission(ctx, submissionID, fmt.Sprintf("fetching problem: %v", err))
	}

	// Get Judge0 language ID
	languageID, ok := models.Judge0LanguageID[submission.Language]
	if !ok {
		return w.failSubmission(ctx, submissionID, fmt.Sprintf("unsupported language: %s", submission.Language))
	}

	testCases := problem.TestCases
	if len(testCases) == 0 {
		return w.failSubmission(ctx, submissionID, "problem has no test cases")
	}

	// Run all test cases
	result := w.runTestCases(ctx, submission, languageID, testCases)

	// Save final result to DB
	if err := w.saveResult(ctx, submissionID, result, len(testCases)); err != nil {
		log.Printf("Failed to save result for submission #%d: %v", submissionID, err)
		return err
	}

	log.Printf("Submission #%d completed: %s (%d/%d tests passed)",
		submissionID, result.status, result.passed, len(testCases))

	return nil
}

// ── Execution Result ──────────────────────────

type executionResult struct {
	status        models.SubmissionStatus
	verdict       string
	stdout        string
	stderr        string
	compileOutput string
	runtimeMS     *int64
	memoryKB      *int64
	passed        int
}

// runTestCases executes all test cases and returns the aggregate result.
func (w *Worker) runTestCases(
	ctx context.Context,
	submission *models.Submission,
	languageID int,
	testCases models.JSONArray[models.TestCase],
) executionResult {
	var (
		passed    int
		totalMS   float64
		maxMemory int64
		lastResult *executor.Judge0Result
	)

	for i, tc := range testCases {
		// Check context cancellation
		select {
		case <-ctx.Done():
			return executionResult{
				status:  models.StatusInternalError,
				verdict: "execution cancelled",
			}
		default:
		}

		// Submit to Judge0
		token, err := w.judge0.Submit(
			submission.SourceCode,
			languageID,
			tc.Input,
			tc.ExpectedOutput,
		)
		if err != nil {
			log.Printf("Judge0 submit error for submission #%d, test %d: %v",
				submission.ID, i+1, err)
			return executionResult{
				status:  models.StatusInternalError,
				verdict: fmt.Sprintf("judge0 error: %v", err),
			}
		}

		// Poll for result
		result, err := w.judge0.PollUntilDone(token, 30)
		if err != nil {
			return executionResult{
				status:  models.StatusInternalError,
				verdict: fmt.Sprintf("polling error: %v", err),
			}
		}

		lastResult = result

		// Handle compilation error (same for all test cases)
		if result.Status.ID == executor.StatusCompileErr {
			return executionResult{
				status:        models.StatusCompilationError,
				compileOutput: result.CompileOutput,
				stderr:        result.Stderr,
			}
		}

		// Handle runtime errors
		if result.Status.ID >= executor.StatusRuntimeErr && result.Status.ID <= executor.StatusRuntimeErr6 {
			return executionResult{
				status:  models.StatusRuntimeError,
				stderr:  result.Stderr,
				stdout:  result.Stdout,
				verdict: fmt.Sprintf("Runtime error on test case %d: %s", i+1, result.Status.Description),
			}
		}

		// Handle TLE
		if result.Status.ID == executor.StatusTLE {
			return executionResult{
				status:  models.StatusTimeLimitExceeded,
				verdict: fmt.Sprintf("Time limit exceeded on test case %d", i+1),
			}
		}

		// Handle internal error
		if result.Status.ID == executor.StatusInternalErr {
			return executionResult{
				status:  models.StatusInternalError,
				verdict: result.Message,
			}
		}

		// Accumulate metrics
		if ms, err := strconv.ParseFloat(result.Time, 64); err == nil {
			totalMS += ms * 1000 // convert seconds to ms
		}
		if int64(result.Memory) > maxMemory {
			maxMemory = int64(result.Memory)
		}

		// Check result
		if result.Status.ID == executor.StatusAccepted {
			passed++
		} else {
			// Wrong answer — stop running further test cases
			expected := strings.TrimSpace(tc.ExpectedOutput)
			got := strings.TrimSpace(result.Stdout)
			return executionResult{
				status:  models.StatusWrongAnswer,
				verdict: fmt.Sprintf("Test case %d: expected %q, got %q", i+1, expected, got),
				stdout:  result.Stdout,
				passed:  passed,
			}
		}
	}

	// All test cases passed
	avgMS := int64(math.Round(totalMS / float64(len(testCases))))
	stdout := ""
	if lastResult != nil {
		stdout = lastResult.Stdout
	}

	return executionResult{
		status:    models.StatusAccepted,
		stdout:    stdout,
		runtimeMS: &avgMS,
		memoryKB:  &maxMemory,
		passed:    passed,
	}
}

// ── Database Helpers ──────────────────────────

func (w *Worker) getSubmission(ctx context.Context, id int64) (*models.Submission, error) {
	row := w.db.QueryRowContext(ctx, `
		SELECT id, problem_id, language, source_code, status
		FROM submissions WHERE id = $1
	`, id)

	var s models.Submission
	if err := row.Scan(&s.ID, &s.ProblemID, &s.Language, &s.SourceCode, &s.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("submission #%d not found", id)
		}
		return nil, err
	}
	return &s, nil
}

func (w *Worker) getProblem(ctx context.Context, id int64) (*models.Problem, error) {
	row := w.db.QueryRowContext(ctx, `
		SELECT id, test_cases FROM problems WHERE id = $1
	`, id)

	var p models.Problem
	if err := row.Scan(&p.ID, &p.TestCases); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("problem #%d not found", id)
		}
		return nil, err
	}
	return &p, nil
}

func (w *Worker) updateStatus(ctx context.Context, id int64, status models.SubmissionStatus) error {
	_, err := w.db.ExecContext(ctx, `
		UPDATE submissions SET status = $2, updated_at = NOW() WHERE id = $1
	`, id, string(status))
	return err
}

func (w *Worker) failSubmission(ctx context.Context, id int64, reason string) error {
	log.Printf("Submission #%d failed: %s", id, reason)
	_, err := w.db.ExecContext(ctx, `
		UPDATE submissions
		SET status = 'internal_error', verdict = $2, updated_at = NOW()
		WHERE id = $1
	`, id, reason)
	return err
}

func (w *Worker) saveResult(ctx context.Context, id int64, result executionResult, totalTestCases int) error {
	_, err := w.db.ExecContext(ctx, `
		UPDATE submissions SET
			status            = $2,
			verdict           = $3,
			stdout            = $4,
			stderr            = $5,
			compile_output    = $6,
			runtime_ms        = $7,
			memory_kb         = $8,
			test_cases_passed = $9,
			test_cases_total  = $10,
			updated_at        = NOW()
		WHERE id = $1
	`,
		id,
		string(result.status),
		nullString(result.verdict),
		nullString(result.stdout),
		nullString(result.stderr),
		nullString(result.compileOutput),
		result.runtimeMS,
		result.memoryKB,
		result.passed,
		totalTestCases,
	)
	return err
}

// nullString returns nil for empty strings (for nullable DB columns).
func nullString(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

// ── Client (Enqueuer) ─────────────────────────

// Client enqueues tasks to Redis for the worker to process.
type Client struct {
	client *asynq.Client
}

// NewClient creates a new Asynq client for enqueuing tasks.
func NewClient() *Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}

	redisOpt, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	return &Client{
		client: asynq.NewClient(redisOpt),
	}
}

// EnqueueSubmission adds a submission execution task to the queue.
func (c *Client) EnqueueSubmission(submissionID int64) error {
	task, err := NewExecuteSubmissionTask(submissionID)
	if err != nil {
		return fmt.Errorf("creating task: %w", err)
	}

	info, err := c.client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("enqueuing task: %w", err)
	}

	log.Printf("Enqueued submission #%d (task ID: %s, queue: %s)",
		submissionID, info.ID, info.Queue)
	return nil
}

// Close closes the Asynq client connection.
func (c *Client) Close() error {
	return c.client.Close()
}

// ── Retry Helper ──────────────────────────────

// withRetry retries a function up to maxAttempts times with exponential backoff.
func withRetry(maxAttempts int, fn func() error) error {
	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if err := fn(); err != nil {
			lastErr = err
			if attempt < maxAttempts-1 {
				time.Sleep(time.Duration(attempt+1) * 500 * time.Millisecond)
			}
			continue
		}
		return nil
	}
	return lastErr
}
