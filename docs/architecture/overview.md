# Architecture Overview

## Scope

`hello-word-20` is fullstack: Next.js frontend, Go backend, PostgreSQL database. Product scope is one public page that reads one stored greeting through backend API and centers it on white screen.

## Stack

| Part | Choice | Reason | Rejected alternative |
|---|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches repository scaffold and CI; server component page can compose story components. | Static HTML rejected because SRS requires backend/API/DB read. |
| Backend | Go 1.22 HTTP service | Small stdlib server, easy container build. | Node backend rejected to keep default stack and avoid second runtime pattern. |
| Database | PostgreSQL 16 | Required source of greeting row. | Frontend hardcoded copy rejected by HOME-001. |
| Styling | CSS tokens in `app/globals.css` plus Tailwind available | CI checks token usage; design is tiny. | Component-level raw values rejected because hardcoded values fail token gate. |

## Repository layout

```text
code/backend/
  cmd/api/main.go
  internal/migrations/migrations.go
  migrations/*.sql
  .env.example
code/frontend/
  app/layout.tsx
  app/page.tsx
  app/globals.css
  components/
  lib/mock/
docs/architecture/
  overview.md
  erd.md
  services.md
```

`app/page.tsx` stays thin composition root. Stories add imports and elements there. Story components use `export default function ComponentName()`.

## Data flow

1. Browser loads frontend.
2. Frontend story component calls backend path from `NEXT_PUBLIC_API_URL`.
3. Backend reads PostgreSQL through `DATABASE_URL`.
4. Backend returns JSON. Frontend renders value; no greeting literal in frontend display code.

## Backend contracts

Backend reads `DATABASE_URL`, then applies migrations from `code/backend/migrations/` on boot, then listens on `PORT`, falling back to `APP_PORT`, then `8080`. `/healthz` returns 200 only after migrations succeed and `SELECT 1` works.

## Environment variables

### Root compose

- `POSTGRES_USER` — local PostgreSQL user.
- `POSTGRES_PASSWORD` — local PostgreSQL password.
- `POSTGRES_DB` — local PostgreSQL database.
- `BACKEND_PORT` — host port for backend.
- `FRONTEND_PORT` — host port for frontend.
- `NEXT_PUBLIC_API_URL` — browser-visible backend origin.

### Backend

- `DATABASE_URL` — PostgreSQL connection URL, injected by runtime/compose.
- `PORT` — HTTP port, injected by runtime/compose.
- `APP_PORT` — optional legacy port fallback.

### Frontend

- `NEXT_PUBLIC_API_URL` — public backend origin, default `http://localhost:8080`.

## Naming conventions

- API paths use `/v1/...`; no `/api` prefix because deploy proxy strips `/api` before backend.
- Database identifiers use snake_case.
- Go packages use short lowercase names.
- React components use PascalCase files and `export default function` declarations.
- CSS custom properties use semantic token names from `design/design-system.md`.

## Run and checks

Local full stack:

```sh
cp .env.example .env
docker compose --profile local up --build
```

Backend checks from `code/backend`:

```sh
go build ./...
go vet ./...
go test ./...
```

Frontend checks from `code/frontend`:

```sh
npm ci
npm run lint
npm run build
npm test --if-present
```

CI gate is `.github/workflows/ci.yml`; container workflows are precommitted and not edited here.

## Risks and limits

- Missing greeting row has no product-approved UI state. Backend will return error envelope until scope changes.
- No auth, writes, caching, or admin edit route. Add only if SRS changes.
- Migrations are intentionally simple SQL files tracked by filename; enough for one small schema.
