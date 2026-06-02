# Tamarind Theme Specification (Design System)

This document defines the strictly typed "Interface" that **all Tamarind Themes** must implement. It has been derived from analyzing the commonalities across `bird`, `pastel`, `neon`, `gram`, `midnight`, `basic`, and `blue`.

## 0. Philosophy: Style Ownership
Themes are the **sole owners** of the visual presentation. 
*   **No Inline Styles**: Content (Markdown) should never require inline `style="..."` attributes to look correct.
*   **Class-Driven**: Visuals are applied via standard utility classes (e.g., `.card`, `.btn`).
*   **Responsiveness**: The theme handles all responsive behavior. A `<div class="card">` must automatically look good on mobile and desktop without user intervention.

---

## 1. CSS Variable Tokens (The "DNA")
Every theme already shares a robust set of variables. This spec formalizes them.

### Colors
```css
:root {
    /* Branding */
    --primary-color:    #...; /* Main interaction color (Buttons, Links) */
    --secondary-color:  #...; /* Subtle accents or ghosts */
    --accent-color:     #...; /* Highlight/Brand color (distinct from primary if needed) */

    /* Backgrounds */
    --background-color: #...; /* Main page background */
    --card-bg:          #...; /* Component/Article card background */
    --header-bg:        #...; /* Sticky header background */
    
    /* Text */
    --text-color:       #...; /* Main body text */
    --text-secondary:   #...; /* Metadata, captions, footers (lighter) */
    --heading-color:    #...; /* Distinct color for H1-H6 */

    /* Borders */
    --border-color:     #...; /* Separators + Inputs */
    
    /* Advanced (Optional) */
    --glass-bg:         rgba(...); /* For blurring/translucency */
    --nav-hover:        rgba(...); /* For interaction states */
}
```

```yaml requirements
category: variables-colors
items:
  - name: "--primary-color"
    type: css-variable
    scope: ":root"
    required: true
    description: "Main interaction color (Buttons, Links)"
  - name: "--secondary-color"
    type: css-variable
    scope: ":root"
    required: false
    description: "Subtle accents or ghosts"
  - name: "--accent-color"
    type: css-variable
    scope: ":root"
    required: false
    description: "Highlight/Brand color"
  - name: "--background-color"
    type: css-variable
    scope: ":root"
    required: true
    description: "Main page background"
  - name: "--card-bg"
    type: css-variable
    scope: ":root"
    required: true
    description: "Component/Article card background"
  - name: "--header-bg"
    type: css-variable
    scope: ":root"
    required: true
    description: "Sticky header background"
  - name: "--text-color"
    type: css-variable
    scope: ":root"
    required: true
    description: "Main body text color"
  - name: "--text-secondary"
    type: css-variable
    scope: ":root"
    required: true
    description: "Metadata, captions, footers color"
  - name: "--heading-color"
    type: css-variable
    scope: ":root"
    required: false
    description: "Distinct color for H1-H6"
  - name: "--border-color"
    type: css-variable
    scope: ":root"
    required: true
    description: "Separators and input borders"
```

### Typography
```css
:root {
    --font-heading:     '...', sans-serif; /* H1-H6, Buttons */
    --font-body:        '...', sans-serif; /* P, Li, Inputs */
}
```

```yaml requirements
category: variables-typography
items:
  - name: "--font-heading"
    type: css-variable
    scope: ":root"
    required: true
    description: "Heading font family (H1-H6, Buttons)"
  - name: "--font-body"
    type: css-variable
    scope: ":root"
    required: true
    description: "Body font family (P, Li, Inputs)"
```

### Spacing & Architecture
```css
:root {
    /* Layout */
    --max-width:        ...px; /* Main container width (layout-container) */
    --header-height:    ...px; /* Sticky header height */

    /* Shapes */
    --radius-sm:        ...px;
    --radius-md:        ...px;
    --radius-lg:        ...px;
    
    /* Depth */
    --shadow:           ...;
}
```

