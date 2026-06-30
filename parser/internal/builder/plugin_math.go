package builder

import (
	"regexp"
)

type MathPlugin struct {
	pattern *regexp.Regexp
}

func NewMathPlugin() *MathPlugin {
	return &MathPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*math\s*}}(.*?){{\s*/math\s*}}`),
	}
}

func (p *MathPlugin) Name() string            { return "math" }
func (p *MathPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *MathPlugin) Process(match []string, sourceDir string) (string, error) {
	return `<div class="math-block">$$` + match[1] + `$$</div>`, nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewMathPlugin() })
}
