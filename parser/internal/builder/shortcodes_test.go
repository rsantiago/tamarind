package builder

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// createTestImage generates a blank PNG image of specific width/height
func createTestImage(dir, name string, width, height int) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Fill with white
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.White)
		}
	}

	path := filepath.Join(dir, name)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func TestProcessShortcodes_Figure(t *testing.T) {
	// Setup temporary directory for test images
	tmpDir, err := os.MkdirTemp("", "tamarind-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test images with different sizes
	// Breakpoints are 480, 800, 1200

	// 1. Small Image (300w) -> No breakpoints
	if err := createTestImage(tmpDir, "small.png", 300, 300); err != nil {
		t.Fatal(err)
	}

	// 2. Medium Image (600w) -> Only 480w
	if err := createTestImage(tmpDir, "medium.png", 600, 400); err != nil {
		t.Fatal(err)
	}

	// 3. Large Image (1000w) -> 480w, 800w
	if err := createTestImage(tmpDir, "large.png", 1000, 600); err != nil {
		t.Fatal(err)
	}

	// 4. Huge Image (1500w) -> 480w, 800w, 1200w
	if err := createTestImage(tmpDir, "huge.png", 1500, 900); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		input       string
		expected    []string // substrings expected in output
		notExpected []string // substrings that should NOT be in output
	}{
		{
			name:  "Small Image (No Scaling)",
			input: `{{ figure src="small.png" caption="Small" }}`,
			expected: []string{
				`<img src="small.png" alt="Small">`,
				`<figcaption>Small</figcaption>`,
			},
			notExpected: []string{
				"srcset",
			},
		},
		{
			name:  "Medium Image (600px)",
			input: `{{ figure src="medium.png" caption="Medium" }}`,
			expected: []string{
				`srcset="medium-480w.png 480w"`, // ONLY 480w
			},
			notExpected: []string{
				"800w", "1200w",
			},
		},
		{
			name:  "Large Image (1000px)",
			input: `{{ figure src="large.png" caption="Large" }}`,
			expected: []string{
				`large-480w.png 480w`,
				`large-800w.png 800w`,
			},
			notExpected: []string{
				"1200w",
			},
		},
		{
			name:  "Huge Image (1500px)",
			input: `{{ figure src="huge.png" caption="Huge" }}`,
			expected: []string{
				`huge-480w.png 480w`,
				`huge-800w.png 800w`,
				`huge-1200w.png 1200w`,
			},
		},
		{
			name:  "Missing Caption",
			input: `{{ figure src="medium.png" }}`,
			expected: []string{
				`<figure>`,
				`<img src="medium.png"`,
				`alt=""`,
			},
		},
		{
			name:  "External URL (No Srcset)",
			input: `{{ figure src="https://example.com/foo.jpg" caption="External" }}`,
			expected: []string{
				`<img src="https://example.com/foo.jpg" alt="External">`,
			},
			notExpected: []string{
				"srcset",
			},
		},
		{
			name:  "File Not Found (Fallback)",
			input: `{{ figure src="missing.png" caption="Missing" }}`,
			expected: []string{
				`<img src="missing.png" alt="Missing">`,
			},
			notExpected: []string{
				"srcset",
			},
		},
		{
			name:  "With Width",
			input: `{{ figure src="medium.png" width="50%" }}`,
			expected: []string{
				`style="width: 50%; margin: 0 auto; display: block;"`,
				`srcset="medium-480w.png 480w"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processShortcodes(tt.input, tmpDir)

			// Handle "Small Image" specific logic check from earlier comment
			// If 300px < 480px, loop produces 0 sources.
			// Code: if len(sources) > 0 { return ... } return fallback.
			// So small image should match Fallback format (no srcset).

			for _, exp := range tt.expected {
				if !strings.Contains(got, exp) {
					t.Errorf("[%s] Expected output to contain:\n%s\nGot:\n%s", tt.name, exp, got)
				}
			}

			for _, nexp := range tt.notExpected {
				if strings.Contains(got, nexp) {
					t.Errorf("[%s] Expected output NOT to contain:\n%s\nGot:\n%s", tt.name, nexp, got)
				}
			}
		})
	}
}

func TestProcessShortcodes_UIComponents(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Button Primary",
			input: `{{ button href="/about" type="primary" }}About{{ /button }}`,
			expected: []string{
				`<a href="/about" class="btn btn-primary">About</a>`,
			},
		},
		{
			name:  "Button Secondary Small Target",
			input: `{{< button href="https://example.com" type="secondary" size="sm" target="_blank" >}}External{{</ button >}}`,
			expected: []string{
				`<a href="https://example.com" class="btn btn-secondary btn-sm" target="_blank">External</a>`,
			},
		},
		{
			name:  "Card Standard",
			input: `{{ card }}Hello Card{{ /card }}`,
			expected: []string{
				`<div class="card card-padding">Hello Card</div>`,
			},
		},
		{
			name:  "Card No Padding",
			input: `{{< card padding="false" >}}No Padding{{</ card >}}`,
			expected: []string{
				`<div class="card">No Padding</div>`,
			},
		},
		{
			name:  "Alert standard info",
			input: `{{ alert type="info" title="Important info" }}This is content{{ /alert }}`,
			expected: []string{
				`<div class="callout callout-info alert alert-info">`,
				`<div class="callout-title alert-title">Important info</div>`,
				`<div class="callout-content alert-content">This is content</div>`,
			},
		},
		{
			name:  "Alert warn no title",
			input: `{{< alert type="warn" >}}Warn content{{</ alert >}}`,
			expected: []string{
				`<div class="callout callout-warn alert alert-warn">`,
				`<div class="callout-content alert-content">Warn content</div>`,
			},
		},
		{
			name:  "Badge primary",
			input: `{{ badge type="primary" }}Active{{ /badge }}`,
			expected: []string{
				`<span class="badge badge-primary">Active</span>`,
			},
		},
		{
			name:  "Badge standard no type",
			input: `{{< badge >}}New{{</ badge >}}`,
			expected: []string{
				`<span class="badge">New</span>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processShortcodes(tt.input, "")
			for _, exp := range tt.expected {
				if !strings.Contains(got, exp) {
					t.Errorf("[%s] Expected output to contain:\n%s\nGot:\n%s", tt.name, exp, got)
				}
			}
		})
	}
}

