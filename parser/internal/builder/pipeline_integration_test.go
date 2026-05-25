package builder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFullBuildPipeline(t *testing.T) {
	// Create temporary workspace for source, templates, and website output
	tmpDir, err := os.MkdirTemp("", "tamarind-pipeline-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	sourceDir := filepath.Join(tmpDir, "source")
	templateDir := filepath.Join(tmpDir, "templates")
	sharedDir := filepath.Join(tmpDir, "shared")
	websiteDir := filepath.Join(tmpDir, "website")

	// Ensure all directories exist
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}
	if err := os.MkdirAll(templateDir, 0755); err != nil {
		t.Fatalf("Failed to create templates dir: %v", err)
	}
	if err := os.MkdirAll(sharedDir, 0755); err != nil {
		t.Fatalf("Failed to create shared dir: %v", err)
	}

	// Write mock templates (no {{define}} blocks since Go text/template associates the whole file content to the filename template)
	pageTemplateContent := `{{template "site-header" .}}<h1>{{.Title}}</h1><div class="content">{{.Body}}</div>{{template "site-footer" .}}`
	if err := os.WriteFile(filepath.Join(templateDir, "page.mdt"), []byte(pageTemplateContent), 0644); err != nil {
		t.Fatalf("Failed to write page template: %v", err)
	}

	articlesTemplateContent := `{{template "site-header" .}}<h1>Blog Articles</h1>{{range .Articles}}<h2>{{.Title}}</h2>{{end}}{{template "site-footer" .}}`
	if err := os.WriteFile(filepath.Join(templateDir, "articles.mdt"), []byte(articlesTemplateContent), 0644); err != nil {
		t.Fatalf("Failed to write articles template: %v", err)
	}

	headerTemplateContent := `{{define "site-header"}}<html><head><title>{{.Title}}</title></head><body>{{end}}`
	if err := os.WriteFile(filepath.Join(sharedDir, "site-header.mdt"), []byte(headerTemplateContent), 0644); err != nil {
		t.Fatalf("Failed to write header template: %v", err)
	}

	footerTemplateContent := `{{define "site-footer"}}</body></html>{{end}}`
	if err := os.WriteFile(filepath.Join(sharedDir, "site-footer.mdt"), []byte(footerTemplateContent), 0644); err != nil {
		t.Fatalf("Failed to write footer template: %v", err)
	}

	// Write source content
	indexContent := "---\ntitle: Home\n---\nWelcome to the site!"
	if err := os.WriteFile(filepath.Join(sourceDir, "index.md"), []byte(indexContent), 0644); err != nil {
		t.Fatalf("Failed to write index.md: %v", err)
	}

	postDir := filepath.Join(sourceDir, "posts")
	if err := os.MkdirAll(postDir, 0755); err != nil {
		t.Fatalf("Failed to create posts dir: %v", err)
	}

	postContent := "---\ntitle: First Post\ndate: 2026-05-23\ntags: [news]\n---\nThis is the first post."
	if err := os.WriteFile(filepath.Join(postDir, "first-post.md"), []byte(postContent), 0644); err != nil {
		t.Fatalf("Failed to write first-post.md: %v", err)
	}

	// Run Build pipeline
	baseURL := "https://example.com"
	themeConfig := map[string]string{
		"site_name": "Pipeline Test",
	}

	err = Build(sourceDir, templateDir, websiteDir, baseURL, themeConfig, false, false)
	if err != nil {
		t.Fatalf("Build pipeline failed: %v", err)
	}

	// Assertions on generated files
	// 1. Verify index.html exists and is generated
	indexHTMLPath := filepath.Join(websiteDir, "index.html")
	if _, err := os.Stat(indexHTMLPath); os.IsNotExist(err) {
		t.Errorf("Expected index.html to be generated, but it does not exist")
	}

	indexHTMLBytes, err := os.ReadFile(indexHTMLPath)
	if err != nil {
		t.Fatalf("Failed to read index.html: %v", err)
	}
	indexHTMLContent := string(indexHTMLBytes)

	if !strings.Contains(indexHTMLContent, "Welcome to the site!") {
		t.Errorf("index.html does not contain content: %q", indexHTMLContent)
	}
	if !strings.Contains(indexHTMLContent, "<h1>Home</h1>") {
		t.Errorf("index.html does not contain correct title: %q", indexHTMLContent)
	}

	// 2. Verify posts collection index exists (posts.html)
	postsHTMLPath := filepath.Join(websiteDir, "posts.html")
	if _, err := os.Stat(postsHTMLPath); os.IsNotExist(err) {
		t.Errorf("Expected posts.html to be generated, but it does not exist")
	}

	postsHTMLBytes, err := os.ReadFile(postsHTMLPath)
	if err != nil {
		t.Fatalf("Failed to read posts.html: %v", err)
	}
	postsHTMLContent := string(postsHTMLBytes)

	if !strings.Contains(postsHTMLContent, "First Post") {
		t.Errorf("posts.html does not list First Post: %q", postsHTMLContent)
	}

	// 3. Verify single article is generated (posts/first-post.html)
	singlePostPath := filepath.Join(websiteDir, "posts", "first-post.html")
	if _, err := os.Stat(singlePostPath); os.IsNotExist(err) {
		t.Errorf("Expected posts/first-post.html to be generated, but it does not exist")
	}

	singlePostBytes, err := os.ReadFile(singlePostPath)
	if err != nil {
		t.Fatalf("Failed to read posts/first-post.html: %v", err)
	}
	singlePostContent := string(singlePostBytes)

	if !strings.Contains(singlePostContent, "This is the first post.") {
		t.Errorf("single post HTML does not contain content: %q", singlePostContent)
	}
}
