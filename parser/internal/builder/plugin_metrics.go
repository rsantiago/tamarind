package builder

import (
	"fmt"
	"regexp"
)

type MetricsPlugin struct {
	pattern *regexp.Regexp
}

func NewMetricsPlugin() *MetricsPlugin {
	return &MetricsPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*metrics\s*}}(.*?){{\s*/metrics\s*}}`),
	}
}

func (p *MetricsPlugin) Name() string { return "metrics" }
func (p *MetricsPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *MetricsPlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

	reMetric := regexp.MustCompile(`{{\s*metric\s+value="([^"]+)"\s+label="([^"]+)"\s*}}`)
	itemsHtml := ""
	
	metricMatches := reMetric.FindAllStringSubmatch(content, -1)
	for _, itemSubmatch := range metricMatches {
		val := itemSubmatch[1]
		lbl := itemSubmatch[2]
		itemsHtml += fmt.Sprintf(`<div class="metric-card"><div class="metric-value">%s</div><div class="metric-label">%s</div></div>`, val, lbl)
	}

	return fmt.Sprintf(`<div class="metrics-grid">%s</div>`, itemsHtml), nil
}
