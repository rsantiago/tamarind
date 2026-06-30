package builder

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type AccordionPlugin struct {
	pattern *regexp.Regexp
}

func NewAccordionPlugin() *AccordionPlugin {
	return &AccordionPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*accordion\s*}}(.*?){{\s*/accordion\s*}}`),
	}
}

func (p *AccordionPlugin) Name() string            { return "accordion" }
func (p *AccordionPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *AccordionPlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

	reAccordionItem := regexp.MustCompile(`(?s){{\s*accordion-item\s+title="([^"]+)"\s*}}(.*?){{\s*/accordion-item\s*}}`)
	itemsHtml := ""

	itemMatches := reAccordionItem.FindAllStringSubmatch(content, -1)
	for _, itemSubmatch := range itemMatches {
		title := itemSubmatch[1]
		desc := itemSubmatch[2]

		var buf bytes.Buffer
		var htmlDesc string
		if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
			htmlDesc = strings.TrimSpace(buf.String())
		} else {
			htmlDesc = strings.TrimSpace(desc)
		}

		itemsHtml += fmt.Sprintf(`<details class="tamarind-accordion"><summary class="tamarind-accordion-summary">%s</summary><div class="tamarind-accordion-content">%s</div></details>`, title, htmlDesc)
	}

	return fmt.Sprintf(`<div class="tamarind-accordion-container">%s</div>`, itemsHtml), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewAccordionPlugin() })
}
