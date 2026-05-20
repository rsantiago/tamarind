package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// getAssetsDir returns the absolute path to the assets/templates directory.
// It works by navigating from the test file location up to the parser root.
func getAssetsDir() string {
	_, filename, _, _ := runtime.Caller(0)
	// filename is .../parser/internal/builder/theme_verifier_test.go
	// We need .../parser/assets/templates
	builderDir := filepath.Dir(filename)                // .../parser/internal/builder
	parserRoot := filepath.Join(builderDir, "..", "..") // .../parser
	return filepath.Join(parserRoot, "assets", "templates")
}

func TestAllThemesCompliance(t *testing.T) {
	assetsDir := getAssetsDir()

	// 1. Read and parse THEME_SPEC.md
	specPath := filepath.Join(assetsDir, "THEME_SPEC.md")
	specContent, err := os.ReadFile(specPath)
	if err != nil {
		t.Fatalf("Failed to read THEME_SPEC.md at %s: %v", specPath, err)
	}

	requirements, err := ParseSpecRequirements(string(specContent))
	if err != nil {
		t.Fatalf("Failed to parse spec requirements: %v", err)
	}

	if len(requirements) == 0 {
		t.Fatal("No requirements found in THEME_SPEC.md")
	}

	t.Logf("Loaded %d requirements from THEME_SPEC.md", len(requirements))

	// 2. Discover all themes (directories in assets/templates, excluding "shared")
	entries, err := os.ReadDir(assetsDir)
	if err != nil {
		t.Fatalf("Failed to read templates directory: %v", err)
	}

	var themes []string
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "shared" {
			themes = append(themes, entry.Name())
		}
	}

	if len(themes) == 0 {
		t.Fatal("No theme directories found")
	}

	t.Logf("Discovered %d themes: %v", len(themes), themes)

	// Read shared/core.css to merge baseline requirements
	coreCSSPath := filepath.Join(assetsDir, "shared", "core.css")
	coreCSSContent, err := os.ReadFile(coreCSSPath)
	if err != nil {
		t.Fatalf("Failed to read shared/core.css: %v", err)
	}
	coreCSSStr := string(coreCSSContent)

	// 3. Test each theme
	for _, theme := range themes {
		t.Run(theme, func(t *testing.T) {
			cssPath := filepath.Join(assetsDir, theme, "style.css")
			cssContent, err := os.ReadFile(cssPath)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", cssPath, err)
			}

			combinedCSS := coreCSSStr + "\n" + string(cssContent)
			analysis, err := AnalyzeCSS(combinedCSS)
			if err != nil {
				t.Fatalf("Failed to analyze CSS for theme %s: %v", theme, err)
			}

			// Run verification
			failures := VerifyTheme(analysis, requirements)

			if len(failures) > 0 {
				for _, f := range failures {
					t.Errorf("  ✗ %s", f)
				}
				t.Logf("\n  Theme '%s': %d/%d requirements failed", theme, len(failures), countRequired(requirements))
			} else {
				t.Logf("  Theme '%s': All %d requirements passed ✓", theme, countRequired(requirements))
			}
		})
	}
}

func TestCSSParserBasic(t *testing.T) {
	css := `
:root {
    --primary-color: #0095f6;
    --text-color: #262626;
    --radius-sm: 4px;
}

[data-theme="dark"] {
    --primary-color: #e0f1ff;
    --text-color: #f5f5f5;
}

.btn {
    display: inline-block;
    cursor: pointer;
}

.btn-primary {
    background: var(--primary-color);
}

.form-input,
.form-select,
.form-textarea {
    width: 100%;
}

@media (max-width: 768px) {
    .layout-container {
        padding: 0 16px;
    }
}
`
	analysis, err := AnalyzeCSS(css)
	if err != nil {
		t.Fatalf("AnalyzeCSS failed: %v", err)
	}

	// Check variables
	expectedVars := []string{"--primary-color", "--text-color", "--radius-sm"}
	for _, v := range expectedVars {
		if !analysis.Variables[v] {
			t.Errorf("Expected variable %s to be found", v)
		}
	}

	// Check selectors
	expectedSelectors := []string{".btn", ".btn-primary", ".form-input", ".form-select", ".form-textarea", ".layout-container"}
	for _, s := range expectedSelectors {
		if !analysis.Selectors[s] {
			t.Errorf("Expected selector %s to be found", s)
		}
	}

	// Check media queries
	if len(analysis.MediaRules) == 0 {
		t.Error("Expected at least one media query")
	}
	found := false
	for _, rule := range analysis.MediaRules {
		if rule == "(max-width: 768px)" {
			found = true
		}
	}
	if !found {
		t.Errorf("Expected media query '(max-width: 768px)', got: %v", analysis.MediaRules)
	}
}

