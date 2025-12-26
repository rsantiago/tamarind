---
title: LaTeX Math Support
date: 2025-12-23
tags: [features, documentation, math, latex]
description: Guide to using LaTeX for mathematical notation in Tamarind.
---

# LaTeX Math Support

Tamarind uses [KaTeX](https://katex.org/) to render mathematical notation. You can use standard LaTeX syntax for both inline and display math.

## 1. Inline Math

Use single dollar signs `$` to write equations within a line of text.

### Syntax
```markdown
The formula $E = mc^2$ is widely known.
```

### Result
The formula $E = mc^2$ is widely known.

---

## 2. Display Math (Block)

For larger equations that should stand on their own line, use double dollar signs `$$`.

### Syntax
```markdown
$$ \int_{-\infty}^{\infty} e^{-x^2} \,dx = \sqrt{\pi} $$
```

### Result
$$ \int_{-\infty}^{\infty} e^{-x^2} \,dx = \sqrt{\pi} $$

---

## 3. The Math Shortcode

Alternatively, you can use the `math` shortcode.

### Syntax
```markdown
{{{!}}{ math }}
\begin{pmatrix}
   a & b \\
   c & d
\end{pmatrix}
{{{!}}{ /math }}
```

### Result
{{ math }}
\begin{pmatrix}
   a & b \\
   c & d
\end{pmatrix}
{{ /math }}

---

## 4. Complex Example

You can combine text and math freely.

### Syntax
```markdown
Consider the quadratic formula:

$$ x = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a} $$

Where $a$, $b$, and $c$ are coefficients.
```

### Result
Consider the quadratic formula:

$$ x = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a} $$

Where $a$, $b$, and $c$ are coefficients.
