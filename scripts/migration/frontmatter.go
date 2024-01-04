package main

/**
 * Temporary script for modifying Docs front matter.
 * See doc/post-process.md for details.
 */

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

// FrontMatter represents the YAML front matter structure
type FrontMatter struct {
	RedirectTo  string `yaml:"redirect_to,omitempty"`
	RemoveDate  string `yaml:"remove_date,omitempty"`
	Type        string `yaml:"type,omitempty"`
	Stage       string `yaml:"stage,omitempty"`
	Group       string `yaml:"group,omitempty"`
	Info        string `yaml:"info,omitempty"`
	Description string `yaml:"description,omitempty"`
	Title       string `yaml:"title,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run frontmatter.go file1.md [file2.md]")
	}

	// Start worker goroutines
	var wg sync.WaitGroup

	for _, inputFile := range uniqueFilesOnly(os.Args[1:]) {
		wg.Add(1)
		go processFile(&wg, inputFile)
	}

	// Wait for workers to finish
	wg.Wait()
}

func uniqueFilesOnly(files []string) []string {
	slices.Sort(files)

	return slices.Compact(files)
}

func processFile(wg *sync.WaitGroup, inputFile string) {
	defer wg.Done()

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Printf("Error reading file %s: %v\n", inputFile, err)
		return
	}

	updatedContent, err := updateContent(content, inputFile)
	if err != nil {
		log.Printf("Error processing file %s: %v\n", inputFile, err)
		return
	}

	// Write back to the original file
	err = os.WriteFile(inputFile, []byte(updatedContent), 0644)
	if err != nil {
		log.Printf("Error writing to file %s: %v\n", inputFile, err)
	}
}

func updateContent(content []byte, filename string) (string, error) {
	// Split the page front matter and body
	regex := regexp.MustCompile(`(?s)^---\n(.+?)\n---\n(.+)$`)
	matches := regex.FindSubmatch(content)

	if matches == nil {
		return "", fmt.Errorf("[WARN] No valid YAML front matter found")
	}

	frontMatter := string(matches[1])
	body := string(matches[2])

	// Parse the YAML front matter
	var frontMatterData FrontMatter
	err := yaml.Unmarshal([]byte(frontMatter), &frontMatterData)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Error parsing YAML front matter: %v", err)
	}

	// Modify file contents
	frontMatterData, body = updateTitles(frontMatterData, body, filename)

	// Convert content back to a string
	updatedFrontMatter, err := yaml.Marshal(frontMatterData)
	if err != nil {
		return "", fmt.Errorf("[ERROR] Error dumping YAML: %v\n", err)
	}

	updatedContent := fmt.Sprintf("---\n%s---\n\n%s\n", updatedFrontMatter, strings.TrimSpace(body))

	return updatedContent, nil
}

/**
 * Move page titles to front matter
 *
 * Removes the h1 and adds a `title` front matter attribute,
 * containing the original title value.
 *
 * @see https://gohugo.io/content-management/front-matter/
 */
func updateTitles(fm FrontMatter, body string, filename string) (FrontMatter, string) {
	if fm.RedirectTo != "" {
		return fm, body
	}

	regex := regexp.MustCompile(`(?m)^# (.*)`)
	match := regex.FindString(body)

	if len(match) > 0 {
		title := strings.TrimSpace(match[2:])
		body = strings.Replace(body, match, "", 1)
		fm.Title = strings.TrimSpace(title)
	} else {
		fmt.Printf("[WARN] No title found in %s\n", filename)
	}

	return fm, body
}
