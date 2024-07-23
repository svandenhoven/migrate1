package tasks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ModifyDataFiles(files []FileSource) {
	for _, source := range files {
		err := modifyDataFile(source.Destination)
		if err != nil {
			log.Printf("Error modifying file: %s: %v\n", source.Destination, err)
		}
	}
}

func modifyDataFile(filePath string) error {
	var modifiedData string

	// Read the YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Modify the content based on the file name
	switch filepath.Base(filePath) {
	case "navigation.yaml":
		modifiedData = modifyNavigationYAML(string(data))
	default:
		// No modifications needed for this file
		return nil
	}

	// Write the modified content back to the file
	err = os.WriteFile(filePath, []byte(modifiedData), 0200)
	if err != nil {
		return err
	}

	return nil
}

func modifyNavigationYAML(data string) string {
	modifiedData := replaceCommentBlock(data)
	modifiedData = updateYAML(modifiedData)
	modifiedData = fixIndentation(modifiedData)
	return modifiedData
}

var loadTemplateFile = func() (string, error) {
	templateData, err := os.ReadFile("cmd/gldocs/templates/navigation-template.txt")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(templateData)), nil
}

// Replace the comment at the top of navigation.yaml
func replaceCommentBlock(data string) string {
	templateContent, err := loadTemplateFile()
	if err != nil {
		log.Fatal(err)
	}

	// Split the YAML data into lines
	lines := strings.Split(data, "\n")

	// Find the index of the first non-comment line
	startIndex := 0
	for i, line := range lines {
		if !strings.HasPrefix(strings.TrimSpace(line), "#") {
			startIndex = i
			break
		}
	}

	// Replace the comment block with the template content
	modifiedData := templateContent + "\n" + strings.Join(lines[startIndex:], "\n")

	return modifiedData
}

// Update YAML by performing string replacements
func updateYAML(data string) string {
	lines := strings.Split(data, "\n")
	for i, line := range lines {
		// Replace keys
		line = strings.ReplaceAll(line, "sections:", "")
		line = strings.ReplaceAll(line, "section_title:", "title:")
		line = strings.ReplaceAll(line, "category_title:", "title:")
		line = strings.ReplaceAll(line, "doc_title:", "title:")
		line = strings.ReplaceAll(line, "category_url:", "url:")
		line = strings.ReplaceAll(line, "section_url:", "url:")
		line = strings.ReplaceAll(line, "doc_url:", "url:")
		line = strings.ReplaceAll(line, "section_categories:", "submenu:")
		line = strings.ReplaceAll(line, "docs:", "submenu:")

		// Drop the "ee" prefix from URLs
		line = strings.ReplaceAll(line, "'ee/", "'")

		// Modify URLs ending with a slash.
		// When using the "uglyURLs" setting, all URLs end in ".html"
		// See https://github.com/gohugoio/hugo/issues/4428
		if strings.Contains(line, "url:") {
			urlPattern := regexp.MustCompile(`url:\s*['"]?([^'"]+)['"]?`)
			matches := urlPattern.FindStringSubmatch(line)
			if len(matches) > 1 {
				url := matches[1]
				if strings.HasSuffix(url, "/") {
					newURL := strings.TrimSuffix(url, "/") + ".html"
					line = urlPattern.ReplaceAllString(line, fmt.Sprintf("url: '%s'", newURL))
				}
			}
		}

		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

// Fix indentation by removing two spaces from the beginning of each line
// This is needed because we've dropped the top-level "sections" key
func fixIndentation(data string) string {
	r := regexp.MustCompile(`(?m)^  `)
	return r.ReplaceAllString(data, "")
}
