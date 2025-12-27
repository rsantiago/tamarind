package builder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rsantiago/tamarind/parser/internal/models"
)

// GenerateAIResources creates files specifically for AI crawlers and Agents:
// 1. robots.txt (Authorized for AI bots)
// 2. llms.txt (Index of content)
// 3. llms_full.txt (Concatenated full content)
func GenerateAIResources(websiteDir string, articles []models.ArticleMeta, pages []models.PageData, baseURL string, tagsMap map[string][]models.ArticleMeta) error {
	if err := generateRobotsTxt(websiteDir, baseURL); err != nil {
		return err
	}
	if err := generateLLMsTxt(websiteDir, articles, pages, baseURL, tagsMap); err != nil {
		return err
	}
	log.Println("Generated llms.txt")
	if err := generateLLMsFullTxt(websiteDir, articles, pages); err != nil {
		return err
	}
	log.Println("Generated llms_full.txt")
	return nil
}

func ensureTrailingSlash(s string) string {
	if !strings.HasSuffix(s, "/") {
		return s + "/"
	}
	return s
}

func generateRobotsTxt(websiteDir, baseURL string) error {
	baseURL = ensureTrailingSlash(baseURL)
	content := `User-agent: *
Allow: /

# Explicitly allow AI bots for AEO (AI Engine Optimization)
User-agent: GPTBot
Allow: /

User-agent: CCBot
Allow: /

User-agent: Google-Extended
Allow: /

Sitemap: ` + baseURL + `sitemap.xml
`
	return os.WriteFile(filepath.Join(websiteDir, "robots.txt"), []byte(content), 0644)
}

func generateLLMsTxt(websiteDir string, articles []models.ArticleMeta, pages []models.PageData, baseURL string, tagsMap map[string][]models.ArticleMeta) error {
	baseURL = ensureTrailingSlash(baseURL)
	var sb strings.Builder

	sb.WriteString("# Site Context\n")
	sb.WriteString(fmt.Sprintf("Base URL: %s\n\n", baseURL))
	
	sb.WriteString("## Core Pages\n")
	// For "generatedPages" passed from builder, we need to handle their URLs.
	// Since PageData struct doesn't have a generic "URL" field for itself (it has RelPrefix, CanonicalURL, etc which are context dependent),
	// we will rely on the fact that we constructed them in builder.go with specific Titles we can map, OR
	// we should have added the URL to the 'CanonicalURL' field in builder.go and read it here.
	for _, p := range pages {
		url := p.CanonicalURL
		// If CanonicalURL is not set (it usually isn't unless explicitly set), try to infer or skip.
		// In builder.go we will set CanonicalURL to the relative path e.g. "articles.html" or "tags/foo.html"
		if url == "" {
			continue 
		}
		
		fullURL := baseURL + url
		sb.WriteString(fmt.Sprintf("- [%s](%s): %s\n", p.Title, fullURL, p.Description))
	}
	
	sb.WriteString("\n## Articles\n")
	for _, a := range articles {
		if a.Hidden {
			continue
		}
		url := baseURL + a.URL
		desc := a.Description
		if desc == "" {
			desc = "No summary available."
		}
		sb.WriteString(fmt.Sprintf("- [%s](%s): %s\n", a.Title, url, desc))
	}
	
	return os.WriteFile(filepath.Join(websiteDir, "llms.txt"), []byte(sb.String()), 0644)
}

func generateLLMsFullTxt(websiteDir string, articles []models.ArticleMeta, pages []models.PageData) error {
	var sb strings.Builder

	sb.WriteString("# Full Site Content Dump\n\n")

	// Iterate through articles to append their full markdown content
	for _, a := range articles {
		if a.Hidden {
			continue
		}
		
		sb.WriteString(fmt.Sprintf("---\nTitle: %s\nDate: %s\nURL: %s\n---\n\n", a.Title, a.Date, a.URL))
		

		// Read source content
		if a.SourcePath != "" {
			content, err := os.ReadFile(a.SourcePath)
			if err == nil {
				// We strip frontmatter? Usually yes, to keep token count lower and context clean.
				// But retaining it is also fine. Let's strip the delimiter to avoid confusion?
				// Actually, ParseFrontMatter splits it.
				// Let's just dump the raw content as is, it's safer.
				sb.Write(content)
				sb.WriteString("\n\n")
			}
		}
	}
	
	return os.WriteFile(filepath.Join(websiteDir, "llms_full.txt"), []byte(sb.String()), 0644)
}
