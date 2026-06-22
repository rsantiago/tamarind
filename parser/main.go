// Copyright (c) 2026 Rodrigo Santiago. All rights reserved.
// Use of this source code is governed by the Business Source License 1.1
// that can be found in the LICENSE file in the root of this repository.

package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/rsantiago/tamarind/parser/internal/builder"
	"github.com/rsantiago/tamarind/parser/internal/config"
	"github.com/rsantiago/tamarind/parser/internal/server"
	"github.com/rsantiago/tamarind/parser/internal/updater"
	"github.com/rsantiago/tamarind/parser/internal/utils"
)

// --- Embedded Assets ---

//go:embed assets/*
var embeddedAssets embed.FS

// --- Configuration ---

const (
	DefaultStructureDir = "writer-sandbox"
	WebsiteDir          = "website"
)

// --- Main Logic ---

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Asynchronously check for updates (only fires on long-running commands like 'serve')
	updater.CheckForUpdatesAsync()

	switch os.Args[1] {
	case "init":
		initCmd := flag.NewFlagSet("init", flag.ExitOnError)
		reset := initCmd.Bool("reset", false, "Reset the project structure (overwrites existing files)")
		initCmd.Parse(os.Args[2:])
		if err := runInit(*reset, false); err != nil {
			log.Fatalf("Init failed: %v", err)
		}

	case "config":
		if err := config.CreateDefaultConfig(); err != nil {
			log.Fatalf("Config failed: %v", err)
		}

	case "build":
		buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
		theme := buildCmd.String("theme", "", "Theme to use (required)") // Default is empty to check for requirement
		sourcePath := buildCmd.String("source", DefaultStructureDir, "Path to the source directory")
		baseURL := buildCmd.String("url", "http://localhost:8080", "Base URL for the website")
		drafts := buildCmd.Bool("drafts", false, "Include draft content in build")
		buildCmd.Parse(os.Args[2:])

		if *theme == "" {
			fmt.Println("Error: You must specify a theme to build the website.")
			fmt.Println("Usage: tamarind build -theme <name> [-drafts]")
			runThemes() // List themes
			os.Exit(1)
		}

		if err := runBuild(*theme, *sourcePath, *baseURL, *drafts, false); err != nil {
			log.Fatalf("Build failed: %v", err)
		}

	case "serve":
		serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
		port := serveCmd.String("port", "8080", "Port to serve on")
		watch := serveCmd.Bool("watch", false, "Enable live reloading and file watching")
		theme := serveCmd.String("theme", "", "Theme to use (required if watch is enabled)")
		sourcePath := serveCmd.String("source", DefaultStructureDir, "Path to the source directory")
		displayDrafts := serveCmd.Bool("drafts", false, "Include draft content in build (only if watch is enabled)")

		serveCmd.Parse(os.Args[2:])

		if *watch {
			if *theme == "" {
				fmt.Println("Error: You must specify a theme to enable live reloading.")
				fmt.Println("Usage: tamarind serve -watch -theme <name> [-source <dir>] [-drafts]")
				runThemes()
				os.Exit(1)
			}
		}

		if err := runServe(*port, *watch, *theme, *sourcePath, *displayDrafts); err != nil {
			log.Fatalf("Serve failed: %v", err)
		}

	case "themes":
		if err := runThemes(); err != nil {
			log.Fatalf("Themes list failed: %v", err)
		}

	case "quickstart":
		quickstartCmd := flag.NewFlagSet("quickstart", flag.ExitOnError)
		theme := quickstartCmd.String("theme", "gram", "Theme to use (default: gram)")
		port := quickstartCmd.String("port", "8080", "Port to serve on")
		quickstartCmd.Parse(os.Args[2:])

		if err := runQuickstart(*theme, *port); err != nil {
			log.Fatalf("Quickstart failed: %v", err)
		}

	case "update":
		if err := updater.RunUpdate(); err != nil {
			log.Fatalf("Update failed: %v", err)
		}

	case "version":
		fmt.Println("Tamarind Static Site Generator " + config.Version)

	case "doctor":
		if err := runDoctor(); err != nil {
			log.Fatalf("Doctor check failed: %v", err)
		}

	default:
		printUsage()
		os.Exit(1)
	}
}

