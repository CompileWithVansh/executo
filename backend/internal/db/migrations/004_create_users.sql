-- ─────────────────────────────────────────────
--  Migration 004: Create users table + auth
-- ─────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY,
    username        VARCHAR(50) NOT NULL UNIQUE,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   TEXT NOT NULL,
    
    -- Role: 'user' or 'admin'
    role            VARCHAR(10) NOT NULL DEFAULT 'user'
                    CHECK (role IN ('user', 'admin')),
    
    -- Profile
    avatar_url      TEXT,
    bio             VARCHAR(160),
    
    -- Stats (denormalized for fast reads)
    problems_solved INT NOT NULL DEFAULT 0,
    total_submissions INT NOT NULL DEFAULT 0,
    
    -- Account status
    is_active       BOOLEAN NOT NULL DEFAULT true,
    
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

-- Auto-update updated_at
DROP TRIGGER IF EXISTS users_updated_at ON users;
CREATE TRIGGER users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- ─────────────────────────────────────────────
--  Add user_id to submissions (link submissions to users)
-- ─────────────────────────────────────────────

ALTER TABLE submissions
    ADD COLUMN IF NOT EXISTS user_id BIGINT REFERENCES users(id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_submissions_user_id ON submissions (user_id);
