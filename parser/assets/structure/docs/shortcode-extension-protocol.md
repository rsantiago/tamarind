---
title: Shortcode Extension Protocol
---

# Shortcode Extension Protocol

The Tamarind Shortcode Extension Protocol allows developers to build custom shortcode processors in any programming language (Python, Node, Bash, Go) without needing to modify or recompile the core Tamarind engine. 

Instead of writing Go plugins and registering them in the core AST registry, you define an external executable that Tamarind invokes dynamically during the build process.

## 1. Why Create Shortcode Extensions?

The extension protocol opens up infinite possibilities for integrating external systems directly into your markdown without writing Go. Here are five powerful use cases:

1. **Dynamic Content Fetching**: Hit external REST APIs (like GitHub releases, stock prices, or live weather data) at build time and embed the fresh data directly into your pages.
2. **Build-Time LLM Generation**: Send a prompt to an AI provider (like OpenAI or Gemini) and inject dynamically generated summaries, text, or automated code reviews right into your markdown files.
3. **Database Lookups & Query Rendering**: Execute SQL queries against a local database (like SQLite or Postgres) and format the results into clean markdown tables, eliminating the need to manually update data.
4. **Custom Syntax & Math Rendering**: Pre-render complex syntaxes—like converting LaTeX to HTML with KaTeX, or passing proprietary domain-specific languages through a custom formatting script.
5. **Chart & Graphic Generators**: Pipe CSV data directly into Python (using matplotlib or pandas) or Node scripts to dynamically generate Base64 images or inline SVG charts at build time.

## 2. How It Works

During the initialization phase, Tamarind's `BuildPluginRegistry` parses the site's `tamarind.yaml` configuration file. If it finds an `extensions` dictionary, it registers a dynamic `ExtensionPlugin` for each key-value pair.

When the Tamarind parser encounters one of these custom shortcodes in your markdown, it:
1. Shells out to the mapped external executable using `os/exec`.
2. Passes the shortcode's **inner content** via Standard Input (`stdin`).
3. Passes the shortcode's **attributes** as command-line arguments to the script.
4. Captures the executable's Standard Output (`stdout`).
5. Replaces the entire shortcode block with the captured `stdout` string in the Markdown stream, allowing it to be safely rendered by Goldmark in the next phase.

## 3. Sequence Diagram

{{ mermaid }}
sequenceDiagram
    participant Writer as Markdown Author
    participant Engine as Tamarind Engine
    participant Registry as Plugin Registry
    participant Script as External Executable (Bash/Python)
    participant Goldmark as HTML Renderer
    Engine->>Registry: Load tamarind.yaml
    Registry-->>Engine: Find "extensions" map
    Engine->>Registry: Register ExtensionPlugin for each shortcode
    Writer->>Engine: Write {{ custom }} content {{ /custom }}
    Engine->>Registry: ProcessShortcodes()
    Registry->>Script: execute(cmd, args=[attributes], stdin=content)
    Note over Script: Processes stdin & arguments
    Script-->>Registry: returns stdout (HTML or Markdown)
    Registry-->>Engine: Inject output into markdown stream
    Engine->>Goldmark: Convert to HTML
{{ /mermaid }}


## 4. Data Structures

The extension protocol relies on the `tamarind.yaml` mapping and the internal `ExtensionPlugin` struct.

### 4.1 The configuration mapping
In your site's `tamarind.yaml` (located at the root of your project), you define an `extensions` map where the key is the shortcode name, and the value is the relative or absolute path to the executable script.

```yaml
# tamarind.yaml
extensions:
  my_ext: "../scripts/my_ext.sh"
  python_widget: "python3 ../scripts/widget.py"
```

### 3.2 The Go Struct (`plugin_extension.go`)
Internally, Tamarind registers an `ExtensionPlugin` struct that satisfies the `ShortcodePlugin` interface.

```go
type ExtensionPlugin struct {
	name    string
	cmdPath string
}

// Implements Process(match []string, sourceDir string) (string, error)
// Uses os/exec to route match[2] (content) to stdin and match[1] (attributes) to args.
```

### 4.2 Handling attributes and input
The internal `ExtensionPlugin` parses the Shortcode regex, capturing the attributes as `match[1]` and the content as `match[2]`. It pipes the content into the external executable's `stdin` via `strings.NewReader` and passes the attributes in `cmd.Args`.

## 5. How to Declare a New Extension

### Step 5.1: Write the external script
Write a script in any language. It must be executable and capable of reading standard input.

**Example (Bash): `scripts/alert.sh`**
```bash
#!/bin/bash

# Read the inner content from standard input
content=$(cat)

# Read attributes from the first argument (e.g., type="warning")
attributes=$1

# Output HTML directly
echo "<div class=\"custom-alert\" data-attrs=\"$attributes\">"
echo "  <p>$content</p>"
echo "</div>"
```

### Step 5.2: Make it executable
Ensure your script has executable permissions:
```bash
chmod +x scripts/alert.sh
```

### Step 5.3: Register in tamarind.yaml
Add the `extensions` dictionary to your `tamarind.yaml`:

```yaml
extensions:
  custom_alert: "scripts/alert.sh"
```

## 6. How to Use Your Extension

Once registered, you can use the shortcode exactly like any native Tamarind shortcode in your Markdown files.

```markdown
{{{{!}}custom_alert type="danger"}}
This is the inner content that gets passed to stdin!
{{{{!}}/custom_alert}}
```

**Execution Trace:**
When building, Tamarind will execute:
```bash
./scripts/alert.sh 'type="danger"' < (echo "This is the inner content that gets passed to stdin!")
```
And replace the markdown block with the script's `stdout`.
