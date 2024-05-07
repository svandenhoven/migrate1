#!/bin/bash

# Updates Docs content for compatibility with Hugo.
# See doc/post-processing.md for more information.

# Ensure yq is available
if ! command -v yq >/dev/null 2>&1; then
  >&2 printf "${ERROR}ERROR: yq is required but not installed. Exiting.${END}\n"
  exit 1
fi

PRODUCTS_YAML="data/products.yaml"

# Iterate over each repository entry in the YAML file
for ENTRY in $(yq eval '.products | keys | .[]' "$PRODUCTS_YAML"); do
    CLONE_DIR=$(yq eval ".products[\"$ENTRY\"].clone_dir" "$PRODUCTS_YAML")
    DOCS_DIR=$(yq eval ".products[\"$ENTRY\"].docs_dir" "$PRODUCTS_YAML")
    DOCS_PATH="$CLONE_DIR"/"$DOCS_DIR"

    # Rename index files
    printf "Renaming index files in ${CLONE_DIR}...\n"
    find "$DOCS_PATH" -type f -name 'index.md' -exec sh -c 'mv "$1" "${1%/*}/_index.md"' _ {} \;

    # Run content migration scripts
    printf "Updating content in ${CLONE_DIR}...\n"
    go run cmd/gldocs/main.go migrate "$DOCS_PATH"
done

# Fetch YAML data files
printf "Fetching YAML data files...\n"
go run cmd/gldocs/main.go fetch

printf "INFO: Content updates complete!\n"