func TestSpecParser(t *testing.T) {
	specContent := "# Test Spec\n\nSome human readable text.\n\n" +
		"```yaml requirements\n" +
		"category: variables\n" +
		"items:\n" +
		"  - name: \"--primary-color\"\n" +
		"    type: css-variable\n" +
		"    scope: \":root\"\n" +
		"    required: true\n" +
		"    description: \"Main color\"\n" +
		"  - name: \"--optional-var\"\n" +
		"    type: css-variable\n" +
		"    scope: \":root\"\n" +
		"    required: false\n" +
		"    description: \"Optional variable\"\n" +
		"```\n\n" +
		"More text.\n\n" +
		"```yaml requirements\n" +
		"category: components\n" +
		"items:\n" +
		"  - name: \".btn\"\n" +
		"    type: css-selector\n" +
		"    scope: global\n" +
		"    required: true\n" +
		"    description: \"Button base class\"\n" +
		"```\n"

	requirements, err := ParseSpecRequirements(specContent)
	if err != nil {
		t.Fatalf("ParseSpecRequirements failed: %v", err)
	}

	if len(requirements) != 3 {
		t.Fatalf("Expected 3 requirements, got %d", len(requirements))
	}

	// Check first requirement
	if requirements[0].Name != "--primary-color" {
		t.Errorf("Expected first requirement name '--primary-color', got '%s'", requirements[0].Name)
	}
	if requirements[0].Category != "variables" {
		t.Errorf("Expected category 'variables', got '%s'", requirements[0].Category)
	}
	if !requirements[0].Required {
		t.Error("Expected first requirement to be required")
	}

	// Check optional
	if requirements[1].Required {
		t.Error("Expected second requirement to be optional")
	}

	// Check second block
	if requirements[2].Name != ".btn" {
		t.Errorf("Expected third requirement name '.btn', got '%s'", requirements[2].Name)
	}
	if requirements[2].Category != "components" {
		t.Errorf("Expected category 'components', got '%s'", requirements[2].Category)
	}
}

func TestVerifyTheme(t *testing.T) {
	requirements := []Requirement{
		{Name: "--primary-color", Type: "css-variable", Required: true, Category: "variables", Description: "Main color"},
		{Name: "--missing-var", Type: "css-variable", Required: true, Category: "variables", Description: "Missing var"},
		{Name: ".btn", Type: "css-selector", Required: true, Category: "components", Description: "Button"},
		{Name: ".missing-class", Type: "css-selector", Required: true, Category: "components", Description: "Missing class"},
		{Name: "--optional", Type: "css-variable", Required: false, Category: "variables", Description: "Optional"},
	}

	analysis := &CSSAnalysis{
		Variables:  map[string]bool{"--primary-color": true},
		Selectors:  map[string]bool{".btn": true},
		MediaRules: []string{},
	}

	failures := VerifyTheme(analysis, requirements)

	if len(failures) != 2 {
		t.Fatalf("Expected 2 failures, got %d: %v", len(failures), failures)
	}

	// Optional requirement should not generate a failure
	for _, f := range failures {
		if f == "--optional" {
			t.Error("Optional requirement should not generate a failure")
		}
	}
}

// countRequired returns the number of required items in a requirement list.
func countRequired(requirements []Requirement) int {
	count := 0
	for _, r := range requirements {
		if r.Required {
			count++
		}
	}
	return count
}

// Unused import guard
var _ = fmt.Sprintf

func TestGenerateComplianceReport(t *testing.T) {
	assetsDir := getAssetsDir()
	specPath := filepath.Join(assetsDir, "THEME_SPEC.md")

	// Output report to the planning directory
	parserRoot := filepath.Join(filepath.Dir(specPath), "..", "..")
	projectRoot := filepath.Join(parserRoot, "..")
	outputPath := filepath.Join(projectRoot, "planning", "theme-compliance-report.md")

	err := GenerateComplianceReport(specPath, assetsDir, outputPath)
	if err != nil {
		t.Fatalf("GenerateComplianceReport failed: %v", err)
	}

	// Verify report was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatal("Report file was not created")
	}

	content, _ := os.ReadFile(outputPath)
	report := string(content)

	// Basic content checks
	if !strings.Contains(report, "# Tamarind Theme Compliance Report") {
		t.Error("Report missing title")
	}
	if !strings.Contains(report, "gram") {
		t.Error("Report missing gram theme")
	}
	if !strings.Contains(report, "✅ COMPLIANT") {
		t.Error("Report should show at least one compliant theme (gram)")
	}

	t.Logf("Report generated at: %s", outputPath)
}
