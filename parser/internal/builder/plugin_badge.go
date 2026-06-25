package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type BadgePlugin struct {
	pattern *regexp.Regexp
}

func NewBadgePlugin() *BadgePlugin {
	return &BadgePlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*badge(?:\s+type="([^"]*)")?\s*>??}}(.*?){{<?\s*/?\s*badge\s*>??}}`),
	}
}

func (p *BadgePlugin) Name() string { return "badge" }
func (p *BadgePlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *BadgePlugin) Process(match []string, sourceDir string) (string, error) {
	typ := match[1]
	content := match[2]

	classAttr := "badge"
	if typ != "" {
		classAttr += " badge-" + typ
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, classAttr, strings.TrimSpace(content)), nil
}
