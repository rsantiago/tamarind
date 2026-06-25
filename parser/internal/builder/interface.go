package builder

import "regexp"

// ShortcodePlugin defines the contract for all Tamarind shortcodes.
type ShortcodePlugin interface {
	// Name returns the unique identifier for the shortcode.
	Name() string

	// Pattern returns the compiled regex used to find the shortcode.
	Pattern() *regexp.Regexp

	// Process executes the replacement logic for a specific match.
	Process(match []string, sourceDir string) (string, error)
}
