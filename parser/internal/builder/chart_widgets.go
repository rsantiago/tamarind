// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func getChartColors(colorsStr string) []string {
	if colorsStr == "" {
		return []string{
			"var(--chart-1, var(--primary-color, #3b82f6))",
			"var(--chart-2, var(--secondary-color, #10b981))",
			"var(--chart-3, #f59e0b)",
			"var(--chart-4, #ef4444)",
			"var(--chart-5, #8b5cf6)",
			"var(--chart-6, #ec4899)",
			"var(--chart-7, #14b8a6)",
			"var(--chart-8, #f97316)",
			"var(--chart-9, #64748b)",
		}
	}
	parts := strings.Split(colorsStr, ",")
	var result []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	if len(result) == 0 {
		return getChartColors("")
	}
	return result
}

type ChartDataPoint struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

func generateBarChart(sourceDir, filename, title, colors string) string {
	dataDir := filepath.Join(sourceDir, "data")
	filePath := filepath.Join(dataDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("<div class='chart-error'>Error loading data: %s</div>", err.Error())
	}
	return generateBarChartFromJSON(content, title, colors)
}

func generateBarChartFromJSON(content []byte, title, colors string) string {
	var data []ChartDataPoint
	if err := json.Unmarshal(content, &data); err != nil {
		return fmt.Sprintf("<div class='chart-error'>Error parsing JSON: %s</div>", err.Error())
	}

	if len(data) == 0 {
		return "<div class='chart-error'>No data found</div>"
	}

	var max float64
	for _, d := range data {
		if d.Value > max {
			max = d.Value
		}
	}
	if max == 0 {
		max = 1
	}

	html := `<div class="tamarind-barchart" style="margin: 2rem 0; width: 100%;">`
	if title != "" {
		html += fmt.Sprintf(`<h4 class="tamarind-chart-title" style="text-align:center; margin-bottom: 1rem;">%s</h4>`, title)
	}
	html += `<div class="tamarind-barchart-container" style="display:flex; align-items:flex-end; height: 250px; gap: 15px; padding: 10px 0; border-bottom: 2px solid var(--text-secondary); border-left: 2px solid var(--text-secondary); overflow-x: auto;">`

	cList := getChartColors(colors)
	for i, d := range data {
		c := cList[i%len(cList)]
		heightPct := (d.Value / max) * 100.0
		html += fmt.Sprintf(`
		<div class="tamarind-barchart-col" style="flex: 1; display:flex; flex-direction:column; justify-content: flex-end; align-items:center; min-width: 40px; height: 100%%;">
			<div class="tamarind-barchart-val" style="font-size: 0.8rem; margin-bottom: 5px; color: currentColor;">%.1f</div>
			<div class="tamarind-barchart-bar" style="width: 100%%; max-width: 60px; height: %.1f%%%%; background-color: %s; border-radius: 4px 4px 0 0; transition: height 0.5s ease-out;"></div>
			<div class="tamarind-barchart-lbl" style="font-size: 0.8rem; margin-top: 8px; text-align:center; color: currentColor;">%s</div>
		</div>`, d.Value, heightPct, c, d.Label)
	}
	html += `</div></div>`
	return html
}

func generatePieChart(sourceDir, filename, title, colors string) string {
	dataDir := filepath.Join(sourceDir, "data")
	filePath := filepath.Join(dataDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("<div class='chart-error'>Error loading data: %s</div>", err.Error())
	}
	return generatePieChartFromJSON(content, title, colors)
}

