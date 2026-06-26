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

#### Syntax
```markdown
{{{!}}{ mermaid }}
graph LR;
    A[Start] --> B{Is it working?};
    B -- Yes --> C[Great!];
    B -- No --> D[Debug];
    D --> B;
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
graph LR;
    A[Start] --> B{Is it working?};
    B -- Yes --> C[Great!];
    B -- No --> D[Debug];
    D --> B;
{{ /mermaid }}


### Sequence Diagram
Visualize interactions and timelines between systems or objects.

#### Syntax
```markdown
{{{!}}{ mermaid }}
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
{{ /mermaid }}


### Class Diagram
Visualize object-oriented relationships and attributes.

#### Syntax
```markdown
{{{!}}{ mermaid }}
classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal : +int age
    Animal: +isMammal()
    class Duck{
        +String beakColor
        +swim()
    }
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal : +int age
    Animal: +isMammal()
    class Duck{
        +String beakColor
        +swim()
    }
{{ /mermaid }}


### State Diagram
Visualize state transitions of a system over time.

#### Syntax
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

#### Result
{{ mermaid }}
stateDiagram-v2
    [*] --> Still
    Still --> [*]
    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
{{ /mermaid }}


### Entity Relationship Diagram (ERD)
Visualize database schemas and relationships.

#### Syntax
```markdown
{{{!}}{ mermaid }}
erDiagram
    CUSTOMER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
erDiagram
    CUSTOMER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
{{ /mermaid }}

---

## 2. Project Management & Planning

These tools help teams plan schedules, map user flows, and brainstorm hierarchies.

### Mindmap
For brainstorming and organizing hierarchical information visually.

#### Syntax
```markdown
{{{!}}{ mermaid }}
mindmap
  root((Tamarind))
    Design
      Themes
      CSS
    Performance
      Go Binary
      Zero Dependencies
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
mindmap
  root((Tamarind))
    Design
      Themes
      CSS
    Performance
      Go Binary
      Zero Dependencies
{{ /mermaid }}

### Gantt Chart
For project schedules, tracking task durations and dependencies.

#### Syntax
```markdown
{{{!}}{ mermaid }}
gantt
    title A Gantt Diagram
    dateFormat  YYYY-MM-DD
    section Section
    A task           :a1, 2026-01-01, 30d
    Another task     :after a1  , 20d
    section Another
    Task in sec      :2026-01-12  , 12d
    another task      : 24d
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
gantt
    title A Gantt Diagram
    dateFormat  YYYY-MM-DD
    section Section
    A task           :a1, 2026-01-01, 30d
    Another task     :after a1  , 20d
    section Another
    Task in sec      :2026-01-12  , 12d
    another task      : 24d
{{ /mermaid }}


### User Journey
For mapping a user's emotional and practical journey through a product.

#### Syntax
```markdown
{{{!}}{ mermaid }}
journey
    title My working day
    section Go to work
      Make tea: 5: Me
      Go upstairs: 3: Me
      Do work: 1: Me, Cat
    section Go home
      Go downstairs: 5: Me
      Sit down: 5: Me
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
journey
    title My working day
    section Go to work
      Make tea: 5: Me
      Go upstairs: 3: Me
      Do work: 1: Me, Cat
    section Go home
      Go downstairs: 5: Me
      Sit down: 5: Me
{{ /mermaid }}


### Timeline
For plotting chronological events.

#### Syntax
```markdown
{{{!}}{ mermaid }}
timeline
    title History of Social Media
    2002 : LinkedIn
    2004 : Facebook : Orkut
    2005 : YouTube
    2006 : Twitter
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
timeline
    title History of Social Media
    2002 : LinkedIn
    2004 : Facebook : Orkut
    2005 : YouTube
    2006 : Twitter
{{ /mermaid }}


### Requirement Diagram
For mapping out software requirements and test traces.

