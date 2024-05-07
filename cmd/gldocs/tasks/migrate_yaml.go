package tasks

import (
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
	modifiedData = updateYAMLKeys(modifiedData)
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

// Update YAML keys by performing string replacements
func updateYAMLKeys(data string) string {
	modifiedData := strings.ReplaceAll(data, "sections:", "")
	modifiedData = strings.ReplaceAll(modifiedData, "section_title:", "title:")
	modifiedData = strings.ReplaceAll(modifiedData, "category_title:", "title:")
	modifiedData = strings.ReplaceAll(modifiedData, "doc_title:", "title:")
	modifiedData = strings.ReplaceAll(modifiedData, "category_url:", "url:")
	modifiedData = strings.ReplaceAll(modifiedData, "section_url:", "url:")
	modifiedData = strings.ReplaceAll(modifiedData, "doc_url:", "url:")
	modifiedData = strings.ReplaceAll(modifiedData, "section_categories:", "submenu:")
	modifiedData = strings.ReplaceAll(modifiedData, "docs:", "submenu:")

	// Drop the "ee" prefix and ".html" suffix from URLs
	modifiedData = strings.ReplaceAll(modifiedData, "'ee/", "'")
	modifiedData = strings.ReplaceAll(modifiedData, ".html", "")

	return modifiedData
}

// Fix indentation by removing two spaces from the beginning of each line
// This is needed because we've dropped the top-level "sections" key
func fixIndentation(data string) string {
	r := regexp.MustCompile(`(?m)^  `)
	return r.ReplaceAllString(data, "")
}
