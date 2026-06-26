// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
package builder

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var defaultColors = []string{
	"var(--chart-1, var(--primary-color, #3b82f6))",
	"var(--chart-2, var(--secondary-color, #10b981))",
	"var(--chart-3, #93761f)",
	"var(--chart-4, #8ba5c0)",
	"var(--chart-5, #8f233e)",
	"var(--chart-6, #bcd278)",
	"var(--chart-7, #23468f)",
	"var(--chart-8, #db8b6f)",
	"var(--chart-9, #595959)",
}

func generateDonutChart(sourceDir, filename string, args map[string]string) string {
	content, err := os.ReadFile(filepath.Join(sourceDir, "data", filename))
	if err != nil { return err.Error() }
	return generateDonutChartFromJSON(content, args)
}

func generateDonutChartFromJSON(content []byte, args map[string]string) string {
	var data []ChartDataPoint
	if err := json.Unmarshal(content, &data); err != nil { return err.Error() }

	total := 0.0
	for _, d := range data { total += d.Value }
	if total == 0 { total = 1 }

	html := `<div class="tamarind-piechart" style="margin: 2rem 0; padding: 1rem; border: 1px solid var(--border-color); border-radius: 8px;">`
	if args["title"] != "" { html += fmt.Sprintf(`<h4 class="tamarind-chart-title" style="text-align:center; margin-bottom: 1.5rem;">%s</h4>`, args["title"]) }
	html += `<div class="tamarind-piechart-container" style="display: flex; gap: 30px; align-items: center; justify-content: center; flex-wrap: wrap;">`
	html += `<div class="tamarind-donutchart-inner" style="position: relative; width: 250px; height: 250px;">`
	html += `<svg viewBox="0 0 40 40" style="width: 100%; height: 100%; border-radius: 50%; transform: rotate(-90deg);">`

	currentOffset := 0.0
	legendHtml := `<div class="tamarind-piechart-legend" style="display: flex; flex-direction: column; gap: 10px; min-width: 150px;">`
	C := 100.530964915 // 2 * pi * 16

	for i, d := range data {
		color := defaultColors[i%len(defaultColors)]
		dashArray := (d.Value / total) * C
		html += fmt.Sprintf(`<circle r="16" cx="20" cy="20" fill="transparent" stroke="%s" stroke-width="6" stroke-dasharray="%.2f %.2f" stroke-dashoffset="-%.2f" />`, color, dashArray, C, currentOffset)
		legendHtml += fmt.Sprintf(`<div style="display:flex; align-items:center; gap: 10px;"><div style="width: 16px; height: 16px; border-radius:4px; background-color: %s;"></div><span style="font-size: 0.95rem;">%s (%.1f)</span></div>`, color, d.Label, d.Value)
		currentOffset += dashArray
	}
	html += `</svg>`
	html += fmt.Sprintf(`<div style="position:absolute; top:0; left:0; width:100%%; height:100%%; display:flex; align-items:center; justify-content:center; flex-direction:column;"><span style="font-size:2rem; font-weight:bold;">%.0f</span><span style="font-size:0.8rem; opacity:0.7;">Total</span></div>`, total)
	html += `</div>` + legendHtml + `</div></div></div>`
	return html
}

func generateHBarchart(sourceDir, filename string, args map[string]string) string {
	content, err := os.ReadFile(filepath.Join(sourceDir, "data", filename))
	if err != nil { return err.Error() }
	return generateHBarchartFromJSON(content, args)
}

