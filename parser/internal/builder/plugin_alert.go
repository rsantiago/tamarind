package builder

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type AlertPlugin struct {
	pattern *regexp.Regexp
}

func NewAlertPlugin() *AlertPlugin {
	return &AlertPlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*(?:alert|callout)\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*(?:alert|callout)\s*>??}}`),
	}
}

func (p *AlertPlugin) Name() string { return "alert" }
func (p *AlertPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *AlertPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]
	content := match[2]

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
	if typ == "danger" || typ == "error" || typ == "warn" || typ == "warning" {
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

	var buf bytes.Buffer
	var htmlContent string
	if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(content)), &buf); err == nil {
		htmlContent = buf.String()
	} else {
		htmlContent = strings.TrimSpace(content)
	}

	return fmt.Sprintf(`<div class="alert-container alert-%s"><div class="alert-icon-box">%s</div><div class="alert-content">%s<div class="alert-message">%s</div></div></div>`, typ, svgIcon, titleHtml, htmlContent), nil
}