func runQuickstart(theme string, port string) error {
	fmt.Println("🚀 Launching Tamarind Quickstart...")
	fmt.Println("--------------------------------")

	// 1. Initialize Project (Force Reset/Overwrite for demo)
	if err := runInit(true, true); err != nil {
		return fmt.Errorf("init step failed: %w", err)
	}

	// 2. Generate Config (so they have it ready)
	if err := config.CreateDefaultConfig(); err != nil {
		return fmt.Errorf("config step failed: %w", err)
	}

	// 3. Serve with Watcher
	fmt.Println("--------------------------------")
	fmt.Printf("✨ Quickstart Complete! Serving on http://localhost:%s with Live Reloading (Theme: %s)\n", port, theme)
	return runServe(port, true, theme, DefaultStructureDir, false)
}

func runBuild(theme, sourcePath, baseURL string, includeDrafts bool, liveReload bool) error {
	// Load optional config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Warning: Could not load config: %v", err)
	}
	var themeConfig map[string]string
	if cfg != nil {
		themeConfig = cfg.ThemeConfig
	}

	// 1. Create a Master Temp Directory for all templates (theme + shared)
	masterTmpDir, err := os.MkdirTemp("", "tamarind-build-*")
	if err != nil {
		return fmt.Errorf("failed to create master temp dir: %w", err)
	}
	defer os.RemoveAll(masterTmpDir) // Cleanup after build

	// 2. Extract Shared Templates
	sharedPath := filepath.Join("assets", "templates", "shared")
	// Shared might not exist in older versions, checking first
	if _, err := embeddedAssets.ReadDir(sharedPath); err == nil {
		sharedDest := filepath.Join(masterTmpDir, "shared")
		if err := os.MkdirAll(sharedDest, 0755); err != nil {
			return fmt.Errorf("failed to create shared dir: %w", err)
		}
		if err := utils.ExtractDir(embeddedAssets, sharedPath, sharedDest, true); err != nil {
			return fmt.Errorf("failed to extract shared templates: %w", err)
		}
	}

	// 3. Extract Theme Templates
	themeTmpDir := filepath.Join(masterTmpDir, theme)

	assetPath := filepath.Join("assets", "templates", theme)
	// Check if theme exists in embedded assets
	if _, err := embeddedAssets.ReadDir(assetPath); err != nil {
		fmt.Printf("Error: Theme '%s' not found.\n", theme)
		runThemes()
		return fmt.Errorf("theme not found")
	}

	if err := os.MkdirAll(themeTmpDir, 0755); err != nil {
		return fmt.Errorf("failed to create theme dir: %w", err)
	}

	if err := utils.ExtractDir(embeddedAssets, assetPath, themeTmpDir, true); err != nil {
		return fmt.Errorf("failed to extract theme '%s': %w", theme, err)
	}

	// 4. Build using the extracted theme directory
	// builder.go will look for ../shared relative to this themeTmpDir
	return builder.Build(sourcePath, themeTmpDir, WebsiteDir, baseURL, themeConfig, includeDrafts, liveReload)
}

func printUsage() {
	fmt.Println("Usage: tamarind <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  quickstart  Magic command: Init, Build (Gram), and Serve instantly")
	fmt.Println("  init        Initialize a new project structure")
	fmt.Println("  config      Create a configuration file for customization")
	fmt.Println("  build       Build the website (requires -theme, optional -drafts)")
	fmt.Println("  serve       Serve the website locally (optional -watch -theme <name>)")
	fmt.Println("  themes      List available themes")
	fmt.Println("  update      Update Tamarind to the latest version")
	fmt.Println("  version     Display Tamarind version")
	fmt.Println("  doctor      Inspect environment diagnostics")
	fmt.Println("\nRun 'tamarind <command> --help' for more information.")
}

func runThemes() error {
	entries, err := embeddedAssets.ReadDir("assets/templates")
	if err != nil {
		return fmt.Errorf("failed to read embedded templates: %w", err)
	}

	fmt.Println("Available Themes:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("  - " + entry.Name())
		}
	}
	return nil
}

