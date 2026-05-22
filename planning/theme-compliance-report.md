# Tamarind Theme Compliance Report

**Generated**: 2026-05-22 18:53:18

**Total Requirements**: 62 required, 5 optional

## Summary

| Theme | Status | Passed | Failed | Score |
|-------|--------|--------|--------|-------|
| **atlas** | ✅ PASS | 62 | 0 | 100% |
| **basic** | ✅ PASS | 62 | 0 | 100% |
| **bento** | ✅ PASS | 62 | 0 | 100% |
| **bird** | ✅ PASS | 62 | 0 | 100% |
| **blue** | ✅ PASS | 62 | 0 | 100% |
| **brutal** | ✅ PASS | 62 | 0 | 100% |
| **canvas** | ✅ PASS | 62 | 0 | 100% |
| **classic** | ✅ PASS | 62 | 0 | 100% |
| **console** | ✅ PASS | 62 | 0 | 100% |
| **cupertino** | ✅ PASS | 62 | 0 | 100% |
| **eclipse** | ✅ PASS | 62 | 0 | 100% |
| **editorial** | ✅ PASS | 62 | 0 | 100% |
| **forge** | ✅ PASS | 62 | 0 | 100% |
| **gallery** | ✅ PASS | 62 | 0 | 100% |
| **gram** | ✅ PASS | 62 | 0 | 100% |
| **hacker** | ✅ PASS | 62 | 0 | 100% |
| **midnight** | ✅ PASS | 62 | 0 | 100% |
| **neon** | ✅ PASS | 62 | 0 | 100% |
| **network** | ✅ PASS | 62 | 0 | 100% |
| **news** | ✅ PASS | 62 | 0 | 100% |
| **nexus** | ✅ PASS | 62 | 0 | 100% |
| **nordic** | ✅ PASS | 62 | 0 | 100% |
| **overflow** | ✅ PASS | 62 | 0 | 100% |
| **pastel** | ✅ PASS | 62 | 0 | 100% |
| **pod** | ✅ PASS | 62 | 0 | 100% |
| **prose** | ✅ PASS | 62 | 0 | 100% |
| **protocol** | ✅ PASS | 62 | 0 | 100% |
| **scribe** | ✅ PASS | 62 | 0 | 100% |
| **stream** | ✅ PASS | 62 | 0 | 100% |
| **tube** | ✅ PASS | 62 | 0 | 100% |
| **zephyr** | ✅ PASS | 62 | 0 | 100% |

---

## Theme: atlas — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: basic — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: bento — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: bird — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: blue — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: brutal — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: canvas — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: classic — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: console — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: cupertino — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: eclipse — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: editorial — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: forge — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: gallery — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: gram — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: hacker — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: midnight — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: neon — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: network — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: news — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: nexus — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: nordic — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: overflow — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: pastel — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: pod — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: prose — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: protocol — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: scribe — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: stream — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: tube — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

## Theme: zephyr — ✅ COMPLIANT

**Score**: 62/62 (100%)

### ✅ Passed Requirements

#### accessibility

- `mobile-nav-visible` — The site navigation (.nav-links) must not be hidden (display: none) on mobile viewports (<768px)
- `mobile-sidebar-visible` — The primary document sidebar (.sidebar) must not be hidden (display: none) on mobile viewports (<768px)

#### article-content

- `ul` — Unordered lists (ul) must be explicitly styled
- `ol` — Ordered lists (ol) must be explicitly styled
- `li` — List items (li) must be explicitly styled
- `blockquote` — Blockquotes must be explicitly styled
- `code` — Inline code blocks must be explicitly styled
- `pre` — Preformatted code blocks (pre) must be explicitly styled
- `table` — Data tables must be explicitly styled

#### components-badges

- `.badge` — Badge/tag label

#### components-buttons

- `.btn` — Button base class
- `.btn-primary` — Primary action button
- `.btn-secondary` — Secondary action button
- `.btn-ghost` — Ghost/transparent button
- `.btn-sm` — Small button variant
- `.button-group` — Wrapper for multiple buttons

#### components-callouts

- `.callout` — Callout container
- `.callout-info` — Info callout variant
- `.callout-warn` — Warning callout variant
- `.callout-error` — Error callout variant
- `.callout-tip` — Tip callout variant

#### components-forms

- `.form-group` — Form group container
- `.form-label` — Form label
- `.form-input` — Text input field
- `.form-textarea` — Textarea field
- `.form-select` — Select dropdown
- `.form-checkbox` — Checkbox input
- `.form-radio` — Radio input
- `.form-file` — File input

#### layout

- `.layout-container` — Main content container
- `.site-header` — Sticky site header
- `.nav-container` — Navigation wrapper inside header
- `.nav-links` — Navigation links container
- `.site-footer` — Site footer
- `.tamarind-ghost-badge` — Floating Use Tamarind badge element
- `.footer-promo` — Website promotional reference in the footer

#### responsive

- `max-width: 768px` — Mobile breakpoint at 768px

#### shortcodes

- `.video-container` — Responsive video embed container (16:9)
- `.mermaid` — Mermaid diagram container
- `.terminal` — Terminal window frame
- `img` — Responsive image styling to prevent overflow
- `figure` — Responsive figure block formatting

#### templates

- `articles.mdt:tamarind-ghost-badge` — articles.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `articles.mdt:footer-promo` — articles.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `page.mdt:tamarind-ghost-badge` — page.mdt template must contain the 'tamarind-ghost-badge' class or render 'footer.mdt'
- `page.mdt:footer-promo` — page.mdt template must contain the 'footer-promo' class or render 'footer.mdt'
- `menu-no-duplication` — Templates must not duplicate the main menu rendering (must not loop over .Menu more than once)

#### variables-colors

- `--primary-color` — Main interaction color (Buttons, Links)
- `--background-color` — Main page background
- `--card-bg` — Component/Article card background
- `--header-bg` — Sticky header background
- `--text-color` — Main body text color
- `--text-secondary` — Metadata, captions, footers color
- `--border-color` — Separators and input borders

#### variables-spacing

- `--max-width` — Main container width
- `--header-height` — Sticky header height
- `--radius-sm` — Small border radius
- `--radius-md` — Medium border radius
- `--radius-lg` — Large border radius
- `--shadow` — Default box shadow

#### variables-typography

- `--font-heading` — Heading font family (H1-H6, Buttons)
- `--font-body` — Body font family (P, Li, Inputs)

---

