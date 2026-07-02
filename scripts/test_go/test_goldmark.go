package main

import (
	"bytes"
	"fmt"
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

func main() {
	desc := `Run the quickstart command in your terminal. This boots up the development server instantly with zero configuration:
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
</script><div class="terminal">hello</div>
  Your sample website will be compiled and served locally at http://localhost:8000 with automatic hot-reloading!`

	var buf1 bytes.Buffer
	unsafeGoldmark.Convert([]byte(strings.TrimSpace(desc)), &buf1)
	htmlDesc := buf1.String()

	// Simulate timeline plugin wrapper
	timelineHtml := fmt.Sprintf(`<div class="timeline-container"><div class="timeline-item"><div class="timeline-content"><div class="timeline-desc">%s</div></div></div></div>`, htmlDesc)

	// Second pass of Goldmark
	var buf2 bytes.Buffer
	unsafeGoldmark.Convert([]byte(timelineHtml), &buf2)
	
	fmt.Println(buf2.String())
}
