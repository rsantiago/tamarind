---
title: Author Attribution
date: 2025-01-02
tags: [docs, feature]
author: Tamarind Team
---

You can specify an author for each page using the `author` frontmatter field.

## Controlling the Attribution Display Style

Tamarind supports a frontmatter option called `attribution_style` that controls how the publication line (date and author) is displayed at the top of your pages.

### Options

*   `date-and-author` (default): Displays both the date and the author (e.g., *Published on 2026-01-01 by Tamarind Team*).
*   `date-only`: Displays only the publication date (e.g., *Published on 2026-01-01*).
*   `author-only`: Displays only the author (e.g., *Published by Tamarind Team*).
*   `none`: Hides the publication line completely. Ideal for landing pages, homepages, or contact pages.

### Example Frontmatter

```yaml
---
title: The AI-First Web Engine.
date: 2026-01-01
author: Tamarind Team
attribution_style: none
---
```

## Global Configuration

If no author is specified in a page's frontmatter, the default author from `data/info.json` is used.

Edit `data/info.json` to set the global default:

```json
{
  "author": "Your Name",
  "bio": "Your Bio"
}
```
