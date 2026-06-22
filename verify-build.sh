#!/bin/bash
# Exit immediately on error
set -e

echo "=== Tamarind Build Parity Verifier ==="
echo "--------------------------------------"

# 1. Compile host binary
echo "Compiling host binary..."
cd parser && go build -o ../tamarind-host && cd ..

# 2. Compile target binary (Linux AMD64)
echo "Compiling target binary (Linux AMD64)..."
cd parser && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ../tamarind-target && cd ..

# 3. Setup temporary directories
echo "Creating temporary build directories..."
rm -rf tmp-host-build tmp-target-build website
mkdir -p tmp-host-build tmp-target-build

# Initialize sandbox and themes list
THEMES=("blue" "gram" "classic" "bento")

# 4. Compile with Host Binary
echo "Compiling themes using host binary..."
for theme in "${THEMES[@]}"; do
    echo "  Building theme: $theme..."
    rm -rf website
    ./tamarind-host build -theme "$theme" -source writer-sandbox -url "http://localhost:8080"
    mv website tmp-host-build/"$theme"
done

# 5. Compile with Target Binary
echo "Compiling themes using target binary..."
for theme in "${THEMES[@]}"; do
    echo "  Building theme: $theme..."
    rm -rf website
    ./tamarind-target build -theme "$theme" -source writer-sandbox -url "http://localhost:8080"
    mv website tmp-target-build/"$theme"
done

# 6. Perform Diff Comparison
echo "Comparing outputs byte-by-byte..."
if diff -r tmp-host-build tmp-target-build; then
    echo "--------------------------------------"
    echo "Success: Compiled outputs are identical. Parity verified."
    echo "--------------------------------------"
    # Cleanup
    rm -rf tmp-host-build tmp-target-build tamarind-host tamarind-target
    exit 0
else
    echo "--------------------------------------"
    echo "Error: Output mismatch detected between host and target binaries."
    echo "--------------------------------------"
    rm -rf tmp-host-build tmp-target-build tamarind-host tamarind-target
    exit 1
fi
