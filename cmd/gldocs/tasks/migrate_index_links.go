package tasks

import (
	"regexp"
	"strings"
)

func MigrateIndexLinks(files []string) {
	ProcessFiles(files, updateIndexLinks)
}

var indexLinkRegex = regexp.MustCompile(`\[([^\]]+)\]\(([^)]*index\.md[^)]*)\)`)

func updateIndexLinks(content []byte, filename string) (string, error) {
	replaceFunc := func(match []byte) []byte {
		// Extract the link text and URL from the matched substrings
		submatch := indexLinkRegex.FindSubmatch(match)
		if len(submatch) != 3 {
			return match // If submatches are not found, return the original without modification
		}

		linkText := submatch[1]
		linkURL := submatch[2]

		// If the link URL is absolute, return the original without modification
		if strings.HasPrefix(string(linkURL), "http://") || strings.HasPrefix(string(linkURL), "https://") {
			return match
		}

		// Replace "index.md" with "_index.md" in the link URL
		newLinkURL := regexp.MustCompile(`index\.md`).ReplaceAll(linkURL, []byte("_index.md"))

		// Return the replacement markdown link
		return []byte("[" + string(linkText) + "](" + string(newLinkURL) + ")")
	}

	// Perform the replacement
	output := indexLinkRegex.ReplaceAllFunc(content, replaceFunc)

	return string(output), nil
}
