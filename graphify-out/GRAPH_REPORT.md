# Graph Report - executo  (2026-05-15)

## Corpus Check
- 45 files · ~27,641 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 382 nodes · 459 edges · 31 communities (24 shown, 7 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 33 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Community 0|Community 0]]
- [[_COMMUNITY_Community 1|Community 1]]
- [[_COMMUNITY_Community 2|Community 2]]
- [[_COMMUNITY_Community 3|Community 3]]
- [[_COMMUNITY_Community 4|Community 4]]
- [[_COMMUNITY_Community 5|Community 5]]
- [[_COMMUNITY_Community 6|Community 6]]
- [[_COMMUNITY_Community 7|Community 7]]
- [[_COMMUNITY_Community 8|Community 8]]
- [[_COMMUNITY_Community 9|Community 9]]
- [[_COMMUNITY_Community 10|Community 10]]
- [[_COMMUNITY_Community 11|Community 11]]
- [[_COMMUNITY_Community 12|Community 12]]
- [[_COMMUNITY_Community 13|Community 13]]
- [[_COMMUNITY_Community 14|Community 14]]
- [[_COMMUNITY_Community 15|Community 15]]
- [[_COMMUNITY_Community 16|Community 16]]
- [[_COMMUNITY_Community 17|Community 17]]
- [[_COMMUNITY_Community 18|Community 18]]
- [[_COMMUNITY_Community 19|Community 19]]
- [[_COMMUNITY_Community 20|Community 20]]
- [[_COMMUNITY_Community 21|Community 21]]
- [[_COMMUNITY_Community 22|Community 22]]
- [[_COMMUNITY_Community 23|Community 23]]
- [[_COMMUNITY_Community 24|Community 24]]
- [[_COMMUNITY_Community 25|Community 25]]
- [[_COMMUNITY_Community 26|Community 26]]
- [[_COMMUNITY_Community 27|Community 27]]
- [[_COMMUNITY_Community 28|Community 28]]

## God Nodes (most connected - your core abstractions)
1. `dependencies` - 17 edges
2. `compilerOptions` - 15 edges
3. `devDependencies` - 14 edges
4. `Executo — LeetCode-Style Code Execution Platform` - 11 edges
5. `writeJSON()` - 10 edges
6. `writeError()` - 10 edges
7. `Worker` - 10 edges
8. `NewRouter()` - 9 edges
9. `Judge0 CE — Self-Hosted Setup` - 9 edges
10. `Executo` - 8 edges

## Surprising Connections (you probably didn't know these)
- `NewRouter()` --calls--> `NewSubmissionsHandler()`  [INFERRED]
  backend/internal/api/router.go → backend/internal/api/handlers/submissions.go
- `main()` --calls--> `NewRouter()`  [INFERRED]
  backend/cmd/server/main.go → backend/internal/api/router.go
- `NewRouter()` --calls--> `NewProblemsHandler()`  [INFERRED]
  backend/internal/api/router.go → backend/internal/api/handlers/problems.go
- `NewRouter()` --calls--> `NewTestCasesHandler()`  [INFERRED]
  backend/internal/api/router.go → backend/internal/api/handlers/testcases.go
- `NewRouter()` --calls--> `Metrics()`  [INFERRED]
  backend/internal/api/router.go → backend/internal/api/middleware/metrics.go

## Communities (31 total, 7 thin omitted)

### Community 0 - "Community 0"
Cohesion: 0.07
Nodes (37): AcceptanceBar(), DifficultyBadge(), ProblemListProps, OutputBlock(), SubmissionResult(), SubmissionResultProps, clsx, TabId (+29 more)

