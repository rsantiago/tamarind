---
title: Turn markdown files into an gorgeous, AI crawler-friendly website - using a single command line
date: 2025-01-01
description: The AI-First Web Engine.
---
{{ figure src="images/logo.png" width="500px" }}

---

## The AI-First Engine

Tamarind is a static site generator with a single purpose: **To make your content perfectly readable by both Humans and Machines.**

It works like this:
1.  **You write Markdown.** (Just simple text).
2.  **Tamarind builds HTML.** (For humans).
3.  **Tamarind builds Context.** (For Agents).

When you ran `tamarind quickstart` just now, something invisible happened. Look in your `website/` folder. You will see an `index.html`, but you will also see `llms.txt`. 

This file is a clean, structured map of your entire site, stripped of noise and formatted specifically for Large Language Model ingestion. When an AI bot visits this site, it doesn't have to guess. It *knows*.

## Only What Matters

We stripped everything else away.
Most web frameworks are 90% configuration and 10% content. You spend hours fighting Webpack, updating `node_modules`, and debugging hydration errors.

Tamarind has:
*   **No Config**: This site has no configuration file. It just works.
*   **No Dependencies**: It is a single binary. No Node.js. No Python.
*   **No Fluff**: It processes Markdown into HTML in milliseconds.

## Why This Matters

If you are writing documentation, a blog, or a portfolio today, you need to ask yourself: **"Can an AI understand this?"**

If the answer is "I hope so," you are already behind. 
With Tamarind, the answer is "Yes, natively."

### Every website you have ever built is invisible to half the internet.

We usually build for humans. We obsess over CSS, dark mode, and responsive layouts. But while we were refining pixel margins on mobile, something fundamental shifted. The internet is no longer just for eyeballs. It is for **Agents**.

AI models—ChatGPT, Claude, search crawlers—are the new "browsers". And right now, your beautiful React app looks like a garbled mess of `<div>` soup to them. Your documentation is hallucinated. Your blog is ignored.

You are building dead ends for the smartest users on the planet.

**Tamarind fixes this.**

### Try It Now.

This page is just a file called `index.md`.
Open it. Delete this text. Write your own.

You are now building for the Agent Web.
