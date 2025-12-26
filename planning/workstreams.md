# Product Strategy: The AI-First Web Engine

**Vision**: Shift from being "just another static site generator" to the **infrastructure layer for the AI-readable web**.
**Metaphor**: If WordPress was the engine for the "Human Web", Tamarind is the engine for the "Agent Web".

## Strategic Pillar 1: The Enterprise RAG "Source of Truth"
**Value Proposition**: Companies spend millions on internal knowledge management (Notion/Confluence) that is useless for AI agents due to noise and lack of structure. Tamarind safeguards and structures this data.
*   **The 10x Feature**: **Build-Time Vectorization**.
    *   As Tamarind builds the HTML, it also chuncks content, generates embeddings (local embeddings), and outputs a pre-indexed vector database file (sqlite-vss or similar).
    *   **Result**: "Instant RAG". You deploy the website, and you automatically deploy a queryable brain without a complex data pipeline.
*   **Feature**: **Semantic ACLs**.
    *   Tag content chunks as `public`, `internal`, or `confidential`. Generate different versions of the AI context (`llms.txt`) based on the user's/agent's clearance level.

## Strategic Pillar 2: The Autonomous Content Engine (Programmatic SEO)
**Value Proposition**: Content marketing is the highest ROI channel but requires armies of writers. Tamarind becomes a "Gardener" that grows the site, not just builds it.
*   **The 10x Feature**: **"Self-Healing" Content graph**.
    *   Tamarind analyzes the `content/` folder. It identifies "orphan" concepts (topics mentioned but not linked). It uses an LLM to auto-inject internal links and suggest new "stub" articles to fill knowledge gaps.
*   **Feature**: **Synthetic Localization**.
    *   Instead of simple translation, use LLMs to "transcreate" content for different cultural contexts at build time, generating massive long-tail global SEO traffic.

## Strategic Pillar 3: The Federalist Network (Social Hub)
**Value Proposition**: Static sites are lonely islands. Tamarind connects them into an archipelago.
*   **The 10x Feature**: **The Tamarind Hub (Aggregation Protocol)**.
    *   Every Tamarind site exposes a standardized `content.json` or `social.rss` feed.
    *   The central `tamarind.garden` (or similar) website ingests these feeds.
    *   It serves as a "Hacker News" or "Medium" style aggregator exclusively for Tamarind users, driving traffic back to their individual independent sites.
    *   **Incentive**: "Build with Tamarind, get instant distribution."

## Strategic Pillar 4: Agent-Native Interfaces (H2M - Human to Machine)
**Value Proposition**: The future web traffic is Agents, not humans. Websites are currently hostile to Agents. Tamarind sites are "Dual-Mode" by default.
*   **The 10x Feature**: **Structured Action Schemas**.
    *   Every interactive element (forms, buttons) in an Tamarind site auto-generates a corresponding Tool Definition in the `llms.txt`.
    *   Reviewing code? The "Copy" button isn't just JS; it's a semantic tool exposed to visiting Agents.
    *   E-commerce? The "Buy" button exposes a structured transactional schema.

## Execution Backlog (Prioritized by Value)

1.  **Local Embedding Pipeline** (Pillar 1)
    *   *Tech*: Integrate `langchain-go` or direct generic embedding calls during the build process.
    *   *Output*: `knowledge.db` (vector store) generated alongside `index.html`.

2.  **Semantic Link Injection** (Pillar 2)
    *   *Tech*: TF-IDF or Embedding-based similarity check during build.
    *   *Output*: Auto-inserted `[Link]` in markdown text where relevant concepts appear.

3.  **Agent Protocol Implementation (MCP equivalent for Static Sites)** (Pillar 3)
    *   *Tech*: Standardize a schema (JSON-LD extended) that describes *actions* available on the page.

4.  **Diff-Based Content "Refining"**
    *   *Tech*: A command `atman refine <file>` that sends the content to an LLM to fix grammar, tone, and clarity, then presents a `git diff` for the user to accept.

5.  **Data-Driven Page Generation**
    *   *Tech*: Take a dataset (e.g., "1000 SaaS companies"), feeds it to an LLM + Template, and generates 1000 unique, high-value landing pages (Programmatic SEO).

---
*Legacy Features (Deprioritized/Commodity)*:
*   *External Themes (Nice to have, not a differentiator)*
*   *Comments (outsource to 3rd party)*
*   *Admin Dashboard (Visuals are expensive, focus on data)*
