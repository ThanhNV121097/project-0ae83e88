# ERD

## Tables

### greetings

Stores one public greeting for home page.

| Column | Type | Constraints | Notes |
|---|---|---|---|
| `id` | `smallint` | primary key, `id = 1` | Single-row table. |
| `text` | `text` | not null, not empty | Displayed greeting value. |
| `created_at` | `timestamptz` | not null, default `now()` | Seed time. |
| `updated_at` | `timestamptz` | not null, default `now()` | Future-proof read audit only. |

Seed row:

```sql
insert into greetings (id, text) values (1, 'Hello Word')
on conflict (id) do nothing;
```

### schema_migrations

Tracks backend-applied SQL migrations.

| Column | Type | Constraints | Notes |
|---|---|---|---|
| `filename` | `text` | primary key | Migration filename applied. |
| `applied_at` | `timestamptz` | not null, default `now()` | Application time. |

## Relationships

No foreign keys. `greetings` is standalone because product has one public read and no users.

## Decisions

| Decision | Reason | Rejected alternative |
|---|---|---|
| Single `greetings` row keyed by `id = 1` | SRS says exactly one stored value. | Key-value table rejected because only one value exists and adds unused flexibility. |
| Seed in migration | Empty runtime DB must become usable on boot. | Manual seed rejected because deployment creates empty DB. |
| No update endpoint schema additions | Editing is out of scope. | Audit/user columns rejected because no auth or writes exist. |
