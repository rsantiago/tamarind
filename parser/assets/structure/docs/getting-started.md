---
title: Getting Started
subtitle: A comprehensive guide to starting with Tamarind
date: 2025-12-03
tags: [guide, help]
contextual_sidebar: "docs"
---

# User Manual

## Installation

Download the binary for your architecture from the releases page, or build from source:

```bash
go build -o tamarind main.go
```

## Commands

### `init`
Initializes a new project in the current directory.
```bash
./tamarind init
./tamarind init --force  # Overwrites existing files
```

### `build`
Compiles the markdown files into a static website.
```bash
./tamarind build
```
This generating the `website/` folder.

### `serve`
Starts a local development server.
```bash
./tamarind serve --port 3000
```
It automatically handles port conflicts by asking the OS to kill the process on that port.

## Directory Structure

*   `writer-sandbox/`: **Your Content**.
    *   `pages/`: Markdown files that become top-level pages and menu items.
    *   `articles/`: Markdown files that fill the blog section.
*   `templates/`: **Your Design**. Modify `style.css` or `page.mdt` here.
