---
title: Architecture
subtitle: Under the hood of Tamarind
date: 2025-12-02
tags: [technical, go]
---

# System Architecture

Tamarind is designed as a pipeline that transforms raw text into structured HTML.

## The Core Components

### 1. The Embed System (`fs.Embed`)
Tamarind uses Go 1.16+ native embedding to package the default file structure and templates into the executable. This allows `tamarind init` to hydrate a new project directory from the binary itself, without needing to download files from the internet.

### 2. The Parser (`goldmark`)
We utilize `github.com/yuin/goldmark` for markdown parsing. It is compliant, fast, and extensible. We configure it to support GitHub Flavored Markdown (tables, task lists) and auto-generated heading IDs.

### 3. The Template Engine (`html/template`)
Go's standard library template engine is powerful and secure. Tamarind parses `.mdt` files (Markdown Templates) which contain the HTML skeleton. Data structures like `PageData` and `ArticleMeta` are injected into these templates during the build process.

### 4. The Data Model
*   **PageData**: Contains the content (`Body`), metadata (`Title`, `Date`), and navigation data (`Menu`).
*   **Menu**: Generated dynamically by scanning the `pages/` directory.

## Build Process

1.  **Scan**: The program scans `pages/` and `articles/` to build the content graph.
2.  **Parse**: Metadata is extracted via YAML frontmatter.
3.  **Render**: Each markdown file is converted to HTML, wrapped in the appropriate template, and written to the `website/` directory.
