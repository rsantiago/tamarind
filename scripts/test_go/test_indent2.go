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
	input := `<script>console.log();</script><div>hello
    world</div>`
	var buf bytes.Buffer
	md.Convert([]byte(input), &buf)
	fmt.Println("---")
	fmt.Println(buf.String())
	fmt.Println("---")
}
