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

