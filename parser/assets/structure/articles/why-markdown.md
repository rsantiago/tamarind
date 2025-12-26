---
title: Why Markdown?
subtitle: The universal text format
date: 2025-11-30
tags: [writing, tools]
---

# The Case for Markdown

Markdown is a lightweight markup language with plain text formatting syntax. It is designed so that it can be converted to HTML and many other formats using a tool by the same name.

## Readability

The overriding design goal for Markdown's formatting syntax is to make it as readable as possible. The idea is that a Markdown-formatted document should be publishable as-is, as plain text, without looking like it's been marked up with tags or formatting instructions.

## Portability

If you write your content in a database (like WordPress) or a proprietary format (like Notion), you are locked in. Exporting is often painful.

Markdown is just text. You can open it in Notepad, Vim, VS Code, or IA Writer. You can email it. You can stick it in a git repository. It is the closest thing we have to a "universal file format" for the written word.

## Tamarind's Flavor

Tamarind uses strict **CommonMark** parsing via the `goldmark` library. This means we support:

*   **Headers**: `# H1`, `## H2`
*   **Lists**: Ordered `1.` and unordered `-`
*   **Links**: `[text](url)`
*   **Images**: `![alt](url)`
*   **Code Blocks**: Fenced with triple backticks.

We also support **Frontmatter** (YAML style) at the top of the file for metadata.
