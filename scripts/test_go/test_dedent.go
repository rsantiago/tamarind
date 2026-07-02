package main
import (
	"fmt"
	"strings"
)

func dedent(s string) string {
	s = strings.TrimPrefix(s, "\n")
	s = strings.TrimRight(s, " \t\r\n")
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return s
	}
	// find min indent
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

func main() {
	s := `
    $ tamarind themes
    $ tamarind build --theme dark
    `
	fmt.Println(dedent(s))
}
