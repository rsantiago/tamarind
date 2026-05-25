// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

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

	// Metrics Grid: {{ metrics }} ... {{ /metrics }}
	reMetrics := regexp.MustCompile(`(?s){{\s*metrics\s*}}(.*?){{\s*/metrics\s*}}`)
	markdown = reMetrics.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reMetrics.FindStringSubmatch(match)
		content := submatch[1]

		reMetric := regexp.MustCompile(`{{\s*metric\s+value="([^"]+)"\s+label="([^"]+)"\s*}}`)
		itemsHtml := ""
		
		metricMatches := reMetric.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range metricMatches {
			val := itemSubmatch[1]
			lbl := itemSubmatch[2]
			itemsHtml += fmt.Sprintf(`<div class="metric-card"><div class="metric-value">%s</div><div class="metric-label">%s</div></div>`, val, lbl)
		}

		return fmt.Sprintf(`<div class="metrics-grid">%s</div>`, itemsHtml)
	})

	// Features Grid: {{ features }} ... {{ /features }}
	reFeatures := regexp.MustCompile(`(?s){{\s*features\s*}}(.*?){{\s*/features\s*}}`)
	markdown = reFeatures.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reFeatures.FindStringSubmatch(match)
		content := submatch[1]

		reFeature := regexp.MustCompile(`(?s){{\s*feature\s+title="([^"]+)"\s+gradient="([^"]+)"\s+icon="([^"]+)"\s*}}(.*?){{\s*/feature\s*}}`)
		itemsHtml := ""
		
		featureMatches := reFeature.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range featureMatches {
			title := itemSubmatch[1]
			grad := itemSubmatch[2]
			iconName := itemSubmatch[3]
			desc := itemSubmatch[4]
			
			svgIcon := getBuiltInIconSvg(iconName)

			itemsHtml += fmt.Sprintf(`<div class="feature-card"><div class="feature-icon-box gradient-%s">%s</div><h3 class="feature-title">%s</h3><p class="feature-desc">%s</p></div>`, grad, svgIcon, title, strings.TrimSpace(desc))
		}

		return fmt.Sprintf(`<div class="features-grid">%s</div>`, itemsHtml)
	})

	// Capabilities Grid: {{ capabilities-grid }} ... {{ /capabilities-grid }}
	reCapsGrid := regexp.MustCompile(`(?s){{\s*capabilities-grid\s*}}(.*?){{\s*/capabilities-grid\s*}}`)
	markdown = reCapsGrid.ReplaceAllString(markdown, `<div class="capabilities-grid">$1</div>`)

	// Capabilities Checklist Card: {{ capabilities }} ... {{ /capabilities }}
	reCapabilities := regexp.MustCompile(`(?s){{\s*capabilities(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/capabilities\s*}}`)
	markdown = reCapabilities.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reCapabilities.FindStringSubmatch(match)
		title := submatch[1]
		content := submatch[2]

		reCapability := regexp.MustCompile(`{{\s*capability\s+name="([^"]+)"\s+desc="([^"]+)"\s+status="([^"]+)"(?:\s+statusLabel="([^"]*)")?\s*}}`)
		rowsHtml := ""
		
		capMatches := reCapability.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range capMatches {
			name := itemSubmatch[1]
			desc := itemSubmatch[2]
			status := itemSubmatch[3]
			statusLabel := status
			if len(itemSubmatch) > 4 && itemSubmatch[4] != "" {
				statusLabel = itemSubmatch[4]
			}

			rowsHtml += fmt.Sprintf(`<div class="capability-row"><div class="capability-info"><span class="capability-name">%s</span><span class="capability-desc">%s</span></div><span class="capability-status status-%s">%s</span></div>`, name, desc, status, statusLabel)
		}

		headerHtml := ""
		if title != "" {
			headerHtml = fmt.Sprintf(`<div class="capability-header"><div class="capability-card-title">%s</div></div>`, title)
		}

		return fmt.Sprintf(`<div class="capability-card">%s%s</div>`, headerHtml, rowsHtml)
	})

	// Vertical Timeline: {{ timeline }} ... {{ /timeline }}
	reTimeline := regexp.MustCompile(`(?s){{\s*timeline\s*}}(.*?){{\s*/timeline\s*}}`)
	markdown = reTimeline.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reTimeline.FindStringSubmatch(match)
		content := submatch[1]

		reItem := regexp.MustCompile(`(?s){{\s*item\s+title="([^"]+)"(?:\s+number="([^"]*)")?\s*}}(.*?){{\s*/item\s*}}`)
		itemsHtml := ""
		
		itemMatches := reItem.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range itemMatches {
			title := itemSubmatch[1]
			num := itemSubmatch[2]
			desc := itemSubmatch[3]

			badgeHtml := ""
			if num != "" {
				badgeHtml = fmt.Sprintf(`<div class="timeline-badge"><span class="timeline-badge-number">%s</span></div>`, num)
			} else {
				badgeHtml = `<div class="timeline-badge"></div>`
			}

			itemsHtml += fmt.Sprintf(`<div class="timeline-item">%s<div class="timeline-content"><h3 class="timeline-title">%s</h3><p class="timeline-desc">%s</p></div></div>`, badgeHtml, title, strings.TrimSpace(desc))
		}

		return fmt.Sprintf(`<div class="timeline-container">%s</div>`, itemsHtml)
	})

	// Dropdown Selection: {{ dropdown id="X" label="Y" }} ... {{ /dropdown }}
	reDropdown := regexp.MustCompile(`(?s){{\s*dropdown\s+(.*?)\s*}}(.*?){{\s*/dropdown\s*}}`)
	markdown = reDropdown.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reDropdown.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reId := regexp.MustCompile(`id="([^"]+)"`)
		reLabel := regexp.MustCompile(`label="([^"]+)"`)

		id := ""
		if m := reId.FindStringSubmatch(attrs); len(m) > 1 {
			id = m[1]
		}
		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}

		reOption := regexp.MustCompile(`(?s){{\s*option\s+value="([^"]+)"(?:\s+selected="([^"]*)")?\s*}}(.*?){{\s*/option\s*}}`)
		optionsHtml := ""
		
		optMatches := reOption.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range optMatches {
			val := itemSubmatch[1]
			sel := itemSubmatch[2]
			text := itemSubmatch[3]

			selectedAttr := ""
			if sel == "true" {
				selectedAttr = " selected"
			}

			optionsHtml += fmt.Sprintf(`<option value="%s"%s>%s</option>`, val, selectedAttr, strings.TrimSpace(text))
		}

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="tamarind-select-label">%s</label>`, label)
		}

		idAttr := ""
		if id != "" {
			idAttr = fmt.Sprintf(` id="%s"`, id)
		}

		chevronSvg := getBuiltInIconSvg("chevron")

		return fmt.Sprintf(`<div class="tamarind-select-wrapper">%s<div class="tamarind-select-control"><select class="tamarind-select"%s>%s</select><div class="tamarind-select-chevron">%s</div></div></div>`, labelHtml, idAttr, optionsHtml, chevronSvg)
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
		if lang == "" {
			lang = "text"
		}

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
				if start < 1 {
					start = 1
				}
				if end > len(lines) {
					end = len(lines)
				}
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

	// 7. Button (Block): {{ button href="url" ... }}Label{{ /button }}
	reButton := regexp.MustCompile(`(?s){{<?\s*button\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*button\s*>??}}`)
	markdown = reButton.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reButton.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reHref := regexp.MustCompile(`href="([^"]+)"`)
		reType := regexp.MustCompile(`type="([^"]+)"`)
		reSize := regexp.MustCompile(`size="([^"]+)"`)
		reTarget := regexp.MustCompile(`target="([^"]+)"`)

		href := ""
		if m := reHref.FindStringSubmatch(attrs); len(m) > 1 {
			href = m[1]
		}
		typ := "primary"
		if m := reType.FindStringSubmatch(attrs); len(m) > 1 {
			typ = m[1]
		}
		size := ""
		if m := reSize.FindStringSubmatch(attrs); len(m) > 1 {
			size = m[1]
		}
		target := ""
		if m := reTarget.FindStringSubmatch(attrs); len(m) > 1 {
			target = m[1]
		}

		classAttr := "btn"
		if typ != "" {
			classAttr += " btn-" + typ
		}
		if size != "" {
			classAttr += " btn-" + size
		}

		targetAttr := ""
		if target != "" {
			targetAttr = fmt.Sprintf(` target="%s"`, target)
		}

		return fmt.Sprintf(`<a href="%s" class="%s"%s>%s</a>`, href, classAttr, targetAttr, content)
	})

	// 8. Card (Block): {{ card }}...{{ /card }}
	reCard := regexp.MustCompile(`(?s){{<?\s*card(?:\s+padding="([^"]*)")?\s*>??}}(.*?){{<?\s*/?\s*card\s*>??}}`)
	markdown = reCard.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reCard.FindStringSubmatch(match)
		paddingVal := parts[1]
		content := parts[2]

		paddingClass := " card-padding"
		if paddingVal == "false" {
			paddingClass = ""
		}
		return fmt.Sprintf(`<div class="card%s">%s</div>`, paddingClass, content)
	})

	// 9. Alert (Block): {{ alert type="info" title="Title" }}...{{ /alert }} / {{ callout }}
	reAlert := regexp.MustCompile(`(?s){{<?\s*(?:alert|callout)\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*(?:alert|callout)\s*>??}}`)
	markdown = reAlert.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reAlert.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reType := regexp.MustCompile(`type="([^"]+)"`)
		reTitle := regexp.MustCompile(`title="([^"]+)"`)

		typ := "info"
		if m := reType.FindStringSubmatch(attrs); len(m) > 1 {
			typ = m[1]
		}
		title := ""
		if m := reTitle.FindStringSubmatch(attrs); len(m) > 1 {
			title = m[1]
		}

		iconName := typ
		if typ == "danger" || typ == "error" || typ == "warn" {
			iconName = "warning"
		} else if typ == "sparkles" || typ == "magic" || typ == "special" {
			iconName = "sparkles"
		} else if typ == "success" {
			iconName = "success"
		} else {
			iconName = "info"
		}

		svgIcon := getBuiltInIconSvg(iconName)

		titleHtml := ""
		if title != "" {
			titleHtml = fmt.Sprintf(`<h4 class="alert-title">%s</h4>`, title)
		}

		return fmt.Sprintf(`<div class="alert-container alert-%s"><div class="alert-icon-box">%s</div><div class="alert-content">%s<p class="alert-message">%s</p></div></div>`, typ, svgIcon, titleHtml, strings.TrimSpace(content))
	})

	// 10. Badge (Block): {{ badge type="primary" }}...{{ /badge }}
	reBadge := regexp.MustCompile(`(?s){{<?\s*badge(?:\s+type="([^"]*)")?\s*>??}}(.*?){{<?\s*/?\s*badge\s*>??}}`)
	markdown = reBadge.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reBadge.FindStringSubmatch(match)
		typ := parts[1]
		content := parts[2]

		classAttr := "badge"
		if typ != "" {
			classAttr += " badge-" + typ
		}
		return fmt.Sprintf(`<span class="%s">%s</span>`, classAttr, content)
	})

	// Cleanup Escape Token {{!}}
	markdown = strings.ReplaceAll(markdown, "{{!}}", "")

	return markdown
}

func getBuiltInIconSvg(name string) string {
	switch strings.ToLower(name) {
	case "sparkles":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"></path></svg>`
	case "bolt":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z"></path></svg>`
	case "academic":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 14l9-5-9-5-9 5 9 5zm0 0v6M12 21a9.003 9.003 0 008.367-5.633M4.367 15.367A9.003 9.003 0 0012 21"></path></svg>`
	case "shield":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path></svg>`
	case "warning":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path></svg>`
	case "success":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	case "chevron":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"></path></svg>`
	case "info":
		fallthrough
	default:
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	}
}
