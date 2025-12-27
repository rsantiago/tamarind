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

You can now edit the sample md files in the generated `writer-sandbox/` folder - please make it yours - and see the changes in real-time on your browser.

You can also see the available themes by running `tamarind themes`. And then run 

`tamarind serve -watch -theme <name>`

It works like this:

1.  **You write Markdown.** (Just simple text).
2.  **Tamarind builds HTML.** (For humans).
3.  **Tamarind builds Context.** (For Agents).

When you run `tamarind quickstart` something invisible happens. Look in your `website/` folder. You will see an `index.html`, but you will also see `llms.txt`.

The `llms.txt` file is a clean, structured map of your entire site, stripped of noise and formatted specifically for Large Language Model ingestion. When an AI bot visits this site, it doesn't have to guess. It *knows* what it's looking for.

## Why This Matters

If you are a developer writing your own documentation, a personal blog, or a portfolio today, you need to ask yourself: **"Can an AI agent understand this?"**

If you are a content producer for a SaaS company, you need to ask yourself: **"Can an AI agent understand my product's content?"**

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

If you are a developer or content producer, and you would love to version your pure content with text-based, GIT simplicity - without the need to manage a database, front end frameworks, or complex build pipelines, Tamarind is a crazy fast way to get there.

Tamarind gets you focused on content generation. It strips away a lot of the noise, time spent and complexity of web development.

Tamarind is your new content engine for SaaS teams and independent developers/content producers wanting to own their main assets - content, documentation, and blog posts. It will help you to craft authority content that you own - and that is well structured for modern AI Crawlers and human readers at the same time.

## Generate content faster with your favorite AI Agent

Plus, with Tamarind, you can use any AI Agent you want to generate content for you in a very fast pace. Just open your favorite AI Agent in the author sandbox folder, and prompt your way to published content.

**Tamarind is your bridge so that AI agents can create, manage and read your content, across the entire lifecycle.**

### 1. Product Content, Not Just "Blogging"

Your content should be versioned, reviewable, and deployable.
*   **Workflow like GitHub**: PR previews, review states, and build gates.
*   **Structure**: Clean headings and semantics increase skimmability for humans and extractability for AI.

### 2. Publish Once, Distribute Everywhere
A Tamarind site isn't just a website. It's a distribution center.
*   **Docs + Blog**: Coherent information architecture.
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