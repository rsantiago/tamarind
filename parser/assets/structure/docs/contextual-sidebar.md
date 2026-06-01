---
title: Contextual Left Sidebar
subtitle: Dynamic localized hierarchy and cross-category sibling navigation
date: 2026-06-01
tags: [features, navigation, layout]
sidebar: false
---

# Contextual Left Sidebar

Tamarind provides a dynamic, responsive left sidebar navigation area that automatically populates with localized sibling pages based on the page's parent folder context. This prevents complex documentation structures from cluttering the main menu and offers localized navigation for a category or topic.

## Default Subdirectory Behavior

By default, any page inside a subdirectory (nesting depth $\ge 1$) automatically renders a left sidebar containing sibling links from its current folder. There is no manual configuration required to enable localized subdirectory menus. Root-level pages (depth $0$, e.g., the home page) never render a sidebar by default.

## Metadata Overrides

You can explicitly control the sidebar behavior on a per-page basis using `sidebar` or `contextual_sidebar` in the page's markdown frontmatter.

### 1. Disabling the Sidebar Entirely

To completely remove the left sidebar from a subdirectory page and collapse the layout container, set the `sidebar` attribute to `false` in the frontmatter:

```yaml
sidebar: false
```

*(Note: The page you are reading right now has `sidebar: false` set in its frontmatter, which is why the left navigation sidebar has disappeared!)*

### 2. Loading Sibling Links from a Different Folder

If you want a subdirectory page to render navigation links from a different directory (for example, to link to a shared tutorial series or documentation category), specify the target directory path relative to the source directory:

```yaml
sidebar: "another-folder"
```

## Key Features

1. **Automatic Scan & Parse**: The compiler reads all `.md` files in the resolved directory, extracts their frontmatter titles, and formats relative links automatically.
2. **Active Sibling Highlights**: Sibling pages are rendered in a clean stack, with the page currently being viewed automatically receiving the `.active` class highlight.
3. **Canvas Mode Override**: Under no circumstances will a Left Sidebar render on pages set with `canvas: true` or `hide_menu: true`. Canvas mode represents a completely blank canvas (ideal for landing pages) and explicitly overrides and suppresses the contextual sidebar.
4. **Responsive Theme Compliance**: The sidebar stacks below the main content on mobile viewports and remains sticky/fixed on larger screens.

## Styling Compliance Spec

Every active Tamarind theme reserves the responsive grid class `.sidebar-left` or `.context-sidebar` for this capability. The styles automatically inherit core typography, hover transitions, and active border-left highlights from the theme's active CSS variables.
