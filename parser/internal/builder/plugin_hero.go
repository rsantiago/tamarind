// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type HeroPlugin struct {
	pattern *regexp.Regexp
}

func NewHeroPlugin() *HeroPlugin {
	return &HeroPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*hero\s+title="([^"]+)"\s+subtitle="([^"]+)"\s*}}(.*?){{\s*/hero\s*}}`),
	}
}

func (p *HeroPlugin) Name() string            { return "hero" }
func (p *HeroPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *HeroPlugin) Process(match []string, sourceDir string) (string, error) {
	title := match[1]
	subtitle := match[2]
	content := match[3]

	reBtn := regexp.MustCompile(`(?s){{\s*hero_btn\s+label="([^"]+)"\s+link="([^"]+)"\s+primary="([^"]+)"\s*}}`)
	buttonsHtml := ""

	btnMatches := reBtn.FindAllStringSubmatch(content, -1)
	for _, btnMatch := range btnMatches {
		label := btnMatch[1]
		link := btnMatch[2]
		primary := btnMatch[3]

		btnClass := "hero-btn-secondary"
		if primary == "true" {
			btnClass = "hero-btn-primary"
		}

		target := ""
		if strings.HasPrefix(link, "http") {
			target = ` target="_blank"`
		}

		buttonsHtml += fmt.Sprintf(`<a class="%s" href="%s"%s>%s</a>`, btnClass, link, target, label)
	}

	if buttonsHtml != "" {
		buttonsHtml = fmt.Sprintf(`<div class="hero-ctas">%s</div>`, buttonsHtml)
	}

	html := fmt.Sprintf(`
<div class="hero-container">
  <div class="hero-banner">
    <h2 class="hero-title">%s</h2>
    <p class="hero-subtitle">%s</p>
    %s
  </div>
</div>
`, title, subtitle, buttonsHtml)

	return html, nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewHeroPlugin() })
}
