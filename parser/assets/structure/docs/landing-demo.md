---
title: "Tamarind Pro: Full-Bleed Canvas Showcase"
subtitle: "Build premium, high-converting, distraction-free landing pages in pure static markdown"
canvas: true
---

<!-- Hero Section (Naked Layout Showcase) -->
<div style="text-align: center; margin: 4rem auto 2rem auto; max-width: 800px;">
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
    {{ button href="style-guide.html" }}Explore Style Guide{{ /button }}
  </div>
</div>

<!-- Sales Letter Intro Paragraphs -->
<div style="margin: 3rem auto; max-width: 800px; line-height: 1.8; font-size: 1.1rem; color: var(--text-color);">
  <p style="margin-bottom: 1.5rem;">
    Every high-performing sales letter and landing page relies on clean, high-contrast readability. When standard site navigation is stripped away, the reader is guided down a focused conversion path. Tamarind's native Canvas Mode combines raw markdown writing speed with a premium design system that morphs to fit any aesthetic—whether it's Apple's clean minimalist white, modern SaaS slate blue, or a high-contrast terminal console screen.
  </p>
  <p>
    By placing interactive elements, feature grids, and vertical quick starts directly alongside standard paragraphs of copy, you can construct authoritative marketing sites, product launches, or book announcements in under a minute.
  </p>
</div>

---

<!-- 1. Stats & Metrics Grid (Authority Booster) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Key Platform Performance Metrics</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Highlighting real data points and platform speeds establishes immediate trust with new readers.</p>
</div>

<div style="max-width: 800px; margin: 2rem auto; line-height: 1.8; font-size: 1.1rem; color: var(--text-color);">
  <p style="margin-bottom: 1.5rem;">
    Speed is the single most critical factor for conversion rates. Every additional second of load time reduces conversions by up to 20%. Because Tamarind compiles down to fully hydrated, static HTML with zero external script dependencies, it delivers global load speeds under 50ms, earning perfect scores across every performance index.
  </p>
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

<div style="max-width: 800px; margin: 2rem auto; line-height: 1.8; font-size: 1.1rem; color: var(--text-color);">
  <p style="margin-bottom: 1.5rem;">
    Whether you are selling a SaaS subscription, launching a developer tool, or distributing a digital book, readers deserve a transparent feature breakdown. The capability comparison grid highlights key strengths, platform dependencies, and pending features, formatted as a cozy, responsive, and legible checklist card.
  </p>
</div>

{{ capabilities-grid }}
  {{ capabilities title="Core Compiler" }}
    {{ capability name="Static Compiler Binary" desc="Ultra-fast single Go executable compiler" status="success" statusLabel="Yes" }}
    {{ capability name="AST Validation" desc="Validates themes against layout specifications" status="success" statusLabel="Yes" }}
    {{ capability name="Zero-Config CSS" desc="Automatic CSS token harvesting" status="success" statusLabel="Yes" }}
  {{ /capabilities }}

  {{ capabilities title="Platform & Deploy" }}
    {{ capability name="Flat-File CDN" desc="Direct compilation to static HTML/CSS files" status="success" statusLabel="Yes" }}
    {{ capability name="Docker Integration" desc="Secure, hardened container standard" status="success" statusLabel="Yes" }}
    {{ capability name="Local Dev Server" desc="Instant serving with fast reload times" status="success" statusLabel="Yes" }}
  {{ /capabilities }}

  {{ capabilities title="Advanced Features" }}
    {{ capability name="Image Pipeline" desc="Multi-viewport responsive asset generation" status="warning" statusLabel="Partial" }}
    {{ capability name="SEO Validator" desc="Local semantic metadata audit reports" status="success" statusLabel="Yes" }}
    {{ capability name="Telemetry Registry" desc="Opt-in quantitative performance tracing" status="pending" statusLabel="Planned" }}
  {{ /capabilities }}

  {{ capabilities title="Design System" }}
    {{ capability name="16+ Core Themes" desc="Premium, highly aesthetic responsive layouts" status="success" statusLabel="Yes" }}
    {{ capability name="Interactive Widgets" desc="Accordion, carousel, tabs, switcher components" status="success" statusLabel="Yes" }}
    {{ capability name="Layout Constraints" desc="Compiler-enforced alignment safeguards" status="success" statusLabel="Yes" }}
  {{ /capabilities }}
{{ /capabilities-grid }}

