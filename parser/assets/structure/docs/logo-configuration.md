---
title: "Logo Configuration"
subtitle: "Using a Custom Logo Image"
date: "2026-06-26"
tags: ["config", "images", "branding"]
---

Tamarind allows you to replace the default text-based logo in your theme with a custom image logo.

## Setting Up Your Logo

1. **Place your image**: Save your logo image as `logo.png` inside the `images/` directory at the root of your project.
2. **Enable it in config**: Open your `data/info.json` file and set `"use_image_logo": true`.

```json
{
  "name": "My Site",
  "domain": "https://example.com",
  "use_image_logo": true
}
```

Once enabled, Tamarind automatically updates the header of your chosen theme to use your image instead of the text logo.

## Styling the Logo

Each theme handles the logo presentation differently to match its unique aesthetic:
- In minimal themes (like `nordic` or `editorial`), the logo is displayed cleanly in the header.
- In `brutal`, the logo sits inside a brutalist bordered block with a neon yellow background.
- In `canvas`, the logo appears at the top of the sidebar.

Tamarind themes apply a `.logo-image` class to the wrapper when an image logo is used, ensuring it fits seamlessly with the theme's design language without any manual CSS required!