```yaml requirements
category: variables-spacing
items:
  - name: "--max-width"
    type: css-variable
    scope: ":root"
    required: true
    description: "Main container width"
  - name: "--header-height"
    type: css-variable
    scope: ":root"
    required: true
    description: "Sticky header height"
  - name: "--radius-sm"
    type: css-variable
    scope: ":root"
    required: true
    description: "Small border radius"
  - name: "--radius-md"
    type: css-variable
    scope: ":root"
    required: true
    description: "Medium border radius"
  - name: "--radius-lg"
    type: css-variable
    scope: ":root"
    required: true
    description: "Large border radius"
  - name: "--shadow"
    type: css-variable
    scope: ":root"
    required: true
    description: "Default box shadow"
```

---

## 2. Component Contract (The "Atoms")
Every theme MUST provide these utility classes. `bird` already has some, others are missing them.

### Buttons & Actions
**Standard**: `<a href="#" class="btn btn-primary">Action</a>`
*   `.btn`: Inline-block, transition, cursor-pointer, font-weight: bold, text-decoration: none.
*   `.btn-primary`: Background = `--primary-color`, Text = Contrast.
*   `.btn-secondary`: Border = `--border-color`, Text = `--text-color`.
*   `.btn-ghost`: Transparent, Text = `--text-secondary`.
*   `.btn-sm`: Compact padding/font.
*   `.button-group`: Wrapper for multiple buttons (flex, gap).

```yaml requirements
category: components-buttons
items:
  - name: ".btn"
    type: css-selector
    scope: global
    required: true
    description: "Button base class"
  - name: ".btn-primary"
    type: css-selector
    scope: global
    required: true
    description: "Primary action button"
  - name: ".btn-secondary"
    type: css-selector
    scope: global
    required: true
    description: "Secondary action button"
  - name: ".btn-ghost"
    type: css-selector
    scope: global
    required: true
    description: "Ghost/transparent button"
  - name: ".btn-sm"
    type: css-selector
    scope: global
    required: true
    description: "Small button variant"
  - name: ".button-group"
    type: css-selector
    scope: global
    required: true
    description: "Wrapper for multiple buttons"
```

### Forms (New Requirement)
**Standard**: `<div class="form-group"><label class="form-label">Email</label><input class="form-input"></div>`
*   `.form-group`: Margin bottom spacing.
*   `.form-label`: Block, weight bold, margin bottom.
*   `.form-input`, `.form-textarea`, `.form-select`:
    *   Width: 100%.
    *   Padding: ~10-12px.
    *   Border: 1px solid `--border-color`.
    *   Background: Transparent or surface color.
    *   Radius: Theme's radius token.
*   `.form-checkbox`, `.form-radio`:
    *   Accent color: `--primary-color`.
    *   Margin-right: Small spacing.
*   `.form-file`:
    *   Standard padding/border.
    *   Cursor pointer.

```yaml requirements
category: components-forms
items:
  - name: ".form-group"
    type: css-selector
    scope: global
    required: true
    description: "Form group container"
  - name: ".form-label"
    type: css-selector
    scope: global
    required: true
    description: "Form label"
  - name: ".form-input"
    type: css-selector
    scope: global
    required: true
    description: "Text input field"
  - name: ".form-textarea"
    type: css-selector
    scope: global
    required: true
    description: "Textarea field"
  - name: ".form-select"
    type: css-selector
    scope: global
    required: true
    description: "Select dropdown"
  - name: ".form-checkbox"
    type: css-selector
    scope: global
    required: true
    description: "Checkbox input"
  - name: ".form-radio"
    type: css-selector
    scope: global
    required: true
    description: "Radio input"
  - name: ".form-file"
    type: css-selector
    scope: global
    required: true
    description: "File input"
```

### Cards & Surfaces
**Standard**: `<div class="card">Content</div>`
*   `.card`: Background `--card-bg`, Border (optional), Radius `--radius-lg`, Shadow `--shadow`.
*   `.card-padding`: Standard padding (e.g., 24px or 40px).

