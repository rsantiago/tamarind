package builder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// LoadDataFiles scans the "data" directory in the source folder
// and loads YAML/JSON files into a map.
// Example: "data/authors.yaml" -> map["Authors"]...
func LoadDataFiles(sourceDir string) (map[string]interface{}, error) {
	dataDir := filepath.Join(sourceDir, "data")
	result := make(map[string]interface{})

	// check if dir exists
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return result, nil // Return empty map if no data dir
	}

	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("read data dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // only top level files for now
		}

		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))
		baseName := strings.TrimSuffix(name, ext)
		// Capitalize key for template access (e.g. "authors" -> "Authors")
		key := strings.Title(baseName) 

		filePath := filepath.Join(dataDir, name)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("read data file %s: %w", name, err)
		}

		var parsedData interface{}

		if ext == ".yaml" || ext == ".yml" {
			if err := yaml.Unmarshal(content, &parsedData); err != nil {
				return nil, fmt.Errorf("parse yaml %s: %w", name, err)
			}
		} else if ext == ".json" {
			if err := json.Unmarshal(content, &parsedData); err != nil {
				return nil, fmt.Errorf("parse json %s: %w", name, err)
			}
		} else {
			continue // Skip unknown files
		}

		result[key] = parsedData
	}

	return result, nil
}
