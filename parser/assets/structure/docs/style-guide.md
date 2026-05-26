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

**Shortcode Syntax:**
```markdown
{{{{!}}button href="#" type="primary"}}Primary Button{{{{!}}/button}}
{{{{!}}button href="#" type="secondary"}}Secondary Button{{{{!}}/button}}
{{{{!}}button href="#" type="ghost"}}Ghost Button{{{{!}}/button}}
```

**Live Output:**
<div class="button-group" style="margin-top: 1.5rem; margin-bottom: 2.5rem; flex-wrap: wrap; gap: 24px; justify-content: flex-start;">
    {{< button href="#" type="primary" >}}Primary Button{{</ button >}}
    {{< button href="#" type="secondary" >}}Secondary Button{{</ button >}}
    {{< button href="#" type="ghost" >}}Ghost Button{{</ button >}}
</div>

### Compact Actions
Smaller button variations are optimized for card-level actions, sidebars, or inline operations.

**Shortcode Syntax:**
```markdown
{{{{!}}button href="#" type="primary" size="sm"}}Small Primary{{{{!}}/button}}
{{{{!}}button href="#" type="secondary" size="sm"}}Small Secondary{{{{!}}/button}}
```

**Live Output:**
<div class="button-group" style="margin-top: 1.5rem; margin-bottom: 2rem; flex-wrap: wrap; gap: 20px; justify-content: flex-start;">
    {{< button href="#" type="primary" size="sm" >}}Small Primary{{</ button >}}
    {{< button href="#" type="secondary" size="sm" >}}Small Secondary{{</ button >}}
</div>

---

## 3. Forms

These zero-config, highly-styled responsive form shortcodes are built directly into the engine, making it trivial to spin up contact pages, subscription prompts, or query filters.

**Shortcode Syntax:**
```markdown
{{{{!}}form action="#" method="POST"}}
  {{{{!}}form-input label="Email Address" type="email" placeholder="you@example.com"}}
  
  {{{{!}}form-select label="Subject"}}
    {{{{!}}option}}General Inquiry{{{{!}}/option}}
    {{{{!}}option}}Support{{{{!}}/option}}
    {{{{!}}option}}Feedback{{{{!}}/option}}
  {{{{!}}/form-select}}
  
  {{{{!}}form-textarea label="Message" placeholder="Type your message..." rows="4"}}
  
  {{{{!}}form-checkbox label="Subscribe to newsletter"}}
  
  {{{{!}}form-radio-group label="Preferred Contact:"}}
    {{{{!}}form-radio name="contact" label="Email" checked="true"}}
    {{{{!}}form-radio name="contact" label="Phone"}}
  {{{{!}}/form-radio-group}}
  
  {{{{!}}form-file label="Attachment"}}
  
  {{{{!}}button href="#" type="primary"}}Send Message{{{{!}}/button}}
{{{{!}}/form}}
```

**Live Output:**
<div class="card card-padding">
    {{ form action="#" method="POST" }}
        {{ form-input label="Email Address" type="email" placeholder="you@example.com" }}
        {{ form-select label="Subject" }}
            {{ option }}General Inquiry{{ /option }}
            {{ option }}Support{{ /option }}
            {{ option }}Feedback{{ /option }}
        {{ /form-select }}
        {{ form-textarea label="Message" placeholder="Type your message..." rows="4" }}
        {{ form-checkbox label="Subscribe to newsletter" }}
        {{ form-radio-group label="Preferred Contact:" }}
            {{ form-radio name="contact" label="Email" checked="true" }}
            {{ form-radio name="contact" label="Phone" }}
        {{ /form-radio-group }}
        {{ form-file label="Attachment" }}
        {{ button href="#" type="primary" }}Send Message{{ /button }}
    {{ /form }}
</div>

---

## 4. Cards & Badges

Cards are elevated containers for grouping related information, and badges are clean indicators for metadata.

**Shortcode Syntax:**
```markdown
{{{{!}}card}}
  <h3>Feature Card</h3>
  <p>This is a standard card component.</p>
  {{{{!}}badge type="primary"}}New{{{{!}}/badge}}
{{{{!}}/card}}

{{{{!}}card}}
  <h3>Another Card</h3>
  <p>Cards are great for grouping related content in a grid layout.</p>
  {{{{!}}badge}}Legacy{{{{!}}/badge}}
{{{{!}}/card}}
```

**Live Output:**
<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr)); gap: 16px; margin-bottom: 24px;">
{{< card >}}
    <h3>Feature Card</h3>
    <p>This is a standard card component. It uses <code>var(--card-bg)</code>.</p>
    {{< badge type="primary" >}}New{{</ badge >}}
{{</ card >}}
{{< card >}}
    <h3>Another Card</h3>
    <p>Cards are great for grouping related content in a grid layout.</p>
    {{< badge >}}Legacy{{</ badge >}}
{{</ card >}}
</div>

---

## 5. Shortcodes

### Callouts
Callouts are HSL-styled alert notifications with automatic vector icon rendering.

