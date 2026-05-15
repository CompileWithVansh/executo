-- ─────────────────────────────────────────────
--  SQL Queries: Submissions
-- ─────────────────────────────────────────────

-- name: CreateSubmission :one
INSERT INTO submissions (
    problem_id, language, source_code, status,
    test_cases_total
) VALUES (
    $1, $2, $3, 'pending', $4
)
RETURNING id, status, created_at;

-- name: GetSubmissionByID :one
SELECT
    id, problem_id, language, source_code,
    status, verdict, stdout, stderr, compile_output,
    runtime_ms, memory_kb,
    test_cases_passed, test_cases_total,
    created_at, updated_at
FROM submissions
WHERE id = $1;

-- name: UpdateSubmissionResult :exec
UPDATE submissions
SET
    status           = $2,
    verdict          = $3,
    stdout           = $4,
    stderr           = $5,
    compile_output   = $6,
    runtime_ms       = $7,
    memory_kb        = $8,
    test_cases_passed = $9,
    updated_at       = NOW()
WHERE id = $1;

-- name: UpdateSubmissionStatus :exec
UPDATE submissions
SET status = $2, updated_at = NOW()
WHERE id = $1;

-- name: ListSubmissionsByProblem :many
SELECT
    id, problem_id, language, status,
    runtime_ms, memory_kb,
    test_cases_passed, test_cases_total,
    created_at
FROM submissions
WHERE problem_id = $1
ORDER BY created_at DESC
LIMIT 50;

-- name: GetLeaderboard :many
-- Returns top users by problems solved (requires a users table in production).
-- For now, returns aggregate stats per submission session.
SELECT
    problem_id,
    COUNT(DISTINCT CASE WHEN status = 'accepted' THEN problem_id END) AS problems_solved,
    COUNT(*) AS total_submissions
FROM submissions
GROUP BY problem_id
ORDER BY problems_solved DESC
LIMIT 20;
