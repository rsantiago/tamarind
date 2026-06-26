---
title: System Architecture
subtitle: Under the hood of Tamarind's static site generator
date: 2026-06-25
tags: [technical, go, architecture, parser]
---

# Tamarind System Architecture

Tamarind is a highly optimized, zero-dependency Static Site Generator (SSG) built in Go. It is designed as a high-performance pipeline that ingests raw markdown, resolves complex extensible shortcodes, and emits pristine, optimized HTML.

This document serves as a technical overview for developers, theme designers, and open-source contributors looking to understand how Tamarind operates under the hood.

---

## 1. High-Level Pipeline

The build process is orchestrated entirely within the `parser/internal/builder` package and operates in a strict, deterministic sequence:

1. **Configuration Loading**: Reads `tamarind.yaml` and sets up the active theme.
2. **File Scanning**: The `Scanner` traverses the target source directory (e.g., `docs/` or `articles/`), building an in-memory graph of all markdown files and creating a dynamic navigation `Menu`.
3. **Asset Processing**: Static assets (CSS, JS) are mirrored. Images are optionally processed into responsive, multi-resolution variants (`images.go`).
4. **Parsing & Shortcode Resolution**: The `goldmark` engine converts Markdown to HTML, while the `PluginRegistry` evaluates and injects dynamic shortcodes (Charts, Mermaid, Math, etc.).
5. **Template Rendering**: The HTML output is injected into Go `html/template` skeletons (`.mdt` files) located in the active theme folder.
6. **SEO & Discovery Engine**: `sitemap.xml`, `robots.txt`, and `llms.txt` (for AI agents) are automatically generated.

---

## 2. Directory Structure

The repository is modularized to strictly separate the compiler engine from the themes and templates.

```text
parser/
├── main.go                     # CLI entry point (build, serve, update)
├── assets/                     # Packaged via go:embed natively into the binary
│   ├── structure/              # Default scaffold for 'tamarind init'
│   └── templates/              # Core HTML/CSS templates for all built-in themes
├── internal/
│   ├── builder/                # The core build engine and shortcode plugins
│   ├── config/                 # YAML configuration parser
│   ├── models/                 # Shared data structures (PageData, ArticleMeta)
│   ├── seo/                    # XML sitemap and robots generator
│   ├── server/                 # Local development server
│   └── updater/                # OTA self-update mechanism
```

---

## 3. The `go:embed` Virtual Filesystem

Tamarind compiles to a **single static binary**. To achieve this without requiring users to download external dependencies, the entire `assets/` directory is baked directly into the executable using Go's native `embed` package. 

When a user runs `tamarind init`, the CLI reads the embedded `assets/structure/` directory and hydrates a new project scaffold instantly on their local filesystem. Similarly, theme templates are read directly from memory during the build process, ensuring blisteringly fast IO.

---

## 4. The Shortcode & Plugin Registry

Tamarind features a highly extensible shortcode system. Instead of relying on a monolithic parser, specialized components are registered as isolated "Plugins."

### Plugin Architecture
Located in `internal/builder/registry.go`, the `PluginRegistry` allows developers to map string keys to handler functions. 

```go
// Example of how plugins are registered in Tamarind
registry.Register("barchart", generateBarChart)
registry.Register("mermaid", generateMermaidDiagram)
registry.Register("math", generateLaTeX)
```

### Component Isolation
Every feature is encapsulated in its own file (e.g., `plugin_chart.go`, `plugin_tabs.go`, `plugin_terminal.go`). This ensures that if a specific component needs a bug fix, the rest of the compilation pipeline remains entirely untouched.

Tamarind natively supports complex data visualizations (Pie, Bar, Line charts) by injecting structured JSON data directly into the shortcode blocks. These charts are rendered as clean, zero-dependency SVG or CSS grids, ensuring maximum performance without shipping massive JavaScript libraries to the client.

---

## 5. The Data Model

As the scanner reads the file system, it populates shared structs defined in `internal/models/models.go`. The primary data structure injected into the HTML templates is `PageData`:

```go
type PageData struct {
    SiteName      string
    Title         string
    Description   string
    Body          template.HTML // The compiled Markdown
    Menu          []MenuItem    // The auto-generated navigation tree
    Tags          []string
    LastModified  string
    Theme         string
    IsArticle     bool
}
```

Templates (e.g., `page.mdt`) access these variables directly using Go template syntax: `&#123;&#123; .Title &#125;&#125;` and `&#123;&#123; .Body &#125;&#125;`.

---

## 6. Development Server

Tamarind ships with a built-in static server (`tamarind serve`). It mounts the target output directory (usually `website/` or `public/`) to a local port. 

Currently, the server is optimized for static file delivery. When making changes to markdown files or the theme, you must rebuild the site using `tamarind build` to reflect the updates.
