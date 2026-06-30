package builder

import (
	"regexp"
)

// parseShortcodeArgs extracts key="value" pairs from a shortcode argument string.
func parseShortcodeArgs(argStr string) map[string]string {
	args := make(map[string]string)
	re := regexp.MustCompile(`([a-zA-Z0-9_-]+)="([^"]*)"`)
	matches := re.FindAllStringSubmatch(argStr, -1)
	for _, m := range matches {
		args[m[1]] = m[2]
	}
	return args
}

// ChartFilePlugin handles shortcodes like {{ barchart file="..." title="..." }}
type ChartFilePlugin struct {
	name      string
	generator func(sourceDir, file string, args map[string]string) string
	pattern   *regexp.Regexp
}

func NewChartFilePlugin(name string, generator func(sourceDir, file string, args map[string]string) string) *ChartFilePlugin {
	return &ChartFilePlugin{
		name:      name,
		generator: generator,
		pattern:   regexp.MustCompile(`(?s){{\s*` + name + `\s+(.*?)\s*}}`),
	}
}

func (p *ChartFilePlugin) Name() string            { return p.name }
func (p *ChartFilePlugin) Pattern() *regexp.Regexp { return p.pattern }
func (p *ChartFilePlugin) Process(match []string, sourceDir string) (string, error) {
	// Block shortcodes are handled by ChartBlockPlugin. If this match contains a closing tag, ignore it.
	// Actually, the regex `(?s){{\s*name\s+(.*?)\s*}}` will match the opening tag.
	// But it might match too much if not careful. Let's make sure it doesn't cross shortcode boundaries.
	args := parseShortcodeArgs(match[1])
	file := args["file"]
	if file == "" {
		// If there is no file, it's likely a block shortcode opening tag, which we shouldn't process here.
		return match[0], nil
	}
	return p.generator(sourceDir, file, args), nil
}

// ChartBlockPlugin handles shortcodes like {{ barchart title="..." }}...{{ /barchart }}
type ChartBlockPlugin struct {
	name      string
	generator func(jsonData []byte, args map[string]string) string
	pattern   *regexp.Regexp
}

func NewChartBlockPlugin(name string, generator func(jsonData []byte, args map[string]string) string) *ChartBlockPlugin {
	return &ChartBlockPlugin{
		name:      name,
		generator: generator,
		pattern:   regexp.MustCompile(`(?s){{\s*` + name + `(?:\s+(.*?))?\s*}}(.*?){{\s*/` + name + `\s*}}`),
	}
}

func (p *ChartBlockPlugin) Name() string            { return p.name + "Block" }
func (p *ChartBlockPlugin) Pattern() *regexp.Regexp { return p.pattern }
func (p *ChartBlockPlugin) Process(match []string, sourceDir string) (string, error) {
	argStr := match[1]
	jsonData := []byte(match[2])
	args := parseShortcodeArgs(argStr)
	return p.generator(jsonData, args), nil
}
