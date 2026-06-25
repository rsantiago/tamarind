package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type FormPlugin struct {
	pattern *regexp.Regexp
}

func NewFormPlugin() *FormPlugin {
	return &FormPlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*form\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*form\s*>??}}`),
	}
}

func (p *FormPlugin) Name() string { return "form" }
func (p *FormPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]
	content := match[2]

	reAction := regexp.MustCompile(`action="([^"]+)"`)
	reMethod := regexp.MustCompile(`method="([^"]+)"`)

	action := "#"
	if m := reAction.FindStringSubmatch(attrs); len(m) > 1 {
		action = m[1]
	}
	method := "POST"
	if mMethod := reMethod.FindStringSubmatch(attrs); len(mMethod) > 1 {
		method = mMethod[1]
	}

	return fmt.Sprintf(`<form action="%s" method="%s">%s</form>`, action, method, strings.TrimSpace(content)), nil
}

type FormInputPlugin struct {
	pattern *regexp.Regexp
}

func NewFormInputPlugin() *FormInputPlugin {
	return &FormInputPlugin{
		pattern: regexp.MustCompile(`{{<?\s*form-input\s+(.*?)\s*/?>??}}`),
	}
}

func (p *FormInputPlugin) Name() string { return "form-input" }
func (p *FormInputPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormInputPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]

	reLabel := regexp.MustCompile(`label="([^"]+)"`)
	reType := regexp.MustCompile(`type="([^"]+)"`)
	rePlaceholder := regexp.MustCompile(`placeholder="([^"]+)"`)

	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}
	typ := "text"
	if m := reType.FindStringSubmatch(attrs); len(m) > 1 {
		typ = m[1]
	}
	placeholder := ""
	if m := rePlaceholder.FindStringSubmatch(attrs); len(m) > 1 {
		placeholder = m[1]
	}

	labelHtml := ""
	if label != "" {
		labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
	}

	return fmt.Sprintf(`<div class="form-group">%s<input type="%s" class="form-input" placeholder="%s"></div>`, labelHtml, typ, placeholder), nil
}

type FormTextareaPlugin struct {
	pattern *regexp.Regexp
}

func NewFormTextareaPlugin() *FormTextareaPlugin {
	return &FormTextareaPlugin{
		pattern: regexp.MustCompile(`{{<?\s*form-textarea\s+(.*?)\s*/?>??}}`),
	}
}

func (p *FormTextareaPlugin) Name() string { return "form-textarea" }
func (p *FormTextareaPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormTextareaPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]

	reLabel := regexp.MustCompile(`label="([^"]+)"`)
	rePlaceholder := regexp.MustCompile(`placeholder="([^"]+)"`)
	reRows := regexp.MustCompile(`rows="([^"]+)"`)

	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}
	placeholder := ""
	if m := rePlaceholder.FindStringSubmatch(attrs); len(m) > 1 {
		placeholder = m[1]
	}
	rows := "4"
	if m := reRows.FindStringSubmatch(attrs); len(m) > 1 {
		rows = m[1]
	}

	labelHtml := ""
	if label != "" {
		labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
	}

	return fmt.Sprintf(`<div class="form-group">%s<textarea class="form-textarea" rows="%s" placeholder="%s"></textarea></div>`, labelHtml, rows, placeholder), nil
}

type FormSelectPlugin struct {
	pattern *regexp.Regexp
}

func NewFormSelectPlugin() *FormSelectPlugin {
	return &FormSelectPlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*form-select\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*form-select\s*>??}}`),
	}
}

func (p *FormSelectPlugin) Name() string { return "form-select" }
func (p *FormSelectPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormSelectPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]
	content := match[2]

	reLabel := regexp.MustCompile(`label="([^"]+)"`)
	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}

	// Parse simple options: {{ option }}Text{{ /option }} or with value/selected
	reOptionSimple := regexp.MustCompile(`(?s){{\s*option\s*}}(.*?){{\s*/option\s*}}`)
	content = reOptionSimple.ReplaceAllString(content, `<option>$1</option>`)

	reOptionAttr := regexp.MustCompile(`(?s){{\s*option\s+value="([^"]+)"(?:\s+selected="([^"]*)")?\s*}}(.*?){{\s*/option\s*}}`)
	content = reOptionAttr.ReplaceAllStringFunc(content, func(optMatch string) string {
		optParts := reOptionAttr.FindStringSubmatch(optMatch)
		val := optParts[1]
		sel := optParts[2]
		text := optParts[3]

		selectedAttr := ""
		if sel == "true" {
			selectedAttr = " selected"
		}
		return fmt.Sprintf(`<option value="%s"%s>%s</option>`, val, selectedAttr, strings.TrimSpace(text))
	})

	labelHtml := ""
	if label != "" {
		labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
	}

	return fmt.Sprintf(`<div class="form-group">%s<select class="form-select">%s</select></div>`, labelHtml, strings.TrimSpace(content)), nil
}

