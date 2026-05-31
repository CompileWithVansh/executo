# Contributing to Executo

## Dev Setup

1. Clone the repo
2. Copy env files: `cp .env.example .env`
3. Start services: `make dev`
4. Frontend runs at http://localhost:3000
5. Backend API at http://localhost:8080

## Code Style

- **Go**: run `gofmt` before committing
- **TypeScript**: follow existing ESLint config
- **Commits**: keep messages short and descriptive

## Branch Naming

- `feat/feature-name` for new features
- `fix/bug-description` for bug fixes
- `docs/what-changed` for documentation updates

## Pull Requests

- Keep PRs small and focused
- Test locally with `make dev` before pushing
- Update docs if you change API endpoints
