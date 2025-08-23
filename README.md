# Termdo Gateway API

Public-facing API gateway for Termdo. Proxies requests to internal microservices, normalizes responses, and manages browser-token behavior.

This service sits in front of:

- termdo-auth-api: Authentication and JWT issuance
- termdo-tasks-api: Task CRUD domain API
- termdo-web: Frontend clients calling the gateway
- termdo-db: PostgreSQL backing store (used by internal services)
- termdo-infra: Infra and deployment assets

## Features

- Reverse proxy for Auth (`/auth/*`) and Tasks (`/tasks/*`)
- Browser-aware auth: stores JWT in an HTTP-only cookie and strips token from body
- API client mode: returns JWT in response body `token` field
- Automatic token forwarding: cookie -> `Authorization: Bearer <token>` header
- Hostname aggregation: merges upstream `X-Hostname` into response `hostnames`
- Logout endpoint handled locally, clearing the auth cookie

## Tech Stack

- Language: Go 1.25
- Web: Gin
- Reverse proxy: `net/http/httputil`
- Config: environment variables, `godotenv` in dev
- Container: Multi-stage Docker, static binary, scratch runtime

## Getting Started

### Prerequisites

- Go 1.25+
- `.env` file (see `.env.example`)
- Upstream services reachable on the same Docker network (`termdo-net`) or via hostnames

### Environment Variables

- App:
  - `APP_HOST`: Logical hostname label for this gateway (e.g., `termdo-gateway-api`)
  - `APP_PORT`: Port number to listen on (e.g., `3000`).
  - `COOKIE_IS_SECURE`: `true` or `false` (use `true` for HTTPS)
- Auth API upstream:
  - `AUTH_API_PROTOCOL`: `http` or `https`
  - `AUTH_API_HOST`: Hostname of the auth API (container/service name on `termdo-net`)
  - `AUTH_API_PORT`: Port of the auth API
- Tasks API upstream:
  - `TASKS_API_PROTOCOL`: `http` or `https`
  - `TASKS_API_HOST`: Hostname of the tasks API (container/service name on `termdo-net`)
  - `TASKS_API_PORT`: Port of the tasks API

Create `.env` by copying `.env.example` and fill values accordingly.

### Run (Local)

```bash
go run ./source/main.go --dev
# or build
go build -o bin/termdo-gateway-api ./source/main.go
./bin/termdo-gateway-api --dev
```

In `--dev` mode, `.env` is loaded and Gin runs in debug mode.

### Docker

Builds a static binary and runs from `scratch`.

```bash
docker build -t termdo-gateway-api:local .
docker run --rm --env-file .env -p 3000:3000 termdo-gateway-api:local
```

Docker Compose is provided and exposes `3000`:

```bash
docker compose up --build
```

Ensure upstream services (auth, tasks) are also running and accessible at the configured host/ports (preferably on `termdo-net`).

## Routing & Behavior

### Common

- Header `X-Client-Browser: 1` marks a request as a “browser client”.
  - For browser clients: JWT is set in an HTTP-only cookie and removed from response body.
  - For non-browsers: JWT remains in the `token` field in the JSON body.
- If a valid `token` cookie exists for browser clients, the gateway injects `Authorization: Bearer <token>` for upstream calls.
- The gateway aggregates hostnames into `hostnames` object in response body and removes upstream `X-Hostname` header.

### /auth/*

- Proxies all methods and paths to the Auth API.
- Special case `PUT /auth/logout`: handled locally, clears the auth cookie and returns a success body.
- Login/Signup responses: token is captured and handled as per browser/non-browser rules.

Example (API client mode):

```bash
curl -sS -X POST http://localhost:$APP_PORT/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","password":"P@ssw0rd!"}'
```

Example (Browser mode, cookie set; note token not returned in body):

```bash
curl -i -sS -X POST http://localhost:$APP_PORT/auth/login \
  -H 'X-Client-Browser: 1' \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","password":"P@ssw0rd!"}'
```

Logout:

```bash
curl -sS -X PUT http://localhost:$APP_PORT/auth/logout -H 'X-Client-Browser: 1'
```

### /tasks/*

- Requires `Authorization: Bearer <token>` (or browser cookie) on every request.
- Gateway first calls Auth API `/refresh` with the provided token:
  - Ensures validity and obtains `accountId` (and a refreshed token).
  - Rewrites upstream path to prepend `/{accountId}` for the Tasks API.
  - Merges token/hostnames into the final response per the browser/non-browser rules.

Examples (API client mode):

```bash
# List tasks
curl -sS http://localhost:$APP_PORT/tasks \
  -H "Authorization: Bearer $TOKEN"

# Create task
curl -sS -X POST http://localhost:$APP_PORT/tasks \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title":"Study Go","description":"Finish Gin guide","isCompleted":false}'

# Update task
curl -sS -X PUT http://localhost:$APP_PORT/tasks/456 \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title":"Study Go","description":"Proxy logic","isCompleted":true}'

# Delete task
curl -sS -X DELETE http://localhost:$APP_PORT/tasks/456 \
  -H "Authorization: Bearer $TOKEN"
```

Examples (Browser mode):

```bash
# If you have a token cookie from /auth/login with X-Client-Browser: 1
curl -sS http://localhost:$APP_PORT/tasks -H 'X-Client-Browser: 1'
```

## Cookies & Headers

- Cookie name: `token`, path: `/api`, SameSite: Strict, `Secure` controlled by `COOKIE_IS_SECURE`.
- Header `X-Client-Browser: 1` opts into cookie-based flows and token stripping from body.
- Upstream `X-Hostname` is folded into the JSON body under `hostnames` and removed from headers.

## Integration Notes

- Run gateway and upstream services on the same Docker network `termdo-net`.
- Set `AUTH_API_HOST` and `TASKS_API_HOST` to the resolvable container/service names of those services.
- For local testing without Docker, point hosts to `localhost` and expose upstream ports.

## License

MIT — see `LICENSE.md`.
