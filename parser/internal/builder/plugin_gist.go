package builder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type GistPlugin struct {
	pattern *regexp.Regexp
}

func NewGistPlugin() *GistPlugin {
	return &GistPlugin{
		pattern: regexp.MustCompile(`{{\s*gist\s+id="([^"]+)"\s*}}`),
	}
}

func (p *GistPlugin) Name() string            { return "gist" }
func (p *GistPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *GistPlugin) Process(match []string, sourceDir string) (string, error) {
	gistID := match[1]

	resp, err := http.Get(fmt.Sprintf("https://gist.github.com/%s.json", gistID))
	if err != nil {
		return fmt.Sprintf("> **Error loading gist %s**: %v", gistID, err), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("> **Error loading gist %s**: Status %d", gistID, resp.StatusCode), nil
	}

	var data struct {
		Div        string `json:"div"`
		Stylesheet string `json:"stylesheet"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Sprintf("> **Error decoding gist %s**: %v", gistID, err), nil
	}

	// Strip newlines to prevent Markdown from interpreting indented HTML as code blocks
	data.Div = strings.ReplaceAll(data.Div, "\n", "")
	data.Div = strings.ReplaceAll(data.Div, "\r", "")

	return fmt.Sprintf(`<link rel="stylesheet" href="%s">%s`, data.Stylesheet, data.Div), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewGistPlugin() })
}
