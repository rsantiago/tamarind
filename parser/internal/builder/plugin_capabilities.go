package builder

import (
	"fmt"
	"regexp"
)

type CapabilitiesPlugin struct {
	pattern *regexp.Regexp
}

func NewCapabilitiesPlugin() *CapabilitiesPlugin {
	return &CapabilitiesPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*capabilities(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/capabilities\s*}}`),
	}
}

func (p *CapabilitiesPlugin) Name() string            { return "capabilities" }
func (p *CapabilitiesPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *CapabilitiesPlugin) Process(match []string, sourceDir string) (string, error) {
	title := match[1]
	content := match[2]

	reCapabilityStatus := regexp.MustCompile(`{{\s*capability\s+name="([^"]+)"\s+desc="([^"]+)"\s+status="([^"]+)"(?:\s+statusLabel="([^"]*)")?\s*}}`)
	reCapabilityCheck := regexp.MustCompile(`{{\s*capability\s+name="([^"]+)"\s+desc="([^"]+)"\s+check="([^"]+)"\s*}}`)
	rowsHtml := ""

	statusMatches := reCapabilityStatus.FindAllStringSubmatch(content, -1)
	for _, itemSubmatch := range statusMatches {
		name := itemSubmatch[1]
		desc := itemSubmatch[2]
		status := itemSubmatch[3]
		statusLabel := status
		if len(itemSubmatch) > 4 && itemSubmatch[4] != "" {
			statusLabel = itemSubmatch[4]
		}

		rowsHtml += fmt.Sprintf(`<div class="capability-row"><div class="capability-info"><span class="capability-name">%s</span><span class="capability-desc">%s</span></div><span class="capability-status status-%s">%s</span></div>`, name, desc, status, statusLabel)
	}

	if len(statusMatches) == 0 {
		checkMatches := reCapabilityCheck.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range checkMatches {
			name := itemSubmatch[1]
			desc := itemSubmatch[2]
			checkVal := itemSubmatch[3]

			status := "pending"
			statusLabel := "Pending"
			if checkVal == "true" {
				status = "success"
				statusLabel = "Yes"
			} else if checkVal == "warn" {
				status = "warning"
				statusLabel = "Partial"
			} else if checkVal == "false" {
				status = "error"
				statusLabel = "No"
			}

			rowsHtml += fmt.Sprintf(`<div class="capability-row"><div class="capability-info"><span class="capability-name">%s</span><span class="capability-desc">%s</span></div><span class="capability-status status-%s">%s</span></div>`, name, desc, status, statusLabel)
		}
	}

	headerHtml := ""
	if title != "" {
		headerHtml = fmt.Sprintf(`<div class="capability-header"><div class="capability-card-title">%s</div></div>`, title)
	}

	return fmt.Sprintf(`<div class="capability-card">%s%s</div>`, headerHtml, rowsHtml), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewCapabilitiesPlugin() })
}
