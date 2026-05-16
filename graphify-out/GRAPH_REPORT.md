# Graph Report - Executo (2026-05-16)

## Corpus Check
- 47 files · ~28,500 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 395 nodes · 478 edges · 31 communities
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                          User Browser                           │
│                    (Next.js 14 Frontend)                        │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTP (port 80)
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                         Nginx Reverse Proxy                     │
│          /api/* → backend:8080   /  → frontend:3000             │
└──────────────┬──────────────────────────┬───────────────────────┘
               │                          │
               ▼                          ▼
┌──────────────────────┐    ┌─────────────────────────┐
│   Go Backend (8080)  │    │  Next.js Frontend (3000) │
│  - REST API          │    │  - Monaco Editor         │
│  - Asynq Worker      │    │  - Problem list/detail   │
│  - Rate limiting     │    │  - Playground            │
│  - Prometheus metrics│    │  - Leaderboard           │
└──────┬───────────────┘    └─────────────────────────┘
       │
       ├──────────────────────────────────┐
       │                                  │
       ▼                                  ▼
┌─────────────────┐            ┌──────────────────────┐
│  PostgreSQL (5432)│           │    Redis (6379)       │
│  - problems      │           │    - Asynq job queue  │
│  - submissions   │           │    - Rate limit state │
│  - suggested_tc  │           └──────────┬───────────┘
└─────────────────┘                       │
                                          ▼
                               ┌──────────────────────┐
                               │   Asynq Worker (Go)  │
                               │   - Dequeues jobs    │
                               │   - Calls Judge0 API │
                               │   - Updates DB       │
                               └──────────┬───────────┘
                                          │
                                          ▼
                               ┌──────────────────────┐
                               │  Judge0 CE (2358)    │
                               │  - Runs code in      │
                               │    Docker containers │
                               │  - Returns verdict   │
                               └──────────────────────┘

┌─────────────────────────────────────────────────────┐
│              Monitoring Stack                        │
│  Prometheus (9090) ← scrapes backend /metrics       │
│  Grafana (3001)    ← visualizes Prometheus data     │
└─────────────────────────────────────────────────────┘
```

## Data Flow: Code Submission

```
User writes code → POST /submissions → Backend validates →
  → Save to PostgreSQL (status: pending) →
  → Enqueue to Redis (Asynq) →
  → Worker picks up job →
  → Worker calls Judge0 for each test case →
  → Judge0 runs in Docker sandbox →
  → Worker updates PostgreSQL with verdict →
  → Frontend polls GET /submissions/:id →
  → User sees result
```

## Data Flow: Playground Run

```
User writes code → POST /run → Backend forwards to Judge0 →
  → Returns token →
  → Frontend polls GET /run/:token →
  → Judge0 returns stdout/stderr/metrics →
  → User sees output
```

## God Nodes (most connected - core abstractions)
1. `NewRouter()` - 12 edges (wires all handlers + middleware)
2. `Worker` - 10 edges (orchestrates Judge0 execution)
3. `Judge0Client` - 9 edges (code execution interface)
4. `DB` - 8 edges (database connection layer)
5. `writeJSON()` / `writeError()` - 10 edges each (shared response helpers)
6. `ProblemsHandler` - 7 edges (problem CRUD)
7. `SubmissionsHandler` - 7 edges (submission lifecycle)
8. `RunHandler` - 4 edges (playground execution)

## Module Dependency Graph

```
cmd/server/main.go
  ├── internal/api/router.go
  │     ├── internal/api/handlers/problems.go
  │     ├── internal/api/handlers/submissions.go
  │     ├── internal/api/handlers/testcases.go
  │     ├── internal/api/handlers/run.go
  │     ├── internal/api/handlers/health.go
  │     ├── internal/api/middleware/cors.go
  │     ├── internal/api/middleware/metrics.go
  │     ├── internal/api/middleware/ratelimit.go
  │     └── internal/executor/judge0.go
  ├── internal/db/db.go
  ├── internal/executor/judge0.go
  └── internal/queue/worker.go
        ├── internal/queue/tasks.go
        ├── internal/executor/judge0.go
        └── internal/models/*.go
```

## Frontend Component Tree

```
app/layout.tsx (RootLayout)
  ├── components/Navbar.tsx
  ├── app/page.tsx (HomePage - landing)
  ├── app/playground/page.tsx (PlaygroundPage)
  │     └── components/Editor.tsx (Monaco wrapper)
  ├── app/problems/page.tsx (ProblemsPage)
  │     └── components/ProblemList.tsx
  ├── app/problems/[id]/page.tsx (ProblemDetailPage)
  │     ├── components/Editor.tsx
  │     └── components/SubmissionResult.tsx
  └── app/leaderboard/page.tsx (LeaderboardPage)
```

## API Endpoints

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | /health | HealthHandler | Service health check |
| POST | /run | RunHandler.CreateRun | Playground code execution |
| GET | /run/:token | RunHandler.GetRunResult | Poll playground result |
| GET | /problems | ProblemsHandler.ListProblems | List problems (paginated) |
| GET | /problems/:id | ProblemsHandler.GetProblem | Get problem by ID |
| GET | /problems/slug/:slug | ProblemsHandler.GetProblemBySlug | Get problem by slug |
| GET | /problems/:slug/testcases | TestCasesHandler.GetApprovedTestCases | Public test cases |
| POST | /problems/:slug/suggest-testcase | TestCasesHandler.SuggestTestCase | Suggest test case |
| POST | /submissions | SubmissionsHandler.CreateSubmission | Submit code |
| GET | /submissions/:id | SubmissionsHandler.GetSubmission | Poll submission result |
| GET | /submissions?problem_id= | SubmissionsHandler.ListSubmissions | List submissions |
| GET | /leaderboard | SubmissionsHandler.GetLeaderboard | Leaderboard |
| GET | /stats | ProblemsHandler.GetStats | Platform statistics |
| GET | /admin/testcases | TestCasesHandler.AdminListTestCases | Admin: pending suggestions |
| PATCH | /admin/testcases/:id | TestCasesHandler.AdminPatchTestCase | Admin: approve/reject |
| GET | /metrics | promhttp.Handler | Prometheus metrics |

## Database Schema

```
problems
  ├── id (BIGSERIAL PK)
  ├── title, slug, description, difficulty
  ├── examples (JSON), constraints (JSON)
  ├── test_cases (JSON - hidden from users)
  ├── function_signature (JSON)
  ├── lc_number, lc_url
  ├── total_submissions, accepted_submissions
  └── created_at, updated_at

submissions
  ├── id (BIGSERIAL PK)
  ├── problem_id (FK → problems)
  ├── language, source_code, status
  ├── verdict, stdout, stderr, compile_output
  ├── runtime_ms, memory_kb
  ├── test_cases_passed, test_cases_total
  └── created_at, updated_at

suggested_testcases
  ├── id (BIGSERIAL PK)
  ├── problem_id (FK → problems)
  ├── input, expected_output, note
  ├── status (pending/approved/rejected)
  ├── submitted_by, admin_note
  └── created_at, updated_at
```

## Technology Stack

| Layer | Technology | Version |
|-------|-----------|---------|
| Frontend | Next.js | 14.2.3 |
| UI | React + TypeScript | 18.3 / 5.4 |
| Styling | Tailwind CSS | 3.4 |
| Editor | Monaco Editor | 0.48 |
| HTTP Client | Axios | 1.7 |
| Backend | Go (net/http) | 1.22 |
| Queue | Asynq (Redis) | 0.24 |
| Database | PostgreSQL | 15 |
| Cache/Queue | Redis | 7 |
| Code Execution | Judge0 CE | 1.13.1 |
| Reverse Proxy | Nginx | 1.25 |
| Monitoring | Prometheus + Grafana | 2.51 / 10.4 |
| Containerization | Docker Compose | v2 |

## Communities (Key Clusters)

### Backend Core (Community 6, 9, 14)
Router → Handlers → DB → Models
- High cohesion, well-structured dependency chain
- `NewRouter()` is the main bridge connecting all backend modules

### Frontend Pages (Community 0, 19)
Pages → Components → API Client → Types
- Each page is relatively self-contained
- Shared components (Editor, Navbar) provide consistency

### Execution Pipeline (Community 4, 11)
Worker → Judge0Client → Submit/Poll → Results
- Clean separation between queue management and execution
- Judge0Client handles all HTTP communication with Judge0

### Infrastructure (Community 5, 12, 13)
Docker Compose → Services → Config → Monitoring
- Well-documented setup process
- Clear service boundaries

## Bugs Fixed (2026-05-16)

1. **Missing `/run` endpoint** — Playground was calling a non-existent route (CRITICAL)
2. **SubmissionRateLimit creating new limiter per request** — Rate limiting was ineffective
3. **Judge0 database user mismatch** — Judge0 couldn't connect to PostgreSQL
4. **Missing seed.sql** — `make seed` would fail
5. **Deprecated docker-compose version key** — Caused warnings
6. **Playground API URL construction** — Double `/api/api/` prefix in some configs

## Suggested Improvements

- Add user authentication (JWT is configured but not enforced)
- Add WebSocket for real-time submission status (replace polling)
- Add Redis-backed rate limiting for horizontal scaling
- Add database connection health check to /health endpoint
- Add request ID middleware for distributed tracing
- Consider adding a `/api/run` route that accepts non-base64 code for simpler testing
