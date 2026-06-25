package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type DropdownPlugin struct {
	pattern *regexp.Regexp
}

func NewDropdownPlugin() *DropdownPlugin {
	return &DropdownPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*dropdown\s+(.*?)\s*}}(.*?){{\s*/dropdown\s*}}`),
	}
}

func (p *DropdownPlugin) Name() string { return "dropdown" }
func (p *DropdownPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *DropdownPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]
	content := match[2]

	reId := regexp.MustCompile(`id="([^"]+)"`)
	reLabel := regexp.MustCompile(`label="([^"]+)"`)

	id := ""
	if m := reId.FindStringSubmatch(attrs); len(m) > 1 {
		id = m[1]
	}
	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}

	reOption := regexp.MustCompile(`(?s){{\s*option\s+value="([^"]+)"(?:\s+selected="([^"]*)")?\s*}}(.*?){{\s*/option\s*}}`)
	optionsHtml := ""
	
	optMatches := reOption.FindAllStringSubmatch(content, -1)
	for _, itemSubmatch := range optMatches {
		val := itemSubmatch[1]
		sel := itemSubmatch[2]
		text := itemSubmatch[3]

		selectedAttr := ""
		if sel == "true" {
			selectedAttr = " selected"
		}

		optionsHtml += fmt.Sprintf(`<option value="%s"%s>%s</option>`, val, selectedAttr, strings.TrimSpace(text))
	}

	labelHtml := ""
	if label != "" {
		labelHtml = fmt.Sprintf(`<label class="tamarind-select-label">%s</label>`, label)
	}

	idAttr := ""
	if id != "" {
		idAttr = fmt.Sprintf(` id="%s"`, id)
	}

	chevronSvg := getBuiltInIconSvg("chevron")

	return fmt.Sprintf(`<div class="tamarind-select-wrapper">%s<div class="tamarind-select-control"><select class="tamarind-select"%s>%s</select><div class="tamarind-select-chevron">%s</div></div></div>`, labelHtml, idAttr, optionsHtml, chevronSvg), nil
}
