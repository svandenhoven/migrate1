package tasks

import (
	"net/http"
	"net/http/httptest"
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

		err = os.WriteFile(filePath, []byte(""), 0400)
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

func TestIsIgnoredPath(t *testing.T) {
	testCases := []struct {
		path     string
		expected bool
	}{
		{"../gitlab-operator/doc/adr/file.md", true},
		{"../gitlab/drawers/file.md", true},
		{"../gitlab/file.md", false},
		{"../other/file.md", false},
	}

	for _, tc := range testCases {
		result := isIgnoredPath(tc.path)
		require.Equal(t, result, tc.expected)
	}
}

func TestFetchRemoteDataFiles(t *testing.T) {
	testServer, tempDir, err := createTestServer(t)
	require.NoError(t, err)
	defer testServer.Close()

	// Define the test cases
	testCases := []struct {
		name  string
		files []FileSource
	}{
		{
			name: "Single file",
			files: []FileSource{
				{
					URL:         testServer.URL + "/file1.txt",
					Destination: filepath.Join(tempDir, "file1.txt"),
				},
			},
		},
		{
			name: "Multiple files",
			files: []FileSource{
				{
					URL:         testServer.URL + "/file1.txt",
					Destination: filepath.Join(tempDir, "file1.txt"),
				},
				{
					URL:         testServer.URL + "/file2.txt",
					Destination: filepath.Join(tempDir, "file2.txt"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			FetchRemoteDataFiles(tc.files)

			// Check if the files were downloaded successfully
			for _, file := range tc.files {
				_, err := os.Stat(file.Destination)
				require.NoError(t, err)
			}
		})
	}
}

func TestDownloadFile(t *testing.T) {
	testServer, tempDir, err := createTestServer(t)
	require.NoError(t, err)
	defer testServer.Close()

	// Define the test cases
	testCases := []struct {
		name        string
		url         string
		destination string
		expectError bool
	}{
		{
			name:        "Valid file",
			url:         testServer.URL + "/file.txt",
			destination: filepath.Join(tempDir, "file.txt"),
			expectError: false,
		},
		{
			name:        "Invalid URL",
			url:         "invalid-url",
			destination: filepath.Join(tempDir, "invalid.txt"),
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := DownloadFile(tc.url, tc.destination)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				_, err := os.Stat(tc.destination)
				require.NoError(t, err)
			}
		})
	}
}

func createTestServer(t *testing.T) (*httptest.Server, string, error) {
	t.Helper()

	tempDir := t.TempDir()

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("test content"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	return testServer, tempDir, nil
}