```yaml requirements
category: components-cards
items:
  - name: ".card"
    type: css-selector
    scope: global
    required: true
    description: "Card surface component"
  - name: ".card-padding"
    type: css-selector
    scope: global
    required: true
    description: "Card padding utility"
```

### Tags & Badges
**Standard**: `<span class="badge">New</span>`
*   `.badge` (or `.tag-link` / `.post-tag` unified): Inline-flex, small caps, padding, radius.

```yaml requirements
category: components-badges
items:
  - name: ".badge"
    type: css-selector
    scope: global
    required: true
    description: "Badge/tag label"
  - name: ".badge-primary"
    type: css-selector
    scope: global
    required: true
    description: "Primary badge variant"
  - name: ".badge-secondary"
    type: css-selector
    scope: global
    required: true
    description: "Secondary badge variant"
  - name: ".badge-success"
    type: css-selector
    scope: global
    required: true
    description: "Success badge variant"
  - name: ".badge-warning"
    type: css-selector
    scope: global
    required: true
    description: "Warning badge variant"
  - name: ".badge-danger"
    type: css-selector
    scope: global
    required: true
    description: "Danger badge variant"
```

### Alerts & Callouts
**Standard**: `<div class="callout callout-info alert alert-info"><div class="callout-title alert-title">Title</div><div class="callout-content alert-content">Content</div></div>`
*   `.alert`: Container class for user notifications.
*   `.alert-title`: Title font weight bold/margin bottom.
*   `.alert-content`: Inner text wrapper.

```yaml requirements
category: components-alerts
items:
  - name: ".alert"
    type: css-selector
    scope: global
    required: true
    description: "Alert wrapper container"
  - name: ".alert-title"
    type: css-selector
    scope: global
    required: true
    description: "Alert title element"
  - name: ".alert-content"
    type: css-selector
    scope: global
    required: true
    description: "Alert body content"
  - name: ".alert-info"
    type: css-selector
    scope: global
    required: true
    description: "Info alert variant styling"
  - name: ".alert-warn"
    type: css-selector
    scope: global
    required: true
    description: "Warning alert variant styling"
  - name: ".alert-error"
    type: css-selector
    scope: global
    required: true
    description: "Error alert variant styling"
  - name: ".alert-tip"
    type: css-selector
    scope: global
    required: true
    description: "Tip alert variant styling"
```

### Callouts (Already Present in All Themes)
*   `.callout`: Container with border-left.
*   `.callout-info` / `.callout-warn` / `.callout-error` / `.callout-tip`.

```yaml requirements
category: components-callouts
items:
  - name: ".callout"
    type: css-selector
    scope: global
    required: true
    description: "Callout container"
  - name: ".callout-info"
    type: css-selector
    scope: global
    required: true
    description: "Info callout variant"
  - name: ".callout-warn"
    type: css-selector
    scope: global
    required: true
    description: "Warning callout variant"
  - name: ".callout-error"
    type: css-selector
    scope: global
    required: true
    description: "Error callout variant"
  - name: ".callout-tip"
    type: css-selector
    scope: global
    required: true
    description: "Tip callout variant"
```

---

## 3. Layout Contract
*   **.layout-container** (or `.container`): Centered wrapper respecting `--max-width`.
*   **.site-header**: Sticky, contains `.nav-container`.
*   **.nav-links**: Flex container for menu items.
*   **.site-footer**: Centered, text-secondary.

```yaml requirements
category: layout
items:
  - name: ".layout-container"
    type: css-selector
    scope: global
    required: true
    description: "Main content container"
  - name: ".site-header"
    type: css-selector
    scope: global
    required: true
    description: "Sticky site header"
  - name: ".nav-container"
    type: css-selector
    scope: global
    required: true
    description: "Navigation wrapper inside header"
  - name: ".nav-links"
    type: css-selector
    scope: global
    required: true
    description: "Navigation links container"
  - name: ".site-footer"
    type: css-selector
    scope: global
    required: true
    description: "Site footer"
  - name: ".tamarind-ghost-badge"
    type: css-selector
    scope: global
    required: true
    description: "Floating Use Tamarind badge element"
  - name: ".footer-promo"
    type: css-selector
    scope: global
    required: true
    description: "Website promotional reference in the footer"
```

