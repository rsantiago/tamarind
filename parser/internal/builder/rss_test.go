package builder

import (
	"encoding/xml"
	"github.com/rsantiago/tamarind/parser/internal/models"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateRSS(t *testing.T) {
	// Create temp dir for output
	tmpDir, err := os.MkdirTemp("", "tamarind-rss-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	articles := []models.ArticleMeta{
		{
			Title:       "Test Article 1",
			Date:        "2023-10-27",
			URL:         "articles/test-1.html",
			Description: "Description 1",
			Draft:       false,
			Hidden:      false,
		},
		{
			Title: "Test Draft",
			Date:  "2023-10-28",
			URL:   "articles/draft.html",
			Draft: true,
		},
	}

	baseURL := "https://example.com"
	siteName := "Test Site"
	desc := "Test Description"

	err = GenerateRSS(tmpDir, articles, baseURL, siteName, desc)
	if err != nil {
		t.Fatalf("GenerateRSS failed: %v", err)
	}

	// Verify file exists
	feedPath := filepath.Join(tmpDir, "feed.xml")
	if _, err := os.Stat(feedPath); os.IsNotExist(err) {
		t.Fatal("feed.xml was not created")
	}

	// Verify content
	content, err := os.ReadFile(feedPath)
	if err != nil {
		t.Fatal(err)
	}

	// Simple check for key elements
	var rss RSS
	if err := xml.Unmarshal(content, &rss); err != nil {
		t.Fatalf("Failed to unmarshal generated RSS: %v", err)
	}

	if rss.Channel.Title != siteName {
		t.Errorf("Expected title %s, got %s", siteName, rss.Channel.Title)
	}
	if len(rss.Channel.Items) != 1 {
		t.Errorf("Expected 1 item (excluding draft), got %d", len(rss.Channel.Items))
	}
	if rss.Channel.Items[0].Title != "Test Article 1" {
		t.Errorf("Expected item title 'Test Article 1', got %s", rss.Channel.Items[0].Title)
	}
	// Check Date Format (RFC1123Z)
	expectedDateSnippet := "Fri, 27 Oct 2023" // Part of standard string
	if len(rss.Channel.Items[0].PubDate) < len(expectedDateSnippet) {
		t.Errorf("Date seems invalid: %s", rss.Channel.Items[0].PubDate)
	}
}