func generateHBarchartFromJSON(content []byte, args map[string]string) string {
	var data []ChartDataPoint
	if err := json.Unmarshal(content, &data); err != nil { return err.Error() }

	max := 0.0
	for _, d := range data { if d.Value > max { max = d.Value } }
	showX := args["show-x"] != "false"
	showY := args["show-y"] != "false"
	gridX := args["grid-x"] == "true"
	gridY := args["grid-y"] == "true"

	html := `<div class="tamarind-hbarchart" style="margin: 2rem 0; width: 100%;">`
	if args["title"] != "" { html += fmt.Sprintf(`<h4 class="tamarind-chart-title" style="text-align:center; margin-bottom: 1rem;">%s</h4>`, args["title"]) }
	
	bgGradients := []string{}
	bgSizes := []string{}
	if gridY {
		bgGradients = append(bgGradients, `linear-gradient(var(--border-color) 1px, transparent 1px)`)
		bgSizes = append(bgSizes, `100% 20%`)
	}
	if gridX {
		bgGradients = append(bgGradients, `linear-gradient(90deg, var(--border-color) 1px, transparent 1px)`)
		bgSizes = append(bgSizes, `20% 100%`)
	}
	bgStyle := ""
	if len(bgGradients) > 0 {
		bgStyle = fmt.Sprintf(`background-image: %s; background-size: %s;`, strings.Join(bgGradients, ", "), strings.Join(bgSizes, ", "))
	}
	html += fmt.Sprintf(`<div class="tamarind-hbarchart-container" style="display:flex; flex-direction:column; justify-content:center; gap: 15px; padding-left: 10px; border-left: 2px solid var(--border-color); %s">`, bgStyle)

	for i, d := range data {
		widthPct := (d.Value / max) * 100.0
		color := defaultColors[i%len(defaultColors)]
		
		lblHtml := ""
		if showY {
			lblHtml = fmt.Sprintf(`<div style="width: 80px; font-size: 0.85rem; text-align:right;">%s</div>`, d.Label)
		}
		
		valHtml := ""
		if showX {
			valHtml = fmt.Sprintf(`<span style="margin-left: 10px; font-size: 0.8rem;">%.1f</span>`, d.Value)
		}
		
		html += fmt.Sprintf(`
		<div style="display:flex; align-items:center; gap: 10px; flex: 1;">
			%s
			<div style="flex:1; display:flex; align-items:center;">
				<div style="width: %.1f%%; height: 24px; background-color: %s; border-radius: 0 4px 4px 0;"></div>
				%s
			</div>
		</div>`, lblHtml, widthPct, color, valHtml)
	}
	html += `</div></div>`
	return html
}

type MultiSeriesChartData struct {
	Categories []string `json:"categories"`
	Series     []struct {
		Name string    `json:"name"`
		Data []float64 `json:"data"`
	} `json:"series"`
}

func generateMultiLineChart(sourceDir, filename string, args map[string]string) string {
	content, err := os.ReadFile(filepath.Join(sourceDir, "data", filename))
	if err != nil { return err.Error() }
	return generateMultiLineChartFromJSON(content, args)
}

