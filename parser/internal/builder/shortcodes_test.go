package builder

import (
	"strings"
	"testing"
)

func TestProcessShortcodes_Figure(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // substrings expected in output
	}{
		{
			name:  "Responsive JPG",
			input: `{{ figure src="images/photo.jpg" caption="My Photo" }}`,
			expected: []string{
				`<figure>`,
				`<img src="images/photo.jpg"`,
				`srcset="images/photo-480w.jpg 480w, images/photo-800w.jpg 800w, images/photo-1200w.jpg 1200w"`,
				`sizes="(max-width: 480px) 100vw, (max-width: 800px) 100vw, 100vw"`,
				`alt="My Photo"`,
				`<figcaption>My Photo</figcaption>`,
				`</figure>`,
			},
		},
		{
			name:  "Responsive PNG",
			input: `{{ figure src="assets/image.png" caption="PNG Image" }}`,
			expected: []string{
				`srcset="assets/image-480w.png 480w`,
				`assets/image-800w.png 800w`,
			},
		},
		{
			name:  "Non-Responsive GIF",
			input: `{{ figure src="meme.gif" caption="Funny" }}`,
			expected: []string{
				`<img src="meme.gif" alt="Funny">`,
				`<figcaption>Funny</figcaption>`,
			},
		},
		{
			name:  "External URL (Fallback)",
			input: `{{ figure src="https://example.com/foo.jpg" caption="External" }}`,
			// Currently implementation logic is just extension check.
			// filepath.Ext("https://example.com/foo.jpg") is ".jpg".
			// So it WILL try to generate srcset with local suffixes.
			// Ideally we might want to skip remote URLs for srcset unless we know they exist?
			// But for now, let's just verify behavior matches current code.
			// Current code: `ext := filepath.Ext(src)` -> `.jpg` -> generates srcset.
			// This might be a bug/LIMITATION if we don't want to optimize external images (we can't generate the -480w versions for them locally).
			// The optimize logic only works on local resources.
			// So `figure` shortcode should probably NOT add srcset if the src starts with `http`.
			// Let's UPDATE the code to handle this, then update this test.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processShortcodes(tt.input, "")
			for _, exp := range tt.expected {
				if !strings.Contains(got, exp) {
					t.Errorf("Expected output to contain:\n%s\nGot:\n%s", exp, got)
				}
			}
			
			// Additional check for External Case: Ensure NO srcset
			if tt.name == "External URL (Fallback)" && strings.Contains(got, "srcset") {
				t.Errorf("External URL should not have srcset, but got:\n%s", got)
			}
		})
	}
}
