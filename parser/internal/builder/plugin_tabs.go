package builder

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type TabsPlugin struct {
	tabsIndex      int
	pattern        *regexp.Regexp
}

func NewTabsPlugin() *TabsPlugin {
	return &TabsPlugin{
		tabsIndex:      0,
		pattern:        regexp.MustCompile(`(?s){{\s*tabs\s*}}(.*?){{\s*/tabs\s*}}`),
	}
}

func (p *TabsPlugin) Name() string { return "tabs" }
func (p *TabsPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *TabsPlugin) Process(match []string, sourceDir string) (string, error) {
	p.tabsIndex++
	content := match[1]

	reTabItem := regexp.MustCompile(`(?s){{\s*tab\s+title="([^"]+)"\s*}}(.*?){{\s*/tab\s*}}`)
	itemMatches := reTabItem.FindAllStringSubmatch(content, -1)
	if len(itemMatches) == 0 {
		return "", nil
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

		tabId := fmt.Sprintf("tab-%d-%d", p.tabsIndex, i)
		activeBtnClass := ""
		activePaneClass := ""
		if i == 0 {
			activeBtnClass = " active"
			activePaneClass = " active"
		}

		buttonsHtml += fmt.Sprintf(`<button class="tamarind-tab-btn%s" onclick="tamarindSwitchTab(event, '%s')">%s</button>`, activeBtnClass, tabId, title)
		panesHtml += fmt.Sprintf(`<div id="%s" class="tamarind-tab-pane%s">%s</div>`, tabId, activePaneClass, htmlDesc)
	}

	jsScript := `
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

	return fmt.Sprintf(`%s<div class="tamarind-tabs"><div class="tamarind-tabs-bar">%s</div><div class="tamarind-tabs-content">%s</div></div>`, jsScript, buttonsHtml, panesHtml), nil
}
