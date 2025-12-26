---
title: "Feature Idea: Local Admin & Visual Editor"
status: Draft
created: 2025-12-18
tags: [admin, features, future]
---

# Concept
Integrate a visual Markdown editor/CMS directly into the `tamarind serve` command to allow for local editing without external tools.

# Architecture: "The Local Admin"
Current `serve` command only reads. The new version would read/write.

## 1. Backend (`main.go`)
- Update `serve` command to multiplex routes.
- **GET /***: Serves the static website (existing).
- **GET /admin**: Serves the editor interface.
- **GET /api/files**: Returns JSON list of all `.md` files in `writer-sandbox`.
- **POST /api/save**: Accepts `{path, content}`, validates it is within sandbox, writes to disk, triggers site rebuild.

## 2. Frontend (Embedded)
- Embed **EasyMDE** (SimpleMDE fork) CSS/JS into the binary.
- Create `assets/templates/admin/index.html`:
  - Sidebar: List of files (fetched from `/api/files`).
  - Main: Editor implementation.
  - Footer: Status (Save success/fail).

# Benefits
1.  **Zero Dependency**: No Node.js, no Database.
2.  **Portable**: Works perfectly offline/air-gapped.
3.  **Fast**: Direct filesystem IO.

# Implementation Steps
1.  Verify `net/http` routing logic to handle API vs Static files.
2.  Add `listMarkdownFiles` function to scanner.
3.  Add `handleSave` function with security checks (prevent directory traversal).
4.  Embed editor assets.
