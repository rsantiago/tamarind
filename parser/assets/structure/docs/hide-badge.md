---
title: Hiding the Tamarind Badge
date: 2026-06-23
tags: [docs, feature]
author: Tamarind Team
hide_badge: true
---

By default, Tamarind proudly displays an "I Use Tamarind" ghost badge in the bottom right corner of all generated pages. This helps support the independent web and sovereign software development.

However, if you have purchased a Self-Service Commercial or Enterprise license, or if you simply need a clean layout for a specific client landing page, you have the right to strip this attribution.

## How to Hide the Badge

To completely remove the ghost badge from a specific page, simply add the `hide_badge` variable to that page's frontmatter and set it to `true`.

### Example

```yaml
---
title: Client Landing Page
date: 2026-06-23
hide_badge: true
---
```

When you re-run the `tamarind build` or `tamarind serve` command, the compiler will read `hide_badge: true` and completely omit the badge HTML from the final output.

*Note: Since you are currently viewing this documentation page, you might notice that the "I Use Tamarind" badge is missing from the bottom corner! That's because this very page uses `hide_badge: true` as a live demonstration.*
