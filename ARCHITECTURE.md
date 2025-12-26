# Tamarind Codebase & Architecture Report

This document provides a technical overview of the `tamarind` static site generator, intended for engineers looking to understand, debug, or refactor the core system.

## 1. Architecture Overview

Tamarind follows a **"Big Bang" Static Site Generation** architecture. It is designed to be stateless and simple: every time you run `build`, it tears down the previous state and rebuilds the site from scratch (with some OS-level caching).

### The Pipeline
1.  **Init (Binary Start)**: CLI flags are parsed (theme, watch mode, drafts).
2.  **Asset Extraction**: Determine if running on a fresh system; extract embedded templates/assets to memory or temp dirs.
3.  **Data Loading**: Scan `writer-sandbox/data/*.json` into a global map.
4.  **Content Scan**: Walk the file system to build an in-memory index of all content (`articles`, `docs`, etc.).
5.  **Image Optimization**: Walk `images/` folder and generate resized variants (480w, 800w, 1200w) on the fly.
6.  **Page Generation**:
    *   Render Collection Indices (e.g., `articles.html`).
    *   Render Tag Pages (`tags/foo.html`).
    *   Render Individual Pages (combines `Markdown` + `FrontMatter` + `Templates`).
7.  **Final Polish**: Generate `sitemap.xml`, `robots.txt`, `llms.txt`, and RSS feeds.

---

## 2. Main Data Structures (`internal/models/models.go`)

The system relies on three primary structures to pass data around:

### `ArticleMeta`
*   **Purpose**: Lightweight metadata for lists, archives, and loops.
*   **Why**: We need to sort/filter thousands of articles without loading their full Markdown body into memory.
*   **Fields**: `Title`, `Date`, `Tags`, `URL`, `Hidden`, `Draft`.

### `PageData`
*   **Purpose**: The "God Object" passed to HTML templates.
*   **Why**: Templates need access to *everything*: the current page content, the site menu, the full list of articles (for sidebars), site configuration, and custom data files.
*   **Fields**: `Body` (HTML), `Articles` (List), `Data` (Map), `Menu`, `Paginator`.

### `FrontMatter`
*   **Purpose**: The raw configuration at the top of every Markdown file.
*   **Why**: Maps directly to YAML/JSON frontmatter.

---

## 3. Critical Code Paths

### The Orchestrator: `builder.Build()` (`internal/builder/builder.go`)
This function dictates the lifecycle. It loads templates (theme + shared), calls the scanner, and then iterates through collections to trigger generation. It is the sync point for everything.

### Content Discovery: `ScanCollections` (`internal/builder/scanner.go`)
It treats **any subfolder** in the source directory as a Collection.
*   If you make a folder named `notes`, Tamarind automatically treats it as a collection and tries to generate `notes.html`.
*   **Mechanism**: `filepath.Walk` looking for `.md` files.

### Shortcode Processor: `processShortcodes` (`internal/builder/shortcodes.go`)
Tamarind uses **Regex Replacement** *before* Markdown parsing to handle shortcodes.
*   **Pros**: Fast, simple, no AST manipulation needed.
*   **Cons**: Fragile. Syntax errors in user content can break the regex (e.g., nested braces).
*   **Key Logic**: Looks for patterns like `{{ figure src="..." }}` and replaces them with raw HTML.

### Live Reloading (`internal/server/server.go`)
Uses `fsnotify` to watch the file system. When a file setup changes:
1.  Triggers a rebuild.
2.  Sends an event via **Server Sent Events (SSE)** to the browser.
3.  Injected JavaScript in `pages.go` receives the event and reloads the page.

---

## 4. Bottlenecks

1.  **Image Optimization Loop**:
    *   **Issue**: `OptimizeImage` is called during the file walk. While it checks file existence, traversing a large directory of thousands of images and checking dimension rules (`images.go`) can slow down the build loop significantly on slower disks.
    *   **Mitigation**: Currently handled sequentially.

2.  **Full Rebuilds**:
    *   **Issue**: Changing *one* comma in one article triggers a rebuild of the entire site (CSS, Indices, Asset Copying).
    *   **Impact**: Fast for <1000 pages (~200ms), but will scale linearly and eventually become sluggish (seconds) for massive sites.

3.  **Regex Shortcodes**:
    *   **Issue**: Running multiple Regex passes (Mermaid, Math, Includes, Youtube, Gist) on every single Markdown body strings is CPU intensive.

---

## 5. Complexity & "Gotchas"

*   **Relative Path Hell**:
    *   The system calculates `RelPrefix` (e.g., `../../`) manually in `pages.go` to ensure site portability (so it works without a web server). This logic is fragile and hard to get right for deep nesting.
*   **Template Coupling**:
    *   Themes are "embedded" in the binary but extracted to `tmp` folders during build. Debugging which template is actually being used (the embedded one vs. a local override) can be confusing.
*   **Silent Failures**:
    *   If a template key is missing (e.g. `{{ .Data.WrongKey }}`), Go templates often fail silently or print `<no value>`, which confuses users.

---

## 6. Refactoring Opportunities

### 1. **Incremental Builds**
Instead of `Build()`, implement a `Watcher` state that knows *which* file changed.
*   If `.css` changed -> Copy assets only.
*   If `.md` changed -> Re-render that page + Index pages.

### 2. **Interface-Based Scanner**
Decouple `ScanCollections` from `os.ReadDir`.
*   **Why**: Allows testing with an in-memory file system (e.g., `afero`) and makes the builder portable (e.g., running inside WASM in the future).

### 3. **AST-Based Shortcodes**
Move from Regex to a Markdown AST extension (Goldmark extension).
*   **Why**: More robust parsing, better error reporting, and avoids accidental matches inside code blocks.

### 4. **Structured Asset Pipeline**
Currently, `assets.go` just copies files. A proper pipeline structure would allow:
*   Minification (CSS/JS).
*   SASS/SCSS support.
*   Fingerprinting (cache busting).

---

> **Summary**: Tamarind is built for *flow* and *speed*. The code favors straightforward procedural logic over abstract patterns. It is robust for its intended scope (blogs, documentation, portfolios) but would require architectural changes (Incremental builds, AST parsing) to scale to enterprise levels (10k+ pages).
