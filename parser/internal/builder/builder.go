// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rsantiago/tamarind/parser/internal/models"
	"github.com/rsantiago/tamarind/parser/internal/seo"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Build generates the static website
func Build(sourceDir, templateDir, websiteDir, baseURL string, themeConfig map[string]string, includeDrafts bool, liveReload bool) error {
	log.Printf("Building site... Src: %s, Tmpl: %s, Drafts: %v, LiveReload: %v", sourceDir, templateDir, includeDrafts, liveReload)

	if err := os.MkdirAll(websiteDir, 0755); err != nil {
		return fmt.Errorf("failed to create website dir: %w", err)
	}

	// 1. Load Templates
	funcMap := template.FuncMap{
		"sanitizeTag": func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
		},
		"add": func(a, b int) int {
			return a + b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"mod": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a % b
		},
	}

	tmpl := template.New("").Funcs(funcMap)

	// Load Theme Templates AND Shared Templates
	// Note: We need to parse them cumulatively.

	// First, parse theme templates
	themeFiles, err := filepath.Glob(filepath.Join(templateDir, "*.mdt"))
	if err != nil {
		return fmt.Errorf("failed to list theme templates: %w", err)
	}

	// Second, parse shared templates
	// Assuming templateDir is ".../themes/blue", we go up to ".../themes" then "shared"
	baseTemplateDir := filepath.Dir(templateDir) // parser/assets/templates
	sharedDir := filepath.Join(baseTemplateDir, "shared")
	sharedFiles, err := filepath.Glob(filepath.Join(sharedDir, "*.mdt"))
	if err != nil {
		// It's okay if shared doesn't exist? Maybe stricter is better.
		// For now log warning or ignore.
	}

	allFiles := append(themeFiles, sharedFiles...)

	if len(allFiles) > 0 {
		_, err = tmpl.ParseFiles(allFiles...)
		if err != nil {
			return fmt.Errorf("failed to parse templates: %w", err)
		}
	} else {
		return fmt.Errorf("no template files found")
	}

	// 2. Prepare Markdown Parser
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Table, extension.Strikethrough, extension.Linkify, extension.TaskList),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var articles []models.ArticleMeta
	var tagsMap = make(map[string][]models.ArticleMeta)

	// Collect Site Name from config or default
	siteName := "Tamarind"
	if val, ok := themeConfig["site_name"]; ok {
		siteName = val
	}

	// Attempt to read parsing config from index.md
	indexPath := filepath.Join(sourceDir, "index.md")
	if content, err := os.ReadFile(indexPath); err == nil {
		fm, _ := ParseFrontMatter(content)
		if fm.SiteName != "" {
			siteName = fm.SiteName
		}
	}

	// 2a. Load Data Files
	siteData, err := LoadDataFiles(sourceDir)
	if err != nil {
		log.Printf("Warning: Failed to load data files: %v", err)
	}

	// 3. Scan Content
	log.Println("Scanning content...")
	collections, err := ScanCollections(sourceDir, includeDrafts)
	if err != nil {
		return fmt.Errorf("failed to scan content: %w", err)
	}

	// 3b. Generate Dynamic Menu (After scanning collections so we know them)
	menu, err := ScanPagesAndCollections(sourceDir, collections)
	if err != nil {
		log.Printf("Warning: Failed to scan menu: %v", err)
		// Fallback? nil is handled by templates usually (empty menu)
	}

	// Flatten collections into articles list, build tags map, and generate INDICES
	for name, items := range collections {
		articles = append(articles, items...)

		for _, article := range items {
			for _, tag := range article.Tags {
				t := strings.ToLower(tag)
				tagsMap[t] = append(tagsMap[t], article)
			}
		}

		// Generate Index for this collection (e.g. articles.html, docs.html)
		if err := generateCollectionIndex(name, items, tmpl, menu, siteName, baseURL, websiteDir, "", siteData, themeConfig, liveReload); err != nil {
			return fmt.Errorf("failed to generate index for %s: %w", name, err)
		}
	}

	// Sort Articles by Date (desc)
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Date > articles[j].Date
	})

	// Adjust Menu URLs based on context? No, they are relative from root.
	// Wait, relPrefix handles it.

	// Re-walk to generate pages injecting articles list
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		relPath, _ := filepath.Rel(sourceDir, path)
		if strings.HasPrefix(relPath, ".") {
			return nil
		}

		// Skip collection overrides at the root (e.g., constitution.md) to prevent overwriting dynamic collection indices
		if !strings.Contains(relPath, string(os.PathSeparator)) {
			base := strings.TrimSuffix(relPath, ".md")
			if _, isCollection := collections[base]; isCollection {
				return nil
			}
		}

		// If generating docs index or similar, we might want a subset of articles?
		// Currently passing all articles.

		return generatePage(path, sourceDir, websiteDir, md, tmpl, articles, menu, siteName, baseURL, template.CSS(""), siteData, liveReload)
	})
	if err != nil {
		return fmt.Errorf("failed to generate pages: %w", err)
	}

	// Generate Tag Pages
	if err := os.MkdirAll(filepath.Join(websiteDir, "tags"), 0755); err != nil {
		return err
	}
	for tag, taggedArticles := range tagsMap {
		// Sort tagged articles
		sort.Slice(taggedArticles, func(i, j int) bool {
			return taggedArticles[i].Date > taggedArticles[j].Date
		})

		if err := generateCollectionIndex("tags/"+tag, taggedArticles, tmpl, menu, siteName, baseURL, websiteDir, "", siteData, themeConfig, liveReload); err != nil {
			return fmt.Errorf("failed to generate tag page %s: %w", tag, err)
		}
	}

	// 4. Copy Static Assets
	if err := copyAssets(templateDir, websiteDir); err != nil {
		return fmt.Errorf("failed to copy assets: %w", err)
	}
	// 4b. Copy Shared Assets (like core.css)
	sharedDir = filepath.Join(filepath.Dir(templateDir), "shared")
	if err := copyAssets(sharedDir, websiteDir); err != nil {
		// Log warning but don't fail, shared dir might not have css
		log.Printf("Warning: failed to copy shared assets: %v", err)
	}

	// 5. Copy Site Resources
	if err := copySiteResources(sourceDir, websiteDir); err != nil {
		return fmt.Errorf("failed to copy site resources: %w", err)
	}

	// Collect all generated pages for AI resources
	// 1. Static Pages (already handled implicitly? No, we didn't track them.)
	// We need to re-scan or track them.
	// For now, let's manually add the important collection pages we just built.
	var generatedPages []models.PageData

	// Articles Index
	generatedPages = append(generatedPages, models.PageData{
		Title:        "Articles",
		Description:  "Index of all articles.",
		CanonicalURL: "articles.html",
	})

	// Add Tag Pages
	for tag := range tagsMap {
		generatedPages = append(generatedPages, models.PageData{
			Title:        "Tag: " + tag,
			Description:  "Articles tagged with " + tag,
			CanonicalURL: "tags/" + tag + ".html",
		})
	}

	// 6. Generate SEO Resources (Sitemap & Generic Robots)
	if err := seo.GenerateSEOFiles(websiteDir, baseURL); err != nil {
		log.Printf("Warning: failed to generate SEO files: %v", err)
	}

	// 7. Generate AI Resources (llms.txt, robots.txt override)
	// This will overwrite robots.txt with a more specific one for AI.
	if err := GenerateAIResources(websiteDir, articles, generatedPages, baseURL, tagsMap); err != nil {
		log.Printf("Warning: failed to generate AI resources: %v", err)
	}

	// 8. Generate RSS Feed
	siteDescription := "A static site built with Tamarind"
	if val, ok := themeConfig["site_description"]; ok {
		siteDescription = val
	}

	// Attempt to read description from index.md (Re-using logic or just reading it again)
	// We already read index.md at the start of Build function, but we didn't store the description in a variable accessible here.
	// Let's read it again for simplicity/local scoping.
	if content, err := os.ReadFile(filepath.Join(sourceDir, "index.md")); err == nil {
		fm, _ := ParseFrontMatter(content)
		if fm.Description != "" {
			siteDescription = fm.Description
		}
	}
	if err := GenerateRSS(websiteDir, articles, baseURL, siteName, siteDescription); err != nil {
		log.Printf("Warning: failed to generate RSS feed: %v", err)
	}

	return nil
}
