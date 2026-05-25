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
		return (strings.Contains(tplStr, target) || strings.Contains(tplStr, "footer.mdt"))
	}

	return false
}
