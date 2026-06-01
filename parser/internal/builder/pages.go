// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/rsantiago/tamarind/parser/internal/models"
	"github.com/yuin/goldmark"
)

const LiveReloadScript = `
<script>
	(function() {
		var evtSource = new EventSource("/livereload");
		evtSource.onmessage = function(event) {
			location.reload();
		};
		evtSource.onerror = function() {
			console.log("Live reload disconnected. Retrying...");
		};
	})();
</script>
`

func generatePage(srcPath, sourceDir, websiteDir string, md goldmark.Markdown, tmpl *template.Template, articles []models.ArticleMeta, menu []models.MenuItem, siteName, baseURL string, customCSS template.CSS, siteData map[string]interface{}, liveReload bool) error {
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("read file %s: %w", srcPath, err)
	}

	fm, bodyMarkdown := ParseFrontMatter(content)

	// Pre-process Shortcodes (e.g. YouTube)
	// We do this BEFORE template execution so {{< shortcode >}} isn't confused with {{ .Data }}
	// Although currently our shortcodes mostly use {{ figure ... }} which looks like actions.
	// If we run shortcodes first, they become HTML. HTML is safe in templates.
	bodyMarkdown = processShortcodes(bodyMarkdown, sourceDir)

	// Execute Body as Template (to allow {{ .Data.foo }})
	// We create a temporary template from the markdown body
	contentTmpl, err := template.New("content").Parse(bodyMarkdown)
	if err == nil {
		var buf bytes.Buffer
		// Context for the content template
		// Note: We can't pass full PageData because Body is missing.
		// We pass a subset relevant for content writing.
		dataCtx := struct {
			Data     map[string]interface{}
			Title    string
			Date     string
			SiteName string
			BaseURL  string
		}{
			Data:     siteData,
			Title:    fm.Title,
			Date:     fm.Date,
			SiteName: siteName,
			BaseURL:  baseURL,
		}

		if err := contentTmpl.Execute(&buf, dataCtx); err == nil {
			bodyMarkdown = buf.String()
		} else {
			// Warn but don't fail, user might have just written {{ without meaning a template
			log.Printf("Warning: Failed to execute template in content %s: %v", srcPath, err)
		}
	}

	var bodyHTML bytes.Buffer
	if err := md.Convert([]byte(bodyMarkdown), &bodyHTML); err != nil {
		return fmt.Errorf("markdown convert %s: %w", srcPath, err)
	}

	relPath, _ := filepath.Rel(sourceDir, srcPath)
	targetPath := filepath.Join(websiteDir, strings.TrimSuffix(relPath, ".md")+".html")

	depth := strings.Count(relPath, string(os.PathSeparator))
	relPrefix := strings.Repeat("../", depth)

	relPathRelative, _ := filepath.Rel(sourceDir, srcPath)
	webPath := strings.TrimSuffix(relPathRelative, ".md") + ".html"
	canonicalURL := ""
	if baseURL != "" {
		canonicalURL = baseURL + "/" + webPath
	}

	// Determine Author
	author := fm.Author
	if author == "" {
		if info, ok := siteData["info"].(map[string]interface{}); ok {
			if siteAuthor, ok := info["author"].(string); ok {
				author = siteAuthor
			}
		}
	}

	var contextualSidebar []models.SidebarItem
	if fm.ContextualSidebar != "" && !fm.Canvas {
		contextualSidebar = getContextualSidebar(sourceDir, websiteDir, srcPath, fm.ContextualSidebar, relPrefix)
	}

	data := models.PageData{
		Title:             fm.Title,
		Subtitle:          fm.Subtitle,
		Date:              fm.Date,
		Tags:              fm.Tags,
		Body:              template.HTML(bodyHTML.String()),
		RelPrefix:         relPrefix,
		Articles:          articles,
		Menu:              menu,
		SiteName:          siteName,
		BaseURL:           baseURL,
		CanonicalURL:      canonicalURL,
		Description:       fm.Description,
		Image:             fm.Image,
		CustomCSS:         customCSS,
		Hidden:            fm.Hidden,
		Canvas:            fm.Canvas,
		HideMenu:          fm.HideMenu,
		HideFooter:        fm.HideFooter,
		Data:              siteData,
		Author:            author,
		ContextualSidebar: contextualSidebar,
	}

	templateName := "page.mdt"
	if relPath == "articles.md" {
		templateName = "articles.mdt"
	}

	var output bytes.Buffer
	if err := tmpl.ExecuteTemplate(&output, templateName, data); err != nil {
		return fmt.Errorf("template execute %s: %w", srcPath, err)
	}

	htmlStr := output.String()
	htmlStr = postProcessSidebar(htmlStr, data.ContextualSidebar)
	htmlStr = postProcessCanvas(htmlStr, data.Canvas, data.HideMenu, data.HideFooter)

	var outputBytes []byte
	if liveReload {
		outputBytes = []byte(htmlStr + LiveReloadScript)
	} else {
		outputBytes = []byte(htmlStr)
	}

	// Ensure target directory exists for nested pages (e.g., articles/foo.html)
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("mkdir all %s: %w", filepath.Dir(targetPath), err)
	}

	if err := os.WriteFile(targetPath, outputBytes, 0644); err != nil {
		return fmt.Errorf("write file %s: %w", targetPath, err)
	}

	log.Printf("Generated: %s", targetPath)
	return nil
}

