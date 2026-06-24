package builder
import "testing"
import "os"
import "path/filepath"
import "fmt"
func TestDebugContrast(t *testing.T) {
	cssPath := filepath.Join("..", "..", "assets", "templates", "brutal", "style.css")
	content, _ := os.ReadFile(cssPath)
	analysis, _ := AnalyzeCSS(string(content))
	c1 := analysis.LightVars["--chart-1"]
	c2 := analysis.LightVars["--chart-2"]
	fmt.Printf("Brutal Light Chart 1: '%s'\n", c1)
	fmt.Printf("Brutal Light Chart 2: '%s'\n", c2)
	r1, g1, b1, ok1 := ParseColor(ResolveVal(c1, analysis.LightVars))
	fmt.Printf("Parsed 1: %v %v %v %v\n", r1, g1, b1, ok1)
	c1d := analysis.DarkVars["--chart-1"]
	fmt.Printf("Brutal Dark Chart 1: '%s'\n", c1d)
}
