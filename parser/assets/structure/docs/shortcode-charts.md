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

## Customizing Colors

You can override the default theme colors by providing a comma-separated list of colors (names, HEX, or CSS variables) to the `colors` parameter. This gives you absolute control over the visualization while maintaining the minimalist, zero-dependency rendering approach.

```markdown
{{< barchart title="Custom Brand Colors" colors="magenta,cyan,#FF5733" >}}
[
    {"label": "Direct", "value": 40},
    {"label": "Social", "value": 30},
    {"label": "Referral", "value": 20}
]
{{< /barchart >}}
```

{{ barchart title="Custom Brand Colors" colors="magenta,cyan,#FF5733" }}
[
    {"label": "Direct", "value": 40},
    {"label": "Social", "value": 30},
    {"label": "Referral", "value": 20}
]
{{ /barchart }}

### Reordering Theme Colors

You can map specific CSS theme variables to specific data series. This keeps the chart deeply linked to the active theme (dynamically adjusting to light/dark mode switches), while giving you absolute control over which color applies to which series!

For example, to reverse the first three colors so that Series 1 uses color 3, Series 2 uses color 2, and Series 3 uses color 1:

```markdown
{{< barchart title="Reordered Theme Colors" colors="var(--chart-3),var(--chart-2),var(--chart-1)" >}}
[
    {"label": "Series 1", "value": 40},
    {"label": "Series 2", "value": 30},
    {"label": "Series 3", "value": 20}
]
{{< /barchart >}}
```

{{ barchart title="Reordered Theme Colors" colors="var(--chart-3),var(--chart-2),var(--chart-1)" }}
[
    {"label": "Series 1", "value": 40},
    {"label": "Series 2", "value": 30},
    {"label": "Series 3", "value": 20}
]
{{ /barchart }}

---

## Chart Configuration Flags

Tamarind allows you to fine-tune your charts using optional flags directly inside the shortcode. These flags are available for all Cartesian (X-Y axis) charts, such as `barchart`, `linechart`, `multilinechart`, `groupedbarchart`, and `hbarchart`.

*   `show-x="false"`: Hides the X-axis labels (for vertical bars/lines). For horizontal bar charts, this controls the numerical value labels. (Default: `true`)
*   `show-y="false"`: Hides the Y-axis value labels (for vertical bars/lines). For horizontal bar charts, this controls the category labels. (Default: `true`)
*   `show-dots="false"`: Disables the data point circles on line and multiline charts. (Default: `true`)
*   `grid-y="true"`: Enables subtle horizontal background grid lines corresponding to the Y-axis. (Default: `false`)
*   `grid-x="true"`: Enables subtle vertical background grid lines corresponding to the X-axis. (Default: `false`)

### Examples: Configuration Flags

**1. Default Chart (No Flags)**

```markdown
{{!}}{ multilinechart title="Default Chart" }}
{
  "categories": ["Mon", "Tue", "Wed", "Thu", "Fri"],
  "series": [
    { "name": "Active Users", "data": [120, 150, 200, 180, 250] },
    { "name": "New Signups", "data": [45, 60, 80, 75, 110] }
  ]
}
{{!}}{ /multilinechart }}
```

{{ multilinechart title="Default Chart" }}
{
  "categories": ["Mon", "Tue", "Wed", "Thu", "Fri"],
  "series": [
    { "name": "Active Users", "data": [120, 150, 200, 180, 250] },
    { "name": "New Signups", "data": [45, 60, 80, 75, 110] }
  ]
}
{{ /multilinechart }}

**2. Minimalist Chart (Hidden Axes, No Dots, With Grid-Y)**

```markdown
{{!}}{ multilinechart title="Clean & Minimalist" show-x="false" show-y="false" show-dots="false" grid-y="true" }}
{
  "categories": ["Mon", "Tue", "Wed", "Thu", "Fri"],
  "series": [
    { "name": "Active Users", "data": [120, 150, 200, 180, 250] },
    { "name": "New Signups", "data": [45, 60, 80, 75, 110] }
  ]
}
{{!}}{ /multilinechart }}
```

{{ multilinechart title="Clean & Minimalist" show-x="false" show-y="false" show-dots="false" grid-y="true" }}
{
  "categories": ["Mon", "Tue", "Wed", "Thu", "Fri"],
  "series": [
    { "name": "Active Users", "data": [120, 150, 200, 180, 250] },
    { "name": "New Signups", "data": [45, 60, 80, 75, 110] }
  ]
}
{{ /multilinechart }}

**3. Maximum Detail (Everything Shown)**

```markdown
{{!}}{ multilinechart title="Maximum Detail" show-x="true" show-y="true" show-dots="true" grid-x="true" grid-y="true" }}
{
  "categories": ["Mon", "Tue", "Wed", "Thu", "Fri"],
  "series": [
    { "name": "Active Users", "data": [120, 150, 200, 180, 250] },
    { "name": "New Signups", "data": [45, 60, 80, 75, 110] }
  ]
}
{{!}}{ /multilinechart }}
```

{{ multilinechart title="Maximum Detail" show-x="true" show-y="true" show-dots="true" grid-x="true" grid-y="true" }}
{
  "categories": ["Mon", "Tue", "Wed", "Thu", "Fri"],
  "series": [
    { "name": "Active Users", "data": [120, 150, 200, 180, 250] },
    { "name": "New Signups", "data": [45, 60, 80, 75, 110] }
  ]
}
{{ /multilinechart }}

---

## Theme Compliance: Chart Colors

Tamarind natively supports rendering a wide variety of data visualization charts (Pie, Bar, Line, Radar, etc.) directly from markdown. To ensure these charts look perfect on every theme, the parser relies on a standardized set of CSS variables.

**Theme developers MUST define the following CSS variables in their `theme.css` files:**

```css
:root {
  /* Primary and Secondary branding are always used for the first two series */
  --primary-color: #3b82f6;
  --secondary-color: #10b981;

  /* Chart-specific extended color palette */
  --chart-1: var(--primary-color);
  --chart-2: var(--secondary-color);
  --chart-3: #f59e0b;
  --chart-4: #ef4444;
  --chart-5: #8b5cf6;
  --chart-6: #ec4899;
  --chart-7: #14b8a6;
  --chart-8: #f97316;
  --chart-9: #64748b;
}
```

If a theme fails to define these variables, the compiler will safely fall back to the built-in system defaults, but the charts may not perfectly match the theme's intended aesthetic.

### Axis Contrast Requirement

In addition to the primary chart palette, Tamarind utilizes the `--text-secondary` CSS variable to render the structural components of charts, including X/Y axes, radar webs, and chart bounding boxes. 

To ensure charts remain legible, theme developers **must** verify that their `--text-secondary` color maintains a high enough contrast ratio against their `--background-color` and `--card-bg` in both light and dark modes. Failure to provide sufficient contrast will result in invisible or washed-out chart structures.
