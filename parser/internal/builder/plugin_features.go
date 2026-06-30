package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type FeaturesPlugin struct {
	pattern *regexp.Regexp
}

func NewFeaturesPlugin() *FeaturesPlugin {
	return &FeaturesPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*features\s*}}(.*?){{\s*/features\s*}}`),
	}
}

func (p *FeaturesPlugin) Name() string            { return "features" }
func (p *FeaturesPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FeaturesPlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

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

	return fmt.Sprintf(`<div class="features-grid">%s</div>`, itemsHtml), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewFeaturesPlugin() })
}
