package builder

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Requirement represents a single testable requirement from THEME_SPEC.md
type Requirement struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`        // "css-variable", "css-selector", "media-query"
	Scope       string `yaml:"scope"`       // ":root", "global", "@media"
	Required    bool   `yaml:"required"`
	Category    string `yaml:"-"`           // Set from parent block
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

		switch req.Type {
		case "css-variable":
			if !analysis.Variables[req.Name] {
				failures = append(failures, fmt.Sprintf("[%s] Missing CSS variable: %s — %s", req.Category, req.Name, req.Description))
			}
		case "css-selector":
			if !analysis.Selectors[req.Name] {
				failures = append(failures, fmt.Sprintf("[%s] Missing CSS selector: %s — %s", req.Category, req.Name, req.Description))
			}
		case "media-query":
			found := false
			for _, rule := range analysis.MediaRules {
				if strings.Contains(rule, req.Name) {
					found = true
					break
				}
			}
			if !found {
				failures = append(failures, fmt.Sprintf("[%s] Missing media query containing: %s — %s", req.Category, req.Name, req.Description))
			}
		}
	}

	return failures
}