---

<!-- 4. Vertical Quick Start Timeline (Onboarding Path) -->
<div id="quick-start" style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Get Live in Three Steps</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Our mathematically aligned vertical timeline maps a tight quick start flow with no overlaps.</p>
</div>

<div style="max-width: 800px; margin: 2rem auto; line-height: 1.8; font-size: 1.1rem; color: var(--text-color);">
  <p style="margin-bottom: 1.5rem;">
    Setting up a brand-new publishing platform shouldn't require complex configuration files or setting up heavy node package structures. The three-step vertical progress timeline illustrates how easy it is to initialize a sandbox, configure design variables, and generate a fully optimized static directory in seconds.
  </p>
</div>

{{ timeline }}
  {{ item title="Initialize the Workspace" number="1" }}
    Run our zero-config setup command to extract default styles, configurations, and showcase documents:
    {{ terminal }}
      {{ tab title="macOS" }}
      $ tamarind init
      {{ /tab }}
      {{ tab title="Linux" }}
      $ tamarind init
      {{ /tab }}
      {{ tab title="Windows" }}
      PS> .\tamarind.exe init
      {{ /tab }}
    {{ /terminal }}
  {{ /item }}
  {{ item title="Toggle the Canvas Flag" number="2" }}
    Enable the landing page layout by writing `canvas: true` directly in your YAML front matter.
  {{ /item }}
  {{ item title="Compile and Publish" number="3" }}
    Run a lightning-fast build using any of our premium themes and deploy the naked folder to any host:
    {{ terminal }}
      {{ tab title="macOS" }}
      $ tamarind build -theme blue
      {{ /tab }}
      {{ tab title="Linux" }}
      $ tamarind build -theme blue
      {{ /tab }}
      {{ tab title="Windows" }}
      PS> .\tamarind.exe build -theme blue
      {{ /tab }}
    {{ /terminal }}
  {{ /item }}
{{ /timeline }}

---

<!-- 4.1 Interactive Tabs Showcase -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Interactive Tabbed Layout</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Leverage the general tabs shortcode to structure tabbed content tables or multi-language code snippets.</p>
</div>

<div style="max-width: 800px; margin: 0 auto 3rem auto;">
  {{ tabs }}
    {{ tab title="npm" }}
    Install the Tamarind developer CLI tool globally using npm:
    ```bash
    $ npm install -g @tamarind/cli
    ```
    {{ /tab }}
    {{ tab title="yarn" }}
    Install the Tamarind developer CLI tool globally using yarn:
    ```bash
    $ yarn global add @tamarind/cli
    ```
    {{ /tab }}
    {{ tab title="pnpm" }}
    Install the Tamarind developer CLI tool globally using pnpm:
    ```bash
    $ pnpm add -g @tamarind/cli
    ```
    {{ /tab }}
  {{ /tabs }}
</div>

---

<!-- 5. Interactive Plan Selector (Combo Box) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Interactive License Selector</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Test our customized, responsive select menu styling designed to blend perfectly with theme colors.</p>
</div>

<div style="max-width: 600px; margin: 0 auto 3rem auto;">
  {{ dropdown id="plan-selector" label="Select a Tamarind License Plan:" }}
    {{ option value="free" }}Tamarind Community Edition (Open Source, Free){{ /option }}
    {{ option value="pro" selected="true" }}Tamarind Pro Canvas License ($49/one-time){{ /option }}
    {{ option value="enterprise" }}Tamarind Enterprise Bundle (Unlimited Sites){{ /option }}
  {{ /dropdown }}
</div>

---

<!-- 5.1 High-Impact Pricing Grid -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Simple, Transparent Pricing</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Choose the tier that matches your deployment scope. Cancel or switch plans at any time.</p>
</div>

