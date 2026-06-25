// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
	pricingGridCount := 0
	tabsScriptInjected := false
	terminalScriptInjected := false
	// 0. Agent (Comment): {{ agent "instruction" }} -> Removed from output
	reAgent := regexp.MustCompile(`{{\s*agent\s+"(.*?)"\s*}}`)
	markdown = reAgent.ReplaceAllString(markdown, "")

	mermaidScriptInjected := false
	// 1. Mermaid (Block): {{ mermaid }}...{{ /mermaid }}
	reMermaid := regexp.MustCompile(`(?s){{\s*mermaid\s*}}(.*?){{\s*/mermaid\s*}}`)
	markdown = reMermaid.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reMermaid.FindStringSubmatch(match)
		content := submatch[1]

		jsScript := ""
		if !mermaidScriptInjected {
			jsScript = `
<script>
if (typeof tamarindMaximizeMermaid !== 'function') {
	window.tamarindMaximizeMermaid = function(btn) {
		var wrapper = btn.closest('.mermaid-wrapper');
		if (wrapper.classList.contains('maximized')) {
			wrapper.classList.remove('maximized');
			btn.innerText = 'Maximize';
			document.body.style.overflow = '';
		} else {
			wrapper.classList.add('maximized');
			btn.innerText = 'Close';
			document.body.style.overflow = 'hidden';
		}
	};
}
</script>
<style>
.mermaid-wrapper { position: relative; }
.mermaid-maximize-btn {
	position: absolute;
	top: 10px;
	right: 10px;
	z-index: 10;
	padding: 4px 8px;
	background: var(--primary, #333);
	color: #fff;
	border: none;
	border-radius: 4px;
	cursor: pointer;
	font-size: 12px;
	opacity: 0;
	transition: opacity 0.2s;
}
.mermaid-wrapper:hover .mermaid-maximize-btn {
	opacity: 1;
}
.mermaid-wrapper.maximized {
	position: fixed;
	top: 0;
	left: 0;
	width: 100vw;
	height: 100vh;
	background: var(--background-color, #fff);
	z-index: 9999;
	display: flex;
	align-items: center;
	justify-content: center;
	overflow: auto;
	padding: 40px;
}
.mermaid-wrapper.maximized .mermaid {
	width: 100%;
	height: 100%;
	display: flex;
	align-items: center;
	justify-content: center;
}
.mermaid-wrapper.maximized .mermaid svg {
	max-height: 100% !important;
	max-width: 100% !important;
}
</style>
`
			mermaidScriptInjected = true
		}
		
		return fmt.Sprintf(`%s<div class="mermaid-wrapper"><button class="mermaid-maximize-btn" onclick="tamarindMaximizeMermaid(this)">Maximize</button><div class="mermaid">%s</div></div>`, jsScript, content)
	})

	// Interactive Tabs: {{ tabs }} ... {{ /tabs }}
	reTabs := regexp.MustCompile(`(?s){{\s*tabs\s*}}(.*?){{\s*/tabs\s*}}`)
	tabsIndex := 0
	markdown = reTabs.ReplaceAllStringFunc(markdown, func(match string) string {
		tabsIndex++
		submatch := reTabs.FindStringSubmatch(match)
		content := submatch[1]

		reTabItem := regexp.MustCompile(`(?s){{\s*tab\s+title="([^"]+)"\s*}}(.*?){{\s*/tab\s*}}`)
		itemMatches := reTabItem.FindAllStringSubmatch(content, -1)
		if len(itemMatches) == 0 {
			return ""
		}

		buttonsHtml := ""
		panesHtml := ""

		for i, itemSubmatch := range itemMatches {
			title := itemSubmatch[1]
			desc := itemSubmatch[2]

			var buf bytes.Buffer
			var htmlDesc string
			if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
				htmlDesc = buf.String()
			} else {
				htmlDesc = strings.TrimSpace(desc)
			}

			tabId := fmt.Sprintf("tab-%d-%d", tabsIndex, i)
			activeBtnClass := ""
			activePaneClass := ""
			if i == 0 {
				activeBtnClass = " active"
				activePaneClass = " active"
			}

			buttonsHtml += fmt.Sprintf(`<button class="tamarind-tab-btn%s" onclick="tamarindSwitchTab(event, '%s')">%s</button>`, activeBtnClass, tabId, title)
			panesHtml += fmt.Sprintf(`<div id="%s" class="tamarind-tab-pane%s">%s</div>`, tabId, activePaneClass, htmlDesc)
		}

		jsScript := ""
		if !tabsScriptInjected {
			jsScript = `
<script>
if (typeof tamarindSwitchTab !== 'function') {
window.tamarindSwitchTab = function(evt, tabId) {
var i, tabpanes, tablinks;
var container = evt.currentTarget.closest('.tamarind-tabs');
tabpanes = container.getElementsByClassName("tamarind-tab-pane");
for (i = 0; i < tabpanes.length; i++) {
tabpanes[i].className = "tamarind-tab-pane";
}
tablinks = container.getElementsByClassName("tamarind-tab-btn");
for (i = 0; i < tablinks.length; i++) {
tablinks[i].className = tablinks[i].className.replace(" active", "");
}
document.getElementById(tabId).className = "tamarind-tab-pane active";
evt.currentTarget.className += " active";
};
}
</script>
`
			tabsScriptInjected = true
		}

		return fmt.Sprintf(`%s<div class="tamarind-tabs"><div class="tamarind-tabs-bar">%s</div><div class="tamarind-tabs-content">%s</div></div>`, jsScript, buttonsHtml, panesHtml)
	})

	// 3. Terminal (Block): {{ terminal }}...{{ /terminal }}
	reTerm := regexp.MustCompile(`(?s){{\s*terminal\s*}}(.*?){{\s*/terminal\s*}}`)
	terminalIndex := 0
	copyBtnHtml := `<button class="terminal-copy-btn" onclick="tamarindCopyTerminal(event)" aria-label="Copy" title="Copy"><svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button>`
	markdown = reTerm.ReplaceAllStringFunc(markdown, func(match string) string {
		terminalIndex++
		sub := reTerm.FindStringSubmatch(match)
		content := strings.TrimSpace(sub[1])

		jsScript := ""
		if !terminalScriptInjected {
			jsScript = `
<script>
if (typeof tamarindSwitchTerminalTab !== 'function') {
window.tamarindSwitchTerminalTab = function(evt, tabId) {
var i, tabpanes, tablinks;
var container = evt.currentTarget.closest('.terminal');
tabpanes = container.getElementsByClassName("terminal-tab-pane");
for (i = 0; i < tabpanes.length; i++) {
tabpanes[i].className = "terminal-tab-pane";
}
tablinks = container.getElementsByClassName("terminal-tab-btn");
for (i = 0; i < tablinks.length; i++) {
tablinks[i].className = tablinks[i].className.replace(" active", "");
}
document.getElementById(tabId).className = "terminal-tab-pane active";
evt.currentTarget.className += " active";
};

window.tamarindCopyTerminal = function(evt) {
var container = evt.currentTarget.closest('.terminal');
var activePane = container.querySelector('.terminal-tab-pane.active code');
var content = "";
if (activePane) {
content = activePane.innerText;
} else {
var codeBlock = container.querySelector('.terminal-content code');
if (codeBlock) content = codeBlock.innerText;
}

// Strip terminal prompts like '$', 'PS>', or '>' from the beginning of lines
content = content.replace(/^(?:\$|PS>|>)\s*/gm, '').trim();

navigator.clipboard.writeText(content).then(function() {
var btn = evt.currentTarget;
var originalHtml = btn.innerHTML;
btn.innerHTML = '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#10b981" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>';
setTimeout(function() { btn.innerHTML = originalHtml; }, 2000);
});
};
}
</script>
`
			terminalScriptInjected = true
		}

		reTab := regexp.MustCompile(`(?s){{\s*tab\s+title="([^"]+)"\s*}}(.*?){{\s*/tab\s*}}`)
		tabMatches := reTab.FindAllStringSubmatch(content, -1)

		if len(tabMatches) > 0 {
			buttonsHtml := ""
			panesHtml := ""
			for i, tabMatch := range tabMatches {
				title := tabMatch[1]
				tabContent := strings.TrimSpace(tabMatch[2])

				tabId := fmt.Sprintf("term-tab-%d-%d", terminalIndex, i)
				activeBtnClass := ""
				activePaneClass := ""
				if i == 0 {
					activeBtnClass = " active"
					activePaneClass = " active"
				}

				buttonsHtml += fmt.Sprintf(`<button class="terminal-tab-btn%s" onclick="tamarindSwitchTerminalTab(event, '%s')">%s</button>`, activeBtnClass, tabId, title)
				panesHtml += fmt.Sprintf(`<div id="%s" class="terminal-tab-pane%s"><pre class="terminal-content"><code>%s</code></pre></div>`, tabId, activePaneClass, tabContent)
			}

			return fmt.Sprintf(`%s<div class="terminal terminal-has-tabs">%s<div class="terminal-header"><div class="terminal-dots"><span class="dot red"></span><span class="dot yellow"></span><span class="dot green"></span></div><div class="terminal-tabs-bar">%s</div></div>%s</div>`, jsScript, copyBtnHtml, buttonsHtml, panesHtml)
		} else {
			return fmt.Sprintf(`%s<div class="terminal">%s<div class="terminal-header"><span class="dot red"></span><span class="dot yellow"></span><span class="dot green"></span></div><pre class="terminal-content"><code>%s</code></pre></div>`, jsScript, copyBtnHtml, content)
		}
	})

	// Metrics Grid: {{ metrics }} ... {{ /metrics }}
	reMetrics := regexp.MustCompile(`(?s){{\s*metrics\s*}}(.*?){{\s*/metrics\s*}}`)
	markdown = reMetrics.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reMetrics.FindStringSubmatch(match)
		content := submatch[1]

		reMetric := regexp.MustCompile(`{{\s*metric\s+value="([^"]+)"\s+label="([^"]+)"\s*}}`)
		itemsHtml := ""
		
		metricMatches := reMetric.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range metricMatches {
			val := itemSubmatch[1]
			lbl := itemSubmatch[2]
			itemsHtml += fmt.Sprintf(`<div class="metric-card"><div class="metric-value">%s</div><div class="metric-label">%s</div></div>`, val, lbl)
		}

		return fmt.Sprintf(`<div class="metrics-grid">%s</div>`, itemsHtml)
	})

	// Social Ribbon: {{ social_ribbon }} ... {{ /social_ribbon }}
	reSocialRibbon := regexp.MustCompile(`(?s){{\s*social_ribbon\s*}}(.*?){{\s*/social_ribbon\s*}}`)
	markdown = reSocialRibbon.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reSocialRibbon.FindStringSubmatch(match)
		content := submatch[1]

		reTestimonial := regexp.MustCompile(`(?s){{\s*testimonial\s+(.*?)\s*}}(.*?){{\s*/testimonial\s*}}`)
		
		// Parse testimonial cards
		type testimonialData struct {
			stars  string
			avatar string
			author string
			handle string
			quote  string
		}
		
		var testimonials []testimonialData
		reAttr := regexp.MustCompile(`(\w+)="([^"]*)"`)
		
		testimonialMatches := reTestimonial.FindAllStringSubmatch(content, -1)
		for _, tMatch := range testimonialMatches {
			attrStr := tMatch[1]
			quote := strings.TrimSpace(tMatch[2])
			
			attrs := make(map[string]string)
			for _, attrMatch := range reAttr.FindAllStringSubmatch(attrStr, -1) {
				attrs[attrMatch[1]] = attrMatch[2]
			}
			
			stars := attrs["stars"]
			avatar := attrs["avatar"]
			author := attrs["author"]
			handle := attrs["handle"]
			
			testimonials = append(testimonials, testimonialData{
				stars:  stars,
				avatar: avatar,
				author: author,
				handle: handle,
				quote:  quote,
			})
		}
		
		// Render function for cards
		renderCards := func(items []testimonialData) string {
			var sb strings.Builder
			for _, t := range items {
				// Convert stars number (e.g. "5") or raw stars (e.g. "★★★★★") to stars string
				starsDisplay := ""
				if numStars, err := strconv.Atoi(t.stars); err == nil {
					for i := 0; i < numStars; i++ {
						starsDisplay += "★"
					}
				} else {
					starsDisplay = t.stars // if it's already "★★★★★"
				}
				if starsDisplay == "" {
					starsDisplay = "★★★★★" // fallback default
				}
				
				avatarHtml := ""
				if t.avatar != "" {
					avatarHtml = fmt.Sprintf(`<img class="avatar" src="%s" alt="%s">`, t.avatar, t.author)
				}
				
				authorHtml := ""
				if t.author != "" {
					authorHtml = fmt.Sprintf(`<span class="author">%s</span>`, t.author)
				}
				
				handleHtml := ""
				if t.handle != "" {
					handleHtml = fmt.Sprintf(`<span class="handle">%s</span>`, t.handle)
				}
				
				starsHtml := ""
				if starsDisplay != "" {
					starsHtml = fmt.Sprintf(`<div class="stars">%s</div>`, starsDisplay)
				}
				
				profileInfoHtml := ""
				if authorHtml != "" || handleHtml != "" || starsHtml != "" {
					profileInfoHtml = fmt.Sprintf(`<div class="profile-info">%s%s%s</div>`, authorHtml, handleHtml, starsHtml)
				}
				
				profileHtml := ""
				if avatarHtml != "" || profileInfoHtml != "" {
					profileHtml = fmt.Sprintf(`<div class="profile">%s%s</div>`, avatarHtml, profileInfoHtml)
				}
				
				sb.WriteString(fmt.Sprintf(`
<div class="tamarind-social-ribbon-card">
  <div class="quote">“%s”</div>
  %s
</div>`, t.quote, profileHtml))
			}
			return sb.String()
		}
		
		// Render original list
		originalCardsHtml := renderCards(testimonials)
		
		// Duplicate cards inside track for seamless marquee loop
		return fmt.Sprintf(`<div class="tamarind-social-ribbon-container"><div class="tamarind-social-ribbon-track">%s%s</div></div>`, originalCardsHtml, originalCardsHtml)
	})

	// Features Grid: {{ features }} ... {{ /features }}
	reFeatures := regexp.MustCompile(`(?s){{\s*features\s*}}(.*?){{\s*/features\s*}}`)
	markdown = reFeatures.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reFeatures.FindStringSubmatch(match)
		content := submatch[1]

		reFeature := regexp.MustCompile(`(?s){{\s*feature\s+title="([^"]+)"\s+gradient="([^"]+)"\s+icon="([^"]+)"\s*}}(.*?){{\s*/feature\s*}}`)
		itemsHtml := ""
		
		featureMatches := reFeature.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range featureMatches {
			title := itemSubmatch[1]
			grad := itemSubmatch[2]
			iconName := itemSubmatch[3]
			desc := itemSubmatch[4]
			
			svgIcon := getBuiltInIconSvg(iconName)

			itemsHtml += fmt.Sprintf(`<div class="feature-card"><div class="feature-icon-box gradient-%s">%s</div><h3 class="feature-title">%s</h3><p class="feature-desc">%s</p></div>`, grad, svgIcon, title, strings.TrimSpace(desc))
		}

		return fmt.Sprintf(`<div class="features-grid">%s</div>`, itemsHtml)
	})

	// Pricing Grid: {{ pricing monthly_label="..." annual_label="..." discount="..." }} ... {{ /pricing }}
	rePricing := regexp.MustCompile(`(?s){{\s*pricing\s+(.*?)\s*}}(.*?){{\s*/pricing\s*}}`)
	markdown = rePricing.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := rePricing.FindStringSubmatch(match)
		attrString := submatch[1]
		content := submatch[2]

		// Parse grid attributes
		reAttr := regexp.MustCompile(`(\w+)="([^"]*)"`)
		gridAttrs := make(map[string]string)
		for _, attrMatch := range reAttr.FindAllStringSubmatch(attrString, -1) {
			gridAttrs[attrMatch[1]] = attrMatch[2]
		}

		monthlyLabel := gridAttrs["monthly_label"]
		if monthlyLabel == "" {
			monthlyLabel = "Monthly"
		}
		annualLabel := gridAttrs["annual_label"]
		if annualLabel == "" {
			annualLabel = "Annual"
		}
		discount := gridAttrs["discount"]

		pricingGridCount++
		gridID := fmt.Sprintf("pricing-grid-%d", pricingGridCount)

		// Structure to hold plan configuration
		type planData struct {
			title         string
			priceMonthly  string
			priceAnnual   string
			periodMonthly string
			periodAnnual  string
			featuredClass string
			badgeHtml     string
			buttonText    string
			urlMonthly    string
			urlAnnual     string
			bulletsHtml   string
		}

		var plans []planData

		// Parse nested plans
		rePlan := regexp.MustCompile(`(?s){{\s*plan\s+(.*?)\s*}}(.*?){{\s*/plan\s*}}`)
		planMatches := rePlan.FindAllStringSubmatch(content, -1)

		for _, planMatch := range planMatches {
			planAttrStr := planMatch[1]
			planContent := planMatch[2]

			planAttrs := make(map[string]string)
			for _, attrMatch := range reAttr.FindAllStringSubmatch(planAttrStr, -1) {
				planAttrs[attrMatch[1]] = attrMatch[2]
			}

			title := planAttrs["title"]
			price := planAttrs["price"]
			priceMonthly := planAttrs["price_monthly"]
			priceAnnual := planAttrs["price_annual"]
			if priceMonthly == "" {
				priceMonthly = price
			}
			if priceAnnual == "" {
				priceAnnual = price
			}

			period := planAttrs["period"]
			periodMonthly := planAttrs["period_monthly"]
			periodAnnual := planAttrs["period_annual"]
			if periodMonthly == "" {
				periodMonthly = period
			}
			if periodAnnual == "" {
				periodAnnual = period
			}

			url := planAttrs["url"]
			urlMonthly := planAttrs["url_monthly"]
			urlAnnual := planAttrs["url_annual"]
			if urlMonthly == "" {
				urlMonthly = url
			}
			if urlAnnual == "" {
				urlAnnual = url
			}

			featured := planAttrs["featured"]
			badge := planAttrs["badge"]
			button := planAttrs["button"]
			if button == "" {
				button = "Get Started"
			}

			featuredClass := ""
			if featured == "true" {
				featuredClass = " featured"
			}

			badgeHtml := ""
			if badge != "" {
				badgeHtml = fmt.Sprintf(`<div class="card-badge-poc">%s</div>`, badge)
			}

			// Parse plan bullets (markdown list)
			bulletsHtml := ""
			reListItem := regexp.MustCompile(`(?m)^\s*[-\*]\s*(.*?)\s*$`)
			listMatches := reListItem.FindAllStringSubmatch(planContent, -1)
			for _, listMatch := range listMatches {
				bulletsHtml += fmt.Sprintf(`<li>%s</li>`, strings.TrimSpace(listMatch[1]))
			}

			plans = append(plans, planData{
				title:         title,
				priceMonthly:  priceMonthly,
				priceAnnual:   priceAnnual,
				periodMonthly: periodMonthly,
				periodAnnual:  periodAnnual,
				featuredClass: featuredClass,
				badgeHtml:     badgeHtml,
				buttonText:    button,
				urlMonthly:    urlMonthly,
				urlAnnual:     urlAnnual,
				bulletsHtml:   bulletsHtml,
			})
		}

		// Auto-detect if dynamic toggle features are needed
		hasToggle := false
		for _, p := range plans {
			if (p.priceMonthly != "" && p.priceAnnual != "" && p.priceMonthly != p.priceAnnual) ||
				(p.periodMonthly != "" && p.periodAnnual != "" && p.periodMonthly != p.periodAnnual) ||
				(p.urlMonthly != "" && p.urlAnnual != "" && p.urlMonthly != p.urlAnnual) {
				hasToggle = true
				break
			}
		}

		formatPrice := func(p string) string {
			if p == "" {
				return ""
			}
			if strings.HasPrefix(p, "$") {
				return p
			}
			return "$" + p
		}

		// Generate plan cards HTML
		plansHtml := ""
		for _, p := range plans {
			var priceValHtml string
			var periodHtml string
			var buttonHtml string

			if hasToggle {
				priceValHtml = fmt.Sprintf(`<div class="price-val" data-monthly="%s" data-annual="%s">%s</div>`,
					formatPrice(p.priceMonthly), formatPrice(p.priceAnnual), formatPrice(p.priceMonthly))
				periodHtml = fmt.Sprintf(`<div class="price-period" data-monthly="%s" data-annual="%s">%s</div>`,
					p.periodMonthly, p.periodAnnual, p.periodMonthly)

				if p.urlMonthly != "" || p.urlAnnual != "" {
					buttonHtml = fmt.Sprintf(`<a href="%s" class="pricing-btn" data-monthly-url="%s" data-annual-url="%s">%s</a>`,
						p.urlMonthly, p.urlMonthly, p.urlAnnual, p.buttonText)
				} else {
					buttonHtml = fmt.Sprintf(`<button class="pricing-btn">%s</button>`, p.buttonText)
				}
			} else {
				priceValHtml = fmt.Sprintf(`<div class="price-val">%s</div>`, formatPrice(p.priceMonthly))
				periodHtml = fmt.Sprintf(`<div class="price-period">%s</div>`, p.periodMonthly)

				if p.urlMonthly != "" {
					buttonHtml = fmt.Sprintf(`<a href="%s" class="pricing-btn">%s</a>`, p.urlMonthly, p.buttonText)
				} else {
					buttonHtml = fmt.Sprintf(`<button class="pricing-btn">%s</button>`, p.buttonText)
				}
			}

			plansHtml += fmt.Sprintf(`
<div class="price-card%s">
  %s
  <div class="price-header">
    <h4>%s</h4>
    %s
    %s
  </div>
  <ul>
    %s
  </ul>
  %s
</div>`, p.featuredClass, p.badgeHtml, p.title, priceValHtml, periodHtml, p.bulletsHtml, buttonHtml)
		}

		if !hasToggle {
			return fmt.Sprintf(`
<div class="pricing-wrapper" id="%s">
  <div class="pricing-grid-poc">
    %s
  </div>
</div>`, gridID, plansHtml)
		}

		discountBadgeHtml := ""
		if discount != "" {
			discountBadgeHtml = fmt.Sprintf(` <span style="color:#10b981; font-size:0.8rem;">(%s)</span>`, discount)
		}

		return fmt.Sprintf(`
<div class="pricing-wrapper" id="%s">
  <div class="billing-toggle">
    <span>%s</span>
    <label class="switch-poc">
      <input type="checkbox" onchange="togglePricingGrid(this, '%s')">
      <span class="slider-poc"></span>
    </label>
    <span>%s%s</span>
  </div>
  <div class="pricing-grid-poc">
    %s
  </div>
</div>
<script>
if (typeof togglePricingGrid !== 'function') {
  window.togglePricingGrid = function(checkbox, gridId) {
    var container = document.getElementById(gridId);
    if (!container) return;
    var priceElements = container.querySelectorAll('.price-val');
    priceElements.forEach(function(el) {
      var monthly = el.getAttribute('data-monthly');
      var annual = el.getAttribute('data-annual');
      el.innerText = checkbox.checked ? annual : monthly;
    });
    var periodElements = container.querySelectorAll('.price-period');
    periodElements.forEach(function(el) {
      var monthly = el.getAttribute('data-monthly');
      var annual = el.getAttribute('data-annual');
      el.innerText = checkbox.checked ? annual : monthly;
    });
    var buttonElements = container.querySelectorAll('.pricing-btn');
    buttonElements.forEach(function(el) {
      var monthlyUrl = el.getAttribute('data-monthly-url');
      var annualUrl = el.getAttribute('data-annual-url');
      if (monthlyUrl && annualUrl) {
        el.setAttribute('href', checkbox.checked ? annualUrl : monthlyUrl);
      }
    });
  }
}
</script>`, gridID, monthlyLabel, gridID, annualLabel, discountBadgeHtml, plansHtml)
	})

	// Capabilities Grid: {{ capabilities-grid }} ... {{ /capabilities-grid }}
	reCapsGrid := regexp.MustCompile(`(?s){{\s*capabilities-grid\s*}}(.*?){{\s*/capabilities-grid\s*}}`)
	markdown = reCapsGrid.ReplaceAllString(markdown, `<div class="capabilities-grid">$1</div>`)

	// Capabilities Checklist Card: {{ capabilities }} ... {{ /capabilities }}
	reCapabilities := regexp.MustCompile(`(?s){{\s*capabilities(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/capabilities\s*}}`)
	markdown = reCapabilities.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reCapabilities.FindStringSubmatch(match)
		title := submatch[1]
		content := submatch[2]

		// Support status="..." statusLabel="..." OR check="..." style attributes!
		reCapabilityStatus := regexp.MustCompile(`{{\s*capability\s+name="([^"]+)"\s+desc="([^"]+)"\s+status="([^"]+)"(?:\s+statusLabel="([^"]*)")?\s*}}`)
		reCapabilityCheck := regexp.MustCompile(`{{\s*capability\s+name="([^"]+)"\s+desc="([^"]+)"\s+check="([^"]+)"\s*}}`)
		rowsHtml := ""
		
		statusMatches := reCapabilityStatus.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range statusMatches {
			name := itemSubmatch[1]
			desc := itemSubmatch[2]
			status := itemSubmatch[3]
			statusLabel := status
			if len(itemSubmatch) > 4 && itemSubmatch[4] != "" {
				statusLabel = itemSubmatch[4]
			}

			rowsHtml += fmt.Sprintf(`<div class="capability-row"><div class="capability-info"><span class="capability-name">%s</span><span class="capability-desc">%s</span></div><span class="capability-status status-%s">%s</span></div>`, name, desc, status, statusLabel)
		}

		// Fallback to check="..." attributes
		if len(statusMatches) == 0 {
			checkMatches := reCapabilityCheck.FindAllStringSubmatch(content, -1)
			for _, itemSubmatch := range checkMatches {
				name := itemSubmatch[1]
				desc := itemSubmatch[2]
				checkVal := itemSubmatch[3]

				status := "pending"
				statusLabel := "Pending"
				if checkVal == "true" {
					status = "success"
					statusLabel = "Yes"
				} else if checkVal == "warn" {
					status = "warning"
					statusLabel = "Partial"
				} else if checkVal == "false" {
					status = "error"
					statusLabel = "No"
				}

				rowsHtml += fmt.Sprintf(`<div class="capability-row"><div class="capability-info"><span class="capability-name">%s</span><span class="capability-desc">%s</span></div><span class="capability-status status-%s">%s</span></div>`, name, desc, status, statusLabel)
			}
		}

		headerHtml := ""
		if title != "" {
			headerHtml = fmt.Sprintf(`<div class="capability-header"><div class="capability-card-title">%s</div></div>`, title)
		}

		return fmt.Sprintf(`<div class="capability-card">%s%s</div>`, headerHtml, rowsHtml)
	})

	// Vertical Timeline: {{ timeline }} ... {{ /timeline }}
	reTimeline := regexp.MustCompile(`(?s){{\s*timeline\s*}}(.*?){{\s*/timeline\s*}}`)
	markdown = reTimeline.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reTimeline.FindStringSubmatch(match)
		content := submatch[1]

		// Support {{ item title="..." number="..." }} OR {{ timeline-item step="..." title="..." }} syntaxes!
		reItem1 := regexp.MustCompile(`(?s){{\s*item\s+title="([^"]+)"(?:\s+number="([^"]*)")?\s*}}(.*?){{\s*/item\s*}}`)
		reItem2 := regexp.MustCompile(`(?s){{\s*timeline-item\s+step="([^"]+)"\s+title="([^"]+)"\s*}}(.*?){{\s*/timeline-item\s*}}`)
		itemsHtml := ""
		
		item1Matches := reItem1.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range item1Matches {
			title := itemSubmatch[1]
			num := itemSubmatch[2]
			desc := itemSubmatch[3]

			var buf bytes.Buffer
			var htmlDesc string
			if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
				htmlDesc = buf.String()
			} else {
				htmlDesc = strings.TrimSpace(desc)
			}

			badgeHtml := ""
			if num != "" {
				badgeHtml = fmt.Sprintf(`<div class="timeline-badge"><span class="timeline-badge-number">%s</span></div>`, num)
			} else {
				badgeHtml = `<div class="timeline-badge"></div>`
			}

			itemsHtml += fmt.Sprintf(`<div class="timeline-item">%s<div class="timeline-content"><h3 class="timeline-title">%s</h3><div class="timeline-desc">%s</div></div></div>`, badgeHtml, title, htmlDesc)
		}

		// Fallback to timeline-item step="..." title="..."
		if len(item1Matches) == 0 {
			item2Matches := reItem2.FindAllStringSubmatch(content, -1)
			for _, itemSubmatch := range item2Matches {
				num := itemSubmatch[1]
				title := itemSubmatch[2]
				desc := itemSubmatch[3]

				var buf bytes.Buffer
				var htmlDesc string
				if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
					htmlDesc = buf.String()
				} else {
					htmlDesc = strings.TrimSpace(desc)
				}

				badgeHtml := ""
				if num != "" {
					badgeHtml = fmt.Sprintf(`<div class="timeline-badge"><span class="timeline-badge-number">%s</span></div>`, num)
				} else {
					badgeHtml = `<div class="timeline-badge"></div>`
				}

				itemsHtml += fmt.Sprintf(`<div class="timeline-item">%s<div class="timeline-content"><h3 class="timeline-title">%s</h3><div class="timeline-desc">%s</div></div></div>`, badgeHtml, title, htmlDesc)
			}
		}

		return fmt.Sprintf(`<div class="timeline-container">%s</div>`, itemsHtml)
	})

	// Dropdown Selection: {{ dropdown id="X" label="Y" }} ... {{ /dropdown }}
	reDropdown := regexp.MustCompile(`(?s){{\s*dropdown\s+(.*?)\s*}}(.*?){{\s*/dropdown\s*}}`)
	markdown = reDropdown.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reDropdown.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reId := regexp.MustCompile(`id="([^"]+)"`)
		reLabel := regexp.MustCompile(`label="([^"]+)"`)

		id := ""
		if m := reId.FindStringSubmatch(attrs); len(m) > 1 {
			id = m[1]
		}
		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}

		reOption := regexp.MustCompile(`(?s){{\s*option\s+value="([^"]+)"(?:\s+selected="([^"]*)")?\s*}}(.*?){{\s*/option\s*}}`)
		optionsHtml := ""
		
		optMatches := reOption.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range optMatches {
			val := itemSubmatch[1]
			sel := itemSubmatch[2]
			text := itemSubmatch[3]

			selectedAttr := ""
			if sel == "true" {
				selectedAttr = " selected"
			}

			optionsHtml += fmt.Sprintf(`<option value="%s"%s>%s</option>`, val, selectedAttr, strings.TrimSpace(text))
		}

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="tamarind-select-label">%s</label>`, label)
		}

		idAttr := ""
		if id != "" {
			idAttr = fmt.Sprintf(` id="%s"`, id)
		}

		chevronSvg := getBuiltInIconSvg("chevron")

		return fmt.Sprintf(`<div class="tamarind-select-wrapper">%s<div class="tamarind-select-control"><select class="tamarind-select"%s>%s</select><div class="tamarind-select-chevron">%s</div></div></div>`, labelHtml, idAttr, optionsHtml, chevronSvg)
	})

	// Accordion: {{ accordion }} ... {{ /accordion }}
	reAccordion := regexp.MustCompile(`(?s){{\s*accordion\s*}}(.*?){{\s*/accordion\s*}}`)
	markdown = reAccordion.ReplaceAllStringFunc(markdown, func(match string) string {
		submatch := reAccordion.FindStringSubmatch(match)
		content := submatch[1]

		reAccordionItem := regexp.MustCompile(`(?s){{\s*accordion-item\s+title="([^"]+)"\s*}}(.*?){{\s*/accordion-item\s*}}`)
		itemsHtml := ""

		itemMatches := reAccordionItem.FindAllStringSubmatch(content, -1)
		for _, itemSubmatch := range itemMatches {
			title := itemSubmatch[1]
			desc := itemSubmatch[2]

			var buf bytes.Buffer
			var htmlDesc string
			if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf); err == nil {
				htmlDesc = buf.String()
			} else {
				htmlDesc = strings.TrimSpace(desc)
			}

			itemsHtml += fmt.Sprintf(`<details class="tamarind-accordion"><summary class="tamarind-accordion-summary">%s</summary><div class="tamarind-accordion-content">%s</div></details>`, title, htmlDesc)
		}

		return fmt.Sprintf(`<div class="tamarind-accordion-container">%s</div>`, itemsHtml)
	})



	// 4. Code Include: {{ include src="file.go" lines="1-10" lang="go" }}
	reInclude := regexp.MustCompile(`{{\s*include\s+src="([^"]+)"(?:\s+lines="([0-9]+-[0-9]+)")?(?:\s+lang="([^"]+)")?\s*}}`)
	markdown = reInclude.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reInclude.FindStringSubmatch(match)
		src := parts[1]
		linesRange := parts[2]
		lang := parts[3]
		if lang == "" {
			lang = "text"
		}

		var content []byte
		var err error

		if strings.HasPrefix(src, "http") {
			resp, err := http.Get(src)
			if err != nil {
				return fmt.Sprintf("> **Error fetching %s**: %v", src, err)
			}
			defer resp.Body.Close()
			content, err = io.ReadAll(resp.Body)
		} else {
			path := filepath.Join(sourceDir, src)
			content, err = os.ReadFile(path)
		}

		if err != nil {
			return fmt.Sprintf("> **Error including %s**: %v", src, err)
		}

		finalContent := string(content)
		if linesRange != "" {
			lParts := strings.Split(linesRange, "-")
			if len(lParts) == 2 {
				start, _ := strconv.Atoi(lParts[0])
				end, _ := strconv.Atoi(lParts[1])

				lines := strings.Split(finalContent, "\n")
				// Validate bounds (1-based -> 0-based)
				if start < 1 {
					start = 1
				}
				if end > len(lines) {
					end = len(lines)
				}
				if start <= end {
					finalContent = strings.Join(lines[start-1:end], "\n")
				}
			}
		}

		return fmt.Sprintf("```%s\n%s\n```", lang, finalContent)
	})

	// 5. GitHub Gist: {{ gist id="123" }}
	reGist := regexp.MustCompile(`{{\s*gist\s+id="([^"]+)"\s*}}`)
	markdown = reGist.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reGist.FindStringSubmatch(match)
		gistID := parts[1]

		resp, err := http.Get(fmt.Sprintf("https://gist.github.com/%s.json", gistID))
		if err != nil {
			return fmt.Sprintf("> **Error loading gist %s**: %v", gistID, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Sprintf("> **Error loading gist %s**: Status %d", gistID, resp.StatusCode)
		}

		var data struct {
			Div        string `json:"div"`
			Stylesheet string `json:"stylesheet"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return fmt.Sprintf("> **Error decoding gist %s**: %v", gistID, err)
		}

		// Strip newlines to prevent Markdown from interpreting indented HTML as code blocks
		data.Div = strings.ReplaceAll(data.Div, "\n", "")
		data.Div = strings.ReplaceAll(data.Div, "\r", "")

		return fmt.Sprintf(`<link rel="stylesheet" href="%s">%s`, data.Stylesheet, data.Div)
	})

	// 6. Math: {{ math }}...{{ /math }}
	reMath := regexp.MustCompile(`(?s){{\s*math\s*}}(.*?){{\s*/math\s*}}`)
	markdown = reMath.ReplaceAllString(markdown, `<div class="math-block">$$$1$$</div>`)

	// YouTube: {{ youtube ID }}
	reYT := regexp.MustCompile(`{{\s*youtube\s+([a-zA-Z0-9_-]+)\s*}}`)
	markdown = reYT.ReplaceAllString(markdown, `<div class="video-container"><iframe src="https://www.youtube.com/embed/$1" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></div>`)

	// Figure: {{ figure src="url" caption="text" width="50%" }}
	reFig := regexp.MustCompile(`{{\s*figure\s+(.*?)\s*}}`)
	markdown = reFig.ReplaceAllStringFunc(markdown, func(match string) string {
		// Parse attributes manually for flexibility
		content := reFig.FindStringSubmatch(match)[1]

		reSrc := regexp.MustCompile(`src="([^"]+)"`)
		reCap := regexp.MustCompile(`caption="([^"]+)"`)
		reWidth := regexp.MustCompile(`width="([^"]+)"`)

		srcMatch := reSrc.FindStringSubmatch(content)
		if srcMatch == nil {
			return match // Invalid, no src
		}
		src := srcMatch[1]

		caption := ""
		capMatch := reCap.FindStringSubmatch(content)
		if capMatch != nil {
			caption = capMatch[1]
		}

		width := ""
		widthMatch := reWidth.FindStringSubmatch(content)
		if widthMatch != nil {
			width = widthMatch[1] // e.g. "500px" or "50%"
		}

		var figcaptionHTML string
		if caption != "" {
			figcaptionHTML = fmt.Sprintf("<figcaption>%s</figcaption>", caption)
		}

		// Style for width
		styleAttr := ""
		if width != "" {
			styleAttr = fmt.Sprintf(` style="width: %s; margin: 0 auto; display: block;"`, width)
		}

		// Responsive Logic
		// We assume src matches the generated optimized files: name-width.ext
		// AND that it is a local resource (not http/s)
		if !strings.HasPrefix(src, "http") {
			// Resolve path to check dimensions
			localPath := filepath.Join(sourceDir, src)

			// Check if file exists and get dimensions
			f, err := os.Open(localPath)
			if err == nil {
				defer f.Close()
				cfg, _, err := image.DecodeConfig(f)
				if err == nil {
					imgWidth := cfg.Width
					ext := filepath.Ext(src)
					base := strings.TrimSuffix(src, ext)

					lowerExt := strings.ToLower(ext)
					if lowerExt == ".jpg" || lowerExt == ".jpeg" || lowerExt == ".png" {
						var sources []string
						breakpoints := []int{480, 800, 1200}

						for _, bp := range breakpoints {
							if imgWidth >= bp {
								sources = append(sources, fmt.Sprintf("%s-%dw%s %dw", base, bp, ext, bp))
							}
						}

						if len(sources) > 0 {
							srcset := strings.Join(sources, ", ")
							sizes := "(max-width: 480px) 100vw, (max-width: 800px) 100vw, 100vw"
							return fmt.Sprintf(`<figure><img src="%s" srcset="%s" sizes="%s" alt="%s"%s>%s</figure>`,
								src, srcset, sizes, caption, styleAttr, figcaptionHTML)
						}
					}
				}
			}
		}

		// Fallback
		return fmt.Sprintf(`<figure><img src="%s" alt="%s"%s>%s</figure>`, src, caption, styleAttr, figcaptionHTML)
	})

	// 7. Button (Block): {{ button href="url" ... }}Label{{ /button }}
	reButton := regexp.MustCompile(`(?s){{<?\s*button\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*button\s*>??}}`)
	markdown = reButton.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reButton.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reHref := regexp.MustCompile(`href="([^"]+)"`)
		reType := regexp.MustCompile(`type="([^"]+)"`)
		reSize := regexp.MustCompile(`size="([^"]+)"`)
		reTarget := regexp.MustCompile(`target="([^"]+)"`)

		href := ""
		if m := reHref.FindStringSubmatch(attrs); len(m) > 1 {
			href = m[1]
		}
		typ := "primary"
		if m := reType.FindStringSubmatch(attrs); len(m) > 1 {
			typ = m[1]
		}
		size := ""
		if m := reSize.FindStringSubmatch(attrs); len(m) > 1 {
			size = m[1]
		}
		target := ""
		if m := reTarget.FindStringSubmatch(attrs); len(m) > 1 {
			target = m[1]
		}

		classAttr := "btn"
		if typ != "" {
			classAttr += " btn-" + typ
		}
		if size != "" {
			classAttr += " btn-" + size
		}

		targetAttr := ""
		if target != "" {
			targetAttr = fmt.Sprintf(` target="%s"`, target)
		}

		return fmt.Sprintf(`<a href="%s" class="%s"%s>%s</a>`, href, classAttr, targetAttr, strings.TrimSpace(content))
	})

	// 8. Card (Block): {{ card }}...{{ /card }}
	reCard := regexp.MustCompile(`(?s){{<?\s*card(?:\s+padding="([^"]*)")?\s*>??}}(.*?){{<?\s*/?\s*card\s*>??}}`)
	markdown = reCard.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reCard.FindStringSubmatch(match)
		paddingVal := parts[1]
		content := parts[2]

		paddingClass := " card-padding"
		if paddingVal == "false" {
			paddingClass = ""
		}
		return fmt.Sprintf(`<div class="card%s">%s</div>`, paddingClass, strings.TrimSpace(content))
	})

	// 9. Alert (Block): {{ alert type="info" title="Title" }}...{{ /alert }} / {{ callout }}
	reAlert := regexp.MustCompile(`(?s){{<?\s*(?:alert|callout)\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*(?:alert|callout)\s*>??}}`)
	markdown = reAlert.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reAlert.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reType := regexp.MustCompile(`type="([^"]+)"`)
		reTitle := regexp.MustCompile(`title="([^"]+)"`)

		typ := "info"
		if m := reType.FindStringSubmatch(attrs); len(m) > 1 {
			typ = m[1]
		}
		title := ""
		if m := reTitle.FindStringSubmatch(attrs); len(m) > 1 {
			title = m[1]
		}

		iconName := typ
		if typ == "danger" || typ == "error" || typ == "warn" || typ == "warning" {
			iconName = "warning"
		} else if typ == "sparkles" || typ == "magic" || typ == "special" {
			iconName = "sparkles"
		} else if typ == "success" {
			iconName = "success"
		} else {
			iconName = "info"
		}

		svgIcon := getBuiltInIconSvg(iconName)

		titleHtml := ""
		if title != "" {
			titleHtml = fmt.Sprintf(`<h4 class="alert-title">%s</h4>`, title)
		}

		var buf bytes.Buffer
		var htmlContent string
		if err := unsafeGoldmark.Convert([]byte(strings.TrimSpace(content)), &buf); err == nil {
			htmlContent = buf.String()
		} else {
			htmlContent = strings.TrimSpace(content)
		}

		return fmt.Sprintf(`<div class="alert-container alert-%s"><div class="alert-icon-box">%s</div><div class="alert-content">%s<div class="alert-message">%s</div></div></div>`, typ, svgIcon, titleHtml, htmlContent)
	})

	// 10. Badge (Block): {{ badge type="primary" }}...{{ /badge }}
	reBadge := regexp.MustCompile(`(?s){{<?\s*badge(?:\s+type="([^"]*)")?\s*>??}}(.*?){{<?\s*/?\s*badge\s*>??}}`)
	markdown = reBadge.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reBadge.FindStringSubmatch(match)
		typ := parts[1]
		content := parts[2]

		classAttr := "badge"
		if typ != "" {
			classAttr += " badge-" + typ
		}
		return fmt.Sprintf(`<span class="%s">%s</span>`, classAttr, strings.TrimSpace(content))
	})

	// 11. Form (Block): {{ form action="..." method="..." }}...{{ /form }}
	reForm := regexp.MustCompile(`(?s){{<?\s*form\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*form\s*>??}}`)
	markdown = reForm.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reForm.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reAction := regexp.MustCompile(`action="([^"]+)"`)
		reMethod := regexp.MustCompile(`method="([^"]+)"`)

		action := "#"
		if m := reAction.FindStringSubmatch(attrs); len(m) > 1 {
			action = m[1]
		}
		method := "POST"
		if mMethod := reMethod.FindStringSubmatch(attrs); len(mMethod) > 1 {
			method = mMethod[1]
		}

		return fmt.Sprintf(`<form action="%s" method="%s">%s</form>`, action, method, strings.TrimSpace(content))
	})

	// 12. Form Input: {{ form-input label="..." type="..." placeholder="..." }}
	reFormInput := regexp.MustCompile(`{{<?\s*form-input\s+(.*?)\s*/?>??}}`)
	markdown = reFormInput.ReplaceAllStringFunc(markdown, func(match string) string {
		attrs := reFormInput.FindStringSubmatch(match)[1]

		reLabel := regexp.MustCompile(`label="([^"]+)"`)
		reType := regexp.MustCompile(`type="([^"]+)"`)
		rePlaceholder := regexp.MustCompile(`placeholder="([^"]+)"`)

		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}
		typ := "text"
		if m := reType.FindStringSubmatch(attrs); len(m) > 1 {
			typ = m[1]
		}
		placeholder := ""
		if m := rePlaceholder.FindStringSubmatch(attrs); len(m) > 1 {
			placeholder = m[1]
		}

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
		}

		return fmt.Sprintf(`<div class="form-group">%s<input type="%s" class="form-input" placeholder="%s"></div>`, labelHtml, typ, placeholder)
	})

	// 13. Form Textarea: {{ form-textarea label="..." placeholder="..." rows="..." }}
	reFormTextarea := regexp.MustCompile(`{{<?\s*form-textarea\s+(.*?)\s*/?>??}}`)
	markdown = reFormTextarea.ReplaceAllStringFunc(markdown, func(match string) string {
		attrs := reFormTextarea.FindStringSubmatch(match)[1]

		reLabel := regexp.MustCompile(`label="([^"]+)"`)
		rePlaceholder := regexp.MustCompile(`placeholder="([^"]+)"`)
		reRows := regexp.MustCompile(`rows="([^"]+)"`)

		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}
		placeholder := ""
		if m := rePlaceholder.FindStringSubmatch(attrs); len(m) > 1 {
			placeholder = m[1]
		}
		rows := "4"
		if m := reRows.FindStringSubmatch(attrs); len(m) > 1 {
			rows = m[1]
		}

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
		}

		return fmt.Sprintf(`<div class="form-group">%s<textarea class="form-textarea" rows="%s" placeholder="%s"></textarea></div>`, labelHtml, rows, placeholder)
	})

	// 14. Form Select (Block): {{ form-select label="..." }}...{{ /form-select }}
	reFormSelect := regexp.MustCompile(`(?s){{<?\s*form-select\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*form-select\s*>??}}`)
	markdown = reFormSelect.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reFormSelect.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reLabel := regexp.MustCompile(`label="([^"]+)"`)
		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}

		// Parse simple options: {{ option }}Text{{ /option }} or with value/selected
		reOptionSimple := regexp.MustCompile(`(?s){{\s*option\s*}}(.*?){{\s*/option\s*}}`)
		content = reOptionSimple.ReplaceAllString(content, `<option>$1</option>`)

		reOptionAttr := regexp.MustCompile(`(?s){{\s*option\s+value="([^"]+)"(?:\s+selected="([^"]*)")?\s*}}(.*?){{\s*/option\s*}}`)
		content = reOptionAttr.ReplaceAllStringFunc(content, func(optMatch string) string {
			optParts := reOptionAttr.FindStringSubmatch(optMatch)
			val := optParts[1]
			sel := optParts[2]
			text := optParts[3]

			selectedAttr := ""
			if sel == "true" {
				selectedAttr = " selected"
			}
			return fmt.Sprintf(`<option value="%s"%s>%s</option>`, val, selectedAttr, strings.TrimSpace(text))
		})

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
		}

		return fmt.Sprintf(`<div class="form-group">%s<select class="form-select">%s</select></div>`, labelHtml, strings.TrimSpace(content))
	})

	// 15. Form Checkbox: {{ form-checkbox label="..." checked="..." }}
	reFormCheckbox := regexp.MustCompile(`{{<?\s*form-checkbox\s+(.*?)\s*/?>??}}`)
	markdown = reFormCheckbox.ReplaceAllStringFunc(markdown, func(match string) string {
		attrs := reFormCheckbox.FindStringSubmatch(match)[1]

		reLabel := regexp.MustCompile(`label="([^"]+)"`)
		reChecked := regexp.MustCompile(`checked="([^"]+)"`)

		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}
		checked := ""
		if m := reChecked.FindStringSubmatch(attrs); len(m) > 1 && m[1] == "true" {
			checked = " checked"
		}

		return fmt.Sprintf(`<div class="form-group"><label class="form-label"><input type="checkbox" class="form-checkbox"%s> %s</label></div>`, checked, label)
	})

	// 16. Form Radio Group (Block): {{ form-radio-group label="..." }}...{{ /form-radio-group }}
	reFormRadioGroup := regexp.MustCompile(`(?s){{<?\s*form-radio-group\s+(.*?)\s*>??}}(.*?){{<?\s*/?\s*form-radio-group\s*>??}}`)
	markdown = reFormRadioGroup.ReplaceAllStringFunc(markdown, func(match string) string {
		parts := reFormRadioGroup.FindStringSubmatch(match)
		attrs := parts[1]
		content := parts[2]

		reLabel := regexp.MustCompile(`label="([^"]+)"`)
		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}

		// Inside form-radio-group, we parse form-radios: {{ form-radio name="..." label="..." checked="..." }}
		reFormRadio := regexp.MustCompile(`{{<?\s*form-radio\s+(.*?)\s*/?>??}}`)
		parsedContent := reFormRadio.ReplaceAllStringFunc(content, func(radioMatch string) string {
			radioAttrs := reFormRadio.FindStringSubmatch(radioMatch)[1]

			reName := regexp.MustCompile(`name="([^"]+)"`)
			reRadioLabel := regexp.MustCompile(`label="([^"]+)"`)
			reRadioChecked := regexp.MustCompile(`checked="([^"]+)"`)

			name := ""
			if m := reName.FindStringSubmatch(radioAttrs); len(m) > 1 {
				name = m[1]
			}
			radioLabel := ""
			if m := reRadioLabel.FindStringSubmatch(radioAttrs); len(m) > 1 {
				radioLabel = m[1]
			}
			radioChecked := ""
			if m := reRadioChecked.FindStringSubmatch(radioAttrs); len(m) > 1 && m[1] == "true" {
				radioChecked = " checked"
			}

			return fmt.Sprintf(`<label class="form-label"><input type="radio" name="%s" class="form-radio"%s> %s</label>`, name, radioChecked, radioLabel)
		})

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
		}

		return fmt.Sprintf(`<div class="form-group">%s%s</div>`, labelHtml, strings.TrimSpace(parsedContent))
	})

	// 17. Form File: {{ form-file label="..." }}
	reFormFile := regexp.MustCompile(`{{<?\s*form-file\s+(.*?)\s*/?>??}}`)
	markdown = reFormFile.ReplaceAllStringFunc(markdown, func(match string) string {
		attrs := reFormFile.FindStringSubmatch(match)[1]

		reLabel := regexp.MustCompile(`label="([^"]+)"`)
		label := ""
		if m := reLabel.FindStringSubmatch(attrs); len(m) > 1 {
			label = m[1]
		}

		labelHtml := ""
		if label != "" {
			labelHtml = fmt.Sprintf(`<label class="form-label">%s</label>`, label)
		}

		return fmt.Sprintf(`<div class="form-group">%s<input type="file" class="form-file"></div>`, labelHtml)
	})

	// Chart Widgets
	reBarChart := regexp.MustCompile(`(?s){{\s*barchart\s+file="([^"]+)"(?:\s+title="([^"]*)")?(?:\s+colors="([^"]*)")?\s*}}`)
	markdown = reBarChart.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reBarChart.FindStringSubmatch(match)
		file := sub[1]
		title := ""
		colors := ""
		if len(sub) > 2 {
			title = sub[2]
		}
		if len(sub) > 3 {
			colors = sub[3]
		}
		return generateBarChart(sourceDir, file, title, colors)
	})

	reBarChartBlock := regexp.MustCompile(`(?s){{\s*barchart(?:\s+title="([^"]*)")?(?:\s+colors="([^"]*)")?\s*}}(.*?){{\s*/barchart\s*}}`)
	markdown = reBarChartBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reBarChartBlock.FindStringSubmatch(match)
		title := sub[1]
		colors := sub[2]
		jsonData := []byte(sub[3])
		return generateBarChartFromJSON(jsonData, title, colors)
	})

	rePieChart := regexp.MustCompile(`(?s){{\s*piechart\s+file="([^"]+)"(?:\s+title="([^"]*)")?(?:\s+colors="([^"]*)")?\s*}}`)
	markdown = rePieChart.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := rePieChart.FindStringSubmatch(match)
		file := sub[1]
		title := ""
		colors := ""
		if len(sub) > 2 {
			title = sub[2]
		}
		if len(sub) > 3 {
			colors = sub[3]
		}
		return generatePieChart(sourceDir, file, title, colors)
	})

	rePieChartBlock := regexp.MustCompile(`(?s){{\s*piechart(?:\s+title="([^"]*)")?(?:\s+colors="([^"]*)")?\s*}}(.*?){{\s*/piechart\s*}}`)
	markdown = rePieChartBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := rePieChartBlock.FindStringSubmatch(match)
		title := sub[1]
		colors := sub[2]
		jsonData := []byte(sub[3])
		return generatePieChartFromJSON(jsonData, title, colors)
	})

	reLineChartBlock := regexp.MustCompile(`(?s){{\s*linechart(?:\s+title="([^"]*)")?(?:\s+colors="([^"]*)")?\s*}}(.*?){{\s*/linechart\s*}}`)
	markdown = reLineChartBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reLineChartBlock.FindStringSubmatch(match)
		title := sub[1]
		colors := sub[2]
		jsonData := []byte(sub[3])
		return generateLineChartFromJSON(jsonData, title, colors)
	})

	reLineChart := regexp.MustCompile(`(?s){{\s*linechart\s+file="([^"]+)"(?:\s+title="([^"]*)")?(?:\s+colors="([^"]*)")?\s*}}`)
	markdown = reLineChart.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reLineChart.FindStringSubmatch(match)
		filename := sub[1]
		title := ""
		colors := ""
		if len(sub) > 2 {
			title = sub[2]
		}
		if len(sub) > 3 {
			colors = sub[3]
		}
		return generateLineChart(sourceDir, filename, title, colors)
	})

	reDonutChart := regexp.MustCompile(`(?s){{\s*donutchart\s+file="([^"]+)"(?:\s+title="([^"]*)")?\s*}}`)
	markdown = reDonutChart.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reDonutChart.FindStringSubmatch(match)
		filename, title := sub[1], sub[2]
		return generateDonutChart(sourceDir, filename, title)
	})

	reHBarchart := regexp.MustCompile(`(?s){{\s*hbarchart\s+file="([^"]+)"(?:\s+title="([^"]*)")?\s*}}`)
	markdown = reHBarchart.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reHBarchart.FindStringSubmatch(match)
		filename, title := sub[1], sub[2]
		return generateHBarchart(sourceDir, filename, title)
	})

	reMultiLine := regexp.MustCompile(`(?s){{\s*multilinechart\s+file="([^"]+)"(?:\s+title="([^"]*)")?\s*}}`)
	markdown = reMultiLine.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reMultiLine.FindStringSubmatch(match)
		filename, title := sub[1], sub[2]
		return generateMultiLineChart(sourceDir, filename, title)
	})

	reGroupedBar := regexp.MustCompile(`(?s){{\s*groupedbarchart\s+file="([^"]+)"(?:\s+title="([^"]*)")?\s*}}`)
	markdown = reGroupedBar.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reGroupedBar.FindStringSubmatch(match)
		filename, title := sub[1], sub[2]
		return generateGroupedBarChart(sourceDir, filename, title)
	})

	reRadar := regexp.MustCompile(`(?s){{\s*radarchart\s+file="([^"]+)"(?:\s+title="([^"]*)")?\s*}}`)
	markdown = reRadar.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reRadar.FindStringSubmatch(match)
		filename, title := sub[1], sub[2]
		return generateRadarChart(sourceDir, filename, title)
	})

	// Advanced Chart Widgets
	reDonutChartBlock := regexp.MustCompile(`(?s){{\s*donutchart(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/donutchart\s*}}`)
	markdown = reDonutChartBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reDonutChartBlock.FindStringSubmatch(match)
		title := sub[1]
		jsonData := []byte(sub[2])
		return generateDonutChartFromJSON(jsonData, title)
	})

	reHBarchartBlock := regexp.MustCompile(`(?s){{\s*hbarchart(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/hbarchart\s*}}`)
	markdown = reHBarchartBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reHBarchartBlock.FindStringSubmatch(match)
		title := sub[1]
		jsonData := []byte(sub[2])
		return generateHBarchartFromJSON(jsonData, title)
	})

	reMultiLineBlock := regexp.MustCompile(`(?s){{\s*multilinechart(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/multilinechart\s*}}`)
	markdown = reMultiLineBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reMultiLineBlock.FindStringSubmatch(match)
		title := sub[1]
		jsonData := []byte(sub[2])
		return generateMultiLineChartFromJSON(jsonData, title)
	})

	reGroupedBarBlock := regexp.MustCompile(`(?s){{\s*groupedbarchart(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/groupedbarchart\s*}}`)
	markdown = reGroupedBarBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reGroupedBarBlock.FindStringSubmatch(match)
		title := sub[1]
		jsonData := []byte(sub[2])
		return generateGroupedBarChartFromJSON(jsonData, title)
	})

	reRadarBlock := regexp.MustCompile(`(?s){{\s*radarchart(?:\s+title="([^"]*)")?\s*}}(.*?){{\s*/radarchart\s*}}`)
	markdown = reRadarBlock.ReplaceAllStringFunc(markdown, func(match string) string {
		sub := reRadarBlock.FindStringSubmatch(match)
		title := sub[1]
		jsonData := []byte(sub[2])
		return generateRadarChartFromJSON(jsonData, title)
	})

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