func generateMultiLineChartFromJSON(content []byte, args map[string]string) string {
	var data MultiSeriesChartData
	if err := json.Unmarshal(content, &data); err != nil { return err.Error() }

	max := 0.0
	for _, s := range data.Series {
		for _, v := range s.Data {
			if v > max { max = v }
		}
	}
	if max == 0 { max = 1 }

	showX := args["show-x"] != "false"
	showY := args["show-y"] != "false"
	showDots := args["show-dots"] != "false"
	gridX := args["grid-x"] == "true"
	gridY := args["grid-y"] == "true"

	width, height, padding := 600.0, 250.0, 40.0
	html := `<div class="tamarind-multiline" style="margin: 2rem 0; width: 100%;">`
	if args["title"] != "" { html += fmt.Sprintf(`<h4 style="text-align:center; margin-bottom: 1rem;">%s</h4>`, args["title"]) }
	html += fmt.Sprintf(`<svg viewBox="0 0 %.0f %.0f" style="width: 100%%; height: auto; max-height: 300px; display: block; overflow: visible;">`, width, height)
	
	if gridY {
		// Draw mild grid lines for Y axis
		for i := 1; i <= 4; i++ {
			gy := height - padding - (float64(i) * (height - 2*padding) / 5.0)
			html += fmt.Sprintf(`<line x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" stroke="var(--border-color)" stroke-width="1" />`, padding, gy, width-padding, gy)
		}
	}
	
	stepX := 0.0
	if len(data.Categories) > 1 { stepX = (width - 2*padding) / float64(len(data.Categories)-1) }

	if gridX {
		for i := 0; i < len(data.Categories); i++ {
			gx := padding + float64(i)*stepX
			html += fmt.Sprintf(`<line x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" stroke="var(--border-color)" stroke-width="1" />`, gx, padding, gx, height-padding)
		}
	}

	html += fmt.Sprintf(`<line x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" stroke="var(--text-secondary)" stroke-width="2" />`, padding, height-padding, width-padding, height-padding)
	html += fmt.Sprintf(`<line x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" stroke="var(--text-secondary)" stroke-width="2" />`, padding, padding, padding, height-padding)

	if showX {
		for i, c := range data.Categories {
			x := padding + float64(i)*stepX
			html += fmt.Sprintf(`<text x="%.1f" y="%.1f" font-size="10" fill="currentColor" text-anchor="middle">%s</text>`, x, height-padding+15, c)
		}
	}

	legendHtml := `<div style="display:flex; justify-content:center; gap:20px; margin-top:10px;">`
	for sIdx, s := range data.Series {
		color := defaultColors[sIdx%len(defaultColors)]
		points := ""
		for i, v := range s.Data {
			if i >= len(data.Categories) { break }
			x := padding + float64(i)*stepX
			y := height - padding - ((v / max) * (height - 2*padding))
			points += fmt.Sprintf("%.1f,%.1f ", x, y)
			
			if showDots {
				html += fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="4" fill="%s" />`, x, y, color)
			}
			if showY {
				html += fmt.Sprintf(`<text x="%.1f" y="%.1f" font-size="10" fill="currentColor" text-anchor="middle">%.1f</text>`, x, y-10, v)
			}
		}
		html += fmt.Sprintf(`<polyline fill="none" stroke="%s" stroke-width="3" points="%s" />`, color, strings.TrimSpace(points))
		legendHtml += fmt.Sprintf(`<div style="display:flex; align-items:center; gap:5px;"><div style="width:12px;height:12px;background:%s;border-radius:2px;"></div><span style="font-size:0.85rem;">%s</span></div>`, color, s.Name)
	}
	html += `</svg>` + legendHtml + `</div></div>`
	return html
}

func generateGroupedBarChart(sourceDir, filename string, args map[string]string) string {
	content, err := os.ReadFile(filepath.Join(sourceDir, "data", filename))
	if err != nil { return err.Error() }
	return generateGroupedBarChartFromJSON(content, args)
}

func generateGroupedBarChartFromJSON(content []byte, args map[string]string) string {
	var data MultiSeriesChartData
	if err := json.Unmarshal(content, &data); err != nil { return err.Error() }

	max := 0.0
	for _, s := range data.Series {
		for _, v := range s.Data {
			if v > max { max = v }
		}
	}
	if max == 0 { max = 1 }

	showX := args["show-x"] != "false"
	showY := args["show-y"] != "false"
	gridX := args["grid-x"] == "true"
	gridY := args["grid-y"] == "true"

	html := `<div class="tamarind-groupedbar" style="margin: 2rem 0; width: 100%;">`
	if args["title"] != "" { html += fmt.Sprintf(`<h4 style="text-align:center; margin-bottom: 1rem;">%s</h4>`, args["title"]) }
	
	bgGradients := []string{}
	bgSizes := []string{}
	if gridY {
		bgGradients = append(bgGradients, `linear-gradient(var(--border-color) 1px, transparent 1px)`)
		bgSizes = append(bgSizes, `100% 20%`)
	}
	if gridX {
		bgGradients = append(bgGradients, `linear-gradient(90deg, var(--border-color) 1px, transparent 1px)`)
		bgSizes = append(bgSizes, `20% 100%`)
	}
	bgStyle := ""
	if len(bgGradients) > 0 {
		bgStyle = fmt.Sprintf(`background-image: %s; background-size: %s;`, strings.Join(bgGradients, ", "), strings.Join(bgSizes, ", "))
	}
	html += fmt.Sprintf(`<div class="tamarind-groupedbar-container" style="display:flex; align-items:flex-end; height: 250px; padding: 10px 0; border-bottom: 2px solid var(--border-color); border-left: 2px solid var(--border-color); overflow-x: auto; gap: 20px; %s">`, bgStyle)

	for i, cat := range data.Categories {
		html += `<div style="flex:1; display:flex; flex-direction:column; align-items:center; justify-content:flex-end; height:100%;">`
		html += `<div style="display:flex; align-items:flex-end; gap:4px; height:100%;">`
		for sIdx, s := range data.Series {
			if i >= len(s.Data) { continue }
			val := s.Data[i]
			heightPct := (val / max) * 100.0
			color := defaultColors[sIdx%len(defaultColors)]
			html += fmt.Sprintf(`<div style="display:flex; flex-direction:column; align-items:center; justify-content:flex-end; height:100%%;">`)
			if showY {
				html += fmt.Sprintf(`<div style="font-size:0.75rem; margin-bottom:4px; color: currentColor;">%.1f</div>`, val)
			}
			html += fmt.Sprintf(`<div style="width:20px; height:%.1f%%; background-color:%s; border-radius:3px 3px 0 0;" title="%s: %.1f"></div>`, heightPct, color, s.Name, val)
			html += `</div>`
		}
		html += `</div>`
		if showX {
			html += fmt.Sprintf(`<div style="font-size:0.8rem; margin-top:8px;">%s</div>`, cat)
		}
		html += `</div>`
	}
	html += `</div>`
	
	legendHtml := `<div style="display:flex; justify-content:center; gap:20px; margin-top:10px;">`
	for sIdx, s := range data.Series {
		color := defaultColors[sIdx%len(defaultColors)]
		legendHtml += fmt.Sprintf(`<div style="display:flex; align-items:center; gap:5px;"><div style="width:12px;height:12px;background:%s;border-radius:2px;"></div><span style="font-size:0.85rem;">%s</span></div>`, color, s.Name)
	}
	html += legendHtml + `</div></div>`
	return html
}

func generateRadarChart(sourceDir, filename string, args map[string]string) string {
	content, err := os.ReadFile(filepath.Join(sourceDir, "data", filename))
	if err != nil { return err.Error() }
	return generateRadarChartFromJSON(content, args)
}

func generateRadarChartFromJSON(content []byte, args map[string]string) string {
	var data MultiSeriesChartData
	if err := json.Unmarshal(content, &data); err != nil { return err.Error() }

	max := 0.0
	for _, s := range data.Series {
		for _, v := range s.Data {
			if v > max { max = v }
		}
	}
	if max == 0 { max = 1 }

	size, center, radius := 300.0, 150.0, 100.0
	html := `<div class="tamarind-radarchart" style="margin: 2rem 0; width: 100%; display:flex; flex-direction:column; align-items:center;">`
	if args["title"] != "" { html += fmt.Sprintf(`<h4 style="margin-bottom: 1rem;">%s</h4>`, args["title"]) }
	html += fmt.Sprintf(`<svg viewBox="0 0 %.0f %.0f" style="width: 100%%; max-width: 400px; height: auto;">`, size, size)
	
	sides := len(data.Categories)
	if sides < 3 { return "Error: Radar chart requires at least 3 categories." }

	// Draw web
	for r := 1; r <= 5; r++ {
		dist := radius * (float64(r) / 5.0)
		points := ""
		for i := 0; i < sides; i++ {
			angle := (math.Pi * 2 * float64(i) / float64(sides)) - (math.Pi / 2)
			x := center + dist*math.Cos(angle)
			y := center + dist*math.Sin(angle)
			points += fmt.Sprintf("%.1f,%.1f ", x, y)
		}
		html += fmt.Sprintf(`<polygon points="%s" fill="none" stroke="var(--text-secondary)" stroke-opacity="0.5" stroke-width="1" />`, points)
	}

	// Draw axes and labels
	for i := 0; i < sides; i++ {
		angle := (math.Pi * 2 * float64(i) / float64(sides)) - (math.Pi / 2)
		x := center + radius*math.Cos(angle)
		y := center + radius*math.Sin(angle)
		lx := center + (radius+20)*math.Cos(angle)
		ly := center + (radius+20)*math.Sin(angle)
		html += fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="var(--text-secondary)" stroke-opacity="0.5" stroke-width="1" />`, center, center, x, y)
		html += fmt.Sprintf(`<text x="%.1f" y="%.1f" font-size="10" fill="currentColor" text-anchor="middle" alignment-baseline="middle">%s</text>`, lx, ly, data.Categories[i])
	}

	// Draw series
	for sIdx, s := range data.Series {
		color := defaultColors[sIdx%len(defaultColors)]
		points := ""
		for i, v := range s.Data {
			if i >= sides { break }
			angle := (math.Pi * 2 * float64(i) / float64(sides)) - (math.Pi / 2)
			rVal := radius * (v / max)
			x := center + rVal*math.Cos(angle)
			y := center + rVal*math.Sin(angle)
			points += fmt.Sprintf("%.1f,%.1f ", x, y)
			html += fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="3" fill="%s" />`, x, y, color)
			html += fmt.Sprintf(`<text x="%.1f" y="%.1f" font-size="9" fill="currentColor" text-anchor="middle">%.1f</text>`, x, y-6, v)
		}
		html += fmt.Sprintf(`<polygon points="%s" fill="%s" fill-opacity="0.3" stroke="%s" stroke-width="2" />`, points, color, color)
	}
	html += `</svg>`
	
	legendHtml := `<div style="display:flex; justify-content:center; gap:20px; margin-top:10px;">`
	for sIdx, s := range data.Series {
		color := defaultColors[sIdx%len(defaultColors)]
		legendHtml += fmt.Sprintf(`<div style="display:flex; align-items:center; gap:5px;"><div style="width:12px;height:12px;background:%s;border-radius:2px;"></div><span style="font-size:0.85rem;">%s</span></div>`, color, s.Name)
	}
	html += legendHtml + `</div></div>`
	return html
}
