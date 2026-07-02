package main
import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)
func main() {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	input := `Some text before
  <script></script><div id="term-tab-7-0" class="terminal-tab-pane active"><pre class="terminal-content"><code>$ tamarind themes
$ tamarind build --theme dark</code></pre></div>`
	var buf bytes.Buffer
	md.Convert([]byte(input), &buf)
	fmt.Println("---")
	fmt.Println(buf.String())
	fmt.Println("---")
}
