// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// MediaBlock holds a media query condition and its raw CSS block content.
type MediaBlock struct {
	Condition string
	Content   string
}

// CSSAnalysis holds the parsed structure of a CSS file for verification.
type CSSAnalysis struct {
	Variables   map[string]bool   // CSS custom properties found in :root (e.g. "--primary-color")
	LightVars   map[string]string // CSS variable values in light mode
	DarkVars    map[string]string // CSS variable values in dark mode
	Selectors   map[string]bool   // CSS class selectors found (e.g. ".btn-primary")
	MediaRules  []string          // Media query conditions (e.g. "(max-width: 768px)")
	MediaBlocks []MediaBlock      // Raw media query blocks for accessibility check
	ThemeDir    string            // Path to the theme directory for template-based validation
	RawCSS      string            // Raw CSS content for custom markup styling validation
}

// AnalyzeCSS parses a CSS string and extracts variables, selectors, and media queries.
// This is a lightweight parser focused on verification, not a full CSS parser.
func AnalyzeCSS(cssContent string) (*CSSAnalysis, error) {
	analysis := &CSSAnalysis{
		Variables:   make(map[string]bool),
		LightVars:   make(map[string]string),
		DarkVars:    make(map[string]string),
		Selectors:   make(map[string]bool),
		MediaRules:  []string{},
		MediaBlocks: []MediaBlock{},
	}

	// 1. Extract CSS variables from :root blocks (including [data-theme="dark"])
	extractVariables(cssContent, analysis)

	// 2. Extract class selectors
	extractSelectors(cssContent, analysis)

	// 3. Extract media queries
	extractMediaQueries(cssContent, analysis)

	return analysis, nil
}

// extractVariables finds all CSS custom properties (--*) defined in :root or [data-theme] blocks.
func extractVariables(css string, analysis *CSSAnalysis) {
	if analysis.LightVars == nil {
		analysis.LightVars = make(map[string]string)
	}
	if analysis.DarkVars == nil {
		analysis.DarkVars = make(map[string]string)
	}

	// 1. Remove comments
	commentPattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	cleaned := commentPattern.ReplaceAllString(css, "")

	// 2. State machine to extract all blocks (handling potential nesting from media queries)
	var pos int
	for pos < len(cleaned) {
		openBrace := strings.Index(cleaned[pos:], "{")
		if openBrace == -1 {
			break
		}
		openBraceIdx := pos + openBrace
		selector := strings.TrimSpace(cleaned[pos:openBraceIdx])

		// Find the matching closing brace
		startIdx := openBraceIdx + 1
		nBraces := 1
		endIdx := startIdx
		for endIdx < len(cleaned) {
			char := cleaned[endIdx]
			if char == '{' {
				nBraces++
			} else if char == '}' {
				nBraces--
				if nBraces == 0 {
					break
				}
			}
			endIdx++
		}

		if nBraces == 0 && endIdx < len(cleaned) {
			body := cleaned[startIdx:endIdx]
			processBlockVariables(selector, body, analysis)

			// If this is a media query, scan its body for nested selectors
			if strings.HasPrefix(strings.TrimSpace(selector), "@media") {
				extractVariables(body, analysis)
			}

			pos = endIdx + 1
		} else {
			pos = openBraceIdx + 1
		}
	}
}

func processBlockVariables(selector, body string, analysis *CSSAnalysis) {
	selector = strings.ToLower(selector)

	isRoot := strings.Contains(selector, ":root")
	isLight := strings.Contains(selector, "data-theme=\"light\"") || strings.Contains(selector, "data-theme='light'") || strings.Contains(selector, "data-theme=light")
	isDark := strings.Contains(selector, "data-theme=\"dark\"") || strings.Contains(selector, "data-theme='dark'") || strings.Contains(selector, "data-theme=dark")

	if !isRoot && !isLight && !isDark {
		return
	}

	// Match variables
	varPattern := regexp.MustCompile(`(--[\w-]+)\s*:\s*([^;}\n]+)`)
	matches := varPattern.FindAllStringSubmatch(body, -1)
	for _, m := range matches {
		if len(m) > 2 {
			varName := strings.TrimSpace(m[1])
			varVal := strings.TrimSpace(m[2])
			analysis.Variables[varName] = true

			if isDark {
				analysis.DarkVars[varName] = varVal
			} else {
				analysis.LightVars[varName] = varVal
			}
		}
	}
}