**Shortcode Syntax:**
```markdown
{{{{!}}alert type="info" title="Info"}}This is an informational callout.{{{{!}}/alert}}
{{{{!}}alert type="warn" title="Warning"}}This is a warning callout.{{{{!}}/alert}}
{{{{!}}alert type="error" title="Error"}}This is an error callout.{{{{!}}/alert}}
{{{{!}}alert type="tip" title="Tip"}}This is a tip callout.{{{{!}}/alert}}
```

**Live Output:**
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
An interactive terminal block showing commands, scripts, or build logs.

**Shortcode Syntax:**
```markdown
{{{{!}}terminal}}
$ tamarind init
> Creating new project...
> Done!
{{{{!}}/terminal}}
```

**Live Output:**
{{ terminal }}
$ tamarind init
> Creating new project...
> Done!
{{ /terminal }}

---

## 6. Premium UI Components

These premium, theme-aware, zero-config responsive components are built directly into the Tamarind static engine.

### 6.1 Stats & Metrics Grid
Perfect for high-authority SaaS homepages or landing pages displaying key product metrics.

**Shortcode Syntax:**
```markdown
{{{{!}}metrics}}
  {{{{!}}metric value="40+" label="Built-in Skills" }}
  {{{{!}}metric value="12" label="Global Integrations" }}
  {{{{!}}metric value="100%" label="Open Source" }}
  {{{{!}}metric value="BSL" label="Tier Moat License" }}
{{{{!}}/metrics}}
```

**Live Output:**
{{ metrics }}
  {{ metric value="40+" label="Built-in Skills" }}
  {{ metric value="12" label="Global Integrations" }}
  {{ metric value="100%" label="Open Source" }}
  {{ metric value="BSL" label="Tier Moat License" }}
{{ /metrics }}

---

### 6.2 Gradient Icon Feature Cards
Modern, soft-bordered card grid featuring custom SVG icon boxes wrapped in diagonal gradient accent panels.

**Shortcode Syntax:**
```markdown
{{{{!}}features}}
  {{{{!}}feature title="Persistent Memory" gradient="blue-cyan" icon="sparkles" }}
    Remembers preferences, projects, and environment parameters across active sessions.
  {{{{!}}/feature}}
  {{{{!}}feature title="Lightning Core" gradient="purple-pink" icon="bolt" }}
    Ultra-low latency static builds powered by parallel multi-core rendering.
  {{{{!}}/feature}}
  {{{{!}}feature title="Secure Isolation" gradient="orange-red" icon="shield" }}
    Enterprise-grade air-gapped sandboxes running on standard local runtimes.
  {{{{!}}/feature}}
{{{{!}}/features}}
```

**Live Output:**
{{ features }}
  {{ feature title="Persistent Memory" gradient="blue-cyan" icon="sparkles" }}
    Remembers preferences, projects, and environment parameters across active sessions.
  {{ /feature }}
  {{ feature title="Lightning Core" gradient="purple-pink" icon="bolt" }}
    Ultra-low latency static builds powered by parallel multi-core rendering.
  {{ /feature }}
  {{ feature title="Secure Isolation" gradient="orange-red" icon="shield" }}
    Enterprise-grade air-gapped sandboxes running on standard local runtimes.
  {{ /feature }}
{{ /features }}

---

### 6.3 Capabilities Checklist Card
A dense tabular card detailing list capabilities, statuses, and description metadata.

**Shortcode Syntax:**
```markdown
{{{{!}}capabilities-grid}}
  {{{{!}}capabilities title="Compiler Capabilities" }}
    {{{{!}}capability name="Zero-Config CSS" desc="Automatic token harvesting" status="success" statusLabel="Ready" }}
    {{{{!}}capability name="AST Checking" desc="Validates themes against spec" status="success" statusLabel="Ready" }}
    {{{{!}}capability name="Image Optimizer" desc="Vivid multi-viewport generation" status="warning" statusLabel="Beta" }}
    {{{{!}}capability name="WASM Extensions" desc="Custom Go plugins" status="pending" statusLabel="Planned" }}
  {{{{!}}/capabilities}}
{{{{!}}/capabilities-grid}}
```

**Live Output:**
{{ capabilities-grid }}
  {{ capabilities title="Compiler Capabilities" }}
    {{ capability name="Zero-Config CSS" desc="Automatic token harvesting" status="success" statusLabel="Ready" }}
    {{ capability name="AST Checking" desc="Validates themes against spec" status="success" statusLabel="Ready" }}
    {{ capability name="Image Optimizer" desc="Vivid multi-viewport generation" status="warning" statusLabel="Beta" }}
    {{ capability name="WASM Extensions" desc="Custom Go plugins" status="pending" statusLabel="Planned" }}
  {{ /capabilities }}
{{ /capabilities-grid }}

---

### 6.4 Vertical Quick Start Timeline
An interactive vertical progress timeline optimized for tutorials, quick starts, and user guides.

