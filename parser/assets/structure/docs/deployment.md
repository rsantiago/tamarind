---
title: Deployment Guide
subtitle: Getting your site online
date: 2025-12-06
tags: [guide, ops]
---

# How to Deploy

Because Tamarind generates pure static files, you can host your website literally anywhere. You do not need a specialized "app platform" like Heroku or Fly.io. You just need a file server.

## Option 1: Netlify / Vercel (Recommended)

These platforms offer free hosting for static sites with global CDNs.

1.  Push your `writer-sandbox` (or your entire project) to GitHub.
2.  Connect your repository to Vercel/Netlify.
3.  Set the **Build Command** to: `tamarind build` (ensure the binary is committed or downloaded).
    *   *Alternative*: Run `tamarind build` locally and just commit the `website` folder. Then set the **Publish Directory** to `website`.

## Option 2: AWS S3 + CloudFront

For total control and infinite scale:

1.  Create an S3 bucket configured for static website hosting.
2.  Run `tamarind build` locally.
3.  Sync the `website` folder to the bucket:
    ```bash
    aws s3 sync website/ s3://your-bucket-name
    ```

## Option 3: GitHub Pages

1.  Go to your repository Settings > Pages.
2.  Choose "GitHub Actions" as the source.
3.  Use the standard static site upload action to publish the `website` directory.

## Tip: Caching

Since your HTML files are static, you can set aggressive cache headers for your images and CSS, making your site feel instant for repeat visitors.
