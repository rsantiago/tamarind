---
title: Agent Comments
date: 2025-01-01
description: Learn how to leave invisible instructions for AI Agents in your markdown.
tags: [shortcodes, ai, dev]
---

The `agent` shortcode allows you to embed instructions, context, or notes directly in your Markdown files that are **completely stripped** from the final HTML output.

This is powerful for:
1.  **AI Instructions**: Leaving prompts for AI agents generating/editing the content, without polluting the published site.
2.  **Internal Notes**: Editorial comments or TODOs that shouldn't be visible to readers.

## Syntax

```markdown
{{ "{{" }} agent "Your instruction or note here" }}
```

## Example usage

You might want to instruct an AI on how to expand a section later:

```markdown
## Future Features

{{ "{{" }} agent "TODO: Expand this section with upcoming Q3 roadmap items from the engineering wiki" }}

Coming soon...
```

## Active Demo

Below is a secret message that you cannot see in the browser, but exists in the source Markdown:

**Start of Secret Area**
{{ agent "This text is invisible in the HTML output!" }}
**End of Secret Area**

(If you see nothing between the bold lines, it works!)
