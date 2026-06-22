// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	ConfigFileName = "tamarind.yaml"
)

var Version = "v4.16.0"

type Config struct {
	ThemeConfig map[string]string `yaml:"theme_config"`
}

func CreateDefaultConfig() error {
	if _, err := os.Stat(ConfigFileName); err == nil {
		fmt.Printf("Config file '%s' already exists.\n", ConfigFileName)
		return nil
	}

	// We write the raw string to ensure it is commented out by default
	defaultConfigContent := `# Tamarind Configuration File
# ------------------------
# This file allows you to override the default theme settings.
# By default, all values are commented out (#), so the theme's default colors are used.
#
# To customize your site:
# 1. Uncomment (remove '#') the line you want to change.
# 2. Change the value (e.g., set primary-color to "#ff0000").
# 3. Run 'tamarind build -theme <name>' to see the changes.

theme_config:
  # primary-color:    "#3b82f6"   # The main accent color (links, buttons)
  # background-color: "#ffffff"   # Page background
  # text-color:       "#1f2937"   # Main text color
  # heading-color:    "#111827"   # Color for titles (h1, h2, etc.)
  # font-family:      "'Inter', sans-serif"
  # pagination-limit: "10"        # Number of posts per page (default: 10)
`

	if err := os.WriteFile(ConfigFileName, []byte(defaultConfigContent), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Printf("Created '%s'.\n", ConfigFileName)
	fmt.Println("You can now edit this file to customize colors and fonts.")
	fmt.Println("Run 'tamarind build -theme <name>' to apply changes.")
	return nil
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(ConfigFileName)
	if os.IsNotExist(err) {
		return nil, nil // No config file is fine
	}
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Printf("Warning: Failed to parse %s: %v", ConfigFileName, err)
		return nil, nil // Treat as empty config on error to avoid breaking build
	}

	return &cfg, nil
}
