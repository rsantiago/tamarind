// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"regexp"
	"strings"
)

// MediaBlock holds a media query condition and its raw CSS block content.
type MediaBlock struct {
	Condition string
	Content   string
}

// CSSAnalysis holds the parsed structure of a CSS file for verification.
type CSSAnalysis struct {
	Variables   map[string]bool // CSS custom properties found in :root (e.g. "--primary-color")
	Selectors   map[string]bool // CSS class selectors found (e.g. ".btn-primary")
	MediaRules  []string        // Media query conditions (e.g. "(max-width: 768px)")
	MediaBlocks []MediaBlock    // Raw media query blocks for accessibility check
	ThemeDir    string          // Path to the theme directory for template-based validation
	RawCSS      string          // Raw CSS content for custom markup styling validation
}

// AnalyzeCSS parses a CSS string and extracts variables, selectors, and media queries.
// This is a lightweight parser focused on verification, not a full CSS parser.
func AnalyzeCSS(cssContent string) (*CSSAnalysis, error) {
	analysis := &CSSAnalysis{
		Variables:   make(map[string]bool),
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
