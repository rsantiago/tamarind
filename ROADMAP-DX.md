# Tamarind DX Roadmap

This roadmap outlines the path to making `tamarind` the premier single-binary static site generator with an unbeatable developer experience.

## Phase 1: The "Flow" State (Immediate Priorities)
Focus: reducing friction and increasing speed of iteration.

- [ ] **Live Reloading (v4.0)**
    - **Goal**: Instant visual feedback.
    - **Impl**: Use `fsnotify` to watch `source` directory. On change -> Rebuild -> Notify browser via WebSocket/SSE to refresh.
    - **Impact**: Massive DX improvement. No more manual refreshing.

- [ ] **Draft Handling**
    - **Goal**: safe writing workflow.
    - **Impl**: Add `draft: true` to front matter. Build ignores these by default. Add `--drafts` flag to `tamarind build` and `tamarind serve` to view them.

- [ ] **RSS/Atom Feed Generation**
    - **Goal**: Connectivity with the open web.
    - **Impl**: Generate `feed.xml` for `articles` collection using standard Go RSS libraries or templates.

## Phase 2: Power & Flexibility
Focus: giving developers control without complexity.

- [ ] **External Theme Support**
    - **Goal**: Community ecosystem.
    - **Impl**: Check `./themes/` folder first before falling back to embedded binary themes. Allows users to `git clone` themes.

- [ ] **Shortcodes / Custom Components**
    - **Goal**: Rich media support without HTML.
    - **Impl**: Parse custom syntax like `{{< youtube id >}}` or `::note ... ::` and inject HTML snippets.

- [x] **Data Files Support**
    - **Goal**: Hybrid static site capabilities.
    - **Impl**: Automatically scans `source/data/*.{yaml,json}` and exposes them in templates via `{{ .Data.FileName }}` (e.g., `{{ .Data.Authors }}`).

## Phase 3: Performance & Polish
Focus: creating the fastest possible output.

- [x] **Image Optimization Pipeline**
    - **Goal**: 100/100 Lighthouse scores automatically.
    - **Impl**: Automatically resizes images to multiple breakpoints (480w, 800w, 1200w) during the build. (WebP generation skipped for now to avoid CGO).

- [ ] **Client-Side Search**
    - **Goal**: Discoverability for larger docs/blogs.
    - **Impl**: Generate a `search.json` index during build. Add a pre-styled Search UI component to all themes.

- [ ] **CLI Polish**
    - **Goal**: Joyful interaction.
    - **Impl**: Beautiful colored output (charmbracelet/lipgloss), progress bars for large builds, interactive `tamarind init` prompts.

## Phase 4: The "Visual" Frontier (Long Term)
- [ ] **Admin GUI**
    - **Goal**: CMS-like experience.
    - **Impl**: A `/admin` route in `tamarind serve` that provides a visual editor for Markdown files, saving back to disk.
