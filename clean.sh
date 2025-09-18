#!/bin/bash

# Check for CI parameter
CI_MODE=false
if [[ "$1" == "CI=true" ]] || [[ "$CI" == "true" ]]; then
    CI_MODE=true
fi

# Dynamically find Go files not in .openapi-generator/FILES
files=()
for file in *.go; do
    if [ -f "$file" ]; then
        # Check if file is NOT in the FILES list
        if ! grep -qx "$file" .openapi-generator/FILES; then
            files+=("$file")
        fi
    fi
done

# Check if any unlisted files were found
if [ ${#files[@]} -eq 0 ]; then
    echo "All Go files are listed in .openapi-generator/FILES"
    exit 0
fi

echo "Go files not in .openapi-generator/FILES:"
echo "========================================="
for file in "${files[@]}"; do
    echo "  - $file"
done
echo ""
echo "Total: ${#files[@]} file(s)"
echo ""

# If CI mode, automatically delete. Otherwise, ask user
if [ "$CI_MODE" = true ]; then
    echo "CI mode enabled - automatically deleting files..."
    for file in "${files[@]}"; do
        if [ -f "$file" ]; then
            rm "$file"
            echo "  Deleted: $file"
        else
            echo "  Not found: $file"
        fi
    done
    echo "Done!"
else
    # Ask user if they want to delete these files
    read -p "Do you want to delete these files? (y/N): " -n 1 -r
    echo ""

    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "Deleting files..."
        for file in "${files[@]}"; do
            if [ -f "$file" ]; then
                rm "$file"
                echo "  Deleted: $file"
            else
                echo "  Not found: $file"
            fi
        done
        echo "Done!"
    else
        echo "No files were deleted."
    fi
fi
