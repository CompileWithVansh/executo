# Executo

Online code runner and practice platform. Write code from scratch, run it instantly, and practice interview problems without pre-filled templates.

## What it does

- **Playground** — Write and execute code in Python, Java, C++, JavaScript
- **150 Problems** — NeetCode 150 list with links to LeetCode, no starter code
- **Community Test Cases** — Users suggest test cases, admin approves them
- **No hand-holding** — No function signatures, no auto-complete hints. You write everything from `main()` to `stdin` parsing, like a real interview

## Architecture

```
Browser → Nginx → Go Backend → Redis (Asynq) → Judge0 (Docker)
                      ↓
                  PostgreSQL
```

| Service | Port | Purpose |
|---------|------|---------|
| Frontend (Next.js) | 3000 | UI + Monaco editor |
| Backend (Go) | 8080 | REST API + job queue |
| PostgreSQL | 5432 | Problems, submissions, test cases |
| Redis | 6379 | Asynq job queue |
| Judge0 CE | 2358 | Code execution in Docker containers |
| Nginx | 80 | Reverse proxy |

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Frontend | Next.js 14, React, TypeScript, Tailwind, Monaco Editor |
| Backend | Go, net/http, Asynq, lib/pq |
| Database | PostgreSQL 15 |
| Queue | Redis + Asynq |
| Execution | Judge0 CE (Docker-based sandboxed runner) |
| Infra | Docker Compose, Nginx |

## Project Structure

```
executo/
├── frontend/          # Next.js app (playground, problems, admin)
├── backend/           # Go API server + Asynq worker
│   ├── cmd/server/    # Entry point
│   ├── internal/
│   │   ├── api/       # Routes + handlers
│   │   ├── db/        # Postgres connection + migrations
│   │   ├── executor/  # Judge0 client
│   │   ├── models/    # Domain types
│   │   └── queue/     # Asynq tasks + worker
│   └── go.mod
├── nginx/             # Reverse proxy config
├── judge0/            # Judge0 config
├── monitoring/        # Prometheus + Grafana
├── docker-compose.yml # Production
└── docker-compose.dev.yml
```

## Quick Start

```bash
cp .env.example .env
docker compose up -d
```

Frontend: http://localhost:3000
API: http://localhost:8080/health

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /problems | List all 150 problems |
| GET | /problems/slug/:slug | Get problem by slug |
| POST | /api/run | Execute code (playground) |
| GET | /api/run/:token | Poll execution result |
| POST | /problems/:slug/suggest-testcase | Suggest a test case |
| GET | /problems/:slug/testcases | Get approved test cases |
| GET | /admin/testcases | List pending suggestions |
| PATCH | /admin/testcases/:id | Approve/reject suggestion |

## Philosophy

LeetCode gives you pre-filled function signatures. In a real interview, you get a blank editor. Executo trains you the way you'll be tested — no starter code, no hints, full program from scratch.

---

Built by **Vansh Gupta**
