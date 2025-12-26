package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func processShortcodes(markdown, sourceDir string) string {
	// 1. Mermaid (Block): {{ mermaid }}...{{ /mermaid }}
	reMermaid := regexp.MustCompile(`(?s){{\s*mermaid\s*}}(.*?){{\s*/mermaid\s*}}`)
	markdown = reMermaid.ReplaceAllString(markdown, `<div class="mermaid">$1</div>`)

	// 2. Callout (Block): {{ callout type="warn" title="Alert" }}...{{ /callout }}
	reCallout := regexp.MustCompile(`(?s){{\s*callout\s+type="([^"]+)"(?:\s+title="([^"]+)")?\s*}}(.*?){{\s*/callout\s*}}`)
	markdown = reCallout.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reCallout.FindStringSubmatch(match)
		typ := submatch[1]
		title := submatch[2]
		content := submatch[3]
		
		html := fmt.Sprintf(`<div class="callout callout-%s">`, typ)
		if title != "" {
			html += fmt.Sprintf(`<div class="callout-title">%s</div>`, title)
		}
		html += fmt.Sprintf(`<div class="callout-content">%s</div></div>`, content)
		return html
	})

	// 3. Terminal (Block): {{ terminal }}...{{ /terminal }}
	reTerm := regexp.MustCompile(`(?s){{\s*terminal\s*}}(.*?){{\s*/terminal\s*}}`)
	markdown = reTerm.ReplaceAllString(markdown, `<div class="terminal"><div class="terminal-header"><span class="dot red"></span><span class="dot yellow"></span><span class="dot green"></span></div><pre class="terminal-content"><code>$1</code></pre></div>`)

	// 4. Code Include: {{ include src="file.go" lines="1-10" lang="go" }}
	reInclude := regexp.MustCompile(`{{\s*include\s+src="([^"]+)"(?:\s+lines="([0-9]+-[0-9]+)")?(?:\s+lang="([^"]+)")?\s*}}`)
	markdown = reInclude.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reInclude.FindStringSubmatch(match)
		src := parts[1]
		linesRange := parts[2]
		lang := parts[3]
		if lang == "" { lang = "text" }
		
		var content []byte
		var err error

		if strings.HasPrefix(src, "http") {
			resp, err := http.Get(src)
			if err != nil {
				return fmt.Sprintf("> **Error fetching %s**: %v", src, err)
			}
			defer resp.Body.Close()
			content, err = io.ReadAll(resp.Body)
		} else {
			path := filepath.Join(sourceDir, src)
			content, err = os.ReadFile(path)
		}

		if err != nil {
			return fmt.Sprintf("> **Error including %s**: %v", src, err)
		}
		
		finalContent := string(content)
		if linesRange != "" {
			lParts := strings.Split(linesRange, "-")
			if len(lParts) == 2 {
				start, _ := strconv.Atoi(lParts[0])
				end, _ := strconv.Atoi(lParts[1])
				
				lines := strings.Split(finalContent, "\n")
				// Validate bounds (1-based -> 0-based)
				if start < 1 { start = 1 }
				if end > len(lines) { end = len(lines) }
				if start <= end {
					finalContent = strings.Join(lines[start-1:end], "\n")
				}
			}
		}

		return fmt.Sprintf("```%s\n%s\n```", lang, finalContent)
	})

	// 5. GitHub Gist: {{ gist id="123" }}
	reGist := regexp.MustCompile(`{{\s*gist\s+id="([^"]+)"\s*}}`)
	markdown = reGist.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reGist.FindStringSubmatch(match)
		gistID := parts[1]

		resp, err := http.Get(fmt.Sprintf("https://gist.github.com/%s.json", gistID))
		if err != nil {
			return fmt.Sprintf("> **Error loading gist %s**: %v", gistID, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Sprintf("> **Error loading gist %s**: Status %d", gistID, resp.StatusCode)
		}

		var data struct {
			Div        string `json:"div"`
			Stylesheet string `json:"stylesheet"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return fmt.Sprintf("> **Error decoding gist %s**: %v", gistID, err)
		}

		// Strip newlines to prevent Markdown from interpreting indented HTML as code blocks
		data.Div = strings.ReplaceAll(data.Div, "\n", "")
		data.Div = strings.ReplaceAll(data.Div, "\r", "")

		return fmt.Sprintf(`<link rel="stylesheet" href="%s">%s`, data.Stylesheet, data.Div)
	})

	// 6. Math: {{ math }}...{{ /math }}
	reMath := regexp.MustCompile(`(?s){{\s*math\s*}}(.*?){{\s*/math\s*}}`)
	markdown = reMath.ReplaceAllString(markdown, `<div class="math-block">$$$1$$</div>`)

	// YouTube: {{ youtube ID }}
	reYT := regexp.MustCompile(`{{\s*youtube\s+([a-zA-Z0-9_-]+)\s*}}`)
	markdown = reYT.ReplaceAllString(markdown, `<div class="video-container"><iframe src="https://www.youtube.com/embed/$1" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></div>`)

	// Figure: {{ figure src="url" caption="text" }}
	reFig := regexp.MustCompile(`{{\s*figure\s+src="([^"]+)"\s+caption="([^"]+)"\s*}}`)
	markdown = reFig.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reFig.FindStringSubmatch(match)
		src := parts[1]
		caption := parts[2]
		
		// Responsive Logic
		// We assume src matches the generated optimized files: name-width.ext
		// AND that it is a local resource (not http/s)
		if !strings.HasPrefix(src, "http") {
			ext := filepath.Ext(src)
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				base := strings.TrimSuffix(src, ext)
				
			// Construct srcset
				// Example: path/to/img-480w.jpg 480w, ...
				srcset := fmt.Sprintf("%s-480w%s 480w, %s-800w%s 800w, %s-1200w%s 1200w", 
					base, ext, base, ext, base, ext)
				
				// sizes: mobile/tablet = 100vw, desktop = capped at 1200px (or typically content width)
				sizes := "(max-width: 480px) 100vw, (max-width: 800px) 100vw, 100vw"

				return fmt.Sprintf(`<figure><img src="%s" srcset="%s" sizes="%s" alt="%s"><figcaption>%s</figcaption></figure>`,
					src, srcset, sizes, caption, caption)
			}
		}
		
		// Fallback for gifs/svgs/external
		return fmt.Sprintf(`<figure><img src="%s" alt="%s"><figcaption>%s</figcaption></figure>`, src, caption, caption)
	})

	// Cleanup Escape Token {{!}}
	markdown = strings.ReplaceAll(markdown, "{{!}}", "")

	return markdown
}
