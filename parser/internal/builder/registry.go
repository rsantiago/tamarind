package builder

import (
	"fmt"
)

var defaultPlugins []func() ShortcodePlugin

// RegisterDefaultPlugin allows plugins to self-register via init()
func RegisterDefaultPlugin(constructor func() ShortcodePlugin) {
	defaultPlugins = append(defaultPlugins, constructor)
}

// PluginRegistry manages all shortcode plugins and executes them against markdown content.
type PluginRegistry struct {
	plugins []ShortcodePlugin
}

// NewPluginRegistry creates a new instance of PluginRegistry and pre-loads all default plugins.
func NewPluginRegistry() *PluginRegistry {
	r := &PluginRegistry{
		plugins: make([]ShortcodePlugin, 0),
	}
	for _, constructor := range defaultPlugins {
		r.Register(constructor())
	}
	return r
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
