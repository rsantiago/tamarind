# Tamarind Theme Specification (Design System)

This document defines the strictly typed "Interface" that **all Tamarind Themes** must implement. It has been derived from analyzing the commonalities across `bird`, `pastel`, `neon`, `gram`, `midnight`, `basic`, and `blue`.

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

### Typography
```css
:root {
    --font-heading:     '...', sans-serif; /* H1-H6, Buttons */
    --font-body:        '...', sans-serif; /* P, Li, Inputs */
}
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

### Cards & Surfaces
**Standard**: `<div class="card">Content</div>`
*   `.card`: Background `--card-bg`, Border (optional), Radius `--radius-lg`, Shadow `--shadow`.
*   `.card-padding`: Standard padding (e.g., 24px or 40px).

### Tags & Badges
**Standard**: `<span class="badge">New</span>`
*   `.badge` (or `.tag-link` / `.post-tag` unified): Inline-flex, small caps, padding, radius.

### Callouts (Already Present in All Themes)
*   `.callout`: Container with border-left.
*   `.callout-info` / `.callout-warn` / `.callout-error` / `.callout-tip`.

---

## 3. Layout Contract
*   **.layout-container** (or `.container`): Centered wrapper respecting `--max-width`.
*   **.site-header**: Sticky, contains `.nav-container`.
*   **.nav-links**: Flex container for menu items.
*   **.site-footer**: Centered, text-secondary.

### Mobile Behavior (Refactor Implemented)
*   Tablets/Phones (<768px):
    *   `.layout-container`: Padding 0-20px.
    *   Grids: Collapse to 1 column.
    *   Headers: Reduce font sizes (clamp).
    *   Nav: Stack vertically or scroll horizontally (`bird`).

---

## 4. Shortcode Support (Mandatory)
All themes MUST support:
1.  **Videos**: `.video-container` (16:9 responsive).
2.  **Code**: `pre` (scrollable), `code` (inline style).
3.  **Figures**: `figure`, `figcaption` centered.
4.  **Mermaid**: `.mermaid` classes (text visibility fixes).
5.  **Terminal**: `.terminal` window frame.

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

### How to Update an Existing Theme
1.  **Audit**: Check if `style.css` contains the mandatory **Button** and **Form** classes.
2.  **Enforce Tokens**: Ensure all standard variables from Section 1 (e.g., `--font-body`, `--header-bg`) are defined. Remove any legacy or ad-hoc variables.
3.  **Normalize**: Standardize margins/padding using the layout contract to ensure `fix-mobile-rendering` improvements persist.


