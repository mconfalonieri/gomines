#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# The version number needs to be passed
VERSION=$1

if [ -z "$VERSION" ]; then
    echo "Missing version parameter"
    exit 1
fi

# Strips 'v' only if it is at the beginning of the string
CLEAN_VERSION="${VERSION#v}"

# The template file
TEMPLATE_FILE="README.tmpl.md"

# The final output file
OUTPUT_FILE="README.md"

# Use sed to replace a placeholder string with the actual version.
# Note: You MUST ensure that the string '{{ .Version }}' 
# or a custom placeholder (like 'PROJECT_VERSION_PLACEHOLDER') exists 
# in your README.md.gotmpl file.

echo -e "Generating \033[36m$OUTPUT_FILE\033[0m with version: \033[33m$VERSION\033[0m"

# Using 'cp' to create a copy first, then 'sed' to modify it.
cp "$TEMPLATE_FILE" "$OUTPUT_FILE"

# Using sed to substitute variables
sed -i \
  -e "s/{{ \.Version }}/$VERSION/g" \
  -e "s/{{ \.CleanVersion }}/$CLEAN_VERSION/g" \
  "$OUTPUT_FILE"

echo -e "Successfully created \033[36m$OUTPUT_FILE\033[0m."
