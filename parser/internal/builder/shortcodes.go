package builder

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func processShortcodes(markdown, sourceDir string) string {
	// 0. Agent (Comment): {{ agent "instruction" }} -> Removed from output
	reAgent := regexp.MustCompile(`{{\s*agent\s+"(.*?)"\s*}}`)
	markdown = reAgent.ReplaceAllString(markdown, "")

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

	// Figure: {{ figure src="url" caption="text" width="50%" }}
	reFig := regexp.MustCompile(`{{\s*figure\s+(.*?)\s*}}`)
	markdown = reFig.ReplaceAllStringFunc(markdown, func(match string) string {
		// Parse attributes manually for flexibility
		content := reFig.FindStringSubmatch(match)[1]
		
		reSrc := regexp.MustCompile(`src="([^"]+)"`)
		reCap := regexp.MustCompile(`caption="([^"]+)"`)
		reWidth := regexp.MustCompile(`width="([^"]+)"`)
		
		srcMatch := reSrc.FindStringSubmatch(content)
		if srcMatch == nil {
			return match // Invalid, no src
		}
		src := srcMatch[1]
		
		caption := ""
		capMatch := reCap.FindStringSubmatch(content)
		if capMatch != nil {
			caption = capMatch[1]
		}
		
		width := ""
		widthMatch := reWidth.FindStringSubmatch(content)
		if widthMatch != nil {
			width = widthMatch[1] // e.g. "500px" or "50%"
		}
		
		var figcaptionHTML string
		if caption != "" {
			figcaptionHTML = fmt.Sprintf("<figcaption>%s</figcaption>", caption)
		}
		
		// Style for width
		styleAttr := ""
		if width != "" {
			styleAttr = fmt.Sprintf(` style="width: %s; margin: 0 auto; display: block;"`, width)
		}

		// Responsive Logic
		// We assume src matches the generated optimized files: name-width.ext
		// AND that it is a local resource (not http/s)
		if !strings.HasPrefix(src, "http") {
			// Resolve path to check dimensions
			localPath := filepath.Join(sourceDir, src)

			// Check if file exists and get dimensions
			f, err := os.Open(localPath)
			if err == nil {
				defer f.Close()
				cfg, _, err := image.DecodeConfig(f)
				if err == nil {
					imgWidth := cfg.Width
					ext := filepath.Ext(src)
					base := strings.TrimSuffix(src, ext)
					
					lowerExt := strings.ToLower(ext)
					if lowerExt == ".jpg" || lowerExt == ".jpeg" || lowerExt == ".png" {
						var sources []string
						breakpoints := []int{480, 800, 1200}
						
						for _, bp := range breakpoints {
							if imgWidth >= bp {
								sources = append(sources, fmt.Sprintf("%s-%dw%s %dw", base, bp, ext, bp))
							}
						}
						
						if len(sources) > 0 {
							srcset := strings.Join(sources, ", ")
							sizes := "(max-width: 480px) 100vw, (max-width: 800px) 100vw, 100vw"
							return fmt.Sprintf(`<figure><img src="%s" srcset="%s" sizes="%s" alt="%s"%s>%s</figure>`,
								src, srcset, sizes, caption, styleAttr, figcaptionHTML)
						}
					}
				}
			}
		}
		
		// Fallback
		return fmt.Sprintf(`<figure><img src="%s" alt="%s"%s>%s</figure>`, src, caption, styleAttr, figcaptionHTML)
	})

	// Cleanup Escape Token {{!}}
	markdown = strings.ReplaceAll(markdown, "{{!}}", "")

	return markdown
}
