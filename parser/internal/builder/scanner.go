package builder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rsantiago/tamarind/parser/internal/models"
	"gopkg.in/yaml.v3"
)

func ScanCollections(sourceDir string, includeDrafts bool) (map[string][]models.ArticleMeta, error) {
	collections := make(map[string][]models.ArticleMeta)
	
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		// Skip special directories
		if name == "pages" || name == "templates" || name == "." || name == ".." || strings.HasPrefix(name, ".") {
			continue
		}

		// This is a collection directory (e.g. articles, blog, docs)
		items, err := scanDirForMarkdown(filepath.Join(sourceDir, name), name, includeDrafts)
		if err != nil {
			log.Printf("Warning: failed to scan collection %s: %v", name, err)
			continue
		}
		if len(items) > 0 {
			collections[name] = items
		}
	}
	return collections, nil
}

func scanDirForMarkdown(dirPath, collectionName string, includeDrafts bool) ([]models.ArticleMeta, error) {
	var items []models.ArticleMeta
	entries, err := os.ReadDir(dirPath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".md" {
			path := filepath.Join(dirPath, entry.Name())
			absPath, _ := filepath.Abs(path) // Ignore error, best effort
			fmt.Printf("DEBUG SCANNER: Scanned %s -> Abs: %s\n", entry.Name(), absPath)
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			fm, _ := ParseFrontMatter(content)
			
			if fm.Hidden {
				continue
			}

			if fm.Draft && !includeDrafts {
				continue
			}

			items = append(items, models.ArticleMeta{
				Title:    fm.Title,
				Subtitle: fm.Subtitle,
				Date:     fm.Date,
				Tags:     fm.Tags,
				// URL is relative to website root, e.g. "blog/post.html"
				URL:      collectionName + "/" + strings.TrimSuffix(entry.Name(), ".md") + ".html",
				Hidden:   fm.Hidden,
				Draft:    fm.Draft,
				Description: fm.Description,
				SourcePath:  absPath,
                Author:      fm.Author,
			})
		}
	}
	return items, nil
}

func parseFrontMatter(content []byte) (models.FrontMatter, string) {
    return ParseFrontMatter(content)
}

func ParseFrontMatter(content []byte) (models.FrontMatter, string) {
	strContent := string(content)
	var fm models.FrontMatter
	
	if strings.HasPrefix(strContent, "---") {
		parts := strings.SplitN(strContent, "---", 3)
		if len(parts) == 3 {
			if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
				log.Printf("Warning: Frontmatter parse error: %v", err)
			}
			return fm, parts[2]
		}
	}
	return fm, strContent
}

func ScanPagesAndCollections(sourceDir string, collections map[string][]models.ArticleMeta) ([]models.MenuItem, error) {
	var menu []models.MenuItem
	
	// Scan root directory for MD files
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".md" {
			path := filepath.Join(sourceDir, entry.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			fm, _ := ParseFrontMatter(content)
			
			if fm.Hidden {
				continue
			}

            // Determine Label
            label := fm.MenuLabel
            if label == "" {
                // Special case for index.md
                if entry.Name() == "index.md" {
                    label = "Home"
                } else {
                    // Title Case filename
                    base := strings.TrimSuffix(entry.Name(), ".md")
                    label = strings.Title(base)
                }
            }

            // Determine URL
            url := strings.TrimSuffix(entry.Name(), ".md") + ".html"
            
            // Determine Order
            order := fm.MenuOrder
            
            menu = append(menu, models.MenuItem{
                Title: label,
                URL:   url,
                Order: order,
            })
		}
	}

	// Add Collections (Order 99+)
	for name := range collections {
		title := strings.Title(name)
		menu = append(menu, models.MenuItem{
			Title: title,
			URL:   name + ".html", 
			Order: 99, // User specified default high order for collections
		})
	}
    
    // Sort Menu
    sort.Slice(menu, func(i, j int) bool {
        if menu[i].Order != menu[j].Order {
            return menu[i].Order < menu[j].Order
        }
        return menu[i].Title < menu[j].Title
    })

	return menu, nil
}
