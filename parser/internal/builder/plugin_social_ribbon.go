package builder

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type SocialRibbonPlugin struct {
	pattern *regexp.Regexp
}

func NewSocialRibbonPlugin() *SocialRibbonPlugin {
	return &SocialRibbonPlugin{
		pattern: regexp.MustCompile(`(?s){{\s*social_ribbon\s*}}(.*?){{\s*/social_ribbon\s*}}`),
	}
}

func (p *SocialRibbonPlugin) Name() string { return "social_ribbon" }
func (p *SocialRibbonPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *SocialRibbonPlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

	reTestimonial := regexp.MustCompile(`(?s){{\s*testimonial\s+(.*?)\s*}}(.*?){{\s*/testimonial\s*}}`)
	
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
		
		testimonials = append(testimonials, testimonialData{
			stars:  attrs["stars"],
			avatar: attrs["avatar"],
			author: attrs["author"],
			handle: attrs["handle"],
			quote:  quote,
		})
	}
	
	renderCards := func(items []testimonialData) string {
		var sb strings.Builder
		for _, t := range items {
			starsDisplay := ""
			if numStars, err := strconv.Atoi(t.stars); err == nil {
				for i := 0; i < numStars; i++ {
					starsDisplay += "★"
				}
			} else {
				starsDisplay = t.stars
			}
			if starsDisplay == "" {
				starsDisplay = "★★★★★"
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
	
	originalCardsHtml := renderCards(testimonials)
	
	return fmt.Sprintf(`<div class="tamarind-social-ribbon-container"><div class="tamarind-social-ribbon-track">%s%s</div></div>`, originalCardsHtml, originalCardsHtml), nil
}
