package builder

import (
	"fmt"
	"regexp"
	"strings"
)

type TerminalPlugin struct {
	terminalIndex int
	pattern       *regexp.Regexp
}

func NewTerminalPlugin() *TerminalPlugin {
	return &TerminalPlugin{
		terminalIndex: 0,
		pattern:       regexp.MustCompile(`(?s){{\s*terminal\s*}}(.*?){{\s*/terminal\s*}}`),
	}
}

func (p *TerminalPlugin) Name() string            { return "terminal" }
func (p *TerminalPlugin) Pattern() *regexp.Regexp { return p.pattern }

func dedentTerminal(s string) string {
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

func (p *TerminalPlugin) Process(match []string, sourceDir string) (string, error) {
	p.terminalIndex++
	content := match[1]
	copyBtnHtml := `<button class="terminal-copy-btn" onclick="tamarindCopyTerminal(event)" aria-label="Copy" title="Copy"><svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button>`

	jsScript := `<script>
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
</script>`

	reTab := regexp.MustCompile(`(?s){{\s*tab\s+title="([^"]+)"\s*}}(.*?){{\s*/tab\s*}}`)
	tabMatches := reTab.FindAllStringSubmatch(content, -1)

	if len(tabMatches) > 0 {
		buttonsHtml := ""
		panesHtml := ""
		for i, tabMatch := range tabMatches {
			title := tabMatch[1]
			tabContent := dedentTerminal(tabMatch[2])

			tabId := fmt.Sprintf("term-tab-%d-%d", p.terminalIndex, i)
			activeBtnClass := ""
			activePaneClass := ""
			if i == 0 {
				activeBtnClass = " active"
				activePaneClass = " active"
			}

			buttonsHtml += fmt.Sprintf(`<button class="terminal-tab-btn%s" onclick="tamarindSwitchTerminalTab(event, '%s')">%s</button>`, activeBtnClass, tabId, title)
			panesHtml += fmt.Sprintf(`<div id="%s" class="terminal-tab-pane%s"><pre class="terminal-content"><code>%s</code></pre></div>`, tabId, activePaneClass, tabContent)
		}

		return fmt.Sprintf("\n\n%s\n<div class=\"terminal terminal-has-tabs\">%s<div class=\"terminal-header\"><div class=\"terminal-dots\"><span class=\"dot red\"></span><span class=\"dot yellow\"></span><span class=\"dot green\"></span></div><div class=\"terminal-tabs-bar\">%s</div></div>%s</div>\n\n", jsScript, copyBtnHtml, buttonsHtml, panesHtml), nil
	} else {
		content = dedentTerminal(content)
		return fmt.Sprintf("\n\n%s\n<div class=\"terminal\">%s<div class=\"terminal-header\"><span class=\"dot red\"></span><span class=\"dot yellow\"></span><span class=\"dot green\"></span></div><pre class=\"terminal-content\"><code>%s</code></pre></div>\n\n", jsScript, copyBtnHtml, content), nil
	}
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewTerminalPlugin() })
}