#### Syntax
```markdown
{{{!}}{ mermaid }}
requirementDiagram
    requirement test_req {
    id: 1
    text: the test text.
    risk: high
    verifymethod: test
    }
    element test_entity {
    type: simulation
    }
    test_entity - satisfies -> test_req
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
requirementDiagram
    requirement test_req {
    id: 1
    text: the test text.
    risk: high
    verifymethod: test
    }
    element test_entity {
    type: simulation
    }
    test_entity - satisfies -> test_req
{{ /mermaid }}

---

## 3. Specialized & Advanced Visualizations

### Pie Chart
Visualize percentage distributions.

#### Syntax
```markdown
{{{!}}{ mermaid }}
pie title What specific languages do you use?
    "Go" : 40
    "Python" : 30
    "JavaScript" : 20
    "Rust" : 10
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
pie title What specific languages do you use?
    "Go" : 40
    "Python" : 30
    "JavaScript" : 20
    "Rust" : 10
{{ /mermaid }}

### Quadrant Chart
Plot items on an X/Y axis with 4 quadrants.

#### Syntax
```markdown
{{{!}}{ mermaid }}
quadrantChart
    title Reach and engagement of campaigns
    x-axis Low Reach --> High Reach
    y-axis Low Engagement --> High Engagement
    quadrant-1 We should expand
    quadrant-2 Need to promote
    quadrant-3 Re-evaluate
    quadrant-4 May be improved
    Campaign A: [0.3, 0.6]
    Campaign B: [0.45, 0.23]
    Campaign C: [0.57, 0.69]
    Campaign D: [0.78, 0.34]
    Campaign E: [0.40, 0.34]
    Campaign F: [0.35, 0.78]
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
quadrantChart
    title Reach and engagement of campaigns
    x-axis Low Reach --> High Reach
    y-axis Low Engagement --> High Engagement
    quadrant-1 We should expand
    quadrant-2 Need to promote
    quadrant-3 Re-evaluate
    quadrant-4 May be improved
    Campaign A: [0.3, 0.6]
    Campaign B: [0.45, 0.23]
    Campaign C: [0.57, 0.69]
    Campaign D: [0.78, 0.34]
    Campaign E: [0.40, 0.34]
    Campaign F: [0.35, 0.78]
{{ /mermaid }}


### Gitgraph
Visualize git branch histories, commits, and merges.

#### Syntax
```markdown
{{{!}}{ mermaid }}
gitGraph
    commit
    commit
    branch develop
    checkout develop
    commit
    commit
    checkout main
    merge develop
    commit
    commit
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
gitGraph
    commit
    commit
    branch develop
    checkout develop
    commit
    commit
    checkout main
    merge develop
    commit
    commit
{{ /mermaid }}

### C4 Architecture
Define software architecture using the C4 model.

#### Syntax
```markdown
{{{!}}{ mermaid }}
C4Context
    title System Context diagram for Internet Banking System
    Person(customerA, "Banking Customer A", "A customer of the bank, with personal bank accounts.")
    Person(customerB, "Banking Customer B")
    Person_Ext(customerC, "Banking Customer C", "desc")
    System(SystemAA, "Internet Banking System", "Allows customers to view information about their bank accounts, and make payments.")
    Rel(customerA, SystemAA, "Uses")
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
C4Context
    title System Context diagram for Internet Banking System
    Person(customerA, "Banking Customer A", "A customer of the bank, with personal bank accounts.")
    Person(customerB, "Banking Customer B")
    Person_Ext(customerC, "Banking Customer C", "desc")
    System(SystemAA, "Internet Banking System", "Allows customers to view information about their bank accounts, and make payments.")
    Rel(customerA, SystemAA, "Uses")
{{ /mermaid }}

---

## 4. Newer/Experimental Diagrams (Beta)

Mermaid is constantly adding new beta diagrams. All of these work natively via the `{{ mermaid }}` shortcode.

### Sankey Diagram
For flow rates between states.

#### Syntax
```markdown
{{{!}}{ mermaid }}
sankey-beta
    Bio-conversion, Liquid, 0.599
    Bio-conversion, Solid, 0.277
    Liquid, Instruments, 0.599
    Solid, Instruments, 0.277
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
sankey-beta
    Bio-conversion, Liquid, 0.599
    Bio-conversion, Solid, 0.277
    Liquid, Instruments, 0.599
    Solid, Instruments, 0.277
{{ /mermaid }}

### XY Chart
For standard line and bar graphs.

