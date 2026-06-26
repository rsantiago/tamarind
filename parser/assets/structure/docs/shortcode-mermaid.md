---
title: Mermaid Diagrams
tags: [features, demo, visualization]
description: Create flowcharts, sequence diagrams, and more using code.
---

# Mermaid Diagrams

Tamarind integrates [Mermaid.js](https://mermaid.js.org/) to let you create diagrams and visualizations using text and code.

## 1. Flowchart

Standard node and edge graphs.

### Syntax
```markdown
{{{!}}{ mermaid }}
graph TD;
    A[Start] --> B{Is it working?};
    B -- Yes --> C[Great!];
    B -- No --> D[Debug];
    D --> B;
{{{!}}{ /mermaid }}
```

### Result
{{ mermaid }}
graph TD;
    A[Start] --> B{Is it working?};
    B -- Yes --> C[Great!];
    B -- No --> D[Debug];
    D --> B;
{{ /mermaid }}

---

## 2. Sequence Diagram

Visualize interactions and timelines.

### Syntax
```markdown
{{{!}}{ mermaid }}
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    loop Healthcheck
        John->>John: Fight against hypochondria
    end
    Note right of John: Rational thoughts <br/>prevail!
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
{{{!}}{ /mermaid }}
```

### Result
{{ mermaid }}
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    loop Healthcheck
        John->>John: Fight against hypochondria
    end
    Note right of John: Rational thoughts <br/>prevail!
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
{{ /mermaid }}

---

## 3. Class Diagram

Visualize object-oriented relationships.

### Syntax
```markdown
{{{!}}{ mermaid }}
classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal <|-- Zebra
    Animal : +int age
    Animal : +String gender
    Animal: +isMammal()
    Animal: +mate()
    class Duck{
        +String beakColor
        +swim()
        +quack()
    }
    class Fish{
        -int sizeInFeet
        -canEat()
    }
    class Zebra{
        +bool is_wild
        +run()
    }
{{{!}}{ /mermaid }}
```

### Result
{{ mermaid }}
classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal <|-- Zebra
    Animal : +int age
    Animal : +String gender
    Animal: +isMammal()
    Animal: +mate()
    class Duck{
        +String beakColor
        +swim()
        +quack()
    }
    class Fish{
        -int sizeInFeet
        -canEat()
    }
    class Zebra{
        +bool is_wild
        +run()
    }
{{ /mermaid }}

---

## 4. State Diagram

Visualize state transitions.

### Syntax
```markdown
{{{!}}{ mermaid }}
stateDiagram-v2
    [*] --> Still
    Still --> [*]
    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
{{{!}}{ /mermaid }}
```

### Result
{{ mermaid }}
stateDiagram-v2
    [*] --> Still
    Still --> [*]
    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
{{ /mermaid }}

---

## 5. Pie Chart

Visualize percentages.

### Syntax
```markdown
{{{!}}{ mermaid }}
pie title What specific languages do you use?
    "Go" : 40
    "Python" : 30
    "JavaScript" : 20
    "Rust" : 10
{{{!}}{ /mermaid }}
```

### Result
{{ mermaid }}
pie title What specific languages do you use?
    "Go" : 40
    "Python" : 30
    "JavaScript" : 20
    "Rust" : 10
{{ /mermaid }}

---

## 6. Entity Relationship Diagram (ERD)

Visualize database schemas.

### Syntax
```markdown
{{{!}}{ mermaid }}
erDiagram
    CUSTOMER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
    CUSTOMER }|..|{ DELIVERY-ADDRESS : uses
{{{!}}{ /mermaid }}
```

### Result
{{ mermaid }}
erDiagram
    CUSTOMER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
    CUSTOMER }|..|{ DELIVERY-ADDRESS : uses
{{ /mermaid }}

---

## 7. Developer Directives & Best Practices

When writing Mermaid diagrams, especially large architectural maps, observe the following rules to prevent syntax errors that will crash the compilation:

1. **Avoid Special Characters**: The Mermaid parser will instantly break if it encounters unescaped punctuation (`/`, `*`, `<`, `>`, `~`, `(`, `)`, `[`, `]`) inside node labels or sequence diagram messages.
2. **Quote Labels in Graphs**: When building a `graph TD`, **always** quote your node labels if they contain spaces or colons. E.g., `Node["My Label: Something"]`.
3. **Sequence Diagrams**: Strip method signatures of arguments and brackets. Use `Participant->>System: Execute` rather than `Participant->>System: Execute("page.mdt")`. Do NOT use slashes in participant names (e.g., use `html_template` instead of `html/template`).
4. **Class Diagrams**: Type definitions must be alphanumeric. Avoid using generic types with `~` (e.g. `map~string, interface~`) or array bracket notations (e.g. `[]string`). Instead, simplify types to alphanumeric structures like `Map` or `List`.
5. **No HTML Tags**: Never place HTML tags (`<br>`, `<strong>`) inside node labels unless you know the specific diagram type fully supports it (only a subset of Mermaid diagrams do).
6. **NO BLANK LINES**: Never place blank empty lines inside the `{{ mermaid }}` shortcode. The underlying shortcode parser injects a `<div class="mermaid">` block into the Markdown tree *before* Goldmark processes the text. If Goldmark encounters a blank line inside the HTML block, it will escape the HTML context and wrap your indented Mermaid code inside `<pre><code>` blocks, completely breaking the syntax.
