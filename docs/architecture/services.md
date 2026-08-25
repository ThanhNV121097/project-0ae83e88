# Service Design

## Shared rules

- Backend routes are mounted without `/api`; deploy proxy strips `/api` before requests reach service.
- Versioned public paths start with `/v1`.
- Responses use `application/json` except `/healthz` plain text.
- External errors use one envelope shape.

## Error envelope

```json
{
  "error": {
    "code": "internal_error",
    "message": "Request failed"
  }
}
```

Codes:

| Code | HTTP | Meaning |
|---|---:|---|
| `not_found` | 404 | Required greeting row is missing. |
| `internal_error` | 500 | Database or server failure. |

## Endpoints

### `GET /healthz`

Readiness endpoint for Docker and runtime.

Request: none.

Success `200 text/plain`:

```text
ok
```

Failure: non-200 or connection failure. Health only passes after migrations succeed and `SELECT 1` against DB works.

### `GET /v1/greeting`

Returns stored home greeting.

Request: none.

Success `200 application/json`:

```json
{
  "text": "Hello Word"
}
```

Errors:

| Condition | HTTP | Body |
|---|---:|---|
| Greeting row `id = 1` missing | 404 | Error envelope with `not_found`. |
| Query fails | 500 | Error envelope with `internal_error`. |

## Decisions

| Decision | Reason | Rejected alternative |
|---|---|---|
| One read endpoint | SRS has one read-only function. | CRUD endpoints rejected as out of scope. |
| JSON object `{ "text": string }` | Stable named field for frontend. | Bare string rejected because harder to extend and less self-describing. |
| Generic external error message | Avoid leaking DB details. | Raw SQL errors rejected for security. |
