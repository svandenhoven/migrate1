package tasks

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMarkdownFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create some test files
	files := []string{
		"file1.md",
		"file2.txt",
		"subdir/file3.md",
		"subdir/file4.txt",
	}
	for _, file := range files {
		filePath := filepath.Join(tempDir, file)
		dirPath := filepath.Dir(filePath)
		err := os.MkdirAll(dirPath, 0700)
		require.Nil(t, err)

		err = os.WriteFile(filePath, []byte(""), 0600)
		require.Nil(t, err)
	}

	markdownFiles, err := FindMarkdownFiles(tempDir)
	require.Nil(t, err)

	// Check the returned markdown files
	expectedFiles := []string{
		filepath.Join(tempDir, "file1.md"),
		filepath.Join(tempDir, "subdir/file3.md"),
	}
	require.Equal(t, len(expectedFiles), len(markdownFiles))

	for i, file := range markdownFiles {
		require.Equal(t, expectedFiles[i], file)
	}
}
