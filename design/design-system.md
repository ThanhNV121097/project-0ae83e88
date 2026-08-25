# Design System — hello-word-20

> Source of truth: the approved `index.html`.
> Every value below is extracted from it. Changing a value here without
> changing the approved design is a defect.

Last updated: 2025-02-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#ffffff` | Page background |
| `--color-text` | `#000000` | Body text, heading text |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |

### 1.2 Spacing

Base unit: `8px`. Every margin, padding, and gap in product uses one of these.

| Token | Value |
|---|---|
| `--space-6` | `24px` |

### 1.3 Typography

Font families (include fallback stack and how font is loaded):

- Body: `Arial, Helvetica, sans-serif` (system stack)
- Headings: `Arial, Helvetica, sans-serif` (system stack)
- Mono: not used

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-3xl` | `clamp(2.5rem, 8vw, 6rem)` | `1` | `400` | h1 |

Heading levels are used in order and never skipped for visual sizing.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | `0` | No rounding used |
| `--border-width` | `0` | No borders used |
| `--shadow-sm` | `none` | No shadows used |
| `--duration-fast` | `0ms` | No motion used |
| `--duration-base` | `0ms` | No motion used |
| `--easing` | `linear` | No motion used |

Motion respects `prefers-reduced-motion: reduce`: no motion exists to reduce.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | `0px` | `100%` | `1` | `24px` |
| `md` | not used | not used | not used | not used |
| `lg` | not used | not used | not used | not used |
| `xl` | not used | not used | not used | not used |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |

## 2. Components

One subsection per reusable component. Every component lists all states.

### 2.1 Centered greeting screen

**Purpose** — Single full-page screen that centers one line of text. Use for this landing page only; not for interactive flows.

**Anatomy** — `[main] [h1]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--text-3xl`, `--space-6` | Only screen in product |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `100vh` | `24px` | `--text-3xl` |

**States** — every row must be filled in.

| State | Visual change | Tokens |
|---|---|---|
| Default | White full-page background, centered black heading | `--color-bg`, `--color-text`, `--text-3xl`, `--space-6` |
| Hover | No change | none |
| Focus (keyboard) | No interactive focus target | none |
| Active / pressed | No change | none |
| Disabled | No disabled state | none |
| Loading | No loading state | none |
| Error | No error state | none |
| Empty | No empty state; screen always shows greeting text | none |

**Accessibility** — semantic `main` landmark, `aria-label="Hello Word screen"`, centered `h1`, no interactive controls, no hit target requirement.

## 3. Content and formatting

- Voice and tone: plain, neutral, no decoration.
- Date, time, number, and currency formats: not used.
- Capitalization rule for buttons, headings, and labels: heading uses title case as written in content, no buttons or labels used.
- Empty-state and error-message wording pattern: not used.

## 4. Known deviations

Places where approved design does not follow its own rules or anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| `main` container | Fixed full-viewport centering screen with no responsive breakpoints beyond one layout | Stakeholder approved one-screen proof page | None |
| Components | No interactive states exist | Page has no interactive UI | None |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-02-14 | Initial design system for `hello-word-20` | pending |
