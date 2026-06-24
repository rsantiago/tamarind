// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Requirement represents a single testable requirement from THEME_SPEC.md
type Requirement struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`  // "css-variable", "css-selector", "media-query"
	Scope       string `yaml:"scope"` // ":root", "global", "@media"
	Required    bool   `yaml:"required"`
	Category    string `yaml:"-"` // Set from parent block
	Description string `yaml:"description"`
}

// requirementBlock matches the YAML structure inside ```yaml requirements blocks
type requirementBlock struct {
	Category string        `yaml:"category"`
	Items    []Requirement `yaml:"items"`
}

// ParseSpecRequirements reads THEME_SPEC.md content and extracts all
// machine-parseable requirement blocks (```yaml requirements ... ```).
func ParseSpecRequirements(specContent string) ([]Requirement, error) {
	var allRequirements []Requirement

	// Find all ```yaml requirements ... ``` fenced blocks
	pattern := regexp.MustCompile("(?s)```yaml requirements\n(.*?)```")
	matches := pattern.FindAllStringSubmatch(specContent, -1)

	if len(matches) == 0 {
		return nil, fmt.Errorf("no requirement blocks found in spec")
	}

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		yamlContent := strings.TrimSpace(match[1])

		var block requirementBlock
		if err := yaml.Unmarshal([]byte(yamlContent), &block); err != nil {
			return nil, fmt.Errorf("failed to parse requirement block: %w\nContent:\n%s", err, yamlContent)
		}

		// Assign category from the block to each item
		for i := range block.Items {
			block.Items[i].Category = block.Category
		}

		allRequirements = append(allRequirements, block.Items...)
	}

	return allRequirements, nil
}

// VerifyTheme checks a CSSAnalysis against a list of requirements.
// Returns a list of failures (empty means full compliance).
func VerifyTheme(analysis *CSSAnalysis, requirements []Requirement) []string {
	var failures []string

	for _, req := range requirements {
		if !req.Required {
			continue
		}

		if !CheckRequirement(analysis, req) {
			failures = append(failures, fmt.Sprintf("[%s] Failed: %s — %s", req.Category, req.Name, req.Description))
		}
	}

	return failures
}

// CheckRequirement evaluates if a single requirement is met by the given CSSAnalysis.
func CheckRequirement(analysis *CSSAnalysis, req Requirement) bool {
	switch req.Type {
	case "contrast-ratio":
		if req.Name == "background-contrast" {
			return verifyBackgroundContrast(analysis)
		} else if req.Name == "chart-color-contrast" {
			return verifyChartColorsContrast(analysis)
		}
		return true

	case "css-variable":
		return analysis.Variables[req.Name]

	case "css-selector":
		return analysis.Selectors[req.Name]

	case "media-query":
		for _, rule := range analysis.MediaRules {
			if strings.Contains(rule, req.Name) {
				return true
			}
		}
		return false

	case "responsive-nav":
		var targetSelector string
		if req.Name == "mobile-nav-visible" {
			targetSelector = ".nav-links"
		} else if req.Name == "mobile-sidebar-visible" {
			targetSelector = ".sidebar"
		} else {
			return false
		}
		// Check display: none inside media blocks
		hiddenPattern := regexp.MustCompile(regexp.QuoteMeta(targetSelector) + `\s*\{\s*display\s*:\s*none`)
		for _, block := range analysis.MediaBlocks {
			cond := strings.ToLower(block.Condition)
			if strings.Contains(cond, "max-width") && (strings.Contains(cond, "768") || strings.Contains(cond, "600")) {
				if hiddenPattern.MatchString(block.Content) {
					return false
				}
			}
		}
		return true

	case "article-markup":
		if analysis.RawCSS == "" {
			return false
		}
		commentPattern := regexp.MustCompile(`(?s)/\*.*?\*/`)
		cleaned := commentPattern.ReplaceAllString(analysis.RawCSS, "")
		parts := strings.Split(cleaned, "{")
		var selectors []string
		for i := 0; i < len(parts)-1; i++ {
			part := parts[i]
			if idx := strings.LastIndex(part, "}"); idx != -1 {
				part = part[idx+1:]
			}
			selectors = append(selectors, strings.TrimSpace(part))
		}

		hasMatch := func(patterns ...string) bool {
			for _, pat := range patterns {
				r := regexp.MustCompile(pat)
				for _, sel := range selectors {
					normalizedSel := strings.Join(strings.Fields(sel), " ")
					if r.MatchString(normalizedSel) {
						return true
					}
				}
			}
			return false
		}

		tag := req.Name
		return hasMatch(`\b\.article-content\s+`+regexp.QuoteMeta(tag)+`\b`, `\b`+regexp.QuoteMeta(tag)+`\b`)

	case "template-feature":
		if analysis.ThemeDir == "" {
			return false
		}

		if req.Name == "menu-no-duplication" {
			for _, filename := range []string{"articles.mdt", "page.mdt"} {
				tplPath := filepath.Join(analysis.ThemeDir, filename)
				tplContent, err := os.ReadFile(tplPath)
				if err != nil {
					return false
				}
				if strings.Count(string(tplContent), "range .Menu") > 1 {
					return false
				}
			}
			return true
		}

		parts := strings.SplitN(req.Name, ":", 2)
		if len(parts) != 2 {
			return false
		}
		filename := parts[0]
		target := parts[1]

		tplPath := filepath.Join(analysis.ThemeDir, filename)
		tplContent, err := os.ReadFile(tplPath)
		if err != nil {
			return false
		}
		tplStr := string(tplContent)
		if target == "tamarind-ghost-badge" {
			return strings.Contains(tplStr, target)
		}
		return (strings.Contains(tplStr, target) || strings.Contains(tplStr, "footer.mdt"))
	}

	return false
}

