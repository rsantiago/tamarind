// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"os"
	"path/filepath"
	"strings"
)

func copyAssets(templateDir, websiteDir string) error {
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		// Skip template files (.mdt)
		if filepath.Ext(name) == ".mdt" {
			continue
		}

		src := filepath.Join(templateDir, name)
		dst := filepath.Join(websiteDir, name)

		data, err := os.ReadFile(src)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dst, data, 0644); err != nil {
			return err
		}
	}
	return nil
}

func copySiteResources(sourceDir, websiteDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(sourceDir, path)
		if relPath == "." {
			return nil
		}

		// Skip special dirs/files (hidden, templates)
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") || name == "templates" || name == "website" {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip markdown, templates, config, hidden files
		ext := strings.ToLower(filepath.Ext(path))
		name := info.Name()
		if ext == ".md" || ext == ".mdt" || name == "tamarind.yaml" || strings.HasPrefix(name, ".") {
			return nil
		}

		// Copy file (images, pdfs, etc.)
		destPath := filepath.Join(websiteDir, relPath)
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(destPath, data, 0644); err != nil {
			return err
		}

		// Image Optimization Pipeline
		if err := OptimizeImage(destPath, destDir); err != nil {
			// Log warning but don't fail, maybe just not an image we can handle or read error
			// fmt.Printf("Warning: Failed to optimize image %s: %v\n", relPath, err)
		}

		return nil
	})
}
