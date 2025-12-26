package builder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"github.com/rsantiago/tamarind/parser/internal/seo"
	"github.com/rsantiago/tamarind/parser/internal/models"
)

func TestEndToEnd_DraftExclusion(t *testing.T) {
	// Simulate an End-to-End Build process regarding SEO/AI files
	
	// 1. Setup Env
	tmpDir, err := os.MkdirTemp("", "tamarind-e2e-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// We are NOT running full Builder.Build() because that requires complex template setups.
	// Instead, we verify the logic segments that matter:
	// Scanner -> Article Filtering -> SEO/AI Generation logic

	// 2. Setup Dummy Articles (Result of Scanner)
	// Scenario: Scanner has already run with includeDrafts=false
	// This means the slice passed to resources MUST NOT contain the draft.
	
	// However, the USER asks if we are testing that they don't get listed.
	// The critical path is: if I accidentally pass a draft to GenerateAIResources, does it check?
	// Currently NO.
	// But Scanner ensures they are not passed.
	
	// Let's verify the `seo.GenerateSEOFiles` behavior by creating physical files.
	// Because that function scans the DISK.
	
	websiteDir := filepath.Join(tmpDir, "website")
	if err := os.Mkdir(websiteDir, 0755); err != nil {
		t.Fatalf("Mkdir website failed: %v", err)
	}

	// 3. Create Files in 'website' as if Builder generated them
	// Published file
	if err := os.WriteFile(filepath.Join(websiteDir, "published.html"), []byte("<html></html>"), 0644); err != nil {
		t.Fatalf("Write published failed: %v", err)
	}
	// Draft file - This should NOT exist if scanner worked. 
	// But let's assume one exists to check if SEO scanner blindly picks it up (it does).
	// Wait, if the file exists on disk, SEO scanner WILL pick it up.
	// So correctness relies entirely on Builder NOT creating the file.
	
	// So we must test `builder.pages.go:generatePage` or similar?
	// Actually, the main safeguard is `ScanCollections` filtering the list of items to process.
	
	// Let's test the `GenerateAIResources` function specifically to see if it FILTERS drafts if they happen to be passed?
	// Currently checking the code: It does NOT check .Draft field.
	
	// Let's create a test that verifies `GenerateAIResources` does NOT filter .Draft=true articles (documenting current behavior)
	// OR we should Modify the code to filter them effectively as a second safeguard?
	// Let's stick to testing the Integration: "Scanner Output + Resource Gen"
	
	// --- INTEGRATION TEST ---
	// 1. Input: Source Dir with Drafts
	// 2. Action: Scan (Draft=False)
	// 3. Action: Generate Resources
	// 4. Assert: Resources don't have draft
	
	srcDir := filepath.Join(tmpDir, "source")
	if err := os.MkdirAll(filepath.Join(srcDir, "posts"), 0755); err != nil {
		t.Fatalf("Mkdir source failed: %v", err)
	}
	
	// Create Draft Markdown
	if err := os.WriteFile(filepath.Join(srcDir, "posts", "draft.md"), []byte("---\ntitle: Draft\ndraft: true\n---\n"), 0644); err != nil {
		t.Fatalf("Write draft failed")
	}
	
	// Create Published Markdown
	if err := os.WriteFile(filepath.Join(srcDir, "posts", "pub.md"), []byte("---\ntitle: Pub\ndraft: false\n---\n"), 0644); err != nil {
		t.Fatalf("Write pub failed")
	}

	// ACTION 1: Scan
	includeDrafts := false
	collections, err := ScanCollections(srcDir, includeDrafts)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}
	
	var articles []models.ArticleMeta
	for _, items := range collections {
		articles = append(articles, items...)
	}
	
	// Assert Intermediary State
	for _, a := range articles {
		if a.Draft {
			t.Fatalf("Scanner failed to filter draft: %s", a.Title)
		}
	}
	
	// ACTION 2: Generate AI Resources
	// We use the filtered 'articles' list
	baseURL := "https://test.com"
	err = GenerateAIResources(websiteDir, articles, []models.PageData{}, baseURL, nil)
	if err != nil {
		t.Fatalf("GenAIResources failed: %v", err)
	}
	
	// ACTION 3: Generate SEO Resources 
	// (SEO scans websiteDir, but since we didn't generate HTML files in this test, sitemap is empty.
	// But let's verify robots.txt exists)
	err = seo.GenerateSEOFiles(websiteDir, baseURL)
	if err != nil {
		t.Fatalf("SEO Gen failed")
	}

	// ASSERTIONS on OUTPUT FILES
	
	// Check llms.txt
	llmsBytes, _ := os.ReadFile(filepath.Join(websiteDir, "llms.txt"))
	llmsContent := string(llmsBytes)
	
	if strings.Contains(llmsContent, "Draft") {
		t.Errorf("llms.txt contains draft title")
	}
	if !strings.Contains(llmsContent, "Pub") {
		t.Errorf("llms.txt missing published title")
	}
	
	// Check llms_full.txt
	fullBytes, _ := os.ReadFile(filepath.Join(websiteDir, "llms_full.txt"))
	fullContent := string(fullBytes)

	if strings.Contains(fullContent, "Title: Draft") {
		t.Errorf("llms_full.txt contains draft content")
	}
	
	// This confirms that "Drafts are not getting listed" purely because they are excluded upstream.
}
