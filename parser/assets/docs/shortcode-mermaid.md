---
title: Mermaid Diagrams
tags: [features, demo, visualization]
description: Create flowcharts, sequence diagrams, and more using code.
---

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
