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
