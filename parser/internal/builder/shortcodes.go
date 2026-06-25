// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var unsafeGoldmark = goldmark.New(
	goldmark.WithExtensions(extension.GFM, extension.Table, extension.Strikethrough, extension.Linkify, extension.TaskList),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
)

func processShortcodes(markdown, sourceDir string) string {
	// Initialize and execute the new PluginRegistry (Phase 3 & 4 Rollout)
	registry := NewPluginRegistry()
	registerChartPlugins(registry)
	
	// Phase 3 & 5 Migrations
	registry.Register(NewMermaidPlugin())
	registry.Register(NewTabsPlugin())
	registry.Register(NewTerminalPlugin())
	registry.Register(NewMetricsPlugin())
	registry.Register(NewSocialRibbonPlugin())
	registry.Register(NewFeaturesPlugin())
	registry.Register(NewPricingPlugin())
	registry.Register(NewCapabilitiesGridPlugin())
	registry.Register(NewCapabilitiesPlugin())
	registry.Register(NewTimelinePlugin())
	registry.Register(NewDropdownPlugin())
	registry.Register(NewAccordionPlugin())
	registry.Register(NewIncludePlugin())
	registry.Register(NewGistPlugin())
	registry.Register(NewMathPlugin())
	registry.Register(NewYouTubePlugin())
	registry.Register(NewFigurePlugin())
	registry.Register(NewButtonPlugin())
	registry.Register(NewCardPlugin())
	registry.Register(NewAlertPlugin())
	registry.Register(NewBadgePlugin())
	registry.Register(NewFormPlugin())
	registry.Register(NewFormInputPlugin())
	registry.Register(NewFormTextareaPlugin())
	registry.Register(NewFormSelectPlugin())
	registry.Register(NewFormCheckboxPlugin())
	registry.Register(NewFormRadioGroupPlugin())
	registry.Register(NewFormFilePlugin())
	
	markdown = registry.ProcessShortcodes(markdown, sourceDir)
	// 0. Agent (Comment): {{ agent "instruction" }} -> Removed from output
	reAgent := regexp.MustCompile(`{{\s*agent\s+"(.*?)"\s*}}`)
	markdown = reAgent.ReplaceAllString(markdown, "")

	// Chart Widgets have been migrated to the PluginRegistry (plugin_chart.go)


	// Cleanup Escape Token {{!}}
	markdown = strings.ReplaceAll(markdown, "{{!}}", "")

	return markdown
}

func getBuiltInIconSvg(name string) string {
	switch strings.ToLower(name) {
	case "sparkles":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"></path></svg>`
	case "bolt":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z"></path></svg>`
	case "academic":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 14l9-5-9-5-9 5 9 5zm0 0v6M12 21a9.003 9.003 0 008.367-5.633M4.367 15.367A9.003 9.003 0 0012 21"></path></svg>`
	case "shield":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path></svg>`
	case "warning":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path></svg>`
	case "success":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	case "chevron":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"></path></svg>`
	case "globe":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9"></path></svg>`
	case "lock":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path></svg>`
	case "layout":
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v4a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path></svg>`
	case "info":
		fallthrough
	default:
		return `<svg fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
	}
}
