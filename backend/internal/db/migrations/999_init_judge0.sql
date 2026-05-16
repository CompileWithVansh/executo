-- ─────────────────────────────────────────────
--  Create Judge0 database and user
--  This runs during PostgreSQL container init
-- ─────────────────────────────────────────────

-- Create the judge0 user (if not exists)
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'judge0') THEN
        CREATE ROLE judge0 WITH LOGIN PASSWORD 'judge0password';
    END IF;
END
$$;

-- Create the judge0 database (must use separate connection, so we use a workaround)
-- Note: CREATE DATABASE cannot run inside a transaction block in initdb scripts,
-- but PostgreSQL's docker-entrypoint handles this by running .sql files one at a time.
SELECT 'CREATE DATABASE judge0 OWNER judge0'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'judge0')\gexec
