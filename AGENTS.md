# Repository Guidelines

## Project Structure & Module Organization
- `website/`: Published static site (HTML, CSS, JS). Main entry `website/index.html`, article listing `website/articles.html`, articles in `website/articles/`, pages in `website/pages/`, shared assets `style.css` and `theme-toggle.js`.
- `structure/`: Canonical markdown sources with YAML frontmatter; mirrors the site (home, article list, articles, pages). Edit here first, then render to `website/`.
- `template/`: Markdown templates (`*.mdt`) for generating pages (`page.mdt`) and article listings (`articles.mdt`).
- `@todo/`: Task notes and open issues; not served publicly.

## Build, Test, and Development Commands
- Render or copy sources from `structure/` into `website/` using your preferred converter (e.g., pandoc or a future parser). Example workflow:
  - `cp structure/index.md website/index.html`
  - `cp structure/articles.md website/articles.html`
  - `cp structure/pages/about.md website/pages/about.html` (same for contact)
- `open website/index.html` (or run a simple static server) to preview; no build step is required.

## Coding Style & Naming Conventions
- Match existing formatting: 4-space indentation in HTML/CSS/JS.
- Keep semantic HTML structure (header/main/article/footer) and class names to align with `style.css`.
- File names use kebab-case (`digital-minimalism.html`).
- JS follows an IIFE pattern with uppercase constants (`theme-toggle.js`); avoid global leakage.

## Testing Guidelines
- No automated tests; manual checks only.
- After content or style changes, manually open `website/index.html`, `website/articles.html`, and any edited article/page. Verify theme toggle works across pages and persists via `localStorage`.
- If JS changes affect theme, test both light and dark modes and reload to confirm shared state.

## Commit & Pull Request Guidelines
- **CRITICAL**: Do NOT commit changes unless explicitly requested by the user.
- Follow the repo history style: short, imperative commits (e.g., “Add markdown source structure and templates”).
- Keep commits focused (content vs. styling vs. JS changes).
- For PRs: include a brief summary, list affected pages, note manual test steps (browsers/viewports), and add screenshots for UI-visible changes.

## Security & Configuration Tips
- No build tooling or package manager is present; avoid adding dependencies unless necessary.
- Keep external fonts/CDN links as-is; document any new third-party assets. Prefer local assets when feasible.
