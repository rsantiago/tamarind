# Changelog

A complete history of **Tamarind** releases and milestones.

## Semantic Releases (v4.0+)

| Version | Description |
| :--- | :--- |
| **v4.11.0** | **Agent Experience**: Added `{{ agent }}` comment shortcode for invisible instructions; Enhanced `quickstart` to enable **Live Reloading** by default; Updated Landing Page copy. |
| **v4.10.1** | **Polish & Copy**: Vibrant headers for 'Gram' individual pages, removed placeholder initials from collections, fixed data loader casing, and updated landing page copy. |
| **v4.10.0** | **Structural Refinement**: Flattened project structure, added **Author Attribution** docs, enhanced **Dynamic Menu**, and polished **Gram** theme headers. |
| **v4.9.1** | **Polish**: Removed figure box shadows across all themes; Enhanced `figure` shortcode with `width` attribute support. |
| **v4.9.0** | **Feature Suite**: Implemented automatic **Image Optimization** (resize pipeline, srcset) and **Data Files** support (`data/*.json` in templates). |
| **v4.8.0** | **RSS Feeds**: Added `feed.xml` generation for articles and RSS footer link. |
| **v4.7.0** | **Live Reloading**: Implemented dev server with `-watch` flag, file watcher, and automatic browser refresh (SSE). |
| **v4.6.1** | **Brand Polish**: Updated footer with larger, bolder "Proudly built using Tamarind" branding. |
| **v4.6.0** | **Draft Handling**: Added `draft: true` support and `-drafts` CLI flag to securely preview unfinished content. |
| **v4.5.0** | **Theme Polish**: Comprehensive Mermaid dark mode support; Redesigned "Neon" (Cyberpunk Dashboard) & "Pastel" (Sharp Scandinavian) themes; Typography overhaul. |
| **v4.4.1** | **DX & SEO Polish**: Fixed `quickstart` command shared templates bug locally; Added missing `sitemap.xml` generation. |
| **v4.4.0** | **AI Engine Optimization (AEO)**: Added `robots.txt`, `llms.txt`, and `llms_full.txt` generation; supported dynamic collection scanning. |
| **v4.3.0** | **Component Refactor**: Centralized `footer.mdt` partial across all themes with Tamarind branding. |
| **v4.2.0** | **Unified Templating**: Centralized `head.mdt` partial for consistent SEO/Scripts management. |
| **v4.1.0** | **AI Automation**: Added comprehensive AI documentation and Agentic System Prompts. |
| **v4.0.0** | **Modular Core**: Complete refactor of `builder.go` into modular components (`scanner`, `pages`, `assets`, `shortcodes`). |

## Semantic Releases (v3.0 - v3.9)

| Version | Description |
| :--- | :--- |
| **v3.9.0** | **Developer Suite**: Added Mermaid, Callouts, Terminal, Code Includes, Gists, and Math support. |
| **v3.8.1** | **Cleanup & Local Assets**: Enabled local image support, generic asset copying, and theme cleanup. |
| **v3.8.0** | **Figure/Image Shortcode**: Added support for semantic, captioned images. |
| **v3.7.0** | **YouTube Shortcode & Configuration**: Added video embedding and centralized version management. |
| **v3.6.0** | **Hidden Content**: Added support for unlisted pages (`hidden: true`) and improved theme link visibility. |
| **v3.5.0** | **Quickstart**: Added the `tamarind quickstart` magic command. |
| **v3.4.1** | **Pagination UI**: Enhanced pagination with a "sliding window" interface. |
| **v3.4.0** | **Pagination Core**: Implemented core pagination logic with configurable limits. |
| **v3.3.0** | **Tags**: Added support for auto-generating Tag pages. |
| **v3.2.0** | **Universal Theme Config**: Standardized CSS variables and added heading-color support. |
| **v3.1.0** | **SEO**: Added meta tags, canonical URLs, and sitemap generation. |
| **v3.0.0** | **Initial Stable Release**: First major release of the SSG. |

## Early Development (v1.0 - v2.8)

| Version | Description |
| :--- | :--- |
| **v2.8.0-bird-refined** | Refined "Bird" theme and build process cleanup. |
| **v2.7.0-gram-refined** | Refined "Gram" theme (spacing, bullets, header). |
| **v2.6.0-cleanup** | Added `themes` command and removed unused templates. |
| **v2.5.0-gram** | Completed "Gram" theme, dynamic SiteTitle, and auto-kill port logic. |
| **v2.4.0** | Upgrade to temp build dirs, safe init backups, and Blue theme refinement. |
| **v2.3.1** | Enhanced content and isolated init. |
| **v2.3.0** | Arbitrary Collections support & Project Restructure. |
| **v2.2.0-tamarind** | "Project Tamarind" branding release. |
| **v1.0.0** | First stable release of the SSG. |

## Pre-Release Milestones

| Tag | Achievement |
| :--- | :--- |
| `theme-sync-fix-v1` | Shared theme state across pages. |
| `site-restructure-v1` | Restructure with `articles/` directory and dark theme refinements. |
| `project-restructure` | Split into `website` and `structure` directories. |
| `parser-v1` | Complete Go parser implementation with local server. |
| `parser-scaffold-v1` | Initial Go parser scaffold. |
| `parser-release-script-v1` | Release packaging script. |
| `mobile-polish-v1` | Mobile typography improvements. |
| `markdown-structure-v1` | Markdown source structure and template system. |
| `logo-bebas-neue` | Added Bebas Neue font for logo. |
| `local-server-v1` | Added simple Go web server. |
| `dark-theme-v1` | Implemented dark theme with gradient headers. |
| `contributors-guide-v1` | Added `AGENTS.md` contributor guide. |
