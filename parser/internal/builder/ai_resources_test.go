// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rsantiago/tamarind/parser/internal/models"
)

func TestGenerateAIResources(t *testing.T) {
	// Setup temporary directory for test output
	tmpDir, err := os.MkdirTemp("", "tamarind-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create dummy source file for full text test
	sourceContent := "This is the raw markdown content of the article."
	sourceFile := filepath.Join(tmpDir, "article.md")
	if err := os.WriteFile(sourceFile, []byte(sourceContent), 0644); err != nil {
		t.Fatalf("Failed to write dummy source file: %v", err)
	}

	baseURL := "https://example.com"
	articles := []models.ArticleMeta{
		{
			Title:       "Test Article 1",
			URL:         "articles/test-1.html",
			Description: "Description 1",
			SourcePath:  sourceFile,
		},
		{
			Title:       "Test Article 2",
			URL:         "articles/test-2.html",
			Description: "Description 2",
			SourcePath:  sourceFile, // Reuse same source for simplicity
		},
	}

	// Mock pages (Tags are currently not passed as Models to this function, but Pages are)
	// We need to verify if Tags pages should be in llms.txt.
	// Usually LLMs prefer content, but navigation nodes (tags) are useful.
	pages := []models.PageData{
		{
			Title:       "About",
			Description: "About page description",
		},
	}

	// EXECUTE
	err = GenerateAIResources(tmpDir, articles, pages, baseURL, nil)
	if err != nil {
		t.Fatalf("GenerateAIResources failed: %v", err)
	}

	// VERIFY ROBOTS.TXT
	robotsBytes, err := os.ReadFile(filepath.Join(tmpDir, "robots.txt"))
	if err != nil {
		t.Fatalf("Failed to read robots.txt: %v", err)
	}
	robotsContent := string(robotsBytes)
	if !strings.Contains(robotsContent, "User-agent: GPTBot") {
		t.Errorf("robots.txt missing GPTBot permission")
	}
	if !strings.Contains(robotsContent, "Sitemap: https://example.com/sitemap.xml") {
		t.Errorf("robots.txt missing correct Sitemap URL")
	}

	// VERIFY LLMS.TXT
	llmsBytes, err := os.ReadFile(filepath.Join(tmpDir, "llms.txt"))
	if err != nil {
		t.Fatalf("Failed to read llms.txt: %v", err)
	}
	llmsContent := string(llmsBytes)

	expectedURLs := []string{
		"https://example.com/articles/test-1.html",
		"https://example.com/articles/test-2.html",
	}
	for _, url := range expectedURLs {
		if !strings.Contains(llmsContent, url) {
			t.Errorf("llms.txt missing URL: %s", url)
		}
	}

	if !strings.Contains(llmsContent, "Description 1") {
		t.Errorf("llms.txt missing descriptions")
	}

	// VERIFY LLMS_FULL.TXT
	fullBytes, err := os.ReadFile(filepath.Join(tmpDir, "llms_full.txt"))
	if err != nil {
		t.Fatalf("Failed to read llms_full.txt: %v", err)
	}
	fullContent := string(fullBytes)

	if !strings.Contains(fullContent, "Title: Test Article 1") {
		t.Errorf("llms_full.txt missing article title metadata")
	}

	// Critical Check: Does it contain the Body Content?
	if !strings.Contains(fullContent, sourceContent) {
		t.Errorf("llms_full.txt missing raw markdown body content. Got:\n%s", fullContent)
	}
}
