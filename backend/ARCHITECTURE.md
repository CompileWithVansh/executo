# Backend Architecture

## Current vs Proposed Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go              # Entry point — wires everything together
│
├── internal/
│   ├── config/                  # NEW — centralized configuration
│   │   └── config.go           # Reads env vars once, validates, exposes typed struct
│   │
│   ├── domain/                  # NEW — pure domain types (no dependencies)
│   │   ├── problem.go          # Problem, Example, TestCase types
│   │   ├── submission.go       # Submission, Language, Status types
│   │   └── errors.go           # Domain-specific error types (NotFound, Validation, etc.)
│   │
│   ├── repository/             # NEW — data access layer (interfaces + implementations)
│   │   ├── interfaces.go       # ProblemRepo, SubmissionRepo interfaces
│   │   ├── postgres/           # PostgreSQL implementations
│   │   │   ├── problem.go
│   │   │   ├── submission.go
│   │   │   └── testcase.go
│   │   └── db.go               # Connection pool setup
│   │
│   ├── service/                # NEW — business logic layer
│   │   ├── problem.go          # ProblemService (list, get, create, stats)
│   │   ├── submission.go       # SubmissionService (create, poll, history)
│   │   ├── playground.go       # PlaygroundService (run code directly)
│   │   └── leaderboard.go      # LeaderboardService
│   │
│   ├── handler/                # HTTP handlers — thin, just parse request → call service → write response
│   │   ├── problem.go
│   │   ├── submission.go
│   │   ├── playground.go
│   │   ├── testcase.go
│   │   ├── health.go
│   │   └── response.go         # Shared JSON response helpers
│   │
│   ├── middleware/             # HTTP middleware (unchanged, already clean)
│   │   ├── cors.go
│   │   ├── metrics.go
│   │   ├── ratelimit.go
│   │   └── auth.go             # Future: JWT validation
│   │
│   ├── router/                 # NEW — route registration separated from handlers
│   │   └── router.go
│   │
│   ├── executor/               # Judge0 client (unchanged, already clean)
│   │   └── judge0.go
│   │
│   └── queue/                  # Asynq worker (unchanged, already clean)
│       ├── tasks.go
│       └── worker.go
│
├── migrations/                 # MOVED — top-level for clarity
│   ├── 001_create_problems.sql
│   ├── 002_create_submissions.sql
│   ├── 003_add_lc_fields.sql
│   └── seed.sql
│
├── go.mod
├── go.sum
├── Dockerfile
├── Dockerfile.dev
├── .air.toml
└── sqlc.yaml
```

## Layer Rules

```
Handler → Service → Repository → Database
   ↓         ↓          ↓
 (HTTP)  (Business)  (Data Access)
```

1. **Handlers** only do: parse request, validate input format, call service, write response
2. **Services** contain business logic: orchestration, validation rules, calling multiple repos
3. **Repositories** only do: SQL queries, return domain types
4. **Domain** types have zero dependencies — they're just structs and interfaces

## Why This Structure?

| Benefit | How |
|---------|-----|
| **Easy to test** | Mock the repository interface, test service logic in isolation |
| **Easy to scale** | Split into microservices later by extracting a domain + service + repo |
| **Easy to onboard** | New dev? "Handlers are HTTP, services are logic, repos are DB" |
| **Easy to swap DB** | Implement a new repo (e.g. MongoDB) without touching services |
| **No circular deps** | Domain has no imports, each layer only imports the one below it |

## Migration Path (from current to proposed)

1. Create `internal/config/config.go` — move all `os.Getenv` calls here
2. Create `internal/domain/` — move types from `models/` (rename package)
3. Create `internal/repository/interfaces.go` — define interfaces
4. Create `internal/repository/postgres/` — extract DB queries from handlers
5. Create `internal/service/` — extract business logic from handlers
6. Slim down handlers to just HTTP parsing + service calls
7. Move `internal/db/migrations/` to top-level `migrations/`
