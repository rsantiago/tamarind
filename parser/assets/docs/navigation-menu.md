---
title: Navigation Menus & Custom Ordering
subtitle: Configure top-level site navigation and custom folder collection overrides
date: 2026-06-01
tags: [features, navigation, config]
---

Tamarind features a dynamic top navigation menu that is automatically constructed at build time. By combining automated directory scans with granular metadata overrides in your page frontmatter, you can construct clean, structured navigation menus without writing custom HTML or configuration files.

---

## 1. How the Menu is Built

During the compilation scan phase, the Tamarind engine performs two operations:
1. **Scans Root Pages**: It searches the root of your source directory for all `.md` files (excluding hidden files or build directories). Each of these pages is added to the menu.
2. **Scans Folder Collections**: It searches for subdirectories (excluding internal folders starting with a dot, `website`, or `data`). Each subdirectory is recognized as an active content collection and is automatically appended to the top menu.

---

## 2. Granular Page Menu Overrides

You can control how individual root pages appear in the menu by defining specific metadata tags in the page's YAML frontmatter:

### `menu_label`
Overrides the text displayed in the navigation link. If omitted, the menu title defaults to the page's `title` frontmatter, or a Title-Cased version of the file's base name.
```yaml
---
title: "Reach Out to Our Core Engineering Team"
menu_label: "Contact"
---
```

### `menu_order`
Explicitly defines the sorting position of the menu item (sorted ascending, meaning lower numbers appear first).
```yaml
---
title: "Tamarind Feature List"
menu_order: 1
---
```

### `hidden`
Completely excludes the page from appearing in both the top-level navigation menus and the generated `sitemap.xml` file, while keeping it compiled and accessible via its direct link.
```yaml
---
title: "Internal Developer Playbook"
hidden: true
---
```

---

## 3. Dynamic Folder Collection Overrides

By default, any discovered folder collections (such as `/blog/` or `/docs/`) are appended to the end of the top navigation menu using a default sorting order of `99`. Multiple collections are sorted alphabetically by title.

To customize the order or change the menu label of a collection:

1. **Create an Override File**: In the root of your source directory, create a `.md` file with the exact name of the target collection folder (for example, `/roadmap/items.md` for the `/roadmap/items/` folder collection).
2. **Set the Frontmatter Keys**: Specify `menu_order` and `menu_label` inside its frontmatter:
   ```yaml
   ---
   title: "Roadmap Items"
   menu_label: "Items"
   menu_order: 20
   ---
   ```

The Tamarind compiler detects the matching override file, extracts the customized menu order and labels, and safely prevents any duplicate entries from rendering in your navigation bar. The actual content index for the collection remains dynamically generated and populated automatically.

---

## 4. Custom Brand Logo Image

Instead of rendering a default text title and SVG icon, you can replace the upper-left website name with a custom brand logo image (like `images/logo.png`).

### Configuration Steps:
1.  **Place the Logo File**: Put your brand logo image at `images/logo.png` inside your source directory (e.g. `writer-sandbox/images/logo.png`).
2.  **Enable the Image Logo**: Add the `"use_image_logo": true` property to your global data file `data/info.json`:
    ```json
    {
      "use_image_logo": true
    }
    ```

The Tamarind compiler automatically copies the logo image, optimizes it for various screen resolutions, and replaces the header text with your brand logo image.
