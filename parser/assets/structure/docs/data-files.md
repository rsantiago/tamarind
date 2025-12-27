---
title: Data Files
description: Learn how to use data files to store structured content.
date: 2025-01-01
tags: [features, data, yaml, json]
---

# Data Files

Data files allow you to verify structured data (like list of authors, social links, or products) separate from your content and layout. This data is automatically loaded by Tamarind and made available in your templates.

## Organizing Data

Create a `data` directory inside your content folder (default: `writer-sandbox`). Add `YAML` or `JSON` files to it.

**Example Structure:**

```text
writer-sandbox/
  data/
    authors.yaml
    social.json
  posts/
    ...
```

### Scenario 1: Single Object (e.g. Site Info)

Use this when you have a single set of key-value pairs, like configuration or social links.

**File:** `writer-sandbox/data/info.json`
```json
{
  "version": "1.0",
  "author": "Tamarind User",
  "contact_email": "hello@example.com"
}
```

**Usage:**
Access fields directly using the key names.

```markdown
**Version:** {{ "{{" }} .Data.info.version }}

**Contact:** [{{ "{{" }} .Data.info.contact_email }}](mailto:{{ "{{" }} .Data.info.contact_email }})
```

#### Active Demo (Live Result)

**Version:** {{ .Data.info.version }}

**Contact:** [{{ .Data.info.contact_email }}](mailto:{{ .Data.info.contact_email }})

### Scenario 2: List of Objects (e.g. Products)

Use this when you have a collection of items to iterate over.

**File:** `writer-sandbox/data/products.json`
```json
[
  {
    "name": "Super Widget",
    "price": "$19.99"
  },
  {
    "name": "Mega Widget",
    "price": "$29.99"
  }
]
```

**Usage:**
Use `range` to loop through the list.

```markdown
{{ "{{" }} range .Data.products }}
- **{{ "{{" }} .name }}**: {{ "{{" }} .price }}
{{ "{{" }} end }}
```

#### Active Demo (Live Result)

{{ range .Data.products }}
- **{{ .name }}**: {{ .price }}
{{ end }}

### Scenario 3: Accessing Specific Items in a List

If you need to grab just the *second* product from the list above, you can use the `index` function. Note that indexes are 0-based.

```markdown
**Featured:** {{ "{{" }} (index .Data.products 1).name }}
```

#### Active Demo (Live Result)

**Featured:** {{ (index .Data.products 1).name }}

## Why use Data Files?

1.  **Separation of Concerns**: Keep content clean and logic-free.
2.  **Reusability**: Define data once, use it in multiple templates (Footer, Header, About Page).
3.  **Maintainability**: Updating a link in one JSON file updates it everywhere.
