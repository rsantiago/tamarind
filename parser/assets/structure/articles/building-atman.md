---
title: Building Tamarind
subtitle: The story behind the code
date: 2025-11-20
tags: [devlog, story]
---

# The Origin Story

It started on a Friday night. I was trying to update my personal blog, which was built with Gatsby. I ran `npm install`.

**1,432 vulnerabilities found.**

I spent the next three hours fixing dependency conflicts, updating React versions, and fighting with GraphQL queries just to fix a typo in a blog post.

## Enough.

I realized that for a simple blog, a Single Page Application (SPA) was overkill. I didn't need client-side routing. I didn't need state management. I just needed HTML.

I opened my code editor and started writing a Go program.

## The First Prototype

The first version was 50 lines of code. It read a directory of text files and used `strings.Replace` to inject content into an HTML header/footer string. It was crude, but it was fast. It built the site in 4ms.

## Evolution

Over the next few weeks, I added:

1.  **Goldmark**: For proper Markdown parsing.
2.  **Frontmatter**: To handle titles and dates.
3.  **HTML Templates**: To separate logic from view.

The result is what you see today: Tamarind. A tool born from frustration, refined by minimalism.
