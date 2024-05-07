package tasks

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/internal/workerpool"
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

// Utility methods for downloading remote files
type FileSource struct {
	URL         string
	Destination string
}

func FetchRemoteDataFiles(files []FileSource) {
	maxWorkers := 3
	jobsCh := make(chan FileSource)
	resultsCh := make(chan *workerpool.JobResult)

	enqueueJobsFn := func(jobsCh chan FileSource) {
		for _, source := range files {
			jobsCh <- source
		}
		close(jobsCh)
	}

	executeWorkFn := func(job FileSource) *workerpool.JobResult {
		return &workerpool.JobResult{Error: DownloadFile(job.URL, job.Destination)}
	}

	workerpool.EnqueueJobs[FileSource](maxWorkers, jobsCh, resultsCh, enqueueJobsFn, executeWorkFn)

	// Wait for all jobs to complete
	for i := 0; i < len(files); i++ {
		r := <-resultsCh
		if r.Error != nil {
			log.Printf("Error downloading file: %v\n", r.Error)
		}
	}
}

func DownloadFile(url, filepath string) error {
	// Fetch the remote file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the local file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to the local file
	_, err = io.Copy(out, resp.Body)
	return err
}