func runServe(port string, watch bool, theme string, sourcePath string, displayDrafts bool) error {
	if watch {
		if theme == "" {
			return fmt.Errorf("theme is required for live reloading")
		}

		// Initial Build
		fmt.Println("Performing initial build...")
		if err := runBuild(theme, sourcePath, "http://localhost:"+port, displayDrafts, true); err != nil {
			return fmt.Errorf("build failed: %w", err)
		}

		// Start Watcher in Goroutine
		go func() {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}
			defer watcher.Close()

			// Add paths to watch
			// Watch source files
			if err := watcher.Add(sourcePath); err != nil {
				log.Printf("Warning: failed to watch source path: %v", err)
			}

			// Recursively watch subdirectories
			filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return watcher.Add(path)
				}
				return nil
			})

			log.Println("Watching for file changes...")
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Remove == fsnotify.Remove {
						log.Printf("File modified: %s. Rebuilding...", event.Name)

						// Debounce
						time.Sleep(100 * time.Millisecond)

						if err := runBuild(theme, sourcePath, "http://localhost:"+port, displayDrafts, true); err != nil {
							log.Printf("Rebuild failed: %v", err)
						} else {
							log.Println("Rebuild complete. Reloading browsers...")
							server.NotifyReload()
						}
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Println("error:", err)
				}
			}
		}()

		return server.Start(port, WebsiteDir, true)
	} else {
		return server.Start(port, WebsiteDir, false)
	}
}

// --- Init Command ---

func runInit(reset bool, force bool) error {
	log.Println("Initializing new project... (" + config.Version + ")")

	// Target directory (Current Directory)
	structurePath := DefaultStructureDir

	// Check collision
	if _, err := os.Stat(structurePath); err == nil {
		if !reset && !force {
			return fmt.Errorf("structure directory '%s' already exists. Use --reset to overwrite", structurePath)
		}

		if !force {
			// Prompt for confirmation
			fmt.Printf("\nWARNING: The directory '%s' exists.\n", structurePath)
			fmt.Print("Do you want to reset it? Current contents will be backed up. (Y/N): ")

			var response string
			fmt.Scanln(&response) // Wait for user input

			response = strings.ToUpper(strings.TrimSpace(response))
			if response != "Y" && response != "YES" {
				fmt.Println("Operation cancelled.")
				return nil
			}
		}

		// Backup existing directory
		timestamp := time.Now().Format("20060102-150405")
		backupPath := fmt.Sprintf("%s-backup-%s", structurePath, timestamp)

		log.Printf("Backing up existing directory to: %s", backupPath)
		if err := os.Rename(structurePath, backupPath); err != nil {
			return fmt.Errorf("failed to backup existing directory: %w", err)
		}
	}

	// Extract Structure (to writer-sandbox in CWD)
	if err := utils.ExtractDir(embeddedAssets, "assets/structure", structurePath, true); err != nil {
		return err
	}

	log.Printf("Project initialized successfully in '%s'!\n", structurePath)
	log.Println("To build the site, run:")
	log.Println("  ./tamarind build -theme blue")
	return nil
}

func runDoctor() error {
	fmt.Println("Tamarind Diagnostic Tool")
	fmt.Println("------------------------")
	fmt.Printf("OS:             %s\n", runtime.GOOS)
	fmt.Printf("Architecture:   %s\n", runtime.GOARCH)
	fmt.Printf("Path Separator: %q\n", string(filepath.Separator))

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("CWD:            Error: %v\n", err)
	} else {
		fmt.Printf("CWD:            %s\n", cwd)
		tempFile := filepath.Join(cwd, ".tamarind-doctor-test")
		if err := os.WriteFile(tempFile, []byte("test"), 0644); err != nil {
			fmt.Printf("Write Access:   Error: %v\n", err)
		} else {
			fmt.Println("Write Access:   OK")
			_ = os.Remove(tempFile)
		}
	}

	fmt.Println("\nEnvironment Variables:")
	cgo := os.Getenv("CGO_ENABLED")
	if cgo == "" {
		cgo = "Not Set"
	}
	fmt.Printf("  CGO_ENABLED:  %s\n", cgo)

	goos := os.Getenv("GOOS")
	if goos == "" {
		goos = "Not Set"
	}
	fmt.Printf("  GOOS:         %s\n", goos)

	goarch := os.Getenv("GOARCH")
	if goarch == "" {
		goarch = "Not Set"
	}
	fmt.Printf("  GOARCH:       %s\n", goarch)

	return nil
}

