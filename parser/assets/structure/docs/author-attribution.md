---
title: Author Attribution
date: 2025-01-02
tags: [docs, feature]
author: Tamarind Team
---

# Author Attribution

You can specify an author for each page using the `author` frontmatter field.

## Usage

In your markdown file:

```yaml
---
title: My Post
author: Different Person
---
```

If no author is specified, the default author from `data/info.json` is used.

## Configuration

Edit `data/info.json` to set the global default:

```json
{
  "author": "Your Name",
  "bio": "Your Bio"
}
```
