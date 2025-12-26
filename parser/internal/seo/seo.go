package seo

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

// GenerateSEOFiles generates Sitemap and Robots.txt
func GenerateSEOFiles(websiteDir, baseURL string) error {
	if baseURL == "" {
		return nil // Skip if no base URL
	}

	// 1. Robots.txt
	robotsContent := fmt.Sprintf("User-agent: *\nAllow: /\nSitemap: %s/sitemap.xml\n", baseURL)
	if err := os.WriteFile(filepath.Join(websiteDir, "robots.txt"), []byte(robotsContent), 0644); err != nil {
		return err
	}
	log.Println("Generated: website/robots.txt")

	// 2. Sitemap.xml
	var urls []string
	err := filepath.WalkDir(websiteDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".html" {
			rel, _ := filepath.Rel(websiteDir, path)
			// Normalize path separators to forward slashes for URLs
			rel = filepath.ToSlash(rel)
			
			// Handle index.html nicely (e.g. root)
			if rel == "index.html" {
				urls = append(urls, baseURL+"/")
			} else {
				urls = append(urls, baseURL+"/"+rel)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	sitemapTmpl := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
{{ range . }}
   <url>
      <loc>{{ . }}</loc>
      <lastmod>{{ now }}</lastmod>
      <changefreq>daily</changefreq>
      <priority>0.8</priority>
   </url>
{{ end }}
</urlset>`

	funcMap := template.FuncMap{
		"now": func() string {
			return time.Now().Format("2006-01-02")
		},
	}

	tmpl, err := template.New("sitemap").Funcs(funcMap).Parse(sitemapTmpl)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(websiteDir, "sitemap.xml"))
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tmpl.Execute(f, urls); err != nil {
		return err
	}
	log.Println("Generated: website/sitemap.xml")

	return nil
}
