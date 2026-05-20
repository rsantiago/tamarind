package builder

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDataFiles(t *testing.T) {
	// 1. Setup Temp Source Dir
	tmpDir, err := os.MkdirTemp("", "tamarind-data-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	dataDir := filepath.Join(tmpDir, "data")
	os.Mkdir(dataDir, 0755)

	// 2. Create YAML file: authors.yaml
	yamlContent := `
alice:
  name: Alice Smith
  role: Editor
bob:
  name: Bob Jones
`
	os.WriteFile(filepath.Join(dataDir, "authors.yaml"), []byte(yamlContent), 0644)

	// 3. Create JSON file: links.json
	jsonContent := `
{
  "twitter": "https://x.com/tamarind",
  "github": "https://github.com/tamarind"
}
`
	os.WriteFile(filepath.Join(dataDir, "links.json"), []byte(jsonContent), 0644)

	// 4. Run LoadDataFiles
	data, err := LoadDataFiles(tmpDir)
	if err != nil {
		t.Fatalf("LoadDataFiles failed: %v", err)
	}

	// 5. Verify Results
	// Check "authors"
	authors, ok := data["authors"]
	if !ok {
		t.Fatal("Expected 'authors' key from authors.yaml")
	}

	authorsMap, ok := authors.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected Authors to be map, got %T", authors)
	}

	// Verify nested data
	alice := authorsMap["alice"].(map[string]interface{})
	if alice["name"] != "Alice Smith" {
		t.Errorf("Expected Alice's name to be 'Alice Smith', got %v", alice["name"])
	}

	// Check "links"
	links, ok := data["links"]
	if !ok {
		t.Fatal("Expected 'links' key from links.json")
	}
	linksMap := links.(map[string]interface{})
	if linksMap["twitter"] != "https://x.com/tamarind" {
		t.Errorf("Expected twitter link, got %v", linksMap["twitter"])
	}
}