// ResolveVal recursively resolves var(--name) calls using the provided variable map.
func ResolveVal(val string, vars map[string]string) string {
	val = strings.TrimSpace(val)
	if !strings.Contains(val, "var(") {
		return val
	}

	varRegex := regexp.MustCompile(`var\(\s*(--[\w-]+)\s*(?:,\s*([^)]+))?\s*\)`)

	for i := 0; i < 5; i++ {
		matches := varRegex.FindStringSubmatch(val)
		if len(matches) == 0 {
			break
		}

		name := matches[1]
		defVal := ""
		if len(matches) > 2 {
			defVal = strings.TrimSpace(matches[2])
		}

		resolved, exists := vars[name]
		if !exists {
			if defVal != "" {
				resolved = defVal
			} else {
				break
			}
		}

		val = strings.Replace(val, matches[0], resolved, 1)
		val = strings.TrimSpace(val)
		if !strings.Contains(val, "var(") {
			break
		}
	}

	return val
}

// ParseColor parses standard CSS color representations (hex, rgb, rgba) into R, G, B floats (0.0 to 1.0).
func ParseColor(val string) (float64, float64, float64, bool) {
	val = strings.TrimSpace(strings.ToLower(val))
	if val == "" {
		return 0, 0, 0, false
	}

	// 1. Handle hex formats: #rgb, #rgba, #rrggbb, #rrggbbaa
	if strings.HasPrefix(val, "#") {
		hex := val[1:]
		if idx := strings.Index(hex, "/*"); idx != -1 {
			hex = strings.TrimSpace(hex[:idx])
		}
		if idx := strings.Index(hex, " "); idx != -1 {
			hex = strings.TrimSpace(hex[:idx])
		}

		var r, g, b uint64
		var err error
		switch len(hex) {
		case 3: // #rgb
			r, err = strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8)
			g, _ = strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8)
			b, _ = strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8)
		case 4: // #rgba
			r, err = strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8)
			g, _ = strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8)
			b, _ = strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8)
		case 6: // #rrggbb
			r, err = strconv.ParseUint(hex[0:2], 16, 8)
			g, _ = strconv.ParseUint(hex[2:4], 16, 8)
			b, _ = strconv.ParseUint(hex[4:6], 16, 8)
		case 8: // #rrggbbaa
			r, err = strconv.ParseUint(hex[0:2], 16, 8)
			g, _ = strconv.ParseUint(hex[2:4], 16, 8)
			b, _ = strconv.ParseUint(hex[4:6], 16, 8)
		default:
			return 0, 0, 0, false
		}
		if err != nil {
			return 0, 0, 0, false
		}
		return float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0, true
	}

	// 2. Handle rgb() / rgba() formats
	if strings.HasPrefix(val, "rgb") {
		s := val
		if strings.HasPrefix(s, "rgba") {
			s = s[4:]
		} else {
			s = s[3:]
		}
		s = strings.Trim(s, "() ")

		s = strings.ReplaceAll(s, ",", " ")
		s = strings.ReplaceAll(s, "/", " ")
		fields := strings.Fields(s)
		if len(fields) < 3 {
			return 0, 0, 0, false
		}

		parseChan := func(f string) (float64, bool) {
			f = strings.TrimSpace(f)
			if strings.HasSuffix(f, "%") {
				pct, err := strconv.ParseFloat(strings.TrimSuffix(f, "%"), 64)
				if err != nil {
					return 0, false
				}
				return pct / 100.0, true
			}
			val, err := strconv.ParseFloat(f, 64)
			if err != nil {
				return 0, false
			}
			return val / 255.0, true
		}

		r, ok1 := parseChan(fields[0])
		g, ok2 := parseChan(fields[1])
		b, ok3 := parseChan(fields[2])
		if !ok1 || !ok2 || !ok3 {
			return 0, 0, 0, false
		}
		return r, g, b, true
	}

	// 3. Handle basic color names
	switch val {
	case "white":
		return 1.0, 1.0, 1.0, true
	case "black":
		return 0.0, 0.0, 0.0, true
	case "transparent":
		return 0.0, 0.0, 0.0, true
	}

	return 0, 0, 0, false
}

