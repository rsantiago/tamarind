---
title: The AI-First Web Engine.
date: 2026-01-01
subtitle: Generate your AI & Human friendly website in seconds, not days.
site_name: Tamarind
menu_label: Home
menu_order: 1
---
{{ figure src="images/logo.png" width="300px" }}

## Generate. Edit. Deploy.

Stop fighting with complex frameworks. Get a live local development server running instantly, ready for you to add your content. Go from zero to a live site in seconds:

{{ timeline }}
  {{ item title="Run the Quickstart" number="1" }}
  Run the quickstart command in your terminal. This boots up the development server instantly with zero configuration:
  {{ terminal }}
    {{ tab title="macOS" }}
    $ tamarind quickstart
    {{ /tab }}
    {{ tab title="Linux" }}
    $ tamarind quickstart
    {{ /tab }}
    {{ tab title="Windows" }}
    PS> .\tamarind.exe quickstart
    {{ /tab }}
  {{ /terminal }}
  Your sample website will be compiled and served locally at `http://localhost:8000` with automatic hot-reloading!
  {{ /item }}

  {{ item title="Edit the Content" number="2" }}
  Use your favorite AI Agent or text editor to mass-edit the `.md` files in the `writer-sandbox/` folder. See your changes render instantly in the browser.
  
  *Example prompts:*
  
  • "Turn the files in `writer-sandbox/` into md files for a website about [topic]"
  
  • "Add a new page about [topic] to this site structure"
  {{ /item }}

  {{ item title="Try Different Themes" number="3" }}
  List available themes and build/compile with your theme of choice:
  {{ terminal }}
    {{ tab title="macOS" }}
    $ tamarind themes
    $ tamarind build --theme dark
    {{ /tab }}
    {{ tab title="Linux" }}
    $ tamarind themes
    $ tamarind build --theme dark
    {{ /tab }}
    {{ tab title="Windows" }}
    PS> .\tamarind.exe themes
    PS> .\tamarind.exe build --theme dark
    {{ /tab }}
  {{ /terminal }}
  {{ /item }}

  {{ item title="Compile for Production & Deploy" number="4" }}
  Run the compiler to build highly optimized, responsive pages:
  {{ terminal }}
    {{ tab title="macOS" }}
    $ tamarind build
    {{ /tab }}
    {{ tab title="Linux" }}
    $ tamarind build
    {{ /tab }}
    {{ tab title="Windows" }}
    PS> .\tamarind.exe build
    {{ /tab }}
  {{ /terminal }}
  Tamarind outputs a pure static folder in `website/` with zero runtime dependencies. Deploy it to Netlify, Vercel, S3, or GitHub Pages.
  {{ /item }}
{{ /timeline }}

It's that simple. No more blocks because you can't find the right CSS property. No more fighting with broken framework dependencies. 

Just focus on your message, your content, and your ideas.

But yes, you can still edit the style.css file in the website/assets/ folder to change the style of the generated site (you can use your favorite AI Agent to do it for you).

## The Web Has Two Audiences Now.
**Humans** read with their eyes.
**AI Agents** read with context.

Most sites are built only for humans. They are invisible to the smartest users on the planet (LLMs).

**Tamarind builds for both.**

When you build with Tamarind, you get a beautiful HTML site intended for humans. But you also get a perfect semantic structure intended for AI. Your site becomes a "Source of Truth" that ChatGPT, Claude, and Google Gemini can actually understand.

Every website you have ever built is invisible to half the internet.
We usually build for humans. We obsess over CSS, dark mode, and responsive layouts. But while we were refining pixel margins on mobile, something fundamental shifted. The internet is no longer just for eyeballs. It is for **Agents**.

AI models—ChatGPT, Claude, search crawlers—are the new "browsers". And right now, your beautiful React app looks like a garbled mess of `<div>` soup to them. Your documentation is hallucinated. Your blog is ignored.

You are building dead ends for the smartest users on the planet.

**Tamarind fixes this.**

## Write faster with your favorite AI Agent

Instead of spending hours using old blog interfaces, logging in to admin panels and editing and saving posts one by one, you can now use your favorite AI Agent to mass-edit the md files in the `writer-sandbox/` folder.

Your pages are just regular text files called markdown files. You can edit them with any text editor you want, including Claude, ChatGPT, Gemini, or any other AI Agent.

And by the way, you can edit the style of the site by editing the generated `style.css` file.

## Ship Content Like Code

If you pitch yourself as a "blogger", you also inherit the "SEO is dying" anxiety.

**Tamarind is a content engine for SaaS teams.**

### 1. Product Content, Not Just "Blogging"
Your content should be versioned, reviewable, and deployable.
*   **Your Content Edition workflow can use GitHub now**: PR previews, review states, and build gates.
*   **Structure**: Clean headings and semantics increase skimmability for humans and extractability for AI.

### 2. Publish Once, Distribute Everywhere
A Tamarind site isn't just a website. It's a distribution center.
*   **Pages, Docs and Blog**: You can structure your content in a way that makes sense for you. You can have a blog section, a documentation section, landing pages, and any other section you want - including hidden pages.
*   **AI-Ready Publishing**: Auto-generated `llms.txt`, consistent answer blocks, and strong metadata defaults.

## Only What Matters

We stripped everything else away.

Most web frameworks are 90% configuration, infrastructure and 10% content. You spend hours fighting Webpack, updating `node_modules`, logging in heavy admin panels to edit pages one by one, and fighting with CSS to make it look the way you want.

Tamarind has:
*   **No Config**: This site has no configuration file. It just works.
*   **No Dependencies**: It is a single binary. No Node.js. No Python. Becuase good, fast content is static.
*   **No Fluff**: It processes Markdown into HTML in milliseconds.

### Try It Now.

Once you use Tamarind quickstart, this page is generated on your `writer-sandbox/` folder as a file called `index.md`.

Open it. Delete this text. Write your own.

You are now building for the Agent Web.