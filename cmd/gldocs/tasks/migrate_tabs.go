package tasks

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	tabsRegex     = regexp.MustCompile(`(?m)^(\s*)::Tabs\s*$`)
	endTabsRegex  = regexp.MustCompile(`(?m)^(\s*)::EndTabs\s*$`)
	tabTitleRegex = regexp.MustCompile(`(?m)^(\s*):::TabTitle\s+(.+)\s*$`)
)

func MigrateTabs(files []string) {
	ProcessFiles(files, updateTabs)
}

func updateTabs(content []byte, filename string) (string, error) {
	str, codeBlocks := ExtractCodeBlocks(string(content))

	// Define the patterns and their corresponding replacements
	patterns := map[*regexp.Regexp]struct {
		replacement string
		handler     func(match []string, replacement string) (string, error)
	}{
		tabsRegex: {
			replacement: "{{< tabs >}}",
			handler:     defaultTabHandler,
		},
		endTabsRegex: {
			replacement: "{{< /tabs >}}",
			handler:     defaultTabHandler,
		},
		tabTitleRegex: {
			replacement: "{{< tab title=\"%s\" >}}",
			handler:     tabTitleHandler,
		},
	}

	// Iterate over the patterns and replace them
	for regex, config := range patterns {
		str = regex.ReplaceAllStringFunc(str, func(match string) string {
			parts := regex.FindStringSubmatch(match)
			result, err := config.handler(parts, config.replacement)
			if err != nil {
				fmt.Printf("Error in file %s: %v\n", filename, err)
				return match // Return the original match if an error occurred
			}
			return result
		})
	}

	// Check for unmatched patterns
	for _, line := range strings.Split(str, "\n") {
		if strings.HasPrefix(line, "::Tabs") || strings.HasPrefix(line, "::EndTabs") || strings.HasPrefix(line, ":::TabTitle") {
			fmt.Printf("Malformed tab structure in %s: %s\n", filename, line)
		}
	}

	// Reinsert the code blocks
	str = ReinsertCodeBlocks(str, codeBlocks)

	return str, nil
}

func defaultTabHandler(parts []string, replacement string) (string, error) {
	if len(parts) < 2 {
		return "", fmt.Errorf("insufficient parts")
	}
	indentation := parts[1]
	return indentation + replacement + "\n", nil
}

func tabTitleHandler(parts []string, replacement string) (string, error) {
	if len(parts) < 3 {
		return "", fmt.Errorf("insufficient parts")
	}
	indentation := parts[1]
	title := parts[2]
	return indentation + strings.TrimSpace(fmt.Sprintf(replacement, title)) + "\n", nil
}
