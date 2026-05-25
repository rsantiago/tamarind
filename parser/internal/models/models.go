// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

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
	BaseURL      string       // SEO
	CanonicalURL string       // SEO
	Image        string       // SEO (OG Image)
	CustomCSS    template.CSS // Config
	Paginator    Paginator
	Hidden       bool
	Data         map[string]interface{} // Data Files Support
	Author       string                 // Page-specific or Global Author
}

type Paginator struct {
	CurrentPage  int
	TotalPages   int
	HasPrev      bool
	HasNext      bool
	PrevURL      string
	NextURL      string
	VisiblePages []PageLink
}

type PageLink struct {
	Number    int
	URL       string
	IsCurrent bool
}

type ArticleMeta struct {
	Title       string
	Subtitle    string
	Date        string
	URL         string
	Tags        []string
	Hidden      bool
	Draft       bool // New Draft Field
	Description string
	SourcePath  string
	Author      string // Metadata
}

type MenuItem struct {
	Title string
	URL   string
	Order int // Sorting order
}

type FrontMatter struct {
	Title       string   `yaml:"title"`
	Subtitle    string   `yaml:"subtitle"`
	Date        string   `yaml:"date"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"` // SEO
	Image       string   `yaml:"image"`       // SEO
	Hidden      bool     `yaml:"hidden"`
	Draft       bool     `yaml:"draft"`      // New Draft Field
	SiteName    string   `yaml:"site_name"`  // Override Site Name
	MenuLabel   string   `yaml:"menu_label"` // Custom Menu Title
	MenuOrder   int      `yaml:"menu_order"` // Menu Sorting Order
	Author      string   `yaml:"author"`     // Override Author
}
