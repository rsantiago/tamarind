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

### Standard Actions
The standard button components are designed for primary calls to action, secondary operations, and discrete navigation.

<div class="button-group" style="margin-top: 1.5rem; margin-bottom: 2.5rem; flex-wrap: wrap; gap: 24px; justify-content: flex-start;">
    {{< button href="#" type="primary" >}}Primary Button{{</ button >}}
    {{< button href="#" type="secondary" >}}Secondary Button{{</ button >}}
    {{< button href="#" type="ghost" >}}Ghost Button{{</ button >}}
</div>

### Compact Actions
Smaller button variations are optimized for card-level actions, sidebars, or inline operations.

<div class="button-group" style="margin-top: 1.5rem; margin-bottom: 2rem; flex-wrap: wrap; gap: 20px; justify-content: flex-start;">
    {{< button href="#" type="primary" size="sm" >}}Small Primary{{</ button >}}
    {{< button href="#" type="secondary" size="sm" >}}Small Secondary{{</ button >}}
</div>

---

## 3. Forms

<div class="card card-padding">
    <form action="#" method="POST">
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
        <div class="form-group">
             <label class="form-label">
                <input type="checkbox" class="form-checkbox"> Subscribe to newsletter
             </label>
        </div>
        <div class="form-group">
             <label class="form-label">Preferred Contact:</label>
             <label class="form-label">
                <input type="radio" name="contact" class="form-radio" checked> Email
             </label>
             <label class="form-label">
                <input type="radio" name="contact" class="form-radio"> Phone
             </label>
        </div>
        <div class="form-group">
            <label class="form-label">Attachment</label>
            <input type="file" class="form-file">
        </div>
        <button type="submit" class="btn btn-primary">Send Message</button>
    </form>
</div>

---

## 4. Cards & Badges

{{< card >}}
    <h3>Feature Card</h3>
    <p>This is a standard card component. It uses <code>var(--card-bg)</code>.</p>
    {{< badge type="primary" >}}New{{</ badge >}}
{{</ card >}}
<br>
{{< card >}}
    <h3>Another Card</h3>
    <p>Cards are great for grouping related content in a grid layout.</p>
    {{< badge >}}Legacy{{</ badge >}}
{{</ card >}}

---

## 5. Shortcodes

### Callouts
{{< alert type="info" title="Info" >}}
This is an informational callout.
{{</ alert >}}

{{< alert type="warn" title="Warning" >}}
This is a warning callout.
{{</ alert >}}

{{< alert type="error" title="Error" >}}
This is an error callout.
{{</ alert >}}

{{< alert type="tip" title="Tip" >}}
This is a tip callout.
{{</ alert >}}

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
