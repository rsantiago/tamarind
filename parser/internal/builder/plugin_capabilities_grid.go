package builder

import (
	"regexp"
)

type CapabilitiesGridPlugin struct {
	pattern *regexp.Regexp
}

func NewCapabilitiesGridPlugin() *CapabilitiesGridPlugin {
	return &CapabilitiesGridPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*capabilities-grid\s*}}(.*?){{\s*/capabilities-grid\s*}}`),
	}
}

func (p *CapabilitiesGridPlugin) Name() string            { return "capabilities-grid" }
func (p *CapabilitiesGridPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *CapabilitiesGridPlugin) Process(match []string, sourceDir string) (string, error) {
	// The original code uses ReplaceAllString, so we just wrap match[1] in the div.
	return `<div class="capabilities-grid">` + match[1] + `</div>`, nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewCapabilitiesGridPlugin() })
}
