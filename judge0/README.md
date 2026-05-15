# Judge0 CE — Self-Hosted Setup

Judge0 CE is the open-source code execution engine that powers Executo. It runs submitted code in isolated Docker containers with strict resource limits.

## How Judge0 Works

```
Your Code → Judge0 API → isolate (sandbox) → Docker container → Result
```

1. You POST source code + language ID to Judge0's REST API
2. Judge0 queues the job in Redis
3. A Judge0 worker picks up the job
4. The worker uses `isolate` (a Linux sandbox) to run the code
5. Resource limits (CPU, memory, time) are enforced
6. stdout, stderr, and metrics are captured
7. The result is stored and returned via GET

## Quick Start (via Docker Compose)

Judge0 is included in the main `docker-compose.yml`. Just run:

```bash
make dev
```

This starts:
- `judge0-server` — the REST API on port 2358
- `judge0-workers` — the execution workers

## Verify Judge0 is Running

```bash
# Check the API
curl http://localhost:2358/about

# List supported languages
curl http://localhost:2358/languages | jq '.[].name'

# Test a submission manually
curl -X POST http://localhost:2358/submissions \
  -H "Content-Type: application/json" \
  -d '{
    "source_code": "print(\"Hello, World!\")",
    "language_id": 71,
    "stdin": ""
  }'
# Returns: {"token": "abc123..."}

# Get the result
curl http://localhost:2358/submissions/abc123
```

## Language IDs Used by Executo

| Language   | Judge0 ID |
|------------|-----------|
| Python 3   | 71        |
| Java       | 62        |
| C++        | 54        |
| JavaScript | 63        |

Full list: `curl http://localhost:2358/languages`

## Configuration

Edit `judge0/judge0.conf` to adjust:

- `WORKERS` — number of concurrent execution workers (default: 2)
- `CPU_TIME_LIMIT` — max CPU seconds per submission (default: 5)
- `MEMORY_LIMIT` — max memory in KB (default: 262144 = 256MB)
- `WALL_TIME_LIMIT` — max wall clock time (default: 10s)

## Using RapidAPI Judge0 (Alternative)

If you don't want to self-host Judge0, you can use the hosted version:

1. Sign up at https://rapidapi.com/judge0-official/api/judge0-ce
2. Get your API key
3. Set in `.env`:
   ```
   JUDGE0_URL=https://judge0-ce.p.rapidapi.com
   JUDGE0_API_KEY=your_rapidapi_key_here
   ```
4. Remove the `judge0-server` and `judge0-workers` services from docker-compose

## Troubleshooting

### Judge0 won't start
```bash
# Check logs
docker compose logs judge0-server

# Common issue: needs privileged mode for isolate
# Ensure docker-compose.yml has: privileged: true
```

### Submissions stuck in queue
```bash
# Check workers are running
docker compose logs judge0-workers

# Check Redis connection
docker compose exec redis redis-cli ping
```

### "isolate: error" in logs
Judge0 requires Linux kernel features (cgroups, namespaces). It won't work on:
- macOS (use Docker Desktop with Linux VM — should work)
- Windows WSL1 (use WSL2 instead)
- Some cloud VMs with restricted kernels

### Memory issues
If Judge0 workers crash with OOM errors, reduce `MEMORY_LIMIT` in `judge0.conf`
or increase Docker's memory allocation.

## Security Notes

- Judge0 uses `isolate` for sandboxing, which provides strong isolation
- Each submission runs in a separate process with resource limits
- Network access is disabled inside the sandbox
- File system access is restricted to a temporary directory
- For production, enable `AUTHN_TOKEN` in judge0.conf to require authentication
