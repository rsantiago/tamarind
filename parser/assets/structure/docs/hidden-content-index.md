---
title: "Tutorial: Managing Hidden & Unlisted Content"
date: 2025-12-23
description: "Learn how to use the 'hidden: true' front matter tag to publish private, unlisted, or under-construction pages in Tamarind."
tags: [guide, features, layout]
---

Tamarind allows you to publish pages that are fully compiled and accessible via their direct URL, but are **completely excluded** from all automated feeds, navigation menus, and search indexes. 

This is incredibly useful for:
1.  **Unlisted Articles**: Premium newsletters, invite-only landing pages, or private beta documentation.
2.  **Drafts & Previews**: Staging content to show clients/collaborators before public launch.
3.  **A/B Landing Page Testing**: Creating conversion variations that should not be discovered via standard site search or site navigation.

---

## How It Works: The `hidden: true` Tag

To hide any page from the public index, simply add `hidden: true` to the front matter of the page's Markdown file:

```markdown
---
title: "My Secret Page"
date: 2026-06-16
hidden: true
---
```

When the compiler detects this flag:
*   **Feeds Exclusion**: The page is skipped during the generation of the RSS feed (`feed.xml`).
*   **Sitemap Exclusion**: The URL is omitted from the search-engine index mapping (`sitemap.xml`).
*   **Collection Exclusions**: The page does *not* appear in automatic doc or article lists (like the main `docs.html` listing page).
*   **Tag Index Exclusions**: The page's tags are ignored, meaning it will not appear under pages like `tags/secret.html`.

---

## Active Verification Showcase

Here is a list of hidden pages currently compiled in this project. You can click them to verify they exist and render perfectly, and then check the index pages to confirm they are absent.

### Compiled Hidden Pages (Direct Access Links):
*   [Hidden Article (docs/hidden-secret.html)](hidden-secret.html) — A mock unlisted article.
*   [Classified Documentation (docs/hidden-manual.html)](hidden-manual.html) — An unlisted internal guide.
*   [Secret Landing Page (secret-landing.html)](../secret-landing.html) — A full-bleed marketing landing page canvas.

### Verification Steps:
1.  **Feeds Check**: View the [RSS Feed File](../feed.xml) and search for "The Hidden Article" or "Secret Landing Page". (They will not be listed).
2.  **Sitemap Check**: View the [Sitemap](../sitemap.xml) and check the list of `<loc>` tags. (The URLs for these pages are absent).
3.  **Collection Check**: Go back to the [Docs Index](../docs.html) and verify that none of these three pages are listed.
