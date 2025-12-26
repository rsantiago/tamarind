---
title: Image Optimization
description: Learn how Tamarind automatically resizes your images for maximum performance.
date: 2025-01-01
tags: [images, performance, features]
---

# Image Optimization

Tamarind includes a built-in image optimization pipeline that ensures your website loads instantly, even on mobile devices.

## How it Works

When you build your site, Tamarind automatically detects any **JPEG** or **PNG** images in your content directory (e.g. `writer-sandbox`). It then creates three optimized versions of each image at specific breakpoints:

- **480w**: For mobile devices.
- **800w**: For tablets and small laptops.
- **1200w**: For desktop displays.

These files are saved alongside your original image with the corresponding suffix (e.g., `photo-480w.jpg`).

## Usage

You don't need to do anything special to trigger generating these files. Just place your images in your content folders (e.g., `articles/images/my-photo.jpg`) and run `tamarind build` or `tamarind serve`.

### Using in Content

The easiest way to use optimized images is with the **Figure Shortcode**:

```markdown
{{ "{{" }} figure src="images/photo.jpg" caption="My Optimized Photo" }}
```

Tamarind will automatically generate the `srcset` attribute for you, ensuring the browser loads the correct size (480w, 800w, or 1200w).

#### Manual Example (Raw HTML)

If you prefer raw HTML or need custom sizes:

```html
<img 
  src="images/photo.jpg" 
  srcset="images/photo-480w.jpg 480w, images/photo-800w.jpg 800w, images/photo-1200w.jpg 1200w" 
  sizes="(max-width: 480px) 100vw, (max-width: 800px) 100vw, 100vw" 
  alt="My Optimized Photo"
>
```

## Active Demo

Below is a real image processed by Tamarind. If you inspect it in your browser, you will see the `srcset` attribute generated automatically.

{{ figure src="../images/landscape.jpg" caption="Demo: Automatically Optimized JPG" }}

## How it works (Internally)

Imagine you have a file named **`scenery.jpg`** (5MB, 4000px wide).

After running `tamarind build`, your output folder will contain:

1.  `scenery.jpg` (Original, untouched)
2.  `scenery-1200w.jpg` (~150KB, resized to 1200px width)
3.  `scenery-800w.jpg` (~80KB, resized to 800px width)
4.  `scenery-480w.jpg` (~30KB, resized to 480px width)

This ensures a mobile user only downloads 30KB instead of the full 5MB!