type FormCheckboxPlugin struct {
	pattern *regexp.Regexp
}

func NewFormCheckboxPlugin() *FormCheckboxPlugin {
	return &FormCheckboxPlugin{
		pattern: regexp.MustCompile(`{{<?\s*form-checkbox\s+(.*?)\s*/?>??}}`),
	}
}

func (p *FormCheckboxPlugin) Name() string { return "form-checkbox" }
func (p *FormCheckboxPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormCheckboxPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]

	reLabel := regexp.MustCompile(`label="([^"]+)"`)
	reChecked := regexp.MustCompile(`checked="([^"]+)"`)

	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}
	checked := ""
	if m := reChecked.FindStringSubmatch(attrs); len(m) > 1 && m[1] == "true" {
		checked = " checked"
	}

	return fmt.Sprintf(`<div class="form-group"><label class="form-label"><input type="checkbox" class="form-checkbox"%s> %s</label></div>`, checked, label), nil
}

type FormRadioGroupPlugin struct {
	pattern *regexp.Regexp
}

func NewFormRadioGroupPlugin() *FormRadioGroupPlugin {
	return &FormRadioGroupPlugin{
		pattern: regexp.MustCompile(`(?s){{<?\s*form-radio-group\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*form-radio-group\s*>??}}`),
	}
}

func (p *FormRadioGroupPlugin) Name() string { return "form-radio-group" }
func (p *FormRadioGroupPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormRadioGroupPlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]
	content := match[2]

	reLabel := regexp.MustCompile(`label="([^"]+)"`)
	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}

	// Inside form-radio-group, we parse form-radios: {{ form-radio name="..." label="..." checked="..." }}
	reFormRadio := regexp.MustCompile(`{{<?\s*form-radio\s+(.*?)\s*/?>??}}`)
	parsedContent := reFormRadio.ReplaceAllStringFunc(content, func(radioMatch string) string {
		radioAttrs := reFormRadio.FindStringSubmatch(radioMatch)[1]

		reName := regexp.MustCompile(`name="([^"]+)"`)
		reRadioLabel := regexp.MustCompile(`label="([^"]+)"`)
		reRadioChecked := regexp.MustCompile(`checked="([^"]+)"`)

		name := ""
		if m := reName.FindStringSubmatch(radioAttrs); len(m) > 1 {
			name = m[1]
		}
		radioLabel := ""
		if m := reRadioLabel.FindStringSubmatch(radioAttrs); len(m) > 1 {
			radioLabel = m[1]
		}
		radioChecked := ""
		if m := reRadioChecked.FindStringSubmatch(radioAttrs); len(m) > 1 && m[1] == "true" {
			radioChecked = " checked"
		}

		return fmt.Sprintf(`<label class="form-label"><input type="radio" name="%s" class="form-radio"%s> %s</label>`, name, radioChecked, radioLabel)
	})

	labelHtml := ""
	if label != "" {
		labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
	}

	return fmt.Sprintf(`<div class="form-group">%s%s</div>`, labelHtml, strings.TrimSpace(parsedContent)), nil
}

type FormFilePlugin struct {
	pattern *regexp.Regexp
}

func NewFormFilePlugin() *FormFilePlugin {
	return &FormFilePlugin{
		pattern: regexp.MustCompile(`{{<?\s*form-file\s+(.*?)\s*/?>??}}`),
	}
}

func (p *FormFilePlugin) Name() string { return "form-file" }
func (p *FormFilePlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FormFilePlugin) Process(match []string, sourceDir string) (string, error) {
	attrs := match[1]

	reLabel := regexp.MustCompile(`label="([^"]+)"`)
	label := ""
	if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
		label = m[1]
	}

	labelHtml := ""
	if label != "" {
		labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
	}

	return fmt.Sprintf(`<div class="form-group">%s<input type="file" class="form-file"></div>`, labelHtml), nil
}
