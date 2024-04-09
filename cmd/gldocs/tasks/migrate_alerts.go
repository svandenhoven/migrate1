package tasks

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	noteRegex       = regexp.MustCompile(`(?m)^(\s*)NOTE:\n([\s\S]*?\n)\n`)
	flagRegex       = regexp.MustCompile(`(?m)^(\s*)FLAG:\n([\s\S]*?\n)\n`)
	warningRegex    = regexp.MustCompile(`(?m)^(\s*)WARNING:\n([\s\S]*?\n)\n`)
	disclaimerRegex = regexp.MustCompile(`(?m)^(\s*)DISCLAIMER:\n([\s\S]*?\n)\n`)
)

func MigrateAlerts(files []string) {
	ProcessFiles(files, updateAlerts)
}

func updateAlerts(content []byte, filename string) (string, error) {
	str := string(content)
	invalidDisclaimers := false

	// Define the patterns and their corresponding replacements
	patterns := map[*regexp.Regexp]struct {
		replacement string
		handler     func(match []string, replacement string) string
	}{
		noteRegex: {
			replacement: "{{< alert type=\"note\" >}}",
			handler:     defaultAlertHandler,
		},
		flagRegex: {
			replacement: "{{< alert type=\"flag\" >}}",
			handler:     defaultAlertHandler,
		},
		warningRegex: {
			replacement: "{{< alert type=\"warning\" >}}",
			handler:     defaultAlertHandler,
		},
		disclaimerRegex: {
			replacement: "{{< alert type=\"disclaimer\" />}}",
			handler:     disclaimerAlertHandler(&invalidDisclaimers),
		},
	}

	// Iterate over the patterns and replace them
	for regex, config := range patterns {
		str = regex.ReplaceAllStringFunc(str, func(match string) string {
			parts := regex.FindStringSubmatch(match)
			return config.handler(parts, config.replacement)
		})
	}

	if invalidDisclaimers {
		return str, fmt.Errorf("Disclaimer alert(s) found with invalid inner content")
	}

	return str, nil
}

func defaultAlertHandler(parts []string, replacement string) string {
	indentation := parts[1]
	innerContent := parts[2]
	alertBlock := indentation + replacement + "\n\n" +
		innerContent +
		indentation + "{{< /alert >}}\n\n"
	return alertBlock
}

func disclaimerAlertHandler(invalidDisclaimers *bool) func([]string, string) string {
	return func(parts []string, replacement string) string {
		indentation := parts[1]
		innerContent := parts[2]

		// Check if the disclaimer contains valid content.
		// There are some files where people have used the disclaimer alert type with different text,
		// so we need to flag these to manually fix them.
		expectedStart := "This page contains information related to upcoming products, features, and functionality"
		if !strings.HasPrefix(strings.TrimSpace(innerContent), expectedStart) {
			*invalidDisclaimers = true
		}

		alertBlock := indentation + replacement + "\n\n"
		return alertBlock
	}
}
