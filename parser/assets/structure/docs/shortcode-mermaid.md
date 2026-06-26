---
title: Mermaid Diagrams
tags: [features, demo, visualization]
description: Create flowcharts, sequence diagrams, and more using code.
---

# Mermaid Diagrams

Tamarind integrates [Mermaid.js](https://mermaid.js.org/) to let you create diagrams and visualizations using text and code. Because Tamarind hands off the parsed blocks to the Mermaid renderer in the browser, all official diagram types are fully supported out-of-the-box.

---

## 1. Core Structural & Behavioral Diagrams

These diagrams are essential for visualizing algorithms, interactions, code structures, and state transitions.

### Flowchart / Graph
Standard node and edge graphs for process flows.
```markdown
{{!}}{{ mermaid }}
graph LR;
    A[Start] --> B{Is it working?};
    B -- Yes --> C[Great!];
    B -- No --> D[Debug];
    D --> B;
{{!}}{{ /mermaid }}
```

### Sequence Diagram
Visualize interactions and timelines between systems or objects.
```markdown
{{!}}{{ mermaid }}
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
{{!}}{{ /mermaid }}
```

### Class Diagram
Visualize object-oriented relationships and attributes.
```markdown
{{!}}{{ mermaid }}
classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal : +int age
    Animal: +isMammal()
    class Duck{
        +String beakColor
        +swim()
    }
{{!}}{{ /mermaid }}
```

### State Diagram
Visualize state transitions of a system over time.
```markdown
{{!}}{{ mermaid }}
stateDiagram-v2
    [*] --> Still
    Still --> [*]
    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
{{!}}{{ /mermaid }}
```

### Entity Relationship Diagram (ERD)
Visualize database schemas and relationships.
```markdown
{{!}}{{ mermaid }}
erDiagram
    CUSTOMER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
{{!}}{{ /mermaid }}
```

---

## 2. Project Management & Planning

These tools help teams plan schedules, map user flows, and brainstorm hierarchies.

### Mindmap
For brainstorming and organizing hierarchical information visually.
```markdown
{{!}}{{ mermaid }}
mindmap
  root((Tamarind))
    Design
      Themes
      CSS
    Performance
      Go Binary
      Zero Dependencies
{{!}}{{ /mermaid }}
```

### Gantt Chart
For project schedules, tracking task durations and dependencies. (Use `gantt`).

### User Journey
For mapping a user's emotional and practical journey through a product. (Use `journey`).

### Timeline
For plotting chronological events. (Use `timeline`).

### Requirement Diagram
For mapping out software requirements and test traces. (Use `requirementDiagram`).

---

## 3. Specialized & Advanced Visualizations

### Pie Chart
Visualize percentage distributions.
```markdown
{{!}}{{ mermaid }}
pie title What specific languages do you use?
    "Go" : 40
    "Python" : 30
    "JavaScript" : 20
    "Rust" : 10
{{!}}{{ /mermaid }}
```

### Quadrant Chart
Plot items on an X/Y axis with 4 quadrants (e.g., Risk vs. Reward). (Use `quadrantChart`).

### Gitgraph
Visualize git branch histories, commits, and merges. (Use `gitGraph`).

### C4 Architecture
Define software architecture using the C4 model. (Use `C4Context`, `C4Container`, etc.).

---

## 4. Newer/Experimental Diagrams (Beta)

Mermaid is constantly adding new beta diagrams. All of these work natively via the `{{ mermaid }}` shortcode.

* **Sankey Diagram** (`sankey-beta`): For flow rates between states.
* **XY Chart** (`xychart-beta`): For standard line and bar graphs.
* **Block Diagram** (`block-beta`): For spatial block-based layouts.
* **Architecture Diagram** (`architecture-beta`): For cloud/network topology architecture maps.
* **Packet Diagram** (`packet-beta`): For visualizing network packet bytes/bits headers.

---

## 5. Developer Directives & Best Practices

When writing Mermaid diagrams, especially large architectural maps, observe the following rules to prevent syntax errors that will crash the compilation:

1. **Avoid Special Characters**: The Mermaid parser will instantly break if it encounters unescaped punctuation (`/`, `*`, `<`, `>`, `~`, `(`, `)`, `[`, `]`) inside node labels or sequence diagram messages.
2. **Quote Labels in Graphs**: When building a `graph TD`, **always** quote your node labels if they contain spaces or colons. E.g., `Node["My Label: Something"]`.
3. **Sequence Diagrams**: Strip method signatures of arguments and brackets. Use `Participant->>System: Execute` rather than `Participant->>System: Execute("page.mdt")`. Do NOT use slashes in participant names (e.g., use `html_template` instead of `html/template`).
4. **Class Diagrams**: Type definitions must be alphanumeric. Avoid using generic types with `~` (e.g. `map~string, interface~`) or array bracket notations (e.g. `[]string`). Instead, simplify types to alphanumeric structures like `Map` or `List`.
5. **No HTML Tags**: Never place HTML tags (`<br>`, `<strong>`) inside node labels unless you know the specific diagram type fully supports it (only a subset of Mermaid diagrams do).
6. **NO BLANK LINES**: Never place blank empty lines inside the `{{ mermaid }}` shortcode. The underlying shortcode parser injects a `<div class="mermaid">` block into the Markdown tree *before* Goldmark processes the text. If Goldmark encounters a blank line inside the HTML block, it will escape the HTML context and wrap your indented Mermaid code inside `<pre><code>` blocks, completely breaking the syntax.
