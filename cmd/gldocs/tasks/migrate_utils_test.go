package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractCodeBlocks(t *testing.T) {
	testCases := []struct {
		name               string
		input              string
		expectedContent    string
		expectedCodeBlocks []string
	}{
		{
			name: "Single code block",
			input: `
# Heading

Some text.

` + "```go" + `
func main() {
	fmt.Println("Hello, World!")
}
` + "```" + `

More text.
`,
			expectedContent: `
# Heading

Some text.

{CODE_BLOCK}

More text.
`,
			expectedCodeBlocks: []string{
				"```go\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n```",
			},
		},
		{
			name: "Multiple code blocks",
			input: `
# Heading

` + "```go" + `
func foo() {
	// Code block 1
}
` + "```" + `

Some text.

` + "```python" + `
def bar():
	# Code block 2
` + "```" + `

More text.
`,
			expectedContent: `
# Heading

{CODE_BLOCK}

Some text.

{CODE_BLOCK}

More text.
`,
			expectedCodeBlocks: []string{
				"```go\nfunc foo() {\n\t// Code block 1\n}\n```",
				"```python\ndef bar():\n\t# Code block 2\n```",
			},
		},
		{
			name: "No code blocks",
			input: `
# Heading

Just some regular text without any code blocks.
`,
			expectedContent: `
# Heading

Just some regular text without any code blocks.
`,
			expectedCodeBlocks: []string(nil),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			content, codeBlocks := ExtractCodeBlocks(tc.input)
			require.Equal(t, tc.expectedContent, content)
			require.Equal(t, tc.expectedCodeBlocks, codeBlocks)
		})
	}
}

func TestReinsertCodeBlocks(t *testing.T) {
	testCases := []struct {
		name       string
		content    string
		codeBlocks []string
		expected   string
	}{
		{
			name: "Single code block",
			content: `
# Heading

Some text.

{CODE_BLOCK}

More text.
`,
			codeBlocks: []string{
				"```go\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n```",
			},
			expected: `
# Heading

Some text.

` + "```go" + `
func main() {
	fmt.Println("Hello, World!")
}
` + "```" + `

More text.
`,
		},
		{
			name: "Multiple code blocks",
			content: `
# Heading

{CODE_BLOCK}

Some text.

{CODE_BLOCK}

More text.
`,
			codeBlocks: []string{
				"```go\nfunc foo() {\n\t// Code block 1\n}\n```",
				"```python\ndef bar():\n\t# Code block 2\n```",
			},
			expected: `
# Heading

` + "```go" + `
func foo() {
	// Code block 1
}
` + "```" + `

Some text.

` + "```python" + `
def bar():
	# Code block 2
` + "```" + `

More text.
`,
		},
		{
			name: "No code blocks",
			content: `
# Heading

Just some regular text without any code blocks.
`,
			codeBlocks: []string{},
			expected: `
# Heading

Just some regular text without any code blocks.
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReinsertCodeBlocks(tc.content, tc.codeBlocks)
			require.Equal(t, tc.expected, result)
		})
	}
}
