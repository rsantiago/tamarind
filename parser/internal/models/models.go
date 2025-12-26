package models

import "html/template"

type PageData struct {
	Title        string
	Subtitle     string
	Description  string // SEO
	Date         string
	Tags         []string
	Body         template.HTML
	RelPrefix    string
	Articles     []ArticleMeta
	Menu         []MenuItem
	SiteName     string
	BaseURL      string // SEO
	CanonicalURL string // SEO
	Image        string // SEO (OG Image)
	CustomCSS    template.CSS // Config
	Paginator    Paginator
	Hidden       bool
	Data         map[string]interface{} // Data Files Support
}

type Paginator struct {
	CurrentPage int
	TotalPages  int
	HasPrev     bool
	HasNext     bool
	PrevURL     string
	NextURL     string
	VisiblePages []PageLink
}

type PageLink struct {
	Number    int
	URL       string
	IsCurrent bool
}

type ArticleMeta struct {
	Title    string
	Subtitle string
	Date     string
	URL      string
	Tags     []string
	Hidden   bool
	Draft    bool // New Draft Field
	Description string
	SourcePath  string
}

type MenuItem struct {
	Title string
	URL   string
}

type FrontMatter struct {
	Title       string   `yaml:"title"`
	Subtitle    string   `yaml:"subtitle"`
	Date        string   `yaml:"date"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"` // SEO
    Image       string   `yaml:"image"`       // SEO
	Hidden      bool     `yaml:"hidden"`
	Draft       bool     `yaml:"draft"`       // New Draft Field
}
