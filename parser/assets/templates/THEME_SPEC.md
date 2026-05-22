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
    required: false
    description: "Card surface component"
  - name: ".card-padding"
    type: css-selector
    scope: global
    required: false
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

## 5. Bootstrapping / Updating Guide

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