func generateCollectionIndex(name string, items []models.ArticleMeta, tmpl *template.Template, menu []models.MenuItem, siteName, baseURL, websiteDir string, customCSS template.CSS, siteData map[string]interface{}, themeConfig map[string]string, liveReload bool) error {
	pageSize := 10
	if val, ok := themeConfig["pagination-limit"]; ok {
		if limit, err := strconv.Atoi(val); err == nil && limit > 0 {
			pageSize = limit
		}
	}

	totalItems := len(items)
	totalPages := (totalItems + pageSize - 1) / pageSize
	if totalPages == 0 {
		totalPages = 1
	}

	// Adjust depth for nested pages (like tags/foo.html)
	depth := strings.Count(name, "/")
	relPrefix := strings.Repeat("../", depth)
	baseName := filepath.Base(name)

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * pageSize
		end := start + pageSize
		if end > totalItems {
			end = totalItems
		}

		var pageItems []models.ArticleMeta
		if totalItems > 0 {
			pageItems = items[start:end]
		}

		// Determine output filename relative to websiteDir
		// Page 1: name.html (e.g., tags/coding.html)
		// Page N: name-page-N.html (e.g., tags/coding-page-2.html)
		var outputRelPath string
		if page == 1 {
			outputRelPath = name + ".html"
		} else {
			outputRelPath = fmt.Sprintf("%s-page-%d.html", name, page)
		}

		targetPath := filepath.Join(websiteDir, outputRelPath)

		// Paginator Logic
		paginator := models.Paginator{
			CurrentPage: page,
			TotalPages:  totalPages,
			HasPrev:     page > 1,
			HasNext:     page < totalPages,
		}

		if paginator.HasPrev {
			if page == 2 {
				paginator.PrevURL = baseName + ".html"
			} else {
				paginator.PrevURL = fmt.Sprintf("%s-page-%d.html", baseName, page-1)
			}
		}

		if paginator.HasNext {
			paginator.NextURL = fmt.Sprintf("%s-page-%d.html", baseName, page+1)
		}

		// Calculate Visible Pages (Window of 5)
		startPage := page - 2
		endPage := page + 2

		if startPage < 1 {
			endPage += (1 - startPage)
			startPage = 1
		}
		if endPage > totalPages {
			startPage -= (endPage - totalPages)
			if startPage < 1 {
				startPage = 1
			}
			endPage = totalPages
		}

		for p := startPage; p <= endPage; p++ {
			url := fmt.Sprintf("%s-page-%d.html", baseName, p)
			if p == 1 {
				url = baseName + ".html"
			}
			paginator.VisiblePages = append(paginator.VisiblePages, models.PageLink{
				Number:    p,
				URL:       url,
				IsCurrent: p == page,
			})
		}

		data := models.PageData{
			Title:        strings.Title(baseName),
			Subtitle:     fmt.Sprintf("Index of %s (Page %d)", baseName, page),
			Articles:     pageItems,
			Menu:         menu,
			RelPrefix:    relPrefix,
			SiteName:     siteName,
			BaseURL:      baseURL,
			CanonicalURL: baseURL + "/" + outputRelPath,
			Description:  fmt.Sprintf("Browse %s on %s - Page %d", baseName, siteName, page),
			CustomCSS:    customCSS,
			Paginator:    paginator,
			Data:         siteData,
		}

		var output bytes.Buffer
		if err := tmpl.ExecuteTemplate(&output, "articles.mdt", data); err != nil {
			return err
		}

		if liveReload {
			output.WriteString(LiveReloadScript)
		}

		if err := os.WriteFile(targetPath, output.Bytes(), 0644); err != nil {
			return err
		}

		log.Printf("Generated Collection Page: %s", targetPath)
	}
	return nil
}

