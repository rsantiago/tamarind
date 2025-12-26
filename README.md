# Tamarind

![Version](https://img.shields.io/badge/version-v1.0.0-blue.svg)
![Go](https://img.shields.io/badge/built_with-Go-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

> **Tamarind** is the **AI-First Web Engine**. 
> It transforms static content into a living, queryable, and distributed knowledge graph.

**Website**: [usetamarind.com](https://usetamarind.com)

---

## 🔮 The Vision

Static site generators were built for the "Human Web" — optimized for browsers and eyeballs.
**Tamarind** is built for the **Agent Web** — optimized for LLMs, RAG pipelines, and autonomous indexing.

We believe your website shouldn't just be a collection of HTML files. It should be:
1.  **A Knowledge Base**: Instantly queryable by AI.
2.  **A Social Node**: Connected to a decentralized decentralized federation of growers.
3.  **Self-Healing**: Capable of identifying gaps and fixing its own links.

---

## ✨ AI-Native Features

### 1. Build-Time Vectorization (Instant RAG)
Tamarind doesn't just compile HTML. It chunks your content, generates local embeddings, and outputs a `knowledge.db` alongside your site. 
*   **Benefits**: Your site is "RAG-ready" the moment it deploys. No external vector DB required.

### 2. The "Digital Gardener"
Tamarind acts as a gardener for your content.
*   **Self-Healing Links**: Identifies "orphan" concepts and suggests internal links (or auto-injects them).
*   **Synthetic Localization**: Uses LLMs to "transcreate" content for global audiences, not just translate it.

### 3. Agent-Native Schemas
Every interactive element on a Tamarind site exposes a structured semantic definition (MCP-style).
*   Visiting Agents understand "this is a Buy button" or "this is a Code snippet" systematically, not via DOM scraping.

---

## 🔌 The "Tamarind Garden" (Social Hub)

Tamarind sites are designed to be part of a greater whole. By enabling the **Garden Protocol**, your content is automatically aggregated into **Tamarind Garden** — a decentralized hub for content growers.

> "Plant your seed, grow your forest, join the garden."

---

## 🚀 Quick Start

Get your garden growing in seconds:

```bash
# 1. Install & Init
git clone https://github.com/rsantiago/tamarind.git
cd tamarind
go build -o tamarind parser/main.go
./tamarind init

# 2. Build & Serve (Watch Mode)
./tamarind serve -watch -theme pastel
```

---

## 📂 Project Structure

```text
my-garden/
├── tamarind            # The engine
├── tamarind.yaml       # Configuration (AI settings, Theme)
├── writer-sandbox/     # Your Content (The Soil)
│   ├── index.md
│   ├── articles/
│   ├── docs/
│   └── knowledge/      # Agent-specific context
└── website/            # The Result (The Fruit)
    ├── index.html
    ├── llms.txt        # Agent Context
    └── knowledge.db    # Vector Embeddings
```

---

## 🤝 Contributing

We welcome "Growers" to help us build the infrastructure of the AI Web.
Check `AGENTS.md` for our AI-assisted contribution guidelines.

**License**: MIT
