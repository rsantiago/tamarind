package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type ButtonPlugin struct {
	pattern *regexp.Regexp
}

func NewButtonPlugin() *ButtonPlugin {
	return &ButtonPlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*button\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*button\s*>??}}`),
	}
}

func (p *ButtonPlugin) Name() string            { return "button" }
func (p *ButtonPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *ButtonPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]
	content := match[2]

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

	return fmt.Sprintf(`<a href="%s" class="%s"%s>%s</a>`, href, classAttr, targetAttr, strings.TrimSpace(content)), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewButtonPlugin() })
}
