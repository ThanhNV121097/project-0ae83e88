# SRS — home

Module: `home`
Last updated: 2025-02-14
Design: [View the approved design](http://localhost:8080/design/0ae83e88-25ed-4410-846e-9a028369815e)
Design system: `design/design-system.md`

## 1. Purpose

`home` delivers the only screen in `hello-word-20`: a plain centered "Hello Word" page whose text comes from the backend and PostgreSQL, not from frontend source. If this module does not exist, the project no longer proves the end-to-end pipeline: UI, API, and DB row all become disconnected or hardcoded.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Any visitor of the public page | View the centered text |
| System | Next.js page, backend API, and PostgreSQL together | Read stored text and render it on the page |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Render centered hello word

**Out of scope**

- Editing the displayed text — deliberately not built; the product only proves read path end to end.
- Any extra UI, animation, palette, or navigation — omitted by design.
- Authentication, user accounts, and permissions — belong to other modules only if added later.

## 4. Functional requirements

### 4.1 Render centered hello word

**Requirement HOME-001 — Display stored greeting**

*As a* Guest, *I want to* see the greeting from stored data, *so that* the page proves frontend, backend, and PostgreSQL are connected.

Behaviour:

1. The guest opens the home page.
2. The system reads one greeting value through the backend API.
3. The system renders that value as a single line of black text centered horizontally and vertically on a white page.
4. The frontend does not contain the greeting text as a hardcoded display value.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/home/test-cases/render-centered-hello-word.md`. Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | the stored greeting row contains `Hello Word` | the guest loads the page | the page shows `Hello Word` |
| AC-2 | the page is loaded | the guest views the screen | text is centered horizontally and vertically on white background with black text |
| AC-3 | the stored greeting row changes to a different value | the guest reloads the page | the page shows the stored value, not a frontend hardcoded copy |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Not applicable | Single public read with no writes and no role differences | Not applicable: approved design shows one default screen only; no loading, empty, or error state is part of the design |

**Data touched** — the fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | One stored value, readable by the home page |

## 5. Screens

The design is the source of truth for appearance; this section maps functions onto it so nothing in the design is unaccounted for and nothing specified here is missing from the design.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Home screen | Single centered Hello Word page | HOME-001 | default |

## 6. Non-functional requirements

Only what is real for this module. Delete rows that do not apply rather than inventing a number nobody will check.

| Area | Requirement |
|---|---|
| Performance | With backend API already responding, the home page renders within 1 second after initial browser request |
| Accessibility | Text contrast is at least 4.5:1 against the white background |
| Responsive | Page fits 320px width and up with no horizontal scroll |
| Localisation | Copy is English only |

## 7. Dependencies and assumptions

- **Depends on:** Next.js frontend, backend API, and PostgreSQL, for reading and rendering the stored greeting.
- **Assumption:** The database contains exactly one greeting row for this page; if that changes, this module needs list or selection behavior.
- **Assumption:** No error or empty state is part of the approved design; if upstream read failure must be shown later, the design and this SRS both need revision.

| Open question | Proposed default | Who decides |
|---|---|---|
| Should the greeting ever be editable in product scope? | No | Stakeholder |
| Should any fallback text appear if the stored row is missing? | No fallback; treat as unsupported until design changes | Stakeholder / TL |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Render centered hello word | HOME-001 | `test-cases/render-centered-hello-word.md` |
