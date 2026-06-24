---
title: Working with Drafts
subtitle: Safe writing workflows in Tamarind
date: 2025-12-24
tags: [feature, guide, workflow]
description: Learn how to mark content as drafts and preview them locally before publishing.
---

Tamarind provides a built-in workflow for managing draft content, allowing you to work on unfinished articles without accidentally publishing them to your live site or exposing them to search engines and AI agents.

## Marking Content as Draft

To mark any article or page as a draft, simply add the `draft: true` field to your Markdown frontmatter:

```yaml
---
title: My Work In Progress
date: 2025-12-24
draft: true
---

This content will not be visible in the default build.
```

## Default Behavior

By default, when you run `tamarind build` or `tamarind serve`, **all content marked as `draft: true` is ignored**.

This means:
1.  **Invisible**: The pages are not generated.
2.  **No Leaks**: They do not appear in `articles.html`, tag pages, or RSS feeds.
3.  **Secure**: They are explicitly excluded from `robots.txt`, `sitemap.xml`, `llms.txt`, and `llms_full.txt`.

## Previewing Drafts

To preview your draft content locally, use the `-drafts` flag with the build or serve commands:

### Serve with Drafts
```bash
./tamarind serve -drafts
```

### Build with Drafts
```bash
./tamarind build -theme neon -drafts
```

When this flag is active, all draft content will be rendered and visible in lists, just like published content. This allows you to verify formatting and layout before "shipping" the post by removing the `draft: true` line.
