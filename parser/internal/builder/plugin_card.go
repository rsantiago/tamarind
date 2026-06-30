package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type CardPlugin struct {
	pattern *regexp.Regexp
}

func NewCardPlugin() *CardPlugin {
	return &CardPlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*card(?:\s+padding="([^"]*)")?\s*>??}}(.*?){{<?\s*/?\s*card\s*>??}}`),
	}
}

func (p *CardPlugin) Name() string            { return "card" }
func (p *CardPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *CardPlugin) Process(match []string, sourceDir string) (string, error) {
	paddingVal := match[1]
	content := match[2]

	paddingClass := " card-padding"
	if paddingVal == "false" {
		paddingClass = ""
	}
	return fmt.Sprintf("<div class=\"card%s\">\n\n%s\n\n</div>", paddingClass, strings.TrimSpace(content)), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewCardPlugin() })
}
