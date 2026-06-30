package builder

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type TabsPlugin struct {
	tabsIndex int
	pattern   *regexp.Regexp
}

func NewTabsPlugin() *TabsPlugin {
	return &TabsPlugin{
		tabsIndex: 0,
		pattern:   regexp.MustCompile(`(?s){{\s*tabs\s*}}(.*?){{\s*/tabs\s*}}`),
	}
}

func (p *TabsPlugin) Name() string            { return "tabs" }
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
		if err := unsafeGoldmark.Convert([]byte(dedentTabs(desc)), &buf); err == nil {
			htmlDesc = strings.TrimSpace(buf.String())
		} else {
			htmlDesc = dedentTabs(desc)
		}

		tabId := fmt.Sprintf("tab-%d-%d", p.tabsIndex, i)
		activeBtnClass := ""
		activePaneClass := ""
		if i == 0 {
			activeBtnClass = " active"
			activePaneClass = " active"
		}

		buttonsHtml += fmt.Sprintf(`<button class="tamarind-tab-btn%s" onclick="tamarindSwitchTab(event, '%s')">%s</button>`, activeBtnClass, tabId, title)
		panesHtml += fmt.Sprintf("<div id=\"%s\" class=\"tamarind-tab-pane%s\">\n%s\n</div>\n", tabId, activePaneClass, htmlDesc)
	}

	jsScript := `<script>
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
</script>`

	return fmt.Sprintf("\n\n%s\n<div class=\"tamarind-tabs\">\n<div class=\"tamarind-tabs-bar\">%s</div>\n<div class=\"tamarind-tabs-content\">\n%s\n</div>\n</div>\n\n", jsScript, buttonsHtml, panesHtml), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewTabsPlugin() })
}

func dedentTabs(s string) string {
	s = strings.TrimPrefix(s, "\n")
	s = strings.TrimRight(s, " \t\r\n")
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return s
	}
	minIndent := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		indent := len(line) - len(strings.TrimLeft(line, " \t"))
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}
	if minIndent <= 0 {
		return s
	}
	for i, line := range lines {
		if len(line) >= minIndent {
			lines[i] = line[minIndent:]
		}
	}
	return strings.Join(lines, "\n")
}
