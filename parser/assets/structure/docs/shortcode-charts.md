---
title: Data Visualization Charts
tags: [features, demo, visualization, charts]
description: Create native, zero-dependency SVG charts using JSON data and shortcodes.
---


{{ callout type="info" title="Architecture Note" }}
Curious about how Tamarind mathematically guarantees elegant, theme-adaptive, and highly legible color contrast sequences across all 31 themes? Check out the [Hidden Chart Color Theory](/docs/hidden-color-theory.html) architecture document.
{{ /callout }}

Tamarind includes a powerful, zero-dependency chart rendering engine. Instead of relying on massive JavaScript libraries like Chart.js or D3, Tamarind statically compiles your data directly into highly optimized, responsive SVG elements at build time.

## Data Sources

All charts can be populated using two different methods:

### 1. External JSON Files
Perfect for large datasets or sharing data across multiple pages. Place your `.json` file in the `data/` directory and reference it:

```markdown
{{!}}{ barchart file="my_data.json" title="Sales Data" }}
```

### 2. Inline JSON Blocks
Perfect for small datasets or rapid prototyping. Place the JSON array directly between the opening and closing shortcode tags:

```markdown
{{!}}{ barchart title="Inline Data" }}
[
  {"label": "Alpha", "value": 45},
  {"label": "Beta", "value": 75}
]
{{!}}{ /barchart }}
```

---

## Single-Series Charts

Single-series charts expect a flat JSON array of objects with `label` and `value` keys.

### Vertical Bar Chart (`barchart`)

```markdown
{{!}}{ barchart title="Revenue Q1" }}
[
  {"label": "Jan", "value": 450},
  {"label": "Feb", "value": 600},
  {"label": "Mar", "value": 850}
]
{{!}}{ /barchart }}
```

{{ barchart title="Revenue Q1" }}
[
  {"label": "Jan", "value": 450},
  {"label": "Feb", "value": 600},
  {"label": "Mar", "value": 850}
]
{{ /barchart }}

### Horizontal Bar Chart (`hbarchart`)

```markdown
{{!}}{ hbarchart title="Top Frameworks" }}
[
  {"label": "React", "value": 80},
  {"label": "Vue", "value": 65},
  {"label": "Svelte", "value": 40}
]
{{!}}{ /hbarchart }}
```

{{ hbarchart title="Top Frameworks" }}
[
  {"label": "React", "value": 80},
  {"label": "Vue", "value": 65},
  {"label": "Svelte", "value": 40}
]
{{ /hbarchart }}

### Pie Chart (`piechart`)

```markdown
{{!}}{ piechart title="Traffic Sources" }}
[
  {"label": "Organic", "value": 60},
  {"label": "Direct", "value": 25},
  {"label": "Referral", "value": 15}
]
{{!}}{ /piechart }}
```

{{ piechart title="Traffic Sources" }}
[
  {"label": "Organic", "value": 60},
  {"label": "Direct", "value": 25},
  {"label": "Referral", "value": 15}
]
{{ /piechart }}

### Donut Chart (`donutchart`)

```markdown
{{!}}{ donutchart title="Revenue by Tier" }}
[
  {"label": "Pro", "value": 12000},
  {"label": "Enterprise", "value": 45000},
  {"label": "Solo", "value": 4000}
]
{{!}}{ /donutchart }}
```

{{ donutchart title="Revenue by Tier" }}
[
  {"label": "Pro", "value": 12000},
  {"label": "Enterprise", "value": 45000},
  {"label": "Solo", "value": 4000}
]
{{ /donutchart }}

### Line Chart (`linechart`)

```markdown
{{!}}{ linechart title="Monthly Active Users" }}
[
  {"label": "Jan", "value": 10},
  {"label": "Feb", "value": 25},
  {"label": "Mar", "value": 40},
  {"label": "Apr", "value": 60}
]
{{!}}{ /linechart }}
```

{{ linechart title="Monthly Active Users" }}
[
  {"label": "Jan", "value": 10},
  {"label": "Feb", "value": 25},
  {"label": "Mar", "value": 40},
  {"label": "Apr", "value": 60}
]
{{ /linechart }}

---

## Multi-Series Charts

Multi-series charts expect a JSON object containing `categories` (an array of X-axis labels) and `series` (an array of objects containing `name` and `data`).

### Multi-Line Chart (`multilinechart`)

