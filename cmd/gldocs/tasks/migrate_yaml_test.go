package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceCommentBlock(t *testing.T) {
	// Mock the loadTemplateFile function for testing
	loadTemplateFile = func() (string, error) {
		return "This is the template content", nil
	}

	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name: "Replace comment block with template content",
			input: `# Comment line 1
# Comment line 2
sections:
  - section_title: Section 1
    section_url: 'ee/section1.html'
`,
			expectedOutput: `This is the template content
sections:
  - section_title: Section 1
    section_url: 'ee/section1.html'
`,
		},
		{
			name: "No comment block, return original data",
			input: `sections:
  - section_title: Section 1
    section_url: 'ee/section1.html'
`,
			expectedOutput: `This is the template content
sections:
  - section_title: Section 1
    section_url: 'ee/section1.html'
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replaceCommentBlock(tc.input)
			require.Equal(t, tc.expectedOutput, result)
		})
	}
}

func TestUpdateYAMLKeys(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name: "Update YAML keys and perform string replacements",
			input: `sections:
  - section_title: Section 1
    section_url: 'ee/section1.html'
    section_categories:
      - category_title: Category 1
        category_url: 'ee/category1.html'
`,
			expectedOutput: `
  - title: Section 1
    url: 'section1'
    submenu:
      - title: Category 1
        url: 'category1'
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := updateYAMLKeys(tc.input)
			require.Equal(t, tc.expectedOutput, result)
		})
	}
}

func TestFixIndentation(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name: "Remove two spaces from the beginning of each line",
			input: `  - title: Section 1
    url: 'section1'
    submenu:
      - title: Category 1
        url: 'category1'
`,
			expectedOutput: `- title: Section 1
  url: 'section1'
  submenu:
    - title: Category 1
      url: 'category1'
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fixIndentation(tc.input)
			require.Equal(t, tc.expectedOutput, result)
		})
	}
}
