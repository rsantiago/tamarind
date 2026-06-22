package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rsantiago/tamarind/parser/internal/config"
)

func CheckForUpdatesAsync() {
	if config.Version == "v0.0.0" {
		return // Development build, do not nag
	}
	
	go func() {
		// Wait to ensure all server boot logs have printed so this lands at the absolute bottom.
		time.Sleep(1 * time.Second)

		resp, err := http.Get("https://api.github.com/repos/rsantiago/tamarind/releases/latest")
		if err != nil {
			return
		}
		defer resp.Body.Close()

		var release struct {
			TagName string `json:"tag_name"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&release); err == nil {
			if release.TagName != "" && release.TagName != config.Version {
				// Dim Yellow text (\033[33;2m)
				fmt.Printf("\n\033[33;2mUpdate available (%s). Run 'tamarind update'.\033[0m\n", release.TagName)
			}
		}
	}()
}

func RunUpdate() error {
	fmt.Println("🚀 Checking for Tamarind updates...")

	// 1. Get latest release
	resp, err := http.Get("https://api.github.com/repos/rsantiago/tamarind/releases/latest")
	if err != nil {
		return fmt.Errorf("failed to check for updates: %w", err)
	}
	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return fmt.Errorf("failed to parse release info: %w", err)
	}

	if release.TagName == "" {
		return fmt.Errorf("could not determine latest release version")
	}

	if release.TagName == config.Version {
		fmt.Printf("✅ You are already running the latest version (%s)\n", config.Version)
		return nil
	}

	fmt.Printf("📦 Downloading %s...\n", release.TagName)

	// 2. Construct binary name
	binaryName := fmt.Sprintf("tamarind-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	downloadURL := fmt.Sprintf("https://github.com/rsantiago/tamarind/releases/download/%s/%s", release.TagName, binaryName)

	// 3. Download the binary
	downResp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download update: %w", err)
	}
	defer downResp.Body.Close()

	if downResp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d (URL: %s)", downResp.StatusCode, downloadURL)
	}

	// 4. Atomic Replace
	executable, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to locate current executable: %w", err)
	}
	executable, err = filepath.EvalSymlinks(executable)
	if err != nil {
		return err
	}

	tempFile := executable + ".tmp"
	oldFile := executable + ".old"

	// Create temp file
	out, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}

	_, err = io.Copy(out, downResp.Body)
	out.Close()
	if err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to write update data: %w", err)
	}

	// Windows requires moving the running executable out of the way first
	if err := os.Rename(executable, oldFile); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to backup current executable (are you running with permissions?): %w", err)
	}

	// Move new executable into place
	if err := os.Rename(tempFile, executable); err != nil {
		// Try to restore old executable
		os.Rename(oldFile, executable)
		return fmt.Errorf("failed to replace executable: %w", err)
	}

	// Cleanup old file asynchronously (may fail on Windows if still locked)
	go os.Remove(oldFile)

	fmt.Printf("✅ Tamarind successfully updated to %s!\n", release.TagName)
	return nil
}
