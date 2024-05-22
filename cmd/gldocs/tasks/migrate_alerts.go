package tasks

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	noteRegex       = regexp.MustCompile(`(?m)^(\s*)NOTE:\n([\s\S]*?\n)\n`)
	flagRegex       = regexp.MustCompile(`(?m)^(\s*)FLAG:\n([\s\S]*?\n)\n`)
	warningRegex    = regexp.MustCompile(`(?m)^(\s*)WARNING:\n([\s\S]*?\n)\n`)
	disclaimerRegex = regexp.MustCompile(`(?m)^(\s*)DISCLAIMER:\n([\s\S]*?\n)\n`)
	detailsRegex    = regexp.MustCompile(`(?m)^(\s*)DETAILS:\n([\s\S]*?\n)\n`)
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
		handler     func(match []string, replacement string) (string, error)
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
		detailsRegex: {
			replacement: "{{< details >}}",
			handler:     detailsAlertHandler,
		},
	}

	// Iterate over the patterns and replace them
	for regex, config := range patterns {
		str = regex.ReplaceAllStringFunc(str, func(match string) string {
			parts := regex.FindStringSubmatch(match)
			result, err := config.handler(parts, config.replacement)
			if err != nil {
				log.Printf("Warning: Unable to migrate alert in %s: %v", filename, err)
				return match // Return the original match if an error occurred
			}
			return result
		})
	}

	if invalidDisclaimers {
		return str, fmt.Errorf("Disclaimer alert(s) found with invalid inner content")
	}

	return str, nil
}

func defaultAlertHandler(parts []string, replacement string) (string, error) {
	if err := checkParts(parts); err != nil {
		return "", err
	}

	indentation := parts[1]
	innerContent := parts[2]

	alertBlock := indentation + replacement + "\n\n" +
		innerContent +
		indentation + "{{< /alert >}}\n\n"
	return alertBlock, nil
}

func disclaimerAlertHandler(invalidDisclaimers *bool) func([]string, string) (string, error) {
	return func(parts []string, replacement string) (string, error) {
		if err := checkParts(parts); err != nil {
			return "", err
		}

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
		return alertBlock, nil
	}
}

func detailsAlertHandler(parts []string, replacement string) (string, error) {
	if err := checkParts(parts); err != nil {
		return "", err
	}

	indentation := parts[1]
	innerContent := parts[2]

	// Remove bold tags from innerContent
	innerContent = strings.ReplaceAll(innerContent, "**", "")

	// Add bullets to each line of innerContent
	lines := strings.Split(innerContent, "\n")
	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			lines[i] = "- " + trimmedLine
		}
	}
	innerContent = strings.Join(lines, "\n")

	availabilityBlock := indentation + replacement + "\n\n" +
		innerContent +
		indentation + "{{< /details >}}\n\n"
	return availabilityBlock, nil
}

// Validate the submatches captured by the regexes in the alert handler functions.
// This ensures that the expected submatches (indentation and inner content) are present before accessing them.
func checkParts(parts []string) error {
	if len(parts) < 3 {
		return fmt.Errorf("Invalid number of parts: expected at least 3, got %d", len(parts))
	}
	return nil
}
