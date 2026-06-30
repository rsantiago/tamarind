package builder

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type TimelinePlugin struct {
	pattern *regexp.Regexp
}

func NewTimelinePlugin() *TimelinePlugin {
	return &TimelinePlugin{
		pattern: regexp.MustCompile(`(?s){{\s*timeline\s*}}(.*?){{\s*/timeline\s*}}`),
	}
}

func (p *TimelinePlugin) Name() string            { return "timeline" }
func (p *TimelinePlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *TimelinePlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

	reItem1 := regexp.MustCompile(`(?s){{\s*item\s+title="([^"]+)"(?:\s+number="([^"]*)")?\s*}}(.*?){{\s*/item\s*}}`)
	reItem2 := regexp.MustCompile(`(?s){{\s*timeline-item\s+step="([^"]+)"\s+title="([^"]+)"\s*}}(.*?){{\s*/timeline-item\s*}}`)
	itemsHtml := ""

	item1Matches := reItem1.FindAllStringSubmatch(content, -1)
	for _, itemSubmatch := range item1Matches {
		title := itemSubmatch[1]
		num := itemSubmatch[2]
		desc := itemSubmatch[3]

		var buf bytes.Buffer
		var htmlDesc string
		if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
			htmlDesc = buf.String()
		} else {
			htmlDesc = strings.TrimSpace(desc)
		}

		badgeHtml := ""
		if num != "" {
			badgeHtml = fmt.Sprintf(`<div class="timeline-badge"><span class="timeline-badge-number">%s</span></div>`, num)
		} else {
			badgeHtml = `<div class="timeline-badge"></div>`
		}

		itemsHtml += fmt.Sprintf(`<div class="timeline-item">%s<div class="timeline-content"><h3 class="timeline-title">%s</h3><div class="timeline-desc">%s</div></div></div>`, badgeHtml, title, htmlDesc)
	}

	if len(item1Matches) == 0 {
		item2Matches := reItem2.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range item2Matches {
			num := itemSubmatch[1]
			title := itemSubmatch[2]
			desc := itemSubmatch[3]

			var buf bytes.Buffer
			var htmlDesc string
			if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
				htmlDesc = buf.String()
			} else {
				htmlDesc = strings.TrimSpace(desc)
			}

			badgeHtml := ""
			if num != "" {
				badgeHtml = fmt.Sprintf(`<div class="timeline-badge"><span class="timeline-badge-number">%s</span></div>`, num)
			} else {
				badgeHtml = `<div class="timeline-badge"></div>`
			}

			itemsHtml += fmt.Sprintf(`<div class="timeline-item">%s<div class="timeline-content"><h3 class="timeline-title">%s</h3><div class="timeline-desc">%s</div></div></div>`, badgeHtml, title, htmlDesc)
		}
	}

	return fmt.Sprintf(`<div class="timeline-container">%s</div>`, itemsHtml), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewTimelinePlugin() })
}
