# Theme Refactor Tracker

Compliance status of each theme against [THEME_SPEC.md](file:///home/rsantiago/Documents/atman-multi-agents/tamarind/parser/assets/templates/THEME_SPEC.md).

## Status Legend
- ✅ Implemented
- ❌ Missing
- 🔶 Partial (has some elements but not all required by spec)

## Compliance Matrix

| Criteria                  | gram | basic | bird | blue | midnight | neon | pastel |
|---------------------------|------|-------|------|------|----------|------|--------|
| **CSS Variable Tokens**   |      |       |      |      |          |      |        |
| Colors (primary, card-bg) | ✅   | 🔶    | 🔶   | 🔶   | 🔶       | 🔶   | 🔶     |
| Typography (font vars)    | ✅   | 🔶    | 🔶   | 🔶   | 🔶       | 🔶   | 🔶     |
| Radius tokens (sm/md/lg)  | ✅   | ❌    | ❌   | ❌   | ❌       | ❌   | ❌     |
| Shadow token              | ✅   | ❌    | ❌   | ❌   | ❌       | ❌   | ❌     |
| **Component Contract**    |      |       |      |      |          |      |        |
| Buttons (.btn-*)          | ✅   | ❌    | ❌   | ❌   | ❌       | ❌   | ❌     |
| Forms (.form-*)           | ✅   | ❌    | ❌   | ❌   | ❌       | ❌   | ❌     |
| Badges (.badge)           | ✅   | ❌    | ❌   | ❌   | ❌       | ❌   | ❌     |
| Cards (.card)             | ✅   | 🔶    | 🔶   | 🔶   | 🔶       | 🔶   | 🔶     |
| Callouts (.callout-*)     | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| **Layout Contract**       |      |       |      |      |          |      |        |
| .layout-container         | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| .site-header (sticky)     | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| .nav-links                | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| Mobile responsive (<768)  | ✅   | 🔶    | 🔶   | 🔶   | 🔶       | 🔶   | 🔶     |
| **Shortcode Support**     |      |       |      |      |          |      |        |
| .video-container          | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| pre/code blocks           | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| figure/figcaption         | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| .terminal                 | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| .mermaid                  | ✅   | ✅    | ✅   | ✅   | ✅       | ✅   | ✅     |
| **Dark Mode**             | ✅   | ❌    | ❌   | ❌   | N/A      | ❌   | ❌     |

## Overall Progress

| Theme    | Status         | Estimated Effort |
|----------|----------------|------------------|
| gram     | ✅ Complete     | Done             |
| basic    | ❌ Not started  | ~2 hours         |
| bird     | ❌ Not started  | ~2 hours         |
| blue     | ❌ Not started  | ~2 hours         |
| midnight | ❌ Not started  | ~2 hours         |
| neon     | ❌ Not started  | ~2 hours         |
| pastel   | ❌ Not started  | ~2 hours         |

## Key Gaps (All Non-Gram Themes)

1. **Buttons** — No `.btn`, `.btn-primary`, `.btn-secondary`, `.btn-ghost`, `.btn-sm` classes
2. **Forms** — No `.form-group`, `.form-input`, `.form-textarea`, `.form-select` classes
3. **Badges** — No `.badge` class
4. **Radius Tokens** — No `--radius-sm`, `--radius-md`, `--radius-lg` variables
5. **Shadow Token** — No `--shadow` variable
6. **Dark Mode** — Only gram implements `[data-theme="dark"]`

## Refactoring Approach

For each theme, follow the update guide in THEME_SPEC.md §5:
1. Add missing CSS variable tokens to `:root`
2. Copy the Component Contract CSS (buttons, forms, badges) from `gram/style.css`
3. Adjust values (colors, radius, shadows) to match the theme's identity
4. Add `[data-theme="dark"]` block
5. Verify with `tamarind quickstart -theme <name>`