// RelativeLuminance calculates the relative luminance of a color.
func RelativeLuminance(r, g, b float64) float64 {
	adjust := func(c float64) float64 {
		if c <= 0.03928 {
			return c / 12.92
		}
		return math.Pow((c+0.055)/1.055, 2.4)
	}
	rL := adjust(r)
	gL := adjust(g)
	bL := adjust(b)
	return 0.2126*rL + 0.7152*gL + 0.0722*bL
}

// ContrastRatio calculates the contrast ratio between two relative luminance values.
func ContrastRatio(l1, l2 float64) float64 {
	if l1 < l2 {
		l1, l2 = l2, l1
	}
	return (l1 + 0.05) / (l2 + 0.05)
}

// extractSelectors finds all CSS selectors (.class-name or specific tags like img/figure) in the file.
func extractSelectors(css string, analysis *CSSAnalysis) {
	// Remove comments first
	commentPattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	cleaned := commentPattern.ReplaceAllString(css, "")

	// Match class selectors safely (starts with dot, then letter, then words/dashes).
	// This avoids catastrophic backtracking on large strings without brackets.
	selectorPattern := regexp.MustCompile(`\.([a-zA-Z][\w-]*)`)

	matches := selectorPattern.FindAllStringSubmatch(cleaned, -1)
	for _, m := range matches {
		if len(m) > 1 {
			analysis.Selectors["."+m[1]] = true
		}
	}

	// Also extract tag selectors like img and figure
	parts := strings.Split(cleaned, "{")
	for i := 0; i < len(parts)-1; i++ {
		selectorPart := parts[i]
		if idx := strings.LastIndex(selectorPart, "}"); idx != -1 {
			selectorPart = selectorPart[idx+1:]
		}
		selectorPart = strings.TrimSpace(selectorPart)

		// Helper matching word boundaries
		matchWord := func(s, word string) bool {
			r := regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
			return r.MatchString(s)
		}

		if matchWord(selectorPart, "img") {
			analysis.Selectors["img"] = true
		}
		if matchWord(selectorPart, "figure") {
			analysis.Selectors["figure"] = true
		}
	}
}

// extractMediaQueries finds all @media query conditions and their inner contents.
func extractMediaQueries(css string, analysis *CSSAnalysis) {
	// Remove comments first to avoid tracking braces inside comments
	commentPattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	cleaned := commentPattern.ReplaceAllString(css, "")

	pos := 0
	for {
		idx := strings.Index(cleaned[pos:], "@media")
		if idx == -1 {
			break
		}

		mediaPos := pos + idx
		openBraceIdx := strings.Index(cleaned[mediaPos:], "{")
		if openBraceIdx == -1 {
			break
		}

		condition := strings.TrimSpace(cleaned[mediaPos+6 : mediaPos+openBraceIdx])
		analysis.MediaRules = append(analysis.MediaRules, condition)

		// State machine to find the matching closing brace for the media query block
		startIdx := mediaPos + openBraceIdx + 1
		nBraces := 1
		endIdx := startIdx

		for endIdx < len(cleaned) {
			char := cleaned[endIdx]
			if char == '{' {
				nBraces++
			} else if char == '}' {
				nBraces--
				if nBraces == 0 {
					break
				}
			}
			endIdx++
		}

		if nBraces == 0 && endIdx < len(cleaned) {
			content := cleaned[startIdx:endIdx]
			analysis.MediaBlocks = append(analysis.MediaBlocks, MediaBlock{
				Condition: condition,
				Content:   content,
			})
			pos = endIdx + 1
		} else {
			pos = mediaPos + openBraceIdx + 1
		}
	}
}
