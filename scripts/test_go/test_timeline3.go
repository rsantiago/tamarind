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
	reItem1 := regexp.MustCompile(`(?s){{\s*item\s+title="([^"]+)"(?:\s+number="([^"]*)")?\s*}}(.*?){{\s*/item\s*}}`)

	matches := reTimeline.FindAllStringSubmatch(string(content), -1)
	fmt.Printf("Matches for timeline: %d\n", len(matches))
	for i, match := range matches {
		fmt.Printf("Timeline %d:\n", i)
		inner := match[1]
		itemMatches := reItem1.FindAllStringSubmatch(inner, -1)
		fmt.Printf("  Items found: %d\n", len(itemMatches))
	}
}
