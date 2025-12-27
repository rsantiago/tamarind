package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/rsantiago/tamarind/parser/internal/models"
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

	data := models.PageData{
		Title:        fm.Title,
		Subtitle:     fm.Subtitle,
		Date:         fm.Date,
		Tags:         fm.Tags,
		Body:         template.HTML(bodyHTML.String()),
		RelPrefix:    relPrefix,
		Articles:     articles,
		Menu:         menu,
		SiteName:     siteName,
		BaseURL:      baseURL,
		CanonicalURL: canonicalURL,
		Description:  fm.Description,
		Image:        fm.Image,
		CustomCSS:    customCSS,
		Hidden:       fm.Hidden,
		Data:         siteData,
        Author:       author,
	}

	templateName := "page.mdt"
	if relPath == "articles.md" {
		templateName = "articles.mdt"
	}

	var output bytes.Buffer
	if err := tmpl.ExecuteTemplate(&output, templateName, data); err != nil {
		return fmt.Errorf("template execute %s: %w", srcPath, err)
	}

	if liveReload {
		output.WriteString(LiveReloadScript)
	}

	// Ensure target directory exists for nested pages (e.g., articles/foo.html)
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("mkdir all %s: %w", filepath.Dir(targetPath), err)
	}

	if err := os.WriteFile(targetPath, output.Bytes(), 0644); err != nil {
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
