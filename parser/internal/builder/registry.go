package builder

import (
	"fmt"
)

// PluginRegistry manages all shortcode plugins and executes them against markdown content.
type PluginRegistry struct {
	plugins []ShortcodePlugin
}

// NewPluginRegistry creates a new instance of PluginRegistry.
func NewPluginRegistry() *PluginRegistry {
	return &PluginRegistry{
		plugins: make([]ShortcodePlugin, 0),
	}
}

// Register adds a new shortcode plugin to the registry.
func (r *PluginRegistry) Register(plugin ShortcodePlugin) {
	r.plugins = append(r.plugins, plugin)
}

// ProcessShortcodes iterates through all registered plugins and applies them to the markdown.
func (r *PluginRegistry) ProcessShortcodes(markdown string, sourceDir string) string {
	for _, plugin := range r.plugins {
		pattern := plugin.Pattern()
		
		markdown = pattern.ReplaceAllStringFunc(markdown, func(match string) string {
			submatches := pattern.FindStringSubmatch(match)
			
			result, err := plugin.Process(submatches, sourceDir)
			if err != nil {
				return fmt.Sprintf("<!-- Shortcode Error: %s -->", err.Error())
			}
			
			return result
		})
	}
	return markdown
}
