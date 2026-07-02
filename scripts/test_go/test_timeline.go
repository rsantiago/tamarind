package main

import (
	"fmt"
	"regexp"
	"os"
)

func main() {
	content, err := os.ReadFile("parser/assets/structure/docs/style-guide.md")
	if err != nil {
		panic(err)
	}
	reTimeline := regexp.MustCompile(`(?s){{\s*timeline\s*}}(.*?){{\s*/timeline\s*}}`)
	matches := reTimeline.FindAllStringSubmatch(string(content), -1)
	fmt.Printf("Matches for timeline: %d\n", len(matches))
}
