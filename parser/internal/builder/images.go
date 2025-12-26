package builder

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"fmt"

	"golang.org/x/image/draw"
)
// Image Optimization Notes:
// - WebP encoding without CGO is limited in Go. Skipping WebP for now.
// - Focusing on responsive resizing (breakpoints).


func OptimizeImage(srcPath, destDir string) error {
	ext := strings.ToLower(filepath.Ext(srcPath))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return nil // Skip non-image files
	}

	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err // Could not decode, maybe not an image
	}

	// Calculate sizes
	bounds := img.Bounds()
	width := bounds.Dx()
	
	// Define breakpoints (width)
	breakpoints := []int{480, 800, 1200}

	baseName := strings.TrimSuffix(filepath.Base(srcPath), ext)
	
	for _, bp := range breakpoints {
		if width >= bp {
			// Calculate height to maintain aspect ratio
			ratio := float64(bounds.Dy()) / float64(bounds.Dx())
			height := int(float64(bp) * ratio)

			newImg := image.NewRGBA(image.Rect(0, 0, bp, height))
			
			// High quality resampling
			draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

			// Save
			newFileName := fmt.Sprintf("%s-%dw%s", baseName, bp, ext)
			newPath := filepath.Join(destDir, newFileName)
			
			out, err := os.Create(newPath)
			if err != nil {
				return err
			}
			defer out.Close()

			if ext == ".png" {
				if err := png.Encode(out, newImg); err != nil {
					return err
				}
			} else {
				if err := jpeg.Encode(out, newImg, &jpeg.Options{Quality: 85}); err != nil {
					return err
				}
			}
		}
	}
	
	return nil
}