func verifyBackgroundContrast(analysis *CSSAnalysis) bool {
	// 1. Resolve light mode variables
	lightVars := analysis.LightVars
	lightBgVal := ResolveVal(lightVars["--background-color"], lightVars)
	lightCardVal := ResolveVal(lightVars["--card-bg"], lightVars)

	// 2. Resolve dark mode variables (inheriting light mode ones)
	darkVars := make(map[string]string)
	for k, v := range analysis.LightVars {
		darkVars[k] = v
	}
	for k, v := range analysis.DarkVars {
		darkVars[k] = v
	}
	darkBgVal := ResolveVal(darkVars["--background-color"], darkVars)
	darkCardVal := ResolveVal(darkVars["--card-bg"], darkVars)

	// 3. Verify light mode contrast
	r1, g1, b1, ok1 := ParseColor(lightBgVal)
	r2, g2, b2, ok2 := ParseColor(lightCardVal)
	if !ok1 || !ok2 {
		return false
	}
	lLightBg := RelativeLuminance(r1, g1, b1)
	lLightCard := RelativeLuminance(r2, g2, b2)
	lightRatio := ContrastRatio(lLightBg, lLightCard)
	
	if lightRatio < 1.015 {
		return false
	}

	// 4. Verify dark mode contrast
	r3, g3, b3, ok3 := ParseColor(darkBgVal)
	r4, g4, b4, ok4 := ParseColor(darkCardVal)
	if !ok3 || !ok4 {
		return false
	}
	lDarkBg := RelativeLuminance(r3, g3, b3)
	lDarkCard := RelativeLuminance(r4, g4, b4)
	darkRatio := ContrastRatio(lDarkBg, lDarkCard)

	if darkRatio < 1.015 {
		return false
	}

	return true
}

func verifyChartColorsContrast(analysis *CSSAnalysis) bool {
	lightVars := analysis.LightVars
	darkVars := make(map[string]string)
	for k, v := range analysis.LightVars { darkVars[k] = v }
	for k, v := range analysis.DarkVars { darkVars[k] = v }

	getChartColor := func(vars map[string]string, index int) string {
		cName := fmt.Sprintf("--chart-%d", index)
		if val := vars[cName]; val != "" {
			return ResolveVal(val, vars)
		}
		switch index {
		case 1: return ResolveVal(vars["--primary-color"], vars)
		case 2: return ResolveVal(vars["--secondary-color"], vars)
		case 3: return "#58508d"
		case 4: return "#ffa600"
		case 5: return "#2f4b7c"
		case 6: return "#00ba38"
		case 7: return "#1e2e33"
		case 8: return "#ff6361"
		case 9: return "#a05195"
		}
		return ""
	}

	checkContrast := func(vars map[string]string) bool {
		bgVal := vars["--card-bg"]
		if bgVal == "" { bgVal = vars["--background-color"] }
		bgVal = ResolveVal(bgVal, vars)
		br, bg, bb, bgOk := ParseColor(bgVal)
		bgL := RelativeLuminance(br, bg, bb)

		for i := 1; i <= 9; i++ {
			ci := getChartColor(vars, i)
			ri, gi, bi, oki := ParseColor(ci)
			if !oki { continue }
			li := RelativeLuminance(ri, gi, bi)

			if bgOk {
				if ContrastRatio(li, bgL) < 1.35 { return false }
			}

			if i < 9 {
				cj := getChartColor(vars, i+1)
				rj, gj, bj, okj := ParseColor(cj)
				if !okj { continue }
				lj := RelativeLuminance(rj, gj, bj)
				if ContrastRatio(li, lj) < 1.35 { return false }
			}
		}
		return true
	}

	return checkContrast(lightVars) && checkContrast(darkVars)
}
