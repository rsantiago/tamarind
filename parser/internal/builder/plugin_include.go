package builder

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type IncludePlugin struct {
	pattern *regexp.Regexp
}

func NewIncludePlugin() *IncludePlugin {
	return &IncludePlugin{
		pattern: regexp.MustCompile(`{{\s*include\s+src="([^"]+)"(?:\s+lines="([0-9]+-[0-9]+)")?(?:\s+lang="([^"]+)")?\s*}}`),
	}
}

func (p *IncludePlugin) Name() string            { return "include" }
func (p *IncludePlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *IncludePlugin) Process(match []string, sourceDir string) (string, error) {
	src := match[1]
	linesRange := match[2]
	lang := match[3]
	if lang == "" {
		lang = "text"
	}

	var content []byte
	var err error

	if strings.HasPrefix(src, "http") {
		resp, err := http.Get(src)
		if err != nil {
			return fmt.Sprintf("> **Error fetching %s**: %v", src, err), nil
		}
		defer resp.Body.Close()
		content, err = io.ReadAll(resp.Body)
	} else {
		path := filepath.Join(sourceDir, src)
		content, err = os.ReadFile(path)
	}

	if err != nil {
		return fmt.Sprintf("> **Error including %s**: %v", src, err), nil
	}

	finalContent := string(content)
	if linesRange != "" {
		lParts := strings.Split(linesRange, "-")
		if len(lParts) == 2 {
			start, _ := strconv.Atoi(lParts[0])
			end, _ := strconv.Atoi(lParts[1])

			lines := strings.Split(finalContent, "\n")
			// Validate bounds (1-based -> 0-based)
			if start < 1 {
				start = 1
			}
			if end > len(lines) {
				end = len(lines)
			}
			if start <= end {
				finalContent = strings.Join(lines[start-1:end], "\n")
			}
		}
	}

	return fmt.Sprintf("```%s\n%s\n```", lang, finalContent), nil
}

func init() {
	RegisterDefaultPlugin(func() ShortcodePlugin { return NewIncludePlugin() })
}
