# Agent Instructions: Project Setup

**Objective**: Initialize the environment and project structure for the Tamarind Landing Page task.

## Execution Steps

### 1. Setup Environment
*   **Clone Repository**: Clone the remote Tamarind repository.
    *   Command: `git clone https://github.com/rsantiago/tamarind.git ./tamarind`
*   **Checkout Feature Branch**: Switch to the target feature branch.
    *   **Feature Branch Name**: `feature/landing-page`
    *   Command: `cd tamarind && git checkout -b feature/landing-page || git checkout feature/landing-page`
*   **Build Binary**: Compile the latest version of Tamarind.
    *   Command: `cd parser && go build -o ../tamarind` (Run this inside the `tamarind` folder, or `cd tamarind/parser && go build -o ../tamarind` from root).

### 2. Initialize Project
*   **Important**: Return to the root of your workspace (parent of `tamarind` folder).
*   Run: `./tamarind/tamarind init`
    *   This must create the `writer-sandbox` directory at the **root folder level**, NOT inside the `tamarind` cloned repo.
