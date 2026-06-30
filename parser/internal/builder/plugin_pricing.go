package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type PricingPlugin struct {
	pricingGridCount int
	pattern          *regexp.Regexp
}

func NewPricingPlugin() *PricingPlugin {
	return &PricingPlugin{
		pricingGridCount: 0,
		pattern:          regexp.MustCompile(`(?s){{\s*pricing\s+(.*?)\s*}}(.*?){{\s*/pricing\s*}}`),
	}
}

func (p *PricingPlugin) Name() string            { return "pricing" }
func (p *PricingPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *PricingPlugin) Process(match []string, sourceDir string) (string, error) {
	attrString := match[1]
	content := match[2]

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

	p.pricingGridCount++
	gridID := fmt.Sprintf("pricing-grid-%d", p.pricingGridCount)

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

	hasToggle := false
	for _, plan := range plans {
		if (plan.priceMonthly != "" && plan.priceAnnual != "" && plan.priceMonthly != plan.priceAnnual) ||
			(plan.periodMonthly != "" && plan.periodAnnual != "" && plan.periodMonthly != plan.periodAnnual) ||
			(plan.urlMonthly != "" && plan.urlAnnual != "" && plan.urlMonthly != plan.urlAnnual) {
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

	plansHtml := ""
	for _, plan := range plans {
		var priceValHtml string
		var periodHtml string
		var buttonHtml string

		if hasToggle {
			priceValHtml = fmt.Sprintf(`<div class="price-val" data-monthly="%s" data-annual="%s">%s</div>`,
				formatPrice(plan.priceMonthly), formatPrice(plan.priceAnnual), formatPrice(plan.priceMonthly))
			periodHtml = fmt.Sprintf(`<div class="price-period" data-monthly="%s" data-annual="%s">%s</div>`,
				plan.periodMonthly, plan.periodAnnual, plan.periodMonthly)

			if plan.urlMonthly != "" || plan.urlAnnual != "" {
				buttonHtml = fmt.Sprintf(`<a href="%s" class="pricing-btn" data-monthly-url="%s" data-annual-url="%s">%s</a>`,
					plan.urlMonthly, plan.urlMonthly, plan.urlAnnual, plan.buttonText)
			} else {
				buttonHtml = fmt.Sprintf(`<button class="pricing-btn">%s</button>`, plan.buttonText)
			}
		} else {
			priceValHtml = fmt.Sprintf(`<div class="price-val">%s</div>`, formatPrice(plan.priceMonthly))
			periodHtml = fmt.Sprintf(`<div class="price-period">%s</div>`, plan.periodMonthly)

			if plan.urlMonthly != "" {
				buttonHtml = fmt.Sprintf(`<a href="%s" class="pricing-btn">%s</a>`, plan.urlMonthly, plan.buttonText)
			} else {
				buttonHtml = fmt.Sprintf(`<button class="pricing-btn">%s</button>`, plan.buttonText)
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
</div>`, plan.featuredClass, plan.badgeHtml, plan.title, priceValHtml, periodHtml, plan.bulletsHtml, buttonHtml)
	}

	if !hasToggle {
		return fmt.Sprintf(`
<div class="pricing-wrapper" id="%s">
  <div class="pricing-grid-poc">
    %s
  </div>
</div>`, gridID, plansHtml), nil
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
</script>`, gridID, monthlyLabel, gridID, annualLabel, discountBadgeHtml, plansHtml), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewPricingPlugin() })
}