### Mobile Behavior (Refactor Implemented)
*   Tablets/Phones (<768px):
    *   `.layout-container`: Padding 0-20px.
    *   Grids: Collapse to 1 column.
    *   Headers: Reduce font sizes (clamp).
    *   Nav: Stack vertically or scroll horizontally (`bird`).

```yaml requirements
category: responsive
items:
  - name: "max-width: 768px"
    type: media-query
    scope: "@media"
    required: true
    description: "Mobile breakpoint at 768px"
```

---

## 4. Shortcode Support (Mandatory)
All themes MUST support:
1.  **Videos**: `.video-container` (16:9 responsive).
2.  **Code**: `pre` (scrollable), `code` (inline style).
3.  **Figures**: `figure`, `figcaption` centered.
4.  **Mermaid**: `.mermaid` classes (text visibility fixes).
5.  **Terminal**: `.terminal` window frame.

```yaml requirements
category: shortcodes
items:
  - name: ".video-container"
    type: css-selector
    scope: global
    required: true
    description: "Responsive video embed container (16:9)"
  - name: ".mermaid"
    type: css-selector
    scope: global
    required: true
    description: "Mermaid diagram container"
  - name: ".terminal"
    type: css-selector
    scope: global
    required: true
    description: "Terminal window frame"
  - name: "img"
    type: css-selector
    scope: global
    required: true
    description: "Responsive image styling to prevent overflow"
  - name: "figure"
    type: css-selector
    scope: global
    required: true
    description: "Responsive figure block formatting"
```


---

## 5. Accessibility & Mobile Navigation (The "Responsive Contract")
All themes must ensure that primary navigation elements are never hidden on mobile viewports.

```yaml requirements
category: accessibility
items:
  - name: "mobile-nav-visible"
    type: responsive-nav
    required: true
    description: "The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)"
  - name: "mobile-sidebar-visible"
    type: responsive-nav
    required: true
    description: "The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)"
```

---

## 6. Article Content Elements (The "Typographic Contract")
All themes must explicitly style long-form article elements, either globally or scoped under `.article-content`.

```yaml requirements
category: article-content
items:
  - name: "ul"
    type: article-markup
    required: true
    description: "Unordered lists (ul) must be explicitly styled"
  - name: "ol"
    type: article-markup
    required: true
    description: "Ordered lists (ol) must be explicitly styled"
  - name: "li"
    type: article-markup
    required: true
    description: "List items (li) must be explicitly styled"
  - name: "blockquote"
    type: article-markup
    required: true
    description: "Blockquotes must be explicitly styled"
  - name: "code"
    type: article-markup
    required: true
    description: "Inline code blocks must be explicitly styled"
  - name: "pre"
    type: article-markup
    required: true
    description: "Preformatted code blocks (pre) must be explicitly styled"
  - name: "table"
    type: article-markup
    required: true
    description: "Data tables must be explicitly styled"
```

---

## 7. Brand Compliance & Menu Deduplication (The "Brand Contract")
All themes' page and article templates must include core promotional branding and must not duplicate menu loops.

```yaml requirements
category: templates
items:
  - name: "articles.mdt:tamarind-ghost-badge"
    type: template-feature
    required: true
    description: "articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'"
  - name: "articles.mdt:footer-promo"
    type: template-feature
    required: true
    description: "articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'"
  - name: "page.mdt:tamarind-ghost-badge"
    type: template-feature
    required: true
    description: "page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'"
  - name: "page.mdt:footer-promo"
    type: template-feature
    required: true
    description: "page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'"
  - name: "menu-no-duplication"
    type: template-feature
    required: true
    description: "Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)"
```

---

## 9. Premium UI Components (The "Premium Components Contract")
All themes must support the 6 premium layout components built into the Tamarind static engine:

