// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

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
			registry := BuildPluginRegistry()
			got := processShortcodes(registry, tt.input, tmpDir)

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
				`<div class="card card-padding">

Hello Card

</div>`,
			},
		},
		{
			name:  "Card No Padding",
			input: `{{ card padding="false" }}No Padding{{ /card }}`,
			expected: []string{
				`<div class="card">

No Padding

</div>`,
			},
		},
		{
			name:  "Terminal Standard",
			input: `{{ terminal }}go run main.go{{ /terminal }}`,
			expected: []string{
				`<div class="terminal">`,
				`<div class="terminal-header">`,
				`<pre class="terminal-content"><code>go run main.go</code></pre>`,
			},
		},
		{
			name: "Terminal Tabbed",
			input: `{{ terminal }}
{{ tab title="Go" }}
go run main.go
{{ /tab }}
{{ tab title="Python" }}
python main.py
{{ /tab }}
{{ /terminal }}`,
			expected: []string{
				`<div class="terminal terminal-has-tabs">`,
				`<div class="terminal-tabs-bar">`,
				`<button class="terminal-tab-btn active" onclick="tamarindSwitchTerminalTab(event, 'term-tab-`,
				`Go</button>`,
				`<button class="terminal-tab-btn" onclick="tamarindSwitchTerminalTab(event, 'term-tab-`,
				`Python</button>`,
				`class="terminal-tab-pane active"><pre class="terminal-content"><code>go run main.go</code></pre></div>`,
				`class="terminal-tab-pane"><pre class="terminal-content"><code>python main.py</code></pre></div>`,
			},
		},
		{
			name:  "Alert standard info",
			input: `{{ alert type="info" title="Important info" }}This is content{{ /alert }}`,
			expected: []string{
				`<div class="alert-container alert-info">`,
				`<h4 class="alert-title">Important info</h4>`,
				`<div class="alert-message"><p>This is content</p>`,
			},
		},
		{
			name:  "Alert warn no title",
			input: `{{< alert type="warn" >}}Warn content{{</ alert >}}`,
			expected: []string{
				`<div class="alert-container alert-warn">`,
				`<div class="alert-message"><p>Warn content</p>`,
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
		{
			name:  "Form standard",
			input: `{{ form action="/contact" method="POST" }}FormContent{{ /form }}`,
			expected: []string{
				`<form action="/contact" method="POST">FormContent</form>`,
			},
		},
		{
			name:  "Form-input text field",
			input: `{{ form-input label="Full Name" type="text" placeholder="John Doe" }}`,
			expected: []string{
				`<div class="form-group">`,
				`<label class="form-label">Full Name</label>`,
				`<input type="text" class="form-input" placeholder="John Doe">`,
			},
		},
		{
			name:  "Form-textarea custom rows",
			input: `{{ form-textarea label="Bio" placeholder="Tell us about yourself..." rows="6" }}`,
			expected: []string{
				`<div class="form-group">`,
				`<label class="form-label">Bio</label>`,
				`<textarea class="form-textarea" rows="6" placeholder="Tell us about yourself..."></textarea>`,
			},
		},
		{
			name:  "Form-select dropdown options",
			input: `{{ form-select label="Priority" }}{{ option value="high" selected="true" }}High{{ /option }}{{ option }}Low{{ /option }}{{ /form-select }}`,
			expected: []string{
				`<div class="form-group">`,
				`<label class="form-label">Priority</label>`,
				`<select class="form-select">`,
				`<option value="high" selected>High</option>`,
				`<option>Low</option>`,
			},
		},
		{
			name:  "Form-checkbox checked",
			input: `{{ form-checkbox label="I agree" checked="true" }}`,
			expected: []string{
				`<div class="form-group"><label class="form-label"><input type="checkbox" class="form-checkbox" checked> I agree</label></div>`,
			},
		},
		{
			name:  "Form-radio-group with sub-options",
			input: `{{ form-radio-group label="Gender" }}{{ form-radio name="gender" label="Male" checked="true" }}{{ form-radio name="gender" label="Female" }}{{ /form-radio-group }}`,
			expected: []string{
				`<div class="form-group">`,
				`<label class="form-label">Gender</label>`,
				`<label class="form-label"><input type="radio" name="gender" class="form-radio" checked> Male</label>`,
				`<label class="form-label"><input type="radio" name="gender" class="form-radio"> Female</label>`,
			},
		},
		{
			name:  "Form-file upload field",
			input: `{{ form-file label="Resume" }}`,
			expected: []string{
				`<div class="form-group">`,
				`<label class="form-label">Resume</label>`,
				`<input type="file" class="form-file">`,
			},
		},
		{
			name:  "Vertical Timeline with nested HTML console block",
			input: `{{ timeline }}{{ item title="Setup" number="1" }}Initialize <pre><code>tamarind init</code></pre>{{ /item }}{{ /timeline }}`,
			expected: []string{
				`<div class="timeline-container">`,
				`<div class="timeline-item">`,
				`<div class="timeline-badge"><span class="timeline-badge-number">1</span></div>`,
				`<div class="timeline-content">`,
				`<h3 class="timeline-title">Setup</h3>`,
				`<div class="timeline-desc"><p>Initialize <pre><code>tamarind init</code></pre></p>`,
			},
		},
		{
			name: "Pricing Grid with plans (Toggle Mode)",
			input: `{{ pricing monthly_label="Monthly" annual_label="Annual" discount="Save 20%" }}
  {{ plan title="Personal" price_monthly="19" price_annual="15" period_monthly="per month" period_annual="billed annually" button="Get Started" url_monthly="https://checkout.stripe.com/monthly-pro" url_annual="https://checkout.stripe.com/annual-pro" }}
    - 1 Project Site
  {{ /plan }}
{{ /pricing }}`,
			expected: []string{
				`<div class="pricing-wrapper" id="pricing-grid-`,
				`class="billing-toggle"`,
				`<input type="checkbox" onchange="togglePricingGrid(this, 'pricing-grid-`,
				`<div class="pricing-grid-poc">`,
				`<div class="price-card">`,
				`<h4>Personal</h4>`,
				`<div class="price-val" data-monthly="$19" data-annual="$15">$19</div>`,
				`<div class="price-period" data-monthly="per month" data-annual="billed annually">`,
				`<li>1 Project Site</li>`,
				`<a href="https://checkout.stripe.com/monthly-pro" class="pricing-btn" data-monthly-url="https://checkout.stripe.com/monthly-pro" data-annual-url="https://checkout.stripe.com/annual-pro">Get Started</a>`,
			},
		},
		{
			name: "Pricing Grid with plans (Static Mode)",
			input: `{{ pricing }}
  {{ plan title="Personal" price="0" period="Free forever" button="Get Started" url="https://personal.checkout" }}
    - 1 Project Site
  {{ /plan }}
{{ /pricing }}`,
			expected: []string{
				`<div class="pricing-wrapper" id="pricing-grid-`,
				`<div class="pricing-grid-poc">`,
				`<div class="price-card">`,
				`<h4>Personal</h4>`,
				`<div class="price-val">$0</div>`,
				`<div class="price-period">Free forever</div>`,
				`<li>1 Project Site</li>`,
				`<a href="https://personal.checkout" class="pricing-btn">Get Started</a>`,
			},
		},
		{
			name: "Social Proof Ribbon",
			input: `{{ social_ribbon }}
  {{ testimonial stars="5" avatar="../images/avatar_alex.png" author="Alex" handle="@alex_dev" }}
    Setting up Tamarind took less than two minutes.
  {{ /testimonial }}
{{ /social_ribbon }}`,
			expected: []string{
				`<div class="tamarind-social-ribbon-container">`,
				`<div class="tamarind-social-ribbon-track">`,
				`<div class="tamarind-social-ribbon-card">`,
				`<div class="profile">`,
				`<img class="avatar" src="../images/avatar_alex.png" alt="Alex">`,
				`<div class="profile-info">`,
				`<span class="author">Alex</span>`,
				`<span class="handle">@alex_dev</span>`,
				`<div class="stars">★★★★★</div>`,
				`<div class="quote">“Setting up Tamarind took less than two minutes.”</div>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := BuildPluginRegistry()
			got := processShortcodes(registry, tt.input, "")

			// For Static Mode, assert billing-toggle is omitted
			if tt.name == "Pricing Grid with plans (Static Mode)" {
				if strings.Contains(got, "billing-toggle") {
					t.Errorf("[%s] Expected output NOT to contain billing-toggle, but it did", tt.name)
				}
			}

			for _, exp := range tt.expected {
				if !strings.Contains(got, exp) {
					t.Errorf("[%s] Expected output to contain:\n%s\nGot:\n%s", tt.name, exp, got)
				}
			}
		})
	}
}