**Shortcode Syntax:**
```markdown
{{{{!}}timeline}}
  {{{{!}}item title="Initialize Tamarind" number="1" }}
    Spin up your local sandbox by running the initialization command in a fresh workspace directory:
    <pre><code>tamarind init</code></pre>
  {{{{!}}/item}}
  {{{{!}}item title="Configure Cortex" number="2" }}
    Tweak visual design tokens directly in `style.css` using theme variables:
    <pre><code>:root {
  --primary-color: #0047AB;
}</code></pre>
  {{{{!}}/item}}
  {{{{!}}item title="Compile static site" number="3" }}
    Build all production-ready HTML and responsive media assets in seconds:
    <pre><code>tamarind build</code></pre>
  {{{{!}}/item}}
{{{{!}}/timeline}}
```

**Live Output:**
{{ timeline }}
  {{ item title="Initialize Tamarind" number="1" }}
    Spin up your local sandbox by running the initialization command in a fresh workspace directory:
    <pre><code>tamarind init</code></pre>
  {{ /item }}
  {{ item title="Configure Cortex" number="2" }}
    Tweak visual design tokens directly in `style.css` using theme variables:
    <pre><code>:root {
  --primary-color: #0047AB;
}</code></pre>
  {{ /item }}
  {{ item title="Compile static site" number="3" }}
    Build all production-ready HTML and responsive media assets in seconds:
    <pre><code>tamarind build</code></pre>
  {{ /item }}
{{ /timeline }}

---

### 6.5 Premium Semantic Alert Containers
HSL-tinted notices with automated vector icon injections for hints, alerts, tips, and warnings.

**Shortcode Syntax:**
```markdown
{{{{!}}alert type="info" title="Info Banner" }}
A standard informational callout container with automatic vector icon rendering.
{{{{!}}/alert}}

{{{{!}}alert type="warning" title="Security Warning" }}
A cautionary alert specifying a potential configuration bottleneck.
{{{{!}}/alert}}

{{{{!}}alert type="success" title="Recompilation Complete" }}
A successful confirmation notice representing a clean pipeline execution.
{{{{!}}/alert}}

{{{{!}}alert type="sparkles" title="Beta Feature" }}
An elevated sparkles notice emphasizing a brand new premium feature.
{{{{!}}/alert}}
```

**Live Output:**
{{ alert type="info" title="Info Banner" }}
A standard informational callout container with automatic vector icon rendering.
{{ /alert }}

{{ alert type="warning" title="Security Warning" }}
A cautionary alert specifying a potential configuration bottleneck.
{{ /alert }}

{{ alert type="success" title="Recompilation Complete" }}
A successful confirmation notice representing a clean pipeline execution.
{{ /alert }}

{{ alert type="sparkles" title="Beta Feature" }}
An elevated sparkles notice emphasizing a brand new premium feature.
{{ /alert }}

---

### 6.6 Custom Interactive Dropdowns
Theme-aware dropdown selectors with custom chevrons and hover animations.

**Shortcode Syntax:**
```markdown
{{{{!}}dropdown id="doc-language" label="Select Language" }}
  {{{{!}}option value="en" selected="true" }}English (US){{{{!}}/option}}
  {{{{!}}option value="es" }}Español{{{{!}}/option}}
  {{{{!}}option value="pt" }}Português{{{{!}}/option}}
  {{{{!}}option value="ja" }}日本語{{{{!}}/option}}
{{{{!}}/dropdown}}
```

**Live Output:**
{{ dropdown id="doc-language" label="Select Language" }}
  {{ option value="en" selected="true" }}English (US){{ /option }}
  {{ option value="es" }}Español{{ /option }}
  {{ option value="pt" }}Português{{ /option }}
  {{ option value="ja" }}日本語{{ /option }}
{{ /dropdown }}

---

### 6.7 Collapsible Accordions (FAQ Cards)
Theme-aware native details/summary collapsible accordions featuring clean expand icons and interactive reveals.

**Shortcode Syntax:**
```markdown
{{{{!}}accordion}}
  {{{{!}}accordion-item title="Is Tamarind compatible with Astro?"}}
    Yes. Because Tamarind compiles down to fully hydrated, static HTML...
  {{{{!}}/accordion-item}}
  {{{{!}}accordion-item title="Can I customize theme variables?"}}
    Absolutely. Set variable DNA keys under the `:root` pseudo-class...
  {{{{!}}/accordion-item}}
{{{{!}}/accordion}}
```

**Live Output:**
{{ accordion }}
  {{ accordion-item title="Is Tamarind compatible with Astro?" }}
    Yes. Because Tamarind compiles down to fully hydrated, static HTML, you can deploy them directly to self-hosted static environments.
  {{ /accordion-item }}
  {{ accordion-item title="Can I customize theme variables?" }}
    Absolutely. Set variable DNA keys under the `:root` pseudo-class in your style sheet to dynamically rebrand all premium components.
  {{ /accordion-item }}
{{ /accordion }}

