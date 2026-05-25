// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
)

func TestOptimizeImage(t *testing.T) {
	// 1. Setup Temp Dirs
	tmpDir, err := os.MkdirTemp("", "tamarind-img-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	srcDir := filepath.Join(tmpDir, "src")
	destDir := filepath.Join(tmpDir, "dest")
	os.Mkdir(srcDir, 0755)
	os.Mkdir(destDir, 0755)

	// 2. Create Dummy Large Image (2000x1000)
	// We want to verify it scales down to 1200, 800, 480.
	imgName := "test-image.jpg"
	srcPath := filepath.Join(srcDir, imgName)

	// Create a red image
	img := image.NewRGBA(image.Rect(0, 0, 2000, 1000))
	for y := 0; y < 1000; y++ {
		for x := 0; x < 2000; x++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	f, err := os.Create(srcPath)
	if err != nil {
		t.Fatal(err)
	}
	jpeg.Encode(f, img, nil)
	f.Close()

	// 3. Run Optimization
	// Note: OptimizeImage(srcPath, destDir) reads srcPath and writes variants to destDir
	if err := OptimizeImage(srcPath, destDir); err != nil {
		t.Fatalf("OptimizeImage failed: %v", err)
	}

	// 4. Verify Outputs
	expectedWidths := []int{480, 800, 1200}

	for _, w := range expectedWidths {
		variantName := "test-image-" + "480w" + ".jpg" // wait, logic is "%s-%dw%s"
		// Need to match exact format from implementation
		// newFileName := fmt.Sprintf("%s-%dw%s", baseName, bp, ext)
		// ext includes dot.

		variantName = "test-image-" + keyStr(w) + "w.jpg"
		variantPath := filepath.Join(destDir, variantName)

		info, err := os.Stat(variantPath)
		if os.IsNotExist(err) {
			t.Errorf("Expected variant %s not found", variantName)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("Variant %s is empty", variantName)
		}

		// Verify Dimensions
		file, _ := os.Open(variantPath)
		cfg, _, err := image.DecodeConfig(file)
		file.Close()
		if err != nil {
			t.Errorf("Failed to decode variant %s: %v", variantName, err)
			continue
		}

		if cfg.Width != w {
			t.Errorf("Variant %s: expected width %d, got %d", variantName, w, cfg.Width)
		}

		// Expected Height: 2000x1000 (2:1 ratio) => 480 -> 240
		expectedHeight := w / 2
		if cfg.Height != expectedHeight {
			t.Errorf("Variant %s: expected height %d, got %d", variantName, expectedHeight, cfg.Height)
		}
	}
}

// Helper to convert int to string manually to avoid import fmt if not needed,
// strictly speaking fmt is fine in tests.
func keyStr(i int) string {
	// Start basic
	if i == 480 {
		return "480"
	}
	if i == 800 {
		return "800"
	}
	if i == 1200 {
		return "1200"
	}
	return ""
}
