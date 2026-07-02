package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := `  {{ item title="Initialize Tamarind" number="1" }}
Spin up your local sandbox by running the initialization command in a fresh workspace directory:
` + "```bash\n" +
`tamarind init
` + "```\n" +
`  {{ /item }}
  {{ item title="Configure Cortex" number="2" }}
Tweak visual design tokens directly in ` + "`style.css`" + ` using theme variables:
` + "```css\n" +
`:root {
  --primary-color: #0047AB;
}
` + "```\n" +
`  {{ /item }}
`
	reItem1 := regexp.MustCompile(`(?s){{\s*item\s+title="([^"]+)"(?:\s+number="([^"]*)")?\s*}}(.*?){{\s*/item\s*}}`)
	matches := reItem1.FindAllStringSubmatch(content, -1)
	fmt.Printf("Matches: %d\n", len(matches))
}
