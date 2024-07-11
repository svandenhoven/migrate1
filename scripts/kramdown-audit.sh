#!/usr/bin/env bash

# Report on use of Kramdown formatting.
# Pages that match here will need to be manually reviewed and updated.
# See https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/-/issues/16
# for recommended changes.

COLOR_GREEN="\e[32m"
COLOR_YELLOW="\e[33m"
COLOR_RESET="\e[39m"
CMD="grep -r -I -H -n -F"
FORMAT_CMD="sed -E 's/^([^:]+):([0-9]+):.*/\1:\2/'"
DOCS_DIRS=("../gitlab/doc" "../gitlab-operator/doc" "../gitlab-runner/docs" "../charts-gitlab/doc" "../omnibus-gitlab/doc")

# Runs the grep command against each docs directory.
# Args:
# $1 Kramdown pattern to search for
# $2 Human-readable label to print with the output
# $3 Variable name to store the number of matches
check_pattern() {
    local pattern="$1"
    local message="$2"
    local count_var="$3"
    printf "${COLOR_GREEN}Checking for $message...${COLOR_RESET}\n"
    $CMD "$pattern" "${DOCS_DIRS[@]}" | eval "$FORMAT_CMD"
    local count
    count=$($CMD "$pattern" "${DOCS_DIRS[@]}" | wc -l | awk '{print $1}')
    [ "$count" -gt 0 ]
    eval "$count_var=$count"
}

# Kramdown syntax patterns
# See https://kramdown.gettalong.org/quickref.html
check_pattern "{::options" "options tags" "OPTIONS_COUNT"
check_pattern "{: ." "attribute tags" "CLASSES_COUNT"
check_pattern "{::nomarkdown" "nomarkdown tags" "NO_MD_COUNT"
check_pattern "{::comment}" "comment tags" "COMMENT_COUNT"
check_pattern "{: title" "title tags" "TITLE_COUNT"

# Print the summary and exit
echo
printf "${COLOR_YELLOW}Summary:${COLOR_RESET}\n"
echo "- Found $OPTIONS_COUNT options"
echo "- Found $CLASSES_COUNT CSS classes"
echo "- Found $NO_MD_COUNT nomarkdown tags"
echo "- Found $COMMENT_COUNT comment tags"
echo "- Found $TITLE_COUNT title tags"
