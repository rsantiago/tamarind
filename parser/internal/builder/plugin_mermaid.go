package builder

import (
	"fmt"
	"regexp"
)

type MermaidPlugin struct {
	scriptInjected bool
	pattern        *regexp.Regexp
}

func NewMermaidPlugin() *MermaidPlugin {
	return &MermaidPlugin{
		scriptInjected: false,
		pattern:        regexp.MustCompile(`(?s){{\s*mermaid\s*}}(.*?){{\s*/mermaid\s*}}`),
	}
}

func (p *MermaidPlugin) Name() string { return "mermaid" }
func (p *MermaidPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *MermaidPlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

	jsScript := ""
	if !p.scriptInjected {
		jsScript = `
<script>
if (typeof tamarindMaximizeMermaid !== 'function') {
	window.tamarindMaximizeMermaid = function(btn) {
		var wrapper = btn.closest('.mermaid-wrapper');
		var span = btn.querySelector('span');
		var svgPath = btn.querySelector('path');
		if (wrapper.classList.contains('maximized')) {
			wrapper.classList.remove('maximized');
			if(span) span.innerText = 'Maximize';
			if(svgPath) svgPath.setAttribute('d', 'M8 3H5a2 2 0 0 0-2 2v3m18 0V5a2 2 0 0 0-2-2h-3m0 18h3a2 2 0 0 0 2-2v-3M3 16v3a2 2 0 0 0 2 2h3');
			document.body.style.overflow = '';
		} else {
			wrapper.classList.add('maximized');
			if(span) span.innerText = 'Close';
			if(svgPath) svgPath.setAttribute('d', 'M8 3v3a2 2 0 0 1-2 2H3m18 0h-3a2 2 0 0 1-2-2V3m0 18v-3a2 2 0 0 1 2-2h3M3 16h3a2 2 0 0 1 2 2v3');
			document.body.style.overflow = 'hidden';
		}
	};
}
</script>
<style>
.mermaid-wrapper { position: relative; }
.mermaid-maximize-btn {
	position: absolute; top: 12px; right: 12px; z-index: 10;
	display: flex; align-items: center; gap: 6px;
	padding: 6px 14px; background: rgba(128, 128, 128, 0.15);
	backdrop-filter: blur(8px); -webkit-backdrop-filter: blur(8px);
	border: 1px solid rgba(128, 128, 128, 0.2); border-radius: 20px;
	color: currentColor; font-size: 13px; font-weight: 500;
	cursor: pointer; opacity: 0; transform: translateY(-5px);
	transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
	box-shadow: 0 4px 12px rgba(0,0,0,0.05);
}
.mermaid-wrapper:hover .mermaid-maximize-btn {
	opacity: 1; transform: translateY(0);
}
.mermaid-maximize-btn:hover {
	background: var(--primary, #3b82f6); color: #fff; border-color: transparent;
	transform: translateY(-2px); box-shadow: 0 6px 16px rgba(0,0,0,0.2);
}
.mermaid-maximize-btn:active {
	transform: translateY(1px);
}
.mermaid-wrapper.maximized {
	position: fixed; top: 0; left: 0; width: 100vw; height: 100vh;
	background: var(--background-color, #fff); z-index: 9999;
	display: flex; align-items: center; justify-content: center;
	overflow: auto; padding: 40px; margin: 0 !important;
}
.mermaid-wrapper.maximized .mermaid {
	width: 100%; height: 100%; display: flex;
	align-items: center; justify-content: center;
}
.mermaid-wrapper.maximized .mermaid svg {
	max-height: 100% !important; max-width: 100% !important;
}
</style>
`
		p.scriptInjected = true
	}
	
	btnHtml := `<button class="mermaid-maximize-btn" onclick="tamarindMaximizeMermaid(this)">
		<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M8 3H5a2 2 0 0 0-2 2v3m18 0V5a2 2 0 0 0-2-2h-3m0 18h3a2 2 0 0 0 2-2v-3M3 16v3a2 2 0 0 0 2 2h3"/></svg>
		<span>Maximize</span>
		</button>`
	return fmt.Sprintf(`%s<div class="mermaid-wrapper">%s<div class="mermaid">%s</div></div>`, jsScript, btnHtml, content), nil
}
