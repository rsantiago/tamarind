.PHONY: build serve clean build-all serve-all

# Default target
all: build

# Build the website using the parser
build:
	@echo "Building website..."
	@cd parser && go build -o ../tamarind
	@./tamarind build
	@echo "Build complete!"

# Build all themes and generate the gallery index
build-all:
	@chmod +x build-all.sh
	@./build-all.sh

# Serve the website locally
serve:
	@echo "Starting local server..."
	@./tamarind serve

# Serve the gallery of all themes
serve-all: build-all
	@echo "Starting local server for all themes..."
	@./tamarind serve

# Initialize a new project
init:
	@echo "Initializing project..."
	@./tamarind init

# Clean up generated files
clean:
	@rm -rf website/* public-all/*
	@echo "Cleaned website and public-all directories."

# Stop any running server on port 8080
stop:
	@echo "Stopping server on port 8080..."
	@ -fuser -k 8080/tcp > /dev/null 2>&1 || true

# Full refresh: stop server, clean, build, and serve again
refresh: stop clean build serve

# Cross-compilation variables
BINARY_NAME=tamarind
RELEASE_DIR=releases
VERSION=$(shell git rev-parse --short HEAD 2>/dev/null || echo "v0.0.0")
LDFLAGS=-ldflags="-s -w -X github.com/rsantiago/tamarind/parser/internal/config.Version=$(VERSION)"

.PHONY: check-cgo build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-release

check-cgo:
	@echo "Checking for dynamic CGO imports..."
	@if grep -r -n 'import "C"' parser/ --include="*.go" --exclude-dir=vendor; then \
		echo "Error: Dynamic CGO import 'import \"C\"' detected in the codebase."; \
		exit 1; \
	else \
		echo "No dynamic CGO imports found. Codebase is pure Go."; \
	fi

build-linux-amd64: check-cgo
	@echo "Building for Linux AMD64..."
	@mkdir -p $(RELEASE_DIR)
	@cd parser && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ../$(RELEASE_DIR)/$(BINARY_NAME)-linux-amd64

build-linux-arm64: check-cgo
	@echo "Building for Linux ARM64..."
	@mkdir -p $(RELEASE_DIR)
	@cd parser && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o ../$(RELEASE_DIR)/$(BINARY_NAME)-linux-arm64

build-darwin-amd64: check-cgo
	@echo "Building for macOS AMD64..."
	@mkdir -p $(RELEASE_DIR)
	@cd parser && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o ../$(RELEASE_DIR)/$(BINARY_NAME)-darwin-amd64

build-darwin-arm64: check-cgo
	@echo "Building for macOS ARM64 (Apple Silicon)..."
	@mkdir -p $(RELEASE_DIR)
	@cd parser && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o ../$(RELEASE_DIR)/$(BINARY_NAME)-darwin-arm64

build-windows-amd64: check-cgo
	@echo "Building for Windows AMD64..."
	@mkdir -p $(RELEASE_DIR)
	@cd parser && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o ../$(RELEASE_DIR)/$(BINARY_NAME)-windows-amd64.exe

build-release: check-cgo build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64
	@echo "All release binaries compiled successfully."
	@echo "Generating release checksums..."
	@cd $(RELEASE_DIR) && (sha256sum * 2>/dev/null || shasum -a 256 * 2>/dev/null) > SHA256SUMS || true
	@echo "Release checksums written to $(RELEASE_DIR)/SHA256SUMS"



