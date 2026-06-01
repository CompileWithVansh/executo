# ─────────────────────────────────────────────
#  Executo Makefile
# ─────────────────────────────────────────────

.PHONY: help dev build stop clean migrate seed logs test lint

# Default target
help: ## Show this help message
	@echo "Executo — Available Commands"
	@echo "────────────────────────────────────────"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ── Docker Compose ────────────────────────────

dev: ## Start full stack in development mode (with hot reload)
	@echo "Starting Executo in development mode..."
	@cp -n .env.example .env 2>/dev/null || true
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
	@echo ""
	@echo "Services:"
	@echo "  Frontend:   http://localhost:3000"
	@echo "  Backend:    http://localhost:8080"
	@echo "  Nginx:      http://localhost:80"
	@echo "  Grafana:    http://localhost:3001"
	@echo "  Prometheus: http://localhost:9090"
	@echo "  Judge0:     http://localhost:2358"

prod: ## Start full stack in production mode
	@echo "Starting Executo in production mode..."
	docker compose up -d --build
	@echo "Production stack started."

stop: ## Stop all services
	docker compose down

restart: ## Restart all services
	docker compose restart

logs: ## Tail logs from all services
	docker compose logs -f

logs-backend: ## Tail backend logs only
	docker compose logs -f backend

logs-frontend: ## Tail frontend logs only
	docker compose logs -f frontend

ps: ## Show running containers
	docker compose ps

# ── Database ──────────────────────────────────

migrate: ## Run database migrations
	@echo "Running migrations..."
	docker compose exec postgres psql \
		-U $${POSTGRES_USER:-executo} \
		-d $${POSTGRES_DB:-executo_db} \
		-f /docker-entrypoint-initdb.d/001_create_problems.sql \
		-f /docker-entrypoint-initdb.d/002_create_submissions.sql
	@echo "Migrations complete."

seed: ## Seed database with sample problems
	@echo "Seeding database..."
	docker compose exec -T postgres psql \
		-U $${POSTGRES_USER:-executo} \
		-d $${POSTGRES_DB:-executo_db} \
		< backend/internal/db/migrations/seed.sql
	@echo "Seed complete."

db-shell: ## Open PostgreSQL shell
	docker compose exec postgres psql -U $${POSTGRES_USER:-executo} -d $${POSTGRES_DB:-executo_db}

db-reset: ## Drop and recreate the database (DESTRUCTIVE)
	@echo "WARNING: This will delete all data. Press Ctrl+C to cancel, Enter to continue."
	@read confirm
	docker compose exec postgres psql -U $${POSTGRES_USER:-executo} -c "DROP DATABASE IF EXISTS $${POSTGRES_DB:-executo_db};"
	docker compose exec postgres psql -U $${POSTGRES_USER:-executo} -c "CREATE DATABASE $${POSTGRES_DB:-executo_db};"
	$(MAKE) migrate
	$(MAKE) seed

# ── Build ─────────────────────────────────────

build: ## Build all Docker images
	docker compose build

build-backend: ## Build backend Docker image only
	docker compose build backend

build-frontend: ## Build frontend Docker image only
	docker compose build frontend

build-go: ## Build Go binary locally (requires Go installed)
	@echo "Building Go backend..."
	cd backend && go build -o bin/executo ./cmd/server
	@echo "Binary: backend/bin/executo"

build-next: ## Build Next.js app locally (requires Node installed)
	@echo "Building Next.js frontend..."
	cd frontend && npm run build
	@echo "Build complete."

# ── Testing ───────────────────────────────────

test: ## Run all tests
	$(MAKE) test-backend
	$(MAKE) test-frontend

test-backend: ## Run Go tests
	cd backend && go test ./... -v -race -timeout 30s

test-frontend: ## Run Next.js tests
	cd frontend && npm test -- --passWithNoTests

# ── Linting ───────────────────────────────────

lint: ## Lint all code
	$(MAKE) lint-backend
	$(MAKE) lint-frontend

lint-backend: ## Lint Go code
	cd backend && golangci-lint run ./...

lint-frontend: ## Lint TypeScript/React code
	cd frontend && npm run lint

fmt: ## Format all code
	cd backend && gofmt -w .
	cd frontend && npm run format

# ── Utilities ─────────────────────────────────

clean: ## Remove all containers, volumes, and build artifacts
	@echo "WARNING: This removes all data volumes. Press Ctrl+C to cancel."
	@read confirm
	docker compose down -v --remove-orphans
	rm -rf backend/bin
	rm -rf frontend/.next
	rm -rf frontend/node_modules

status: ## Quick status check (containers + ports)
	@echo "── Container Status ──"
	@docker compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}"
	@echo ""
	@echo "── Health Checks ──"
	@curl -sf http://localhost:8080/health > /dev/null 2>&1 && echo "  ✓ Backend (8080)" || echo "  ✗ Backend (8080)"
	@curl -sf http://localhost:3000 > /dev/null 2>&1 && echo "  ✓ Frontend (3000)" || echo "  ✗ Frontend (3000)"
	@curl -sf http://localhost:2358/about > /dev/null 2>&1 && echo "  ✓ Judge0 (2358)" || echo "  ✗ Judge0 (2358)"

health: ## Check health of all services
	@echo "Checking service health..."
	@curl -sf http://localhost:8080/health && echo "✓ Backend" || echo "✗ Backend"
	@curl -sf http://localhost:3000 && echo "✓ Frontend" || echo "✗ Frontend"
	@curl -sf http://localhost:2358/about && echo "✓ Judge0" || echo "✗ Judge0"
	@docker compose exec redis redis-cli ping | grep -q PONG && echo "✓ Redis" || echo "✗ Redis"

install-tools: ## Install development tools (golangci-lint, sqlc, air)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/air-verse/air@latest
	@echo "Tools installed."

sqlc-generate: ## Regenerate sqlc Go code from SQL queries
	cd backend && sqlc generate
