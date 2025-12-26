---
title: Static vs Dynamic
subtitle: Why static usually wins
date: 2025-11-26
tags: [web, architecture]
---

# The Static Renaissance

For the first 15 years of the web, everything was static. You wrote an HTML file, uploaded it via FTP, and that was it.

Then came PHP, Wordpress, creating pages on the fly.
Then came Rails and Django.
Then came the SPA revolution (Angular, React).

We gained interactivity, but we lost simplicity.

## The Cost of Dynamic

Every time a user visits a dynamic site:
1.  The request hits a server.
2.  The server runs code (Python, Node, PHP).
3.  The server queries a database.
4.  The server constructs a page.
5.  The page is sent to the user.

This is expensive. It requires CPU. It adds latency.

## The Static Advantage

With Tamarind:
1.  You run the build **once** on your laptop.
2.  The HTML is uploaded to a CDN.
3.  The user requests a page.
4.  The CDN sends the pre-built file.

**0ms database latency.**
**100% cacheability.**
**Unbreakable security** (you can't hack a clear text HTML file).

For content-driven sites, static is the superior architecture.
