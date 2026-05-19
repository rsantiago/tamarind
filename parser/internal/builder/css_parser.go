package builder

import (
	"regexp"
	"strings"
)

// CSSAnalysis holds the parsed structure of a CSS file for verification.
type CSSAnalysis struct {
	Variables  map[string]bool // CSS custom properties found in :root (e.g. "--primary-color")
	Selectors  map[string]bool // CSS class selectors found (e.g. ".btn-primary")
	MediaRules []string        // Media query conditions (e.g. "(max-width: 768px)")
}

// AnalyzeCSS parses a CSS string and extracts variables, selectors, and media queries.
// This is a lightweight parser focused on verification, not a full CSS parser.
func AnalyzeCSS(cssContent string) (*CSSAnalysis, error) {
	analysis := &CSSAnalysis{
		Variables:  make(map[string]bool),
		Selectors:  make(map[string]bool),
		MediaRules: []string{},
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
	// Match :root { ... } and [data-theme="dark"] { ... } blocks
	blockPattern := regexp.MustCompile(`(?s)(?::root|\[data-theme[^\]]*\])\s*\{([^}]+)\}`)
	blocks := blockPattern.FindAllStringSubmatch(css, -1)

	varPattern := regexp.MustCompile(`(--[\w-]+)\s*:`)

	for _, block := range blocks {
		if len(block) > 1 {
			vars := varPattern.FindAllStringSubmatch(block[1], -1)
			for _, v := range vars {
				if len(v) > 1 {
					analysis.Variables[v[1]] = true
				}
			}
		}
	}
}

// extractSelectors finds all CSS class selectors (.class-name) in the file.
func extractSelectors(css string, analysis *CSSAnalysis) {
	// Remove comments first
	commentPattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	cleaned := commentPattern.ReplaceAllString(css, "")

	// Remove @import and @media blocks' conditions (but keep their content)
	// We want selectors inside media blocks too

	// Match class selectors: .class-name (possibly with pseudo-classes or combinators)
	// We capture just the base class name
	selectorPattern := regexp.MustCompile(`(?m)^[^{]*?(\.[\w-]+)[^{]*?\{`)
	matches := selectorPattern.FindAllStringSubmatch(cleaned, -1)

	for _, m := range matches {
		if len(m) > 1 {
			analysis.Selectors[m[1]] = true
		}
	}

	// Also catch selectors in compound rules like ".form-input, .form-select, .form-textarea {"
	// and selectors that appear mid-line
	multiSelectorPattern := regexp.MustCompile(`(\.[\w-]+)`)
	lines := strings.Split(cleaned, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Only process lines that look like selector lines (contain { or end before {)
		// Skip property lines (contain :) unless they also contain {
		if strings.Contains(trimmed, "{") || (len(trimmed) > 0 && !strings.Contains(trimmed, ":") && !strings.HasPrefix(trimmed, "@") && !strings.HasPrefix(trimmed, "}")) {
			found := multiSelectorPattern.FindAllStringSubmatch(trimmed, -1)
			for _, f := range found {
				if len(f) > 1 {
					analysis.Selectors[f[1]] = true
				}
			}
		}
	}
}

// extractMediaQueries finds all @media query conditions.
func extractMediaQueries(css string, analysis *CSSAnalysis) {
	mediaPattern := regexp.MustCompile(`@media\s*([^{]+)\{`)
	matches := mediaPattern.FindAllStringSubmatch(css, -1)

	for _, m := range matches {
		if len(m) > 1 {
			condition := strings.TrimSpace(m[1])
			analysis.MediaRules = append(analysis.MediaRules, condition)
		}
	}
}
