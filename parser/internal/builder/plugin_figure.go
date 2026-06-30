package builder

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FigurePlugin struct {
	pattern *regexp.Regexp
}

func NewFigurePlugin() *FigurePlugin {
	return &FigurePlugin{
		pattern: regexp.MustCompile(`{{\s*figure\s+(.*?)\s*}}`),
	}
}

func (p *FigurePlugin) Name() string            { return "figure" }
func (p *FigurePlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *FigurePlugin) Process(match []string, sourceDir string) (string, error) {
	content := match[1]

	reSrc := regexp.MustCompile(`src="([^"]+)"`)
	reCap := regexp.MustCompile(`caption="([^"]+)"`)
	reWidth := regexp.MustCompile(`width="([^"]+)"`)

	srcMatch := reSrc.FindStringSubmatch(content)
	if srcMatch == nil {
		return match[0], nil // Invalid, no src
	}
	src := srcMatch[1]

	caption := ""
	capMatch := reCap.FindStringSubmatch(content)
	if capMatch != nil {
		caption = capMatch[1]
	}

	width := ""
	widthMatch := reWidth.FindStringSubmatch(content)
	if widthMatch != nil {
		width = widthMatch[1] // e.g. "500px" or "50%"
	}

	var figcaptionHTML string
	if caption != "" {
		figcaptionHTML = fmt.Sprintf("<figcaption>%s</figcaption>", caption)
	}

	// Style for width
	styleAttr := ""
	if width != "" {
		styleAttr = fmt.Sprintf(` style="width: %s; margin: 0 auto; display: block;"`, width)
	}

	// Responsive Logic
	// We assume src matches the generated optimized files: name-width.ext
	// AND that it is a local resource (not http/s)
	if !strings.HasPrefix(src, "http") {
		// Resolve path to check dimensions
		localPath := filepath.Join(sourceDir, src)

		// Check if file exists and get dimensions
		f, err := os.Open(localPath)
		if err == nil {
			defer f.Close()
			cfg, _, err := image.DecodeConfig(f)
			if err == nil {
				imgWidth := cfg.Width
				ext := filepath.Ext(src)
				base := strings.TrimSuffix(src, ext)

				lowerExt := strings.ToLower(ext)
				if lowerExt == ".jpg" || lowerExt == ".jpeg" || lowerExt == ".png" {
					var sources []string
					breakpoints := []int{480, 800, 1200}

					for _, bp := range breakpoints {
						if imgWidth >= bp {
							sources = append(sources, fmt.Sprintf("%s-%dw%s %dw", base, bp, ext, bp))
						}
					}

					if len(sources) > 0 {
						srcset := strings.Join(sources, ", ")
						sizes := "(max-width: 480px) 100vw, (max-width: 800px) 100vw, 100vw"
						return fmt.Sprintf(`<figure><img src="%s" srcset="%s" sizes="%s" alt="%s"%s>%s</figure>`,
							src, srcset, sizes, caption, styleAttr, figcaptionHTML), nil
					}
				}
			}
		}
	}

	// Fallback
	return fmt.Sprintf(`<figure><img src="%s" alt="%s"%s>%s</figure>`, src, caption, styleAttr, figcaptionHTML), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewFigurePlugin() })
}