#### Syntax
```markdown
{{{!}}{ mermaid }}
xychart-beta
    title "Sales Revenue"
    x-axis [jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec]
    y-axis "Revenue (in $)" 4000 --> 11000
    bar [5000, 6000, 7500, 8200, 9500, 10500, 11000, 10200, 9200, 8500, 7000, 6000]
    line [5000, 6000, 7500, 8200, 9500, 10500, 11000, 10200, 9200, 8500, 7000, 6000]
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
xychart-beta
    title "Sales Revenue"
    x-axis [jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec]
    y-axis "Revenue (in $)" 4000 --> 11000
    bar [5000, 6000, 7500, 8200, 9500, 10500, 11000, 10200, 9200, 8500, 7000, 6000]
    line [5000, 6000, 7500, 8200, 9500, 10500, 11000, 10200, 9200, 8500, 7000, 6000]
{{ /mermaid }}

### Block Diagram
For spatial block-based layouts.

#### Syntax
```markdown
{{{!}}{ mermaid }}
block-beta
    columns 1
    db(("DB"))
    blockArrowId6<["&nbsp;&nbsp;&nbsp;"]>(down)
    block:e:1
        l["Left"]
        m["Middle"]
        r["Right"]
    end
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
block-beta
    columns 1
    db(("DB"))
    blockArrowId6<["&nbsp;&nbsp;&nbsp;"]>(down)
    block:e:1
        l["Left"]
        m["Middle"]
        r["Right"]
    end
{{ /mermaid }}

### Architecture Diagram
For cloud/network topology architecture maps.

#### Syntax
```markdown
{{{!}}{ mermaid }}
architecture-beta
    group api(cloud)[API]
    service db(database)[Database] in api
    service disk(disk)[Storage] in api
    db:L -- R:disk
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
architecture-beta
    group api(cloud)[API]
    service db(database)[Database] in api
    service disk(disk)[Storage] in api
    db:L -- R:disk
{{ /mermaid }}

### Packet Diagram
For visualizing network packet bytes/bits headers.

#### Syntax
```markdown
{{{!}}{ mermaid }}
packet-beta
    0-15: "Source Port"
    16-31: "Destination Port"
    32-63: "Sequence Number"
    64-95: "Acknowledgment Number"
    96-99: "Data Offset"
    100-105: "Reserved"
    106-111: "Flags"
    112-127: "Window Size"
    128-143: "Checksum"
    144-159: "Urgent Pointer"
{{{!}}{ /mermaid }}
```

#### Result
{{ mermaid }}
packet-beta
    0-15: "Source Port"
    16-31: "Destination Port"
    32-63: "Sequence Number"
    64-95: "Acknowledgment Number"
    96-99: "Data Offset"
    100-105: "Reserved"
    106-111: "Flags"
    112-127: "Window Size"
    128-143: "Checksum"
    144-159: "Urgent Pointer"
{{ /mermaid }}

---

## 5. Developer Directives & Best Practices

When writing Mermaid diagrams, especially large architectural maps, observe the following rules to prevent syntax errors that will crash the compilation:

1. **Avoid Special Characters**: The Mermaid parser will instantly break if it encounters unescaped punctuation (`/`, `*`, `<`, `>`, `~`, `(`, `)`, `[`, `]`) inside node labels or sequence diagram messages.
2. **Quote Labels in Graphs**: When building a `graph TD`, **always** quote your node labels if they contain spaces or colons. E.g., `Node["My Label: Something"]`.
3. **Sequence Diagrams**: Strip method signatures of arguments and brackets. Use `Participant->>System: Execute` rather than `Participant->>System: Execute("page.mdt")`. Do NOT use slashes in participant names (e.g., use `html_template` instead of `html/template`).
4. **Class Diagrams**: Type definitions must be alphanumeric. Avoid using generic types with `~` (e.g. `map~string, interface~`) or array bracket notations (e.g. `[]string`). Instead, simplify types to alphanumeric structures like `Map` or `List`.
5. **No HTML Tags**: Never place HTML tags (`<br>`, `<strong>`) inside node labels unless you know the specific diagram type fully supports it (only a subset of Mermaid diagrams do).
6. **NO BLANK LINES**: Never place blank empty lines inside the `{{ mermaid }}` shortcode. The underlying shortcode parser injects a `<div class="mermaid">` block into the Markdown tree *before* Goldmark processes the text. If Goldmark encounters a blank line inside the HTML block, it will escape the HTML context and wrap your indented Mermaid code inside `<pre><code>` blocks, completely breaking the syntax.
