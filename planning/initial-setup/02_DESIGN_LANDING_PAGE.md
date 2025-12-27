# Agent Instructions: Landing Page Design

**Objective**: Create a high-impact landing page for Tamarind highlighting its "Agent-First" philosophy.

## Execution Steps

### 1. Design Landing Page
*   **Target File**: `writer-sandbox/index.md`
*   **Goal**: Replace the default content with a pitch for "The AI-First Web Engine".
*   **Key Features to Highlight**:
    *   Markdown-to-HTML + Context (llms.txt).
    *   Zero config, single binary.
*   **Visuals**: Use `{{ figure src="images/logo.png" width="500px" }}`.

### 2. Build & Verify
*   Build the site using the `midnight` theme (dark mode suits the "Agent/Cyber" aesthetic).
*   Command: `./tamarind/tamarind build -theme midnight`
*   Verify `website/index.html` exists.
