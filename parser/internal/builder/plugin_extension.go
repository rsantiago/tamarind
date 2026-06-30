// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package builder

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type ExtensionPlugin struct {
	name    string
	cmdPath string
	pattern *regexp.Regexp
}

func NewExtensionPlugin(name string, cmdPath string) *ExtensionPlugin {
	// Dynamically build the regex pattern for this shortcode
	// Example: {{ my_ext attribute="value" }} inner content {{ /my_ext }}
	patternStr := fmt.Sprintf(`(?s){{\s*%s(?:\s+([^}]+))?\s*}}(.*?){{\s*/%s\s*}}`, regexp.QuoteMeta(name), regexp.QuoteMeta(name))
	return &ExtensionPlugin{
		name:    name,
		cmdPath: cmdPath,
		pattern: regexp.MustCompile(patternStr),
	}
}

func (p *ExtensionPlugin) Name() string { return p.name }

func (p *ExtensionPlugin) Pattern() *regexp.Regexp { return p.pattern }

func (p *ExtensionPlugin) Process(match []string, sourceDir string) (string, error) {
	attributes := strings.TrimSpace(match[1])
	content := strings.TrimSpace(match[2])

	parts := strings.Fields(p.cmdPath)
	if len(parts) == 0 {
		return "", fmt.Errorf("empty command path for extension %s", p.name)
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	// Pass attributes via command line arguments if there are any
	if attributes != "" {
		cmd.Args = append(cmd.Args, attributes)
	}

	// Pass inner content via standard input
	cmd.Stdin = strings.NewReader(content)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("extension '%s' failed: %v, stderr: %s", p.name, err, stderr.String())
	}

	return out.String(), nil
}
