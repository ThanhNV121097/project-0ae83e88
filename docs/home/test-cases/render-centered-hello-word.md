# Test Cases — Render centered hello word

Risk level: low. Single public read path, no writes, no roles, no error state in approved design.

## HOME-001 — Display stored greeting

### Scenario: AC-1 stored greeting appears on page
**Given** stored greeting row contains `Hello Word`
**When** guest loads home page
**Then** page shows exactly `Hello Word`
**Check**: render_url

### Scenario: AC-2 greeting is centered on white page with black text
**Given** home page is loaded
**When** guest views screen
**Then** visible text is centered horizontally and vertically on white background with black text
**Check**: measure_styles

### Scenario: AC-3 page uses stored value, not hardcoded frontend copy
**Given** stored greeting row changes to `Hello Mars`
**When** guest reloads home page
**Then** page shows `Hello Mars`
**Check**: render_url

## Notes
- No failure, boundary, or permission cases. Approved design and SRS mark unsupported states as out of scope.
- No manual cases. All criteria are observable through render or style measurement.
