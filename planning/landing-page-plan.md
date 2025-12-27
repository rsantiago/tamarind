# Tamarind Landing Page Implementation Plan

## 1. Workspace Configuration
**Objective**: Isolate the build environment to allow agents to work non-destructively.

*   **Root Directory**: `/home/rsantiago/Documents/atman-multi-agents/execution/landing-page-task`
*   **Repo Path**: `/home/rsantiago/Documents/atman-multi-agents/execution/landing-page-task/tamarind`
*   **Site Path**: `/home/rsantiago/Documents/atman-multi-agents/execution/landing-page-task/site`

### Setup Instructions (Agent Workflow)
1.  Create the directory: `mkdir -p execution/landing-page-task`
2.  Clone the repository: `git clone https://github.com/rsantiago/atman-multi-agents execution/landing-page-task/tamarind` (Note: Use local path if remote not available, or copy) -> *Correction*: Since we are local, we should `cp -r tamarind execution/landing-page-task/` to simulate a clone or use `git clone` if the user prefers. The user said "checkout tamarind", implying git. I'll assume local copy or git clone of the local repo. Let's use `cp -r` to ensure we get the *current* state including uncommitted changes if any, or `git clone` for cleanliness. I'll stick to `git clone` of the remote if possible, or usually just copy.
    *   *Refined Step*: Copy the current `tamarind` directory to `execution/landing-page-task/tamarind` to preserve current state.
3.  Build the binary: `cd execution/landing-page-task/tamarind/parser && go build -o ../tamarind`
4.  Initialize Site: `cd execution/landing-page-task && ./tamarind/tamarind init` (Creating `writer-sandbox`).

## 2. Landing Page Design
**Objective**: Showcase the "Agent-First" philosophy.

### Structure (`index.md`)
*   **Hero**: Strong headline about the AI/Agent web shift.
*   **Problem**: "Your site is invisible to LLMs."
*   **Solution**: Tamarind's dual-generation (HTML + Context).
*   **Visual**: Use the `figure` shortcode with the constrained width (500px).
*   **CTA**: "Get the Binary" / "Read the Docs".

### Theme Selection
*   **Primary Choice**: `midnight` (Dark, tech-focused) or `gram` (Clean, editorial).
*   **Customization**: Ensure `box-shadow` is off (already done globally).

## 3. Execution Steps
1.  **Prepare Workspace**: Run the setup script.
2.  **Draft Content**: Create the robust `index.md` emphasizing the USP.
3.  **Build & Verify**: Run `tamarind build` and check output.