```markdown
{{!}}{ multilinechart title="Framework Build Times" }}
{
  "categories": ["Jan", "Feb", "Mar", "Apr", "May"],
  "series": [
    { "name": "Tamarind", "data": [9.0, 7.5, 5.2, 3.1, 1.2] },
    { "name": "Next.js", "data": [3.5, 4.2, 6.0, 8.0, 9.5] }
  ]
}
{{!}}{ /multilinechart }}
```

{{ multilinechart title="Framework Build Times" }}
{
  "categories": ["Jan", "Feb", "Mar", "Apr", "May"],
  "series": [
    { "name": "Tamarind", "data": [9.0, 7.5, 5.2, 3.1, 1.2] },
    { "name": "Next.js", "data": [3.5, 4.2, 6.0, 8.0, 9.5] }
  ]
}
{{ /multilinechart }}

### Color Palette Showcase Chart (`multilinechart`)

```markdown
{{!}}{ multilinechart title="9-Color Palette Showcase" }}
{
  "categories": ["Q1", "Q2", "Q3", "Q4"],
  "series": [
    { "name": "Series 1", "data": [10, 15, 20, 25] },
    { "name": "Series 2", "data": [12, 18, 22, 28] },
    { "name": "Series 3", "data": [14, 20, 24, 30] },
    { "name": "Series 4", "data": [16, 22, 26, 32] },
    { "name": "Series 5", "data": [18, 24, 28, 34] },
    { "name": "Series 6", "data": [20, 26, 30, 36] },
    { "name": "Series 7", "data": [22, 28, 32, 38] },
    { "name": "Series 8", "data": [24, 30, 34, 40] },
    { "name": "Series 9", "data": [26, 32, 36, 42] }
  ]
}
{{!}}{ /multilinechart }}
```

{{ multilinechart title="9-Color Palette Showcase" }}
{
  "categories": ["Q1", "Q2", "Q3", "Q4"],
  "series": [
    { "name": "Series 1", "data": [10, 15, 20, 25] },
    { "name": "Series 2", "data": [12, 18, 22, 28] },
    { "name": "Series 3", "data": [14, 20, 24, 30] },
    { "name": "Series 4", "data": [16, 22, 26, 32] },
    { "name": "Series 5", "data": [18, 24, 28, 34] },
    { "name": "Series 6", "data": [20, 26, 30, 36] },
    { "name": "Series 7", "data": [22, 28, 32, 38] },
    { "name": "Series 8", "data": [24, 30, 34, 40] },
    { "name": "Series 9", "data": [26, 32, 36, 42] }
  ]
}
{{ /multilinechart }}


### Grouped Bar Chart (`groupedbarchart`)

```markdown
{{!}}{ groupedbarchart title="Project Sizes (MB)" }}
{
  "categories": ["Project A", "Project B", "Project C"],
  "series": [
    { "name": "HTML", "data": [1.2, 2.5, 0.8] },
    { "name": "CSS", "data": [0.5, 1.1, 0.3] },
    { "name": "JS", "data": [3.5, 8.2, 1.5] }
  ]
}
{{!}}{ /groupedbarchart }}
```

{{ groupedbarchart title="Project Sizes (MB)" }}
{
  "categories": ["Project A", "Project B", "Project C"],
  "series": [
    { "name": "HTML", "data": [1.2, 2.5, 0.8] },
    { "name": "CSS", "data": [0.5, 1.1, 0.3] },
    { "name": "JS", "data": [3.5, 8.2, 1.5] }
  ]
}
{{ /groupedbarchart }}

### Radar Chart (`radarchart`)

```markdown
{{!}}{ radarchart title="Feature Comparison" }}
{
  "categories": ["Speed", "Simplicity", "Ecosystem", "Customizability", "AI-Readiness"],
  "series": [
    { "name": "Tamarind", "data": [10, 9, 4, 8, 10] },
    { "name": "Competitor", "data": [4, 5, 10, 7, 5] }
  ]
}
{{!}}{ /radarchart }}
```

{{ radarchart title="Feature Comparison" }}
{
  "categories": ["Speed", "Simplicity", "Ecosystem", "Customizability", "AI-Readiness"],
  "series": [
    { "name": "Tamarind", "data": [10, 9, 4, 8, 10] },
    { "name": "Competitor", "data": [4, 5, 10, 7, 5] }
  ]
}
{{ /radarchart }}
