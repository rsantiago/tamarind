#!/bin/sh
set -e

# Tamarind Universal Installer
# Usage: curl -sSL https://usetamarind.com/install.sh | sh

echo "🥭 Installing Tamarind..."

# Detect OS
OS="$(uname -s)"
case "$OS" in
    Linux*)     OS_NAME="linux" ;;
    Darwin*)    OS_NAME="darwin" ;;
    CYGWIN*|MINGW32*|MSYS*|MINGW*) OS_NAME="windows" ;;
    *)          echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Detect Architecture
ARCH="$(uname -m)"
case "$ARCH" in
    x86_64)  ARCH_NAME="amd64" ;;
    amd64)   ARCH_NAME="amd64" ;;
    arm64)   ARCH_NAME="arm64" ;;
    aarch64) ARCH_NAME="arm64" ;;
    *)       echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Construct binary name
BINARY_NAME="tamarind-${OS_NAME}-${ARCH_NAME}"
if [ "$OS_NAME" = "windows" ]; then
    BINARY_NAME="${BINARY_NAME}.exe"
fi

# Define Download URL
DOWNLOAD_URL="https://github.com/rsantiago/tamarind/releases/download/latest/${BINARY_NAME}"

echo "⬇️  Downloading ${BINARY_NAME}..."
if command -v curl >/dev/null 2>&1; then
    curl -# -fLo tamarind "$DOWNLOAD_URL"
elif command -v wget >/dev/null 2>&1; then
    wget -qO tamarind "$DOWNLOAD_URL"
else
    echo "Error: curl or wget is required to download Tamarind."
    exit 1
fi

if [ "$OS_NAME" != "windows" ]; then
    chmod +x tamarind
fi

echo "✅ Tamarind installed successfully in the current directory."
echo ""
echo "🚀 Run './tamarind quickstart' to initialize and boot your server."
