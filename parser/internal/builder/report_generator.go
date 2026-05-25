// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// ThemeReport holds the verification results for a single theme.
type ThemeReport struct {
	Name      string
	Passed    []Requirement
	Failed    []Requirement
	TotalReq  int
	PassedReq int
}

// GenerateComplianceReport runs the verifier against all themes and writes
// an organized markdown report to the specified output path.
func GenerateComplianceReport(specPath, templatesDir, outputPath string) error {
	// 1. Parse spec
	specContent, err := os.ReadFile(specPath)
	if err != nil {
		return fmt.Errorf("failed to read spec: %w", err)
	}

	requirements, err := ParseSpecRequirements(string(specContent))
	if err != nil {
		return fmt.Errorf("failed to parse spec: %w", err)
	}

	// 2. Discover themes
	entries, err := os.ReadDir(templatesDir)
	if err != nil {
		return fmt.Errorf("failed to read templates dir: %w", err)
	}

	var themes []string
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "shared" {
			themes = append(themes, entry.Name())
		}
	}
	sort.Strings(themes)

	// Read shared/core.css to merge baseline requirements
	coreCSSPath := filepath.Join(templatesDir, "shared", "core.css")
	coreCSSContent, err := os.ReadFile(coreCSSPath)
	if err != nil {
		return fmt.Errorf("failed to read shared/core.css: %w", err)
	}
	coreCSSStr := string(coreCSSContent)

	// 3. Verify each theme
	var reports []ThemeReport
	for _, theme := range themes {
		cssPath := filepath.Join(templatesDir, theme, "style.css")
		cssContent, err := os.ReadFile(cssPath)
		if err != nil {
			return fmt.Errorf("failed to read CSS for %s: %w", theme, err)
		}

		combinedCSS := coreCSSStr + "\n" + string(cssContent)
		analysis, err := AnalyzeCSS(combinedCSS)
		if err != nil {
			return fmt.Errorf("failed to analyze CSS for %s: %w", theme, err)
		}
		analysis.ThemeDir = filepath.Join(templatesDir, theme)
		analysis.RawCSS = combinedCSS

		report := ThemeReport{Name: theme}
		required := filterRequired(requirements)
		report.TotalReq = len(required)

		for _, req := range required {
			if checkRequirement(analysis, req) {
				report.Passed = append(report.Passed, req)
			} else {
				report.Failed = append(report.Failed, req)
			}
		}
		report.PassedReq = len(report.Passed)
		reports = append(reports, report)
	}

	// 4. Write report
	var sb strings.Builder
	writeReport(&sb, reports, requirements)

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func checkRequirement(analysis *CSSAnalysis, req Requirement) bool {
	return CheckRequirement(analysis, req)
}

func filterRequired(requirements []Requirement) []Requirement {
	var filtered []Requirement
	for _, r := range requirements {
		if r.Required {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

func writeReport(sb *strings.Builder, reports []ThemeReport, allRequirements []Requirement) {
	sb.WriteString("# Tamarind Theme Compliance Report\n\n")
	sb.WriteString(fmt.Sprintf("**Generated**: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	required := filterRequired(allRequirements)
	sb.WriteString(fmt.Sprintf("**Total Requirements**: %d required, %d optional\n\n",
		len(required), len(allRequirements)-len(required)))

	// Summary table
	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Theme | Status | Passed | Failed | Score |\n")
	sb.WriteString("|-------|--------|--------|--------|-------|\n")

	for _, r := range reports {
		status := "✅ PASS"
		if len(r.Failed) > 0 {
			status = "❌ FAIL"
		}
		score := 0
		if r.TotalReq > 0 {
			score = (r.PassedReq * 100) / r.TotalReq
		}
		sb.WriteString(fmt.Sprintf("| **%s** | %s | %d | %d | %d%% |\n",
			r.Name, status, r.PassedReq, len(r.Failed), score))
	}

	sb.WriteString("\n---\n\n")

	// Detailed per-theme sections
	for _, r := range reports {
		writeThemeSection(sb, r)
	}
}

func writeThemeSection(sb *strings.Builder, report ThemeReport) {
	status := "✅ COMPLIANT"
	if len(report.Failed) > 0 {
		status = "❌ NON-COMPLIANT"
	}
	sb.WriteString(fmt.Sprintf("## Theme: %s — %s\n\n", report.Name, status))

	score := 0
	if report.TotalReq > 0 {
		score = (report.PassedReq * 100) / report.TotalReq
	}
	sb.WriteString(fmt.Sprintf("**Score**: %d/%d (%d%%)\n\n", report.PassedReq, report.TotalReq, score))

	if len(report.Failed) > 0 {
		sb.WriteString("### ❌ Missing Requirements\n\n")

		// Group failures by category
		grouped := groupByCategory(report.Failed)
		categories := sortedKeys(grouped)

		for _, cat := range categories {
			sb.WriteString(fmt.Sprintf("#### %s\n\n", cat))
			for _, req := range grouped[cat] {
				sb.WriteString(fmt.Sprintf("- `%s` — %s\n", req.Name, req.Description))
			}
			sb.WriteString("\n")
		}
	}

	if len(report.Passed) > 0 {
		sb.WriteString("### ✅ Passed Requirements\n\n")

		grouped := groupByCategory(report.Passed)
		categories := sortedKeys(grouped)

		for _, cat := range categories {
			sb.WriteString(fmt.Sprintf("#### %s\n\n", cat))
			for _, req := range grouped[cat] {
				sb.WriteString(fmt.Sprintf("- `%s` — %s\n", req.Name, req.Description))
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("---\n\n")
}

func groupByCategory(requirements []Requirement) map[string][]Requirement {
	grouped := make(map[string][]Requirement)
	for _, r := range requirements {
		grouped[r.Category] = append(grouped[r.Category], r)
	}
	return grouped
}

func sortedKeys(m map[string][]Requirement) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