func generatePieChartFromJSON(content []byte, title, colors string) string {
	var data []ChartDataPoint
	if err := json.Unmarshal(content, &data); err != nil {
		return fmt.Sprintf("<div class='chart-error'>Error parsing JSON: %s</div>", err.Error())
	}

	if len(data) == 0 {
		return "<div class='chart-error'>No data found</div>"
	}

	var total float64
	for _, d := range data {
		total += d.Value
	}
	if total == 0 {
		total = 1
	}

	html := `<div class="tamarind-piechart" style="margin: 2rem 0; padding: 1rem; border: 1px solid var(--border-color); border-radius: 8px;">`
	if title != "" {
		html += fmt.Sprintf(`<h4 class="tamarind-chart-title" style="text-align:center; margin-bottom: 1.5rem;">%s</h4>`, title)
	}

	html += `<div style="display: flex; gap: 30px; align-items: center; justify-content: center; flex-wrap: wrap;">`
	html += `<svg viewBox="0 0 32 32" style="width: 250px; height: 250px; border-radius: 50%; transform: rotate(-90deg);">`

	cList := getChartColors(colors)

	currentOffset := 0.0
	legendHtml := `<div class="tamarind-piechart-legend" style="display: flex; flex-direction: column; gap: 10px; min-width: 150px;">`
	C := 100.530964915

	for i, d := range data {
		color := cList[i%len(cList)]
		valRatio := d.Value / total
		dashArray := valRatio * C
		
		html += fmt.Sprintf(`<circle r="16" cx="16" cy="16" fill="transparent" stroke="%s" stroke-width="32" stroke-dasharray="%.2f %.2f" stroke-dashoffset="-%.2f" style="transition: stroke-dasharray 0.5s ease; cursor: pointer;" />`, color, dashArray, C, currentOffset)
		
		legendHtml += fmt.Sprintf(`<div style="display:flex; align-items:center; gap: 10px;"><div style="width: 16px; height: 16px; border-radius:4px; background-color: %s;"></div><span style="font-size: 0.95rem; color: currentColor;">%s (%.1f)</span></div>`, color, d.Label, d.Value)
		
		currentOffset += dashArray
	}
	html += `</svg>`
	html += legendHtml + `</div></div></div>`

	return html
}

func generateLineChart(sourceDir, filename, title, colors string) string {
	dataDir := filepath.Join(sourceDir, "data")
	filePath := filepath.Join(dataDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("<div class='chart-error'>Error loading data: %s</div>", err.Error())
	}
	return generateLineChartFromJSON(content, title, colors)
}

func generateLineChartFromJSON(content []byte, title, colors string) string {
	var data []ChartDataPoint
	if err := json.Unmarshal(content, &data); err != nil {
		return fmt.Sprintf("<div class='chart-error'>Error parsing JSON: %s</div>", err.Error())
	}

	if len(data) == 0 {
		return "<div class='chart-error'>No data found</div>"
	}

	var max float64
	for _, d := range data {
		if d.Value > max {
			max = d.Value
		}
	}
	if max == 0 {
		max = 1
	}

	width := 600.0
	height := 250.0
	padding := 40.0
	
	html := `<div class="tamarind-linechart" style="margin: 2rem 0; width: 100%;">`
	if title != "" {
		html += fmt.Sprintf(`<h4 class="tamarind-chart-title" style="text-align:center; margin-bottom: 1rem;">%s</h4>`, title)
	}

	html += fmt.Sprintf(`<svg viewBox="0 0 %.0f %.0f" style="width: 100%%; height: auto; max-height: 300px; display: block; overflow: visible;">`, width, height)
	
	// Draw axes
	html += fmt.Sprintf(`<line x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" stroke="var(--text-secondary)" stroke-width="2" />`, padding, height-padding, width-padding, height-padding)
	html += fmt.Sprintf(`<line x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" stroke="var(--text-secondary)" stroke-width="2" />`, padding, padding, padding, height-padding)

	points := ""
	stepX := (width - 2*padding) / float64(len(data)-1)
	if len(data) == 1 {
		stepX = 0
	}

	cList := getChartColors(colors)
	c := cList[0]

	for i, d := range data {
		x := padding + float64(i)*stepX
		y := height - padding - ((d.Value / max) * (height - 2*padding))
		points += fmt.Sprintf("%.1f,%.1f ", x, y)
		
		// Circle for data point
		html += fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="4" fill="%s" />`, x, y, c)
		
		// Label (every N labels to avoid crowding, or just all if small)
		if len(data) <= 10 || i%int(math.Ceil(float64(len(data))/10)) == 0 {
			html += fmt.Sprintf(`<text x="%.1f" y="%.1f" font-size="10" fill="currentColor" text-anchor="middle">%s</text>`, x, height-padding+15, d.Label)
		}
	}

	// Draw line
	html += fmt.Sprintf(`<polyline fill="none" stroke="%s" stroke-width="3" points="%s" />`, c, strings.TrimSpace(points))

	html += `</svg></div>`
	return html
}
