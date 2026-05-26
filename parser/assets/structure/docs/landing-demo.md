---
title: "Tamarind Pro: Full-Bleed Canvas Showcase"
subtitle: "Build premium, high-converting, distraction-free landing pages in pure static markdown"
canvas: true
---

<!-- Hero Section (Naked Layout Showcase) -->
<div style="text-align: center; margin: 4rem auto 3rem auto; max-width: 800px;">
  <div style="display: inline-block; background: linear-gradient(135deg, var(--primary-color, #2563eb), #a855f7); color: #ffffff; padding: 6px 14px; border-radius: 20px; font-size: 13px; font-weight: 600; text-transform: uppercase; letter-spacing: 1px; margin-bottom: 1.5rem;">
    Announcing Canvas Mode
  </div>
  <h1 style="font-size: 3.25rem; font-weight: 800; line-height: 1.15; margin: 0 0 1rem 0; letter-spacing: -1.5px; font-family: var(--font-heading, inherit);">
    Create Gorgeous Landing Pages <br><span style="background: linear-gradient(to right, var(--primary-color, #2563eb), #22c55e); -webkit-background-clip: text; -webkit-text-fill-color: transparent;">Without Distractions</span>
  </h1>
  <p style="font-size: 1.25rem; color: var(--text-secondary, #64748b); line-height: 1.6; margin-bottom: 2rem;">
    Zero distractions, zero menus, and zero complex JS bundles. Write standard markdown and let Tamarind's theme-aware layout compile a premium, air-gapped static landing page out-of-the-box.
  </p>
  <div style="display: flex; gap: 16px; justify-content: center; align-items: center; flex-wrap: wrap;">
    {{ button href="#quick-start" type="primary" }}Deploy in 5 Seconds{{ /button }}
    {{ button href="docs/style-guide.html" }}Explore Style Guide{{ /button }}
  </div>
</div>

---

<!-- 1. Stats & Metrics Grid (Authority Booster) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Key Platform Performance Metrics</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Highlighting real data points and platform speeds establishes immediate trust with new readers.</p>
</div>

{{ metrics }}
  {{ metric value="100%" label="Open Source" }}
  {{ metric value="< 5s" label="Onboarding Speed" }}
  {{ metric value="32" label="Premium Themes" }}
  {{ metric value="99.9" label="Lighthouse Score" }}
{{ /metrics }}

---

<!-- 2. Gradient Icon Feature Cards (Key Benefits) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Premium Features, Zero Effort</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Leverage responsive, beautiful diagonal-gradient cards with built-in vector icons.</p>
</div>

{{ features }}
  {{ feature title="Naked Canvas Mode" gradient="blue-cyan" icon="sparkles" }}
    Strip away standard navigation menus and footers with a single front-matter tag, making it extremely easy to build lead magnets and sales funnels.
  {{ /feature }}
  {{ feature title="Theme Harmony System" gradient="purple-pink" icon="bolt" }}
    All components naturally morph to adopt your active theme's colors, dark-mode gradients, and fonts. Change themes and watch your landing page dynamically re-brand!
  {{ /feature }}
  {{ feature title="Air-Gapped Privacy" gradient="orange-red" icon="shield" }}
    Tamarind compiles to lightweight, standalone HTML with zero unverified CDN assets, guaranteeing full client-side isolation and security compliance.
  {{ /feature }}
{{ /features}}

---

<!-- 3. Capabilities Checklist Card (Features Matrix) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Full Capabilities Comparison</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Illustrate detailed feature sets, licenses, and variants with distinct visual indicators.</p>
</div>

{{ capabilities title="Feature Matrix Comparison" }}
  {{ capability name="Static Compiler Binary" desc="Ultra-fast single Go executable compiler" check="true" }}
  {{ capability name="Premium Component Library" desc="Metrics, Feature cards, Checklists, and Timelines built-in" check="true" }}
  {{ capability name="Naked Canvas Template Mode" desc="Distraction-free single page layout overrides" check="true" }}
  {{ capability name="Custom CDN Dependecies" desc="Third party analytics scripts or web fonts" check="warn" }}
  {{ capability name="Heavy Single Page JS Runtimes" desc="Bypass heavy client-side hydrate frames" check="false" }}
{{ /capabilities }}

---

<!-- 4. Vertical Quick Start Timeline (Onboarding Path) -->
<div id="quick-start" style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Get Live in Three Steps</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Our mathematically aligned vertical timeline maps a tight quick start flow with no overlaps.</p>
</div>

{{ timeline }}
  {{ timeline-item step="1" title="Initialize the Workspace" }}
    Run our zero-config setup command to extract default styles, configurations, and showcase documents:
    {{ terminal }}
    $ tamarind init
    {{ /terminal }}
  {{ /timeline-item }}
  {{ timeline-item step="2" title="Toggle the Canvas Flag" }}
    Enable the landing page layout by writing `canvas: true` directly in your YAML front matter.
  {{ /timeline-item }}
  {{ timeline-item step="3" title="Compile and Publish" }}
    Run a lightning-fast build using any of our premium themes and deploy the naked folder to any host:
    {{ terminal }}
    $ tamarind build -theme blue
    {{ /terminal }}
  {{ /timeline-item }}
{{ /timeline }}

---

<!-- 5. Interactive FAQ Grid (Dropdown Accordions) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Frequently Asked Questions</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Interactive, clean FAQ accordions with smooth micro-animations.</p>
</div>

{{ dropdown title="Frequently Asked Questions Grid" icon="sparkles" }}
  {{ dropdown-item label="Does Canvas Mode support custom styling overrides?" desc="Yes! Standard markdown supports raw HTML/CSS injection, and any custom CSS can be passed in the `custom_css` metadata attribute to override your theme palette directly." url="#" }}
  {{ dropdown-item label="Can I combine Canvas Mode pages with normal pages?" desc="Absolutely. You can have 10 standard pages displaying site chrome (menus/footers) alongside 2 naked Canvas landing pages in the same compiler repository." url="#" }}
  {{ dropdown-item label="Do the buttons support active form actions?" desc="Yes, our standard buttons can be placed inside custom HTML form grids to trigger standard POST actions, mail subscriptions, or sales integrations." url="#" }}
{{ /dropdown }}

---

<!-- Call to Action Info Notification -->
{{ alert type="tip" title="Deployment Tip" }}
Static landing pages are incredibly cheap and performant to host. Deploying Tamarind's compiled single-page output directly to Netlify, Vercel, or GitHub Pages guarantees standard response times under **50ms** anywhere globally.
{{ /alert }}

<div style="text-align: center; margin: 4rem auto 2rem auto; max-width: 600px;">
  <h3 style="font-size: 1.75rem; font-weight: 700; margin-bottom: 1rem; font-family: var(--font-heading, inherit);">Ready to launch your project?</h3>
  {{ button href="#" type="primary" }}Get Started for Free{{ /button }}
</div>