func postProcessCanvas(html string, isCanvas, hideMenu, hideFooter bool) string {
	if isCanvas || hideMenu {
		// Strip Site Header
		reHeader := regexp.MustCompile(`(?s)<header[^>]*class="[^"]*site-header[^"]*"[^>]*>.*?</header>`)
		html = reHeader.ReplaceAllString(html, "")

		// Strip Sidebar
		reSidebar := regexp.MustCompile(`(?s)<aside[^>]*class="[^"]*sidebar[^"]*"[^>]*>.*?</aside>`)
		html = reSidebar.ReplaceAllString(html, "")
	}

	if hideFooter {
		// Strip Footer
		reFooter := regexp.MustCompile(`(?s)<footer[^>]*class="[^"]*site-footer[^"]*"[^>]*>.*?</footer>`)
		html = reFooter.ReplaceAllString(html, "")
	}

	if isCanvas {
		// Inject canvas spacing resets to layout wrapper
		html = strings.ReplaceAll(html, "class=\"layout-container\"", "class=\"layout-container canvas-mode-active\"")
		html = strings.ReplaceAll(html, "class=\"layout-container page-container\"", "class=\"layout-container page-container canvas-mode-active\"")
		html = strings.ReplaceAll(html, "class=\"layout-container window\"", "class=\"layout-container window canvas-mode-active\"")
		html = strings.ReplaceAll(html, "class=\"layout-container page-layout\"", "class=\"layout-container page-layout canvas-mode-active\"")

		// Flatten the main post-article wrappers
		html = strings.ReplaceAll(html, "<article class=\"card\">", "<article class=\"canvas-mode-active\">")
		html = strings.ReplaceAll(html, "<article class=\"post\">", "<article class=\"canvas-mode-active\">")
		html = strings.ReplaceAll(html, "<article class=\"article-container\">", "<article class=\"canvas-mode-active\">")
		html = strings.ReplaceAll(html, "<article class=\"article-content\">", "<article class=\"canvas-mode-active\">")
		html = strings.ReplaceAll(html, "<article class=\"post-content\">", "<article class=\"canvas-mode-active\">")

		// Expand the reading container width limits
		html = strings.ReplaceAll(html, "class=\"post-content\"", "class=\"canvas-width-limit\"")
		html = strings.ReplaceAll(html, "class=\"article-content\"", "class=\"canvas-width-limit\"")
		html = strings.ReplaceAll(html, "class=\"article-content style-guide\"", "class=\"canvas-width-limit style-guide\"")
	}

	return html
}

func getContextualSidebar(sourceDir, websiteDir, currentSrcPath, folderName, relPrefix string) []models.SidebarItem {
	if folderName == "" {
		return nil
	}

	targetDir := filepath.Join(sourceDir, folderName)
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		log.Printf("Warning: Failed to read contextual sidebar folder %s: %v", targetDir, err)
		return nil
	}

	var items []models.SidebarItem
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		filePath := filepath.Join(targetDir, entry.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		fm, _ := ParseFrontMatter(content)
		if fm.Hidden || fm.Draft {
			continue
		}

		title := fm.Title
		if title == "" {
			title = strings.TrimSuffix(entry.Name(), ".md")
		}

		siblingWebPath := filepath.Join(folderName, strings.TrimSuffix(entry.Name(), ".md")+".html")
		url := relPrefix + filepath.ToSlash(siblingWebPath)

		isCurrent := (filepath.Clean(filePath) == filepath.Clean(currentSrcPath))

		items = append(items, models.SidebarItem{
			Title:     title,
			URL:       url,
			IsCurrent: isCurrent,
		})
	}
	return items
}

func postProcessSidebar(htmlStr string, sidebarItems []models.SidebarItem) string {
	if len(sidebarItems) == 0 {
		return htmlStr
	}

	var sb strings.Builder
	sb.WriteString("<aside class=\"sidebar sidebar-left context-sidebar\">\n")
	sb.WriteString("    <nav class=\"sidebar-nav\">\n")
	for _, item := range sidebarItems {
		activeClass := ""
		if item.IsCurrent {
			activeClass = " class=\"active\""
		}
		sb.WriteString(fmt.Sprintf("        <a href=\"%s\"%s>%s</a>\n", item.URL, activeClass, template.HTMLEscapeString(item.Title)))
	}
	sb.WriteString("    </nav>\n")
	sb.WriteString("</aside>\n")
	sidebarHTML := sb.String()

	reSidebar := regexp.MustCompile(`(?s)<aside[^>]*class="[^"]*sidebar[^"]*"[^>]*>.*?</aside>`)
	if reSidebar.MatchString(htmlStr) {
		return reSidebar.ReplaceAllString(htmlStr, sidebarHTML)
	}

	reLayout := regexp.MustCompile(`(?i)(<div[^>]*class="[^"]*layout-container[^"]*"[^>]*>|<main[^>]*class="[^"]*layout-container[^"]*"[^>]*>)`)
	if reLayout.MatchString(htmlStr) {
		return reLayout.ReplaceAllString(htmlStr, "$1\n"+sidebarHTML)
	}

	return htmlStr
}
