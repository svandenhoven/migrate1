package tasks

import (
	"fmt"
	"regexp"
	"strings"
)

func MigrateHistory(files []string) {
	ProcessFiles(files, updateHistory)
}

// Match blockquotes that contain list items and specific keywords
// See https://docs.gitlab.com/ee/development/documentation/versions.html#add-a-history-item
var blockquoteListRegex = regexp.MustCompile(`(?mi)^>\s*([-*])\s*(.*(introduced|added|enabled|deprecated|changed|moved|recommended|removed|renamed).*)$`)

func updateHistory(content []byte, filename string) (string, error) {
	lines := strings.Split(string(content), "\n")
	var result string

	for i := 0; i < len(lines); i++ {
		match := blockquoteListRegex.FindStringSubmatch(lines[i])
		if match != nil {
			result += "{{< history >}}\n\n"
			result += fmt.Sprintf("%s %s", match[1], match[2])

			// Find consecutive matches; these are grouped within a single history shortcode
			for j := i + 1; j < len(lines); j++ {
				nextMatch := blockquoteListRegex.FindStringSubmatch(lines[j])
				if nextMatch != nil {
					result += "\n"
					result += fmt.Sprintf("%s %s", nextMatch[1], nextMatch[2])
					i = j
				} else {
					break
				}
			}

			result += "\n\n{{< /history >}}\n"
		} else {
			result += lines[i]
			if i < len(lines)-1 {
				result += "\n"
			}
		}
	}

	return result, nil
}
