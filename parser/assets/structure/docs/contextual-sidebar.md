---
title: Contextual Left Sidebar
subtitle: Dynamic localized hierarchy and cross-category sibling navigation
date: 2026-06-01
tags: [features, navigation, layout]
contextual_sidebar: "docs"
---

# Contextual Left Sidebar

Tamarind provides a dynamic, responsive left sidebar navigation area that automatically populates with localized sibling pages based on the page's parent folder context. This prevents complex documentation structures from cluttering the main menu and offers localized navigation for a category or topic.

## Configuration

To activate the contextual sidebar on any page, define the `contextual_sidebar` property in your markdown frontmatter specifying the folder name containing the pages you want to show in the left sidebar:

```yaml
contextual_sidebar: "docs"
```

By providing the relative directory name, the Tamarind compiler dynamically scans that folder, parses metadata from all sibling markdown files, and generates a structured list of relative links.

## Key Features

1. **Automatic Scan & Parse**: The compiler reads all `.md` files in the specified directory, extracts their frontmatter titles, and formats relative links automatically.
2. **Active Sibling Highlights**: Sibling pages are rendered in a clean stack, with the page currently being viewed automatically receiving the `.active` class highlight.
3. **Canvas Mode Override**: Under no circumstances will a Left Sidebar render on pages set with `canvas: true` or `hide_menu: true`. Canvas mode represents a completely blank canvas (ideal for landing pages) and explicitly overrides and suppresses the contextual sidebar.
4. **Responsive Theme Compliance**: The sidebar stacks below the main content on mobile viewports and remains sticky/fixed on larger screens.

## Styling Compliance Spec

Every active Tamarind theme reserves the responsive grid class `.sidebar-left` or `.context-sidebar` for this capability. The styles automatically inherit core typography, hover transitions, and active border-left highlights from the theme's active CSS variables.
