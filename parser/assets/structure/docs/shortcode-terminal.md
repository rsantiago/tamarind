---
title: Terminal Window Shortcode
date: 2026-06-16
description: "Simulate styled terminal and CLI console environments directly in your documentation using the terminal shortcode."
tags: [guide, shortcode, styling]
---

The `{{` `terminal` `}}` shortcode provides an elegant way to showcase CLI commands, terminal inputs, compilation logs, or script outputs. Instead of displaying commands inside standard, flat code blocks, Tamarind renders them inside a realistic, themed terminal emulator box complete with macOS-style window controls.

---

## 1. Syntax and Usage

To create a simulated terminal window, wrap your CLI commands and output using the `{{` `terminal` `}}` and `{{` `/terminal` `}}` block tags:

```markdown
{{!}}{ terminal }}
$ tamarind build
Building website...
Compiling 18 markdown files to static HTML
Done in 12ms!
{{!}}{ /terminal }}
```

### Escaping the Shortcode
If you need to document the shortcode itself without executing it inside a code block, prefix it with `{{!}}{` like this:
`{{!}}{ terminal }}` and `{{!}}{ /terminal }}`. For inline text, you can separate the brackets using backticks: `{{` `terminal` `}}`.

---

## 2. Interactive Styling: Input vs. Output

To make your terminal documentation intuitive and readable:
*   **User Inputs / Commands**: Prefix CLI command lines with a dollar sign (`$ `).
*   **System Outputs**: Place command results and log outputs on separate lines directly below without the prompt prefix.

This maintains clear visual separation between commands and their outputs inside the console wrapper.

---

## 3. Auto-Theming and Mac-Style Controls

The `{{` `terminal` `}}` shortcode is designed with **automatic theme adaptation**. You do not need to manually configure colors or layout styles for different websites. 

When you switch between Tamarind themes (like `cupertino`, `nordic`, `brutal`, `neon`, or `zephyr`), the terminal component automatically transforms to match the theme's aesthetic:

*   **macOS-style controls**: In themes like Cupertino, Nordic, and Zephyr, the terminal is adorned with three decorative red, yellow, and green window management buttons at the top left.
*   **Brutalism theme**: In the `brutal` theme, the header is hidden, and the terminal displays as a retro, high-contrast, neon-green monochrome box with a thick black border.
*   **Neon theme**: In the `neon` theme, it renders with a glowing cyan shadow and matching cyber-themed window dots.
*   **Dark-Mode native**: All terminal blocks default to a sleek dark terminal theme regardless of the site's default body theme, ensuring maximum developer readability for logs and code.

---

## 4. Live Terminal Preview

Here is a live rendering of the terminal emulator in action:

{{ terminal }}
$ tamarind init
Initializing new static website project...
Creating directory structure:
  ├── writer-sandbox/
  │   ├── docs/
  │   ├── index.md
  │   └── tamarind.yaml
  └── website/ (compiled output)
Success! Project ready for compilation.
{{ /terminal }}

---

## 5. CSS Reference under the Hood

When compiled, the markdown is transformed into the following structured HTML:

```html
<div class="terminal">
  <div class="terminal-header">
    <span class="dot red"></span>
    <span class="dot yellow"></span>
    <span class="dot green"></span>
  </div>
  <pre class="terminal-content">
    <code>
      $ command input
      command output
    </code>
  </pre>
</div>
```

This clean structure allows custom stylesheets to override or add styles if needed.