{{ pricing monthly_label="Monthly Billing" annual_label="Annual Billing" discount="Save 20%" }}
  {{ plan title="Personal" price_monthly="0" price_annual="0" period_monthly="Free forever" period_annual="Free forever" button="Get Started" }}
    - 1 Project Site
    - Basic Templates
    - Community Support
  {{ /plan }}
  {{ plan title="Developer" price_monthly="19" price_annual="15" period_monthly="per month" period_annual="billed annually" featured="true" badge="Popular" button="Start Pro Trial" }}
    - Unlimited Sites
    - All 31 Pro Themes
    - Contextual Sidebar
    - Priority Email Support
  {{ /plan }}
  {{ plan title="Enterprise" price_monthly="99" price_annual="79" period_monthly="per month" period_annual="billed annually" button="Contact Sales" }}
    - Multi-Seat Licensing
    - WASM Custom Plugins
    - Air-Gapped Deployment
    - 24/7 Phone SLA
  {{ /plan }}
{{ /pricing }}

---

<!-- 6. Frequently Asked Questions (Native Accordion Details) -->
<div style="text-align: center; margin: 4rem 0 1rem 0;">
  <h2 style="font-size: 2rem; font-weight: 700; margin-bottom: 0.5rem; font-family: var(--font-heading, inherit);">Frequently Asked Questions</h2>
  <p style="color: var(--text-secondary, #64748b); max-width: 600px; margin: 0 auto 2rem auto;">Native HTML5 collapsible details cards that look stunning across all premium templates.</p>
</div>

<div style="max-width: 800px; margin: 0 auto 4rem auto;">
  {{ accordion }}
    {{ accordion-item title="Does Canvas Mode support custom styling overrides?" }}
      Yes! Standard markdown supports raw HTML/CSS injection, and any custom CSS can be passed in the `custom_css` metadata attribute to override your theme palette directly.
    {{ /accordion-item }}

    {{ accordion-item title="Can I combine Canvas Mode pages with normal pages?" }}
      Absolutely. You can have 10 standard pages displaying site chrome (menus/footers) alongside 2 naked Canvas landing pages in the same compiler repository.
    {{ /accordion-item }}

    {{ accordion-item title="Do the buttons support active form actions?" }}
      Yes, our standard buttons can be placed inside custom HTML form grids to trigger standard POST actions, mail subscriptions, or sales integrations.
    {{ /accordion-item }}
  {{ /accordion }}
</div>

---

<!-- Call to Action Info Notification -->
{{ alert type="tip" title="Deployment Tip" }}
Static landing pages are incredibly cheap and performant to host. Deploying Tamarind's compiled single-page output directly to Netlify, Vercel, or GitHub Pages guarantees response times under **50ms** anywhere globally.
{{ /alert }}

{{ social_ribbon }}
  {{ testimonial stars="5" avatar="../images/avatar_alex.png" author="Alex" handle="@alex_dev" }}
    Setting up Tamarind took less than two minutes. The compiled static site loads in under 30ms globally, and my Lighthouse score is a perfect 100. Absolutely incredible tool for developers.
  {{ /testimonial }}
  {{ testimonial stars="5" avatar="../images/avatar_sarah.png" author="Sarah" handle="@sarah_creator" }}
    As a designer, I love how easily Tamarind handles custom CSS tokens. I don't have to fight a complex framework; it just compiles my layouts into gorgeous, lightweight pages.
  {{ /testimonial }}
  {{ testimonial stars="5" avatar="../images/avatar_marcus.png" author="Marcus" handle="@marcus_biz" }}
    Implementing the passwordless paywall flow eliminated all customer account friction. We saw a 35% bump in checkouts from day one without writing a single line of database code.
  {{ /testimonial }}
  {{ testimonial stars="5" avatar="../images/avatar_elena.png" author="Elena" handle="@elena_design" }}
    The built-in SEO analyzer and automated Open Graph card generator saved us weeks of marketing setup. Our articles are indexing and ranking faster than ever.
  {{ /testimonial }}
{{ /social_ribbon }}

<div style="text-align: center; margin: 4rem auto 2rem auto; max-width: 600px;">
  <h3 style="font-size: 1.75rem; font-weight: 700; margin-bottom: 1rem; font-family: var(--font-heading, inherit);">Ready to launch your project?</h3>
  {{ button href="#" type="primary" }}Get Started for Free{{ /button }}
</div>
