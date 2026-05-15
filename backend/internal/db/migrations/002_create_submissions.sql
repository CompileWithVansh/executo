-- ─────────────────────────────────────────────
--  Migration 002: Create submissions table
-- ─────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS submissions (
    id                  BIGSERIAL PRIMARY KEY,
    problem_id          BIGINT NOT NULL REFERENCES problems(id) ON DELETE CASCADE,

    -- Language identifier (python3, java, cpp, javascript)
    language            VARCHAR(20) NOT NULL,

    -- The submitted source code
    source_code         TEXT NOT NULL,

    -- Current status of the submission
    status              VARCHAR(30) NOT NULL DEFAULT 'pending'
                        CHECK (status IN (
                            'pending', 'processing', 'accepted',
                            'wrong_answer', 'time_limit_exceeded',
                            'memory_limit_exceeded', 'runtime_error',
                            'compilation_error', 'internal_error'
                        )),

    -- Human-readable verdict detail (e.g. "Expected: [0,1] Got: [1,0]")
    verdict             TEXT,

    -- Captured stdout from the program
    stdout              TEXT,

    -- Captured stderr from the program
    stderr              TEXT,

    -- Compiler output (for compilation errors)
    compile_output      TEXT,

    -- Execution metrics
    runtime_ms          BIGINT,
    memory_kb           BIGINT,

    -- Test case results
    test_cases_passed   INT NOT NULL DEFAULT 0,
    test_cases_total    INT NOT NULL DEFAULT 0,

    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for polling by submission ID (most common query)
CREATE INDEX IF NOT EXISTS idx_submissions_id ON submissions (id);

-- Index for fetching a user's submissions for a problem
CREATE INDEX IF NOT EXISTS idx_submissions_problem_id ON submissions (problem_id);

-- Index for status-based queries (e.g. finding pending submissions)
CREATE INDEX IF NOT EXISTS idx_submissions_status ON submissions (status);

-- Trigger to auto-update updated_at
DROP TRIGGER IF EXISTS submissions_updated_at ON submissions;
CREATE TRIGGER submissions_updated_at
    BEFORE UPDATE ON submissions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update problem stats when a submission is finalized
CREATE OR REPLACE FUNCTION update_problem_stats()
RETURNS TRIGGER AS $$
BEGIN
    -- Only update when status changes to a terminal state
    IF NEW.status != OLD.status AND NEW.status NOT IN ('pending', 'processing') THEN
        UPDATE problems
        SET
            total_submissions = total_submissions + 1,
            accepted_submissions = accepted_submissions + CASE WHEN NEW.status = 'accepted' THEN 1 ELSE 0 END
        WHERE id = NEW.problem_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS submissions_update_problem_stats ON submissions;
CREATE TRIGGER submissions_update_problem_stats
    AFTER UPDATE ON submissions
    FOR EACH ROW
    EXECUTE FUNCTION update_problem_stats();