```yaml requirements
category: premium-components
items:
  # 1. Stats & Metrics Grid
  - name: ".metrics-grid"
    type: css-selector
    scope: global
    required: true
    description: "Container for the metrics grid layout"
  - name: ".metric-card"
    type: css-selector
    scope: global
    required: true
    description: "Individual metric block container"
  - name: ".metric-value"
    type: css-selector
    scope: global
    required: true
    description: "Stylized numeric or main text value inside a metric card"
  - name: ".metric-label"
    type: css-selector
    scope: global
    required: true
    description: "Stylized supporting label underneath the metric value"

  # 2. Gradient Icon Feature Cards
  - name: ".features-grid"
    type: css-selector
    scope: global
    required: true
    description: "Container for the features card grid"
  - name: ".feature-card"
    type: css-selector
    scope: global
    required: true
    description: "Individual feature card block container"
  - name: ".feature-icon-box"
    type: css-selector
    scope: global
    required: true
    description: "Vivid diagonal gradient icon wrapper box"
  - name: ".feature-title"
    type: css-selector
    scope: global
    required: true
    description: "Feature headline font weight and styling"
  - name: ".feature-desc"
    type: css-selector
    scope: global
    required: true
    description: "Feature supporting detail text color and alignment"

  # 3. Capabilities Checklist Card
  - name: ".capabilities-grid"
    type: css-selector
    scope: global
    required: true
    description: "Multi-column layout wrapper for checklist cards"
  - name: ".capability-card"
    type: css-selector
    scope: global
    required: true
    description: "Checklist card wrapper container"
  - name: ".capability-header"
    type: css-selector
    scope: global
    required: true
    description: "Checklist header separator styling"
  - name: ".capability-card-title"
    type: css-selector
    scope: global
    required: true
    description: "Checklist card heading typography styling"
  - name: ".capability-row"
    type: css-selector
    scope: global
    required: true
    description: "Flex row grouping detail columns and status label"
  - name: ".capability-name"
    type: css-selector
    scope: global
    required: true
    description: "Checkbox items title styling"
  - name: ".capability-desc"
    type: css-selector
    scope: global
    required: true
    description: "Checkbox description detail styling"
  - name: ".capability-status"
    type: css-selector
    scope: global
    required: true
    description: "Status indicator badge container"

  # 4. Vertical Quick Start Timeline
  - name: ".timeline-container"
    type: css-selector
    scope: global
    required: true
    description: "Timeline absolute tracking wrapper container"
  - name: ".timeline-item"
    type: css-selector
    scope: global
    required: true
    description: "Relative timeline segment node"
  - name: ".timeline-badge"
    type: css-selector
    scope: global
    required: true
    description: "Circular step sequence bullet container"
  - name: ".timeline-badge-number"
    type: css-selector
    scope: global
    required: true
    description: "Bold numeric indicator text inside timeline badge"
  - name: ".timeline-content"
    type: css-selector
    scope: global
    required: true
    description: "Segment description card container styling"
  - name: ".timeline-title"
    type: css-selector
    scope: global
    required: true
    description: "Segment headline typography styling"
  - name: ".timeline-desc"
    type: css-selector
    scope: global
    required: true
    description: "Segment description detail styling"

  # 5. Semantic Alert Containers
  - name: ".alert-container"
    type: css-selector
    scope: global
    required: true
    description: "Flex alert wrapper banner"
  - name: ".alert-icon-box"
    type: css-selector
    scope: global
    required: true
    description: "SVG symbol flex align wrapper box"
  - name: ".alert-content"
    type: css-selector
    scope: global
    required: true
    description: "Text content block column layout wrapper"
  - name: ".alert-title"
    type: css-selector
    scope: global
    required: true
    description: "Semantic alert heading typography"
  - name: ".alert-message"
    type: css-selector
    scope: global
    required: true
    description: "Supporting details alert body description text"

  # 6. Custom Interactive Dropdowns
  - name: ".tamarind-select-wrapper"
    type: css-selector
    scope: global
    required: true
    description: "Outer wrapper positioning the standard selector controls"
  - name: ".tamarind-select"
    type: css-selector
    scope: global
    required: true
    description: "Stylized select reset button element overrides"
  - name: ".tamarind-select-chevron"
    type: css-selector
    scope: global
    required: true
    description: "Chevron arrow vector flex alignment container"

  # 7. Collapsible Accordions (FAQ Details)
  - name: ".tamarind-accordion"
    type: css-selector
    scope: global
    required: true
    description: "Collapsible HTML5 details accordion card block"
  - name: ".tamarind-accordion-summary"
    type: css-selector
    scope: global
    required: true
    description: "Interactive summary header text inside details card"
  - name: ".tamarind-accordion-content"
    type: css-selector
    scope: global
    required: true
    description: "Expanded text content inside details accordion block"
```

