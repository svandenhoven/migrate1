package tasks

import (
	"fmt"
	"regexp"
	"strings"
)

func MigrateIcons(files []string) {
	ProcessFiles(files, updateIcons)
}

var iconRegex = regexp.MustCompile(`\*\*\{([^}]+)\}\*\*`)

func updateIcons(content []byte, filename string) (string, error) {
	replaceFunc := func(match []byte) []byte {
		// Extract the matched string inside the brackets
		innerString := iconRegex.FindSubmatch(match)[1]

		// Split the inner string by comma and trim whitespace.
		// This drops extra attributes like size and CSS class,
		// which are no longer supported.
		parts := strings.Split(string(innerString), ",")
		iconName := strings.TrimSpace(parts[0])

		// Return the replacement string
		return []byte(fmt.Sprintf(`{{< icon name="%s" >}}`, iconName))
	}

	// Perform the replacement
	output := iconRegex.ReplaceAllFunc(content, replaceFunc)

	return string(output), nil
}
