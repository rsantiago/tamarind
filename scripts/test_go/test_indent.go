package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	
	input := `List available themes:
  <script>console.log();</script><div class="terminal"><pre class="terminal-content"><code>$ tamarind themes
    $ tamarind build --theme dark</code></pre></div><div id="term-tab-7-1" class="terminal-tab-pane"><pre class="terminal-content"><code>$ tamarind themes
    $ tamarind build --theme dark</code></pre></div></div>`

	var buf bytes.Buffer
	md.Convert([]byte(input), &buf)
	fmt.Println(buf.String())
}
