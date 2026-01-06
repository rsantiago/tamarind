---
title: "Theme Style Guide"
date: 2024-01-01
description: "A comprehensive showcase of all Tamarind UI components and styles."
menu_order: 99
---

# Theme Style Guide

This page serves as a test bench for the Tamarind Design System. Every element below should look polished and consistent across all themes.

## 1. Typography

# Heading 1
## Heading 2
### Heading 3
#### Heading 4
##### Heading 5
###### Heading 6

This is a standard paragraph. It should be legible, with a comfortable line length and height. It uses the `var(--font-body)` token.

This is a **bold text**, *italic text*, and `inline code` sample.

<p class="lead">This is a Lead Paragraph (.lead). It is used for introductions and summaries. It should be slightly larger than body text.</p>

<p class="caption">This is a caption height (.caption). Useful for metadata or small notes.</p>

---

## 2. Buttons

<div style="display: flex; gap: 10px; flex-wrap: wrap; margin-bottom: 20px;">
    <a href="#" class="btn btn-primary">Primary Button</a>
    <a href="#" class="btn btn-secondary">Secondary Button</a>
    <a href="#" class="btn btn-ghost">Ghost Button</a>
</div>

<div style="display: flex; gap: 10px; flex-wrap: wrap;">
    <a href="#" class="btn btn-primary btn-sm">Small Primary</a>
    <a href="#" class="btn btn-secondary btn-sm">Small Secondary</a>
</div>

---

## 3. Forms

<div class="card card-padding" style="max-width: 400px;">
    <div class="form-group">
        <label class="form-label">Email Address</label>
        <input type="email" class="form-input" placeholder="you@example.com">
    </div>
    <div class="form-group">
        <label class="form-label">Subject</label>
        <select class="form-select">
            <option>General Inquiry</option>
            <option>Support</option>
            <option>Feedback</option>
        </select>
    </div>
    <div class="form-group">
        <label class="form-label">Message</label>
        <textarea class="form-textarea" rows="4" placeholder="Type your message..."></textarea>
    </div>
    <button class="btn btn-primary" style="width: 100%">Send Message</button>
</div>

---

## 4. Cards & Badges

<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px;">
    <div class="card card-padding">
        <h3>Feature Card</h3>
        <p>This is a standard card component. It uses <code>var(--card-bg)</code>.</p>
        <span class="badge badge-primary">New</span>
    </div>
    <div class="card card-padding">
        <h3>Another Card</h3>
        <p>Cards are great for grouping related content in a grid layout.</p>
        <span class="badge">Legacy</span>
    </div>
</div>

---

## 5. Shortcodes

### Callouts
<div class="callout callout-info">
    <div class="callout-title">Info</div>
    This is an informational callout.
</div>

<div class="callout callout-warn">
    <div class="callout-title">Warning</div>
    This is a warning callout.
</div>

<div class="callout callout-error">
    <div class="callout-title">Error</div>
    This is an error callout.
</div>

<div class="callout callout-tip">
    <div class="callout-title">Tip</div>
    This is a tip callout.
</div>

### Terminal
<div class="terminal">
    <div class="terminal-header">
        <div class="dot red"></div>
        <div class="dot yellow"></div>
        <div class="dot green"></div>
    </div>
    <div class="terminal-content">
        $ tamarind init<br>
        > Creating new project...<br>
        > Done!
    </div>
</div>
