---
title: Turn markdown content into a gorgeous website that's seen by AI Agents. Just run `tamarind quickstart`
date: 2025-01-01
description: The AI-First Web Engine.
site_name: Tamarind
menu_label: Home
menu_order: 1
---
{{ figure src="images/logo.png" width="300px" }}

## The AI-First Static Site Generator

Tamarind has a single purpose: **To quickly turn your content into a fast, beautiful, and perfectly readable website - by both Humans and AI Agents like ChatGPT, Claude, Gemini, and others.**

In order to generate a demo site, get our binary from [here](https://github.com/tamarind/tamarind/releases) and run `tamarind quickstart` on your terminal window.

That's it. 

Tamarind will generate and serve a demo site for you locally (just like this one), so you can see it in action immediately on your browser.

You can now edit the md files in the generated `website/` folder - please make it yours - and see the changes in real-time by running `tamarind build -theme gram && tamarind serve`.

You can also see the available themes by running `tamarind themes`. Just change from `gram` to any other theme name to see it in action.

It works like this:

1.  **You write Markdown.** (Just simple text).
2.  **Tamarind builds HTML.** (For humans).
3.  **Tamarind builds Context.** (For Agents).

When you run `tamarind quickstart` something invisible happens. Look in your `website/` folder. You will see an `index.html`, but you will also see `llms.txt`.

The `llms.txt` file is a clean, structured map of your entire site, stripped of noise and formatted specifically for Large Language Model ingestion. When an AI bot visits this site, it doesn't have to guess. It *knows* what it's looking for.

## Why This Matters

If you are a developer writing your own documentation, a personal blog, or a portfolio today, you need to ask yourself: **"Can an AI understand this?"**

If you are a content producer for a SaaS company, you need to ask yourself: **"Can an AI understand my product's content?"**

If the answer is "I hope AI Agents gets it," you are already behind. 

With Tamarind, the answer is "Yes, my content will be understood natively by AI Agents."

Every website you have ever built is invisible to half the internet.

We usually build websites just for humans. Developers obsess over CSS, dark mode, and responsive layouts. 

But while we were refining pixel margins on mobile, something fundamental shifted. The internet is no longer just for eyeballs. It is for **Agents**.

AI models—ChatGPT, Claude, search crawlers—are the new "browsers". 

And right now, your beautiful React app looks like a garbled mess of `<div>` soup to AI Agents. Your documentation is hallucinated. Your blog is ignored.

Developers everywhere are building dead ends for the smartest users on the planet.

**Tamarind fixes this.**

## Ship Content Like Code

If you pitch yourself as a "blogger", you inherit the "SEO is dying" anxiety.

**Tamarind is a content engine for SaaS teams.**

### 1. Product Content, Not Just "Blogging"

Your content should be versioned, reviewable, and deployable.
*   **Workflow like GitHub**: PR previews, review states, and build gates.
*   **Structure**: Clean headings and semantics increase skimmability for humans and extractability for AI.

### 2. Publish Once, Distribute Everywhere
A Tamarind site isn't just a website. It's a distribution center.
*   **Docs + Blog + Changelog**: Coherent information architecture.
*   **AI-Ready Publishing**: Auto-generated `llms.txt`, consistent answer blocks, and strong metadata defaults.

### 3. The "Money Pages" System
Don't just write posts. Build assets.
*   **Use-case Pages**: Comparison pages, integration guides, and pricing support.
*   **Credibility**: We prioritize assets that feed social and email distribution.

## Only What Matters

We stripped everything else away.
Most web frameworks are 90% configuration and 10% content. You spend hours fighting Webpack, updating `node_modules`, and debugging hydration errors.

Tamarind has:
*   **No Config**: This site has no configuration file. It just works.
*   **No Dependencies**: It is a single binary. No Node.js. No Python.
*   **No Fluff**: It processes Markdown into HTML in milliseconds.

### Try It Now.

This page is just a file called `index.md`.
Open it. Delete this text. Write your own.

You are now building for the Agent Web.