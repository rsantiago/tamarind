package utils

import (

	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// ExtractDir keeps the same logic but accepts an fs.FS interface
func ExtractDir(srcFS fs.FS, srcPrefix, dstRoot string, force bool) error {
	// Check if destination exists
	if _, err := os.Stat(dstRoot); err == nil && !force {
		log.Printf("Directory %s already exists. Use --force to overwrite files.", dstRoot)
	}

	return fs.WalkDir(srcFS, srcPrefix, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Strip the prefix from the path to get the relative destination path
		relPath, err := filepath.Rel(srcPrefix, path)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil
		}

		dstPath := filepath.Join(dstRoot, relPath)

		if d.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		}

		// Check if file exists
		if _, err := os.Stat(dstPath); err == nil && !force {
			return nil
		}

		// Read from embedded FS
		data, err := fs.ReadFile(srcFS, path)
		if err != nil {
			return err
		}

		// Write to disk
		log.Printf("Extracting: %s", dstPath)
		return os.WriteFile(dstPath, data, 0644)
	})
}
