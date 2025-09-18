#!/bin/bash

# Script to check for duplicate filenames (case-insensitive) in the parent directory and all subdirectories

echo "Checking for duplicate filenames (case-insensitive)..."
echo "========================================="

# Get the parent directory (one level up from the tests directory)
PARENT_DIR="$(dirname "$(cd "$(dirname "$0")" && pwd)")"
cd "$PARENT_DIR" || exit 1

echo "Scanning directory: $PARENT_DIR"
echo ""

# Find all files (excluding .git directory) and convert to lowercase for comparison
# Store results in associative arrays to track duplicates
declare -A lowercase_files
declare -A original_files
duplicate_found=false

# Find all files, excluding .git directory
while IFS= read -r -d '' file; do
    # Get just the filename without the path
    filename=$(basename "$file")

    # Convert filename to lowercase for comparison
    lowercase_filename=$(echo "$filename" | tr '[:upper:]' '[:lower:]')

    # Check if we've seen this lowercase filename before
    if [[ -n "${lowercase_files[$lowercase_filename]}" ]]; then
        # Found a duplicate
        if [ "$duplicate_found" = false ]; then
            echo "❌ DUPLICATES FOUND (case-insensitive):"
            echo "----------------------------------------"
            duplicate_found=true
        fi

        # Print both the original and the duplicate
        if [[ "${lowercase_files[$lowercase_filename]}" != "printed" ]]; then
            echo "  Duplicate set for: $lowercase_filename"
            echo "    - ${original_files[$lowercase_filename]}"
            lowercase_files[$lowercase_filename]="printed"
        fi
        echo "    - $file"
    else
        # First time seeing this lowercase filename
        lowercase_files[$lowercase_filename]="found"
        original_files[$lowercase_filename]="$file"
    fi
done < <(find . -type f -not -path "./.git/*" -not -path "./tests/check-duplicate-filenames.sh" -print0)

echo ""

# Print result summary
if [ "$duplicate_found" = true ]; then
    echo "========================================="
    echo "Result: FAILED - Duplicate filenames found (case-insensitive)"
    echo "========================================="
    exit 1
else
    # Count total files checked
    total_files=${#lowercase_files[@]}
    echo "✅ SUCCESS: No duplicate filenames found"
    echo "Total files checked: $total_files"
    echo "========================================="
    exit 0
fi
