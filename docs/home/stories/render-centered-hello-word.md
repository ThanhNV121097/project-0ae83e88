# Story — Render centered hello word

## User story

As a Guest, I want to see the greeting from stored data, so that the page proves frontend, backend, and PostgreSQL are connected.

## In scope

- Read one greeting value through the backend API.
- Render that value on the home page as a single line of black text centered horizontally and vertically on a white background.
- Keep greeting text out of frontend source as a hardcoded display value.
- Use the approved single-screen design only.

## Out of scope

- Editing the stored text.
- Any extra UI, navigation, animation, palette, loading state, empty state, or error state.
- Authentication, user accounts, and permissions.
- More than one greeting row or any selection behavior.

## UI scope

- One screen only: the centered greeting screen from the approved design.
- States: default only.
- No interactive controls, no secondary views, no alternate layouts.

## Acceptance criteria

1. Given the stored greeting row contains `Hello Word`, when the Guest loads the page, then the page shows `Hello Word`.
2. Given the page is loaded, when the Guest views the screen, then text is centered horizontally and vertically on a white background with black text.
3. Given the stored greeting row changes to a different value, when the Guest reloads the page, then the page shows the stored value, not a frontend hardcoded copy.

## Dependencies

- Backend API exists to read the greeting value.
- PostgreSQL contains exactly one greeting row for this page.
- Next.js home page is wired to call the backend API.
- Approved design and design system remain unchanged for this one-screen layout.
