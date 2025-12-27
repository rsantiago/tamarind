---
title: About Tamarind
subtitle: What is this project?
date: 2025-12-01
tags: [about]
menu_label: About
menu_order: 200
---

# What is Tamarind?

Tamarind is a minimalist static site generator (SSG) built in Go. It was created to solve a specific, recurring problem in the software industry: the fragility and over-complexity of modern web development environments.

## The Problem

Have you ever tried to run a 3-year-old React project? Often, it simply fails. `npm` packages are deprecated, node versions stick, breaking changes in dependencies cascade through the system, and the build pipeline collapses under its own weight. The content—your words—becomes held hostage by the tools required to render them.

This is unacceptable for archival content, documentation, or personal blogs. These should be digital monuments, not fragile machines.

## The Solution

Tamarind removes the moving parts. It is a radical simplification of the publishing stack.

1.  **No Node.js**: It runs as a compiled binary. It doesn't care if you have Node, Python, or Ruby installed.
2.  **No External Templates**: The default theme is baking into the binary. You don't need to `git clone` a separate theme repository or manage a submodule.
3.  **Standard Markdown**: Content is written in standard Markdown (CommonMark). Your writing is portable.

We believe that your content should outlive the tools you use to create it. Tamarind is designed to be a durable vessel for your words, producing HTML that arguably has the longest shelf-life of any digital format.
