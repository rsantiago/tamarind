package builder

import (
	"encoding/xml"
	"fmt"
	"github.com/rsantiago/tamarind/parser/internal/models"
	"os"
	"path/filepath"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
}

func GenerateRSS(websiteDir string, articles []models.ArticleMeta, baseURL, siteName, siteDescription string) error {
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       siteName,
			Link:        baseURL,
			Description: siteDescription,
		},
	}

	for _, article := range articles {
		// Skip drafts and hidden posts
		if article.Draft || article.Hidden {
			continue
		}

		pubDate := parseDate(article.Date)
		link := fmt.Sprintf("%s/%s", baseURL, article.URL)

		rss.Channel.Items = append(rss.Channel.Items, Item{
			Title:       article.Title,
			Link:        link,
			Description: article.Description,
			PubDate:     pubDate,
			Guid:        link,
		})
	}

	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal RSS: %w", err)
	}

	header := []byte(xml.Header)
	output = append(header, output...)

	outputPath := filepath.Join(websiteDir, "feed.xml")
	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write RSS feed: %w", err)
	}

	return nil
}

func parseDate(dateStr string) string {
	// Try parsing YYYY-MM-DD which is common in Tamarind
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		// If fails, try to return as is or current time?
		// Ideally valid RSS date: RFC1123Z
		return dateStr
	}
	return t.Format(time.RFC1123Z)
}