---

## 10. Canvas Mode Layouts (The "Canvas Mode Contract")
All themes must support the Canvas Mode layouts provided by the Tamarind engine to render naked landing pages:

```yaml requirements
category: components-canvas
items:
  - name: ".canvas-mode-active"
    type: css-selector
    scope: global
    required: true
    description: "Applied to layout wrappers when canvas landing mode is active to reset shadows, borders, and margins"
  - name: ".canvas-width-limit"
    type: css-selector
    scope: global
    required: true
    description: "Sizing constraints restricting wide-bleed canvas sections to a modern 1200px max-width"
```

---

## 11. Left Sidebar & Responsive Containment (The "Sidebar Contract")
All themes must support the contextual left sidebar navigation layout and ensure strict mobile viewport boundary containment (no horizontal scrolling).

```yaml requirements
category: layout-sidebar
items:
  - name: ".layout-has-sidebar"
    type: css-selector
    scope: global
    required: true
    description: "Grid/Flex wrapper when sidebar layout is active (must establish sidebar columns on desktop and collapse on mobile)"
  - name: ".context-sidebar"
    type: css-selector
    scope: global
    required: true
    description: "Left sidebar container hosting section navigation sibling links (.sidebar-left or .context-sidebar)"
  - name: ".tamarind-sidebar-checkbox"
    type: css-selector
    scope: global
    required: true
    description: "Hidden input control (#tamarind-sidebar-toggle or .tamarind-sidebar-checkbox) managing mobile drawer display state"
  - name: ".tamarind-sidebar-handle"
    type: css-selector
    scope: global
    required: true
    description: "Mobile toggle menu trigger button handle (high visibility, absolute alignment against header background)"
  - name: ".tamarind-sidebar-backdrop"
    type: css-selector
    scope: global
    required: true
    description: "Overlay backdrop covering the viewport when the mobile sidebar drawer is active to prevent under-scrolling"
```

---

## 8. Bootstrapping / Updating Guide

### How to Create a New Theme
1.  **Duplicate**: Copy an existing theme folder (e.g., `cp -r basic my-new-theme`).
    *   *Tip*: Use `basic` for editorial sites, `gram` for grids/visuals, or `bird` for app-like layouts.
2.  **Define Variables**: Open `my-new-theme/style.css` and set your `Cortex` (Colors, Fonts, Radius).
    *   Ensure all standard variables from Section 1 are present.
3.  **Implement Components**: Copy/Paste the "Component Contract" CSS (Buttons, Forms, Badges) and tweak values (border-radius, shadows) to match your vibe.
4.  **Verify Templates**: Check `page.mdt`, `articles.mdt`, `footer.mdt`.
    *   Ensure they use the standard classes (`.site-header`, `.layout-container`, `.card`).
5.  **Run Verifier**: Execute `go test ./internal/builder/ -run TestAllThemesCompliance -v` to validate compliance.

### How to Update an Existing Theme
1.  **Audit**: Check if `style.css` contains the mandatory **Button** and **Form** classes.
2.  **Enforce Tokens**: Ensure all standard variables from Section 1 (e.g., `--font-body`, `--header-bg`) are defined. Remove any legacy or ad-hoc variables.
3.  **Normalize**: Standardize margins/padding using the layout contract to ensure `fix-mobile-rendering` improvements persist.
4.  **Verify**: Run the theme verifier test to confirm compliance.
