// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanCollections_Drafts(t *testing.T) {
	// 1. Setup temporary directory
	tempDir, err := os.MkdirTemp("", "tamarind-test-scan-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 2. Setup collection struct
	articlesDir := filepath.Join(tempDir, "articles")
	if err := os.Mkdir(articlesDir, 0755); err != nil {
		t.Fatalf("Failed to create articles dir: %v", err)
	}

	// 3. Create Test Files

	// File 1: Published (No Draft Field)
	publishedContent := `---
title: Published Article
date: 2023-01-01
---
Content`
	if err := os.WriteFile(filepath.Join(articlesDir, "published.md"), []byte(publishedContent), 0644); err != nil {
		t.Fatalf("Failed to write published.md: %v", err)
	}

	// File 2: Explicitly Draft
	draftContent := `---
title: Draft Article
date: 2023-01-02
draft: true
---
Content`
	if err := os.WriteFile(filepath.Join(articlesDir, "draft.md"), []byte(draftContent), 0644); err != nil {
		t.Fatalf("Failed to write draft.md: %v", err)
	}

	// File 3: Explicitly Published
	explicitContent := `---
title: Explicit Published
date: 2023-01-03
draft: false
---
Content`
	if err := os.WriteFile(filepath.Join(articlesDir, "explicit.md"), []byte(explicitContent), 0644); err != nil {
		t.Fatalf("Failed to write explicit.md: %v", err)
	}

	// 4. Test Scenario A: IncludeDrafts = false (Default Build)
	collections, err := ScanCollections(tempDir, false)
	if err != nil {
		t.Fatalf("ScanCollections failed: %v", err)
	}

	items, ok := collections["articles"]
	if !ok {
		t.Fatalf("Expected 'articles' collection not found")
	}

	// Should have 2 items: published and explicit
	if len(items) != 2 {
		t.Errorf("Expected 2 items when drafts excluded, got %d", len(items))
	}

	for _, item := range items {
		if item.Title == "Draft Article" {
			t.Errorf("Draft article should NOT be present when includeDrafts=false")
		}
	}

	// 5. Test Scenario B: IncludeDrafts = true (tamarind build -drafts)
	collectionsDrafts, err := ScanCollections(tempDir, true)
	if err != nil {
		t.Fatalf("ScanCollections failed: %v", err)
	}

	itemsDrafts, ok := collectionsDrafts["articles"]
	if !ok {
		t.Fatalf("Expected 'articles' collection not found")
	}

	// Should have 3 items: including draft
	if len(itemsDrafts) != 3 {
		t.Errorf("Expected 3 items when drafts included, got %d", len(itemsDrafts))
	}

	foundDraft := false
	for _, item := range itemsDrafts {
		if item.Title == "Draft Article" {
			foundDraft = true
			if !item.Draft {
				t.Errorf("Draft Article metadata 'Draft' should be true")
			}
		}
	}
	if !foundDraft {
		t.Errorf("Draft article expected but not found")
	}
}

func TestScanPagesAndCollections_MenuOrdering(t *testing.T) {
	// 1. Setup temporary directory
	tempDir, err := os.MkdirTemp("", "tamarind-test-menu-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 2. Setup root pages
	// index.md (special case, default Home)
	if err := os.WriteFile(filepath.Join(tempDir, "index.md"), []byte("---\ntitle: Home\n---\n"), 0644); err != nil {
		t.Fatalf("Failed to write index.md: %v", err)
	}
	// about.md with explicit order 5
	aboutContent := `---
title: About Us
menu_label: About
menu_order: 5
---
`
	if err := os.WriteFile(filepath.Join(tempDir, "about.md"), []byte(aboutContent), 0644); err != nil {
		t.Fatalf("Failed to write about.md: %v", err)
	}

	// 3. Setup folder collections and override files
	// docs collection directory
	docsDir := filepath.Join(tempDir, "docs")
	if err := os.Mkdir(docsDir, 0755); err != nil {
		t.Fatalf("Failed to create docs dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(docsDir, "guide.md"), []byte("---\ntitle: Guide\n---\n"), 0644); err != nil {
		t.Fatalf("Failed to write guide.md: %v", err)
	}

	// blog collection directory (no override file)
	blogDir := filepath.Join(tempDir, "blog")
	if err := os.Mkdir(blogDir, 0755); err != nil {
		t.Fatalf("Failed to create blog dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(blogDir, "post.md"), []byte("---\ntitle: Post\n---\n"), 0644); err != nil {
		t.Fatalf("Failed to write post.md: %v", err)
	}

	// docs.md override file at root setting order 10 and a custom label
	docsOverride := `---
title: Documentation Portal
menu_label: Documentation
menu_order: 10
---
`
	if err := os.WriteFile(filepath.Join(tempDir, "docs.md"), []byte(docsOverride), 0644); err != nil {
		t.Fatalf("Failed to write docs.md: %v", err)
	}

	// 4. Run Scanner
	collections, err := ScanCollections(tempDir, false)
	if err != nil {
		t.Fatalf("ScanCollections failed: %v", err)
	}

	menu, err := ScanPagesAndCollections(tempDir, collections)
	if err != nil {
		t.Fatalf("ScanPagesAndCollections failed: %v", err)
	}

	// 5. Assertions
	// Expected Order: Home (0) -> About (5) -> Documentation (10) -> Blog (99)
	if len(menu) != 4 {
		t.Fatalf("Expected exactly 4 menu items, got %d: %+v", len(menu), menu)
	}

	expected := []struct {
		Title string
		URL   string
		Order int
	}{
		{Title: "Home", URL: "index.html", Order: 0},
		{Title: "About", URL: "about.html", Order: 5},
		{Title: "Documentation", URL: "docs.html", Order: 10},
		{Title: "Blog", URL: "blog.html", Order: 99},
	}

	for i, exp := range expected {
		item := menu[i]
		if item.Title != exp.Title {
			t.Errorf("At index %d: expected Title %q, got %q", i, exp.Title, item.Title)
		}
		if item.URL != exp.URL {
			t.Errorf("At index %d: expected URL %q, got %q", i, exp.URL, item.URL)
		}
		if item.Order != exp.Order {
			t.Errorf("At index %d: expected Order %d, got %d", i, exp.Order, item.Order)
		}
	}
}

func TestParseFrontMatter_AttributionStyle(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Default AttributionStyle is empty string",
			content:  "---\ntitle: Test\n---\nContent",
			expected: "",
		},
		{
			name:     "Explicit None AttributionStyle",
			content:  "---\ntitle: Test\nattribution_style: none\n---\nContent",
			expected: "none",
		},
		{
			name:     "Explicit Date Only",
			content:  "---\ntitle: Test\nattribution_style: date-only\n---\nContent",
			expected: "date-only",
		},
		{
			name:     "Explicit Author Only",
			content:  "---\ntitle: Test\nattribution_style: author-only\n---\nContent",
			expected: "author-only",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm, _ := ParseFrontMatter([]byte(tt.content))

			if fm.AttributionStyle != tt.expected {
				t.Errorf("Expected AttributionStyle to be %q, got %q", tt.expected, fm.AttributionStyle)
			}
		})
	}
}
