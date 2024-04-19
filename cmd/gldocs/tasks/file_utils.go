package tasks

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/exp/slices"
)

func ProcessFiles(files []string, processFunc func(content []byte, filename string) (string, error)) {
	var wg sync.WaitGroup
	for _, inputFile := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			content, err := os.ReadFile(file)
			if err != nil {
				log.Printf("Error reading file %s: %v\n", file, err)
				return
			}
			updatedContent, err := processFunc(content, file)
			if err != nil {
				log.Printf("Error processing file %s: %v\n", file, err)
				return
			}
			err = os.WriteFile(file, []byte(updatedContent), 0644)
			if err != nil {
				log.Printf("Error writing to file %s: %v\n", file, err)
			}
		}(inputFile)
	}
	wg.Wait()
}

func FindMarkdownFiles(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			if !isIgnoredPath(path) {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func uniqueFilesOnly(files []string) []string {
	if len(files) == 0 {
		return files
	}
	slices.Sort(files)

	return slices.Compact(files)
}

func isIgnoredPath(path string) bool {
	for _, ignorePath := range IgnorePaths {
		if strings.HasPrefix(path, ignorePath) {
			return true
		}
	}
	return false
}
