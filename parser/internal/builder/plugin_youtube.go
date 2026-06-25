package builder

import (
	"fmt"
	"regexp"
)

type YouTubePlugin struct {
	pattern *regexp.Regexp
}

func NewYouTubePlugin() *YouTubePlugin {
	return &YouTubePlugin{
		pattern: regexp.MustCompile(`{{\s*youtube\s+([a-zA-Z0-9_-]+)\s*}}`),
	}
}

func (p *YouTubePlugin) Name() string { return "youtube" }
func (p *YouTubePlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *YouTubePlugin) Process(match []string, sourceDir string) (string, error) {
	return fmt.Sprintf(`<div class="video-container"><iframe src="https://www.youtube.com/embed/%s" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></div>`, match[1]), nil
}
