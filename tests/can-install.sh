#!/bin/bash

# Check if commit hash parameter is provided
if [ $# -eq 0 ]; then
    echo "Error: Commit hash parameter is required"
    echo "Usage: $0 <commit_hash>"
    exit 1
fi

COMMIT_HASH=$1

echo "Testing installation of opal-go at commit: $COMMIT_HASH"
echo "========================================="

# Create a temporary directory for testing
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR" || exit 1

echo "Working in temporary directory: $TEMP_DIR"

# Initialize a new Go module for testing
echo "Initializing test Go module..."
go mod init test-opal-install > /dev/null 2>&1

# Try to install the package at the specified commit
echo "Running: go get -u github.com/opalsecurity/opal-go@$COMMIT_HASH"
if go get -u "github.com/opalsecurity/opal-go@$COMMIT_HASH"; then
    echo "✅ SUCCESS: Package installed successfully at commit $COMMIT_HASH"

    # Verify the installation by checking go.mod
    echo ""
    echo "Verifying installation in go.mod:"
    grep "github.com/opalsecurity/opal-go" go.mod

    # Clean up
    cd - > /dev/null
    rm -rf "$TEMP_DIR"
    exit 0
else
    echo "❌ FAILED: Could not install package at commit $COMMIT_HASH"

    # Clean up
    cd - > /dev/null
    rm -rf "$TEMP_DIR"
    exit 1
fi