### Community 1 - "Community 1"
Cohesion: 0.07
Nodes (25): 1. Via SQL (recommended for bulk), 2. Via API (POST /api/problems), Adding New Problems, Architecture Overview, Backend → PostgreSQL, Backend → Redis, code:block1 (┌───────────────────────────────────────────────────────────), code:bash (curl -X POST http://localhost:8080/api/problems \) (+17 more)

### Community 2 - "Community 2"
Cohesion: 0.08
Nodes (24): devDependencies, autoprefixer, eslint, eslint-config-next, jest, postcss, prettier, tailwindcss (+16 more)

### Community 3 - "Community 3"
Cohesion: 0.08
Nodes (24): annotations, list, description, editable, fiscalYearStartMonth, graphTooltip, id, __inputs (+16 more)

### Community 4 - "Community 4"
Cohesion: 0.13
Nodes (7): Client, ExecuteSubmissionPayload, executionResult, NewExecuteSubmissionTask(), ParseExecuteSubmissionPayload(), Worker, nullString()

### Community 5 - "Community 5"
Cohesion: 0.1
Nodes (19): code:block1 (Your Code → Judge0 API → isolate (sandbox) → Docker containe), code:bash (make dev), code:bash (# Check the API), code:block4 (JUDGE0_URL=https://judge0-ce.p.rapidapi.com), code:bash (# Check logs), code:bash (# Check workers are running), Configuration, How Judge0 Works (+11 more)

### Community 6 - "Community 6"
Cohesion: 0.17
Nodes (14): NewRouter(), NewProblemsHandler(), NewTestCasesHandler(), CORS(), getAllowedOrigins(), isAllowedOrigin(), ipLimiter, extractIP() (+6 more)

### Community 7 - "Community 7"
Cohesion: 0.11
Nodes (18): compilerOptions, allowJs, esModuleInterop, incremental, isolatedModules, jsx, lib, module (+10 more)

### Community 8 - "Community 8"
Cohesion: 0.11
Nodes (12): CreateProblemRequest, Difficulty, Example, FunctionSignatures, JSONArray, JSONMap, PatchTestCaseRequest, Problem (+4 more)

### Community 9 - "Community 9"
Cohesion: 0.24
Nodes (8): extractIDFromPath(), parseIntParam(), writeError(), writeJSON(), ProblemsHandler, NewSubmissionsHandler(), parseInt64(), SubmissionsHandler

### Community 10 - "Community 10"
Cohesion: 0.12
Nodes (16): dependencies, axios, class-variance-authority, lucide-react, monaco-editor, @monaco-editor/react, next, @radix-ui/react-dialog (+8 more)

### Community 11 - "Community 11"
Cohesion: 0.23
Nodes (6): decodeBase64(), Judge0Client, Judge0Result, Judge0Status, SubmitRequest, SubmitResponse

### Community 12 - "Community 12"
Cohesion: 0.17
Nodes (11): API Endpoints, Architecture, code:block1 (Browser → Nginx → Go Backend → Redis (Asynq) → Judge0 (Docke), code:block2 (executo/), code:bash (cp .env.example .env), Executo, Philosophy, Project Structure (+3 more)

### Community 13 - "Community 13"
Cohesion: 0.18
Nodes (11): 1. Clone and Configure Environment, 2. Start the Full Stack, 3. Run Database Migrations, 4. Seed Sample Problems, 5. Verify Services, code:bash (git clone <your-repo-url> executo), code:bash (# Start everything (first run pulls images, takes 3-5 minute), code:bash (make migrate) (+3 more)

### Community 14 - "Community 14"
Cohesion: 0.29
Nodes (9): DB, getEnv(), MustNew(), New(), WithRetry(), NewJudge0Client(), NewClient(), NewWorker() (+1 more)

### Community 15 - "Community 15"
Cohesion: 0.22
Nodes (9): Backend can't connect to PostgreSQL, code:bash (# Check Judge0 is running), code:bash (# Check postgres is healthy), code:bash (# Check Nginx is routing correctly), code:bash (# Check Redis is running), Frontend shows "Failed to fetch", Judge0 not accepting submissions, Submissions stuck in "pending" (+1 more)

### Community 16 - "Community 16"
Cohesion: 0.5
Nodes (3): jsonError(), jsonOK(), TestCasesHandler

### Community 17 - "Community 17"
Cohesion: 0.36
Nodes (5): isNumeric(), Metrics(), newResponseWriter(), normalizePath(), responseWriter

### Community 18 - "Community 18"
Cohesion: 0.25
Nodes (5): CreateSubmissionRequest, CreateSubmissionResponse, Language, Submission, SubmissionStatus

### Community 19 - "Community 19"
Cohesion: 0.29
Nodes (7): decode(), Editor, LANGUAGES, PlaygroundPage(), RunResult, RunStatus, STATUS_LABELS

### Community 22 - "Community 22"
Cohesion: 0.4
Nodes (3): DEFAULT_OPTIONS, EditorProps, EXECUTO_THEME

### Community 23 - "Community 23"
Cohesion: 0.4
Nodes (4): extends, rules, @typescript-eslint/no-explicit-any, @typescript-eslint/no-unused-vars

## Knowledge Gaps
- **173 isolated node(s):** `HealthResponse`, `ipLimiter`, `DB`, `SubmitRequest`, `SubmitResponse` (+168 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **7 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `NewRouter()` connect `Community 6` to `Community 9`, `Community 14`, `Community 17`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Why does `dependencies` connect `Community 10` to `Community 0`, `Community 2`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **Why does `clsx` connect `Community 0` to `Community 10`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._
- **Are the 4 inferred relationships involving `writeJSON()` (e.g. with `.CreateSubmission()` and `.GetSubmission()`) actually correct?**
  _`writeJSON()` has 4 INFERRED edges - model-reasoned connections that need verification._
- **What connects `HealthResponse`, `ipLimiter`, `DB` to the rest of the system?**
  _173 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Community 0` be split into smaller, more focused modules?**
  _Cohesion score 0.07 - nodes in this community are weakly interconnected._
- **Should `Community 1` be split into smaller, more focused modules?**
  _Cohesion score 0.07 - nodes in this community are weakly interconnected._