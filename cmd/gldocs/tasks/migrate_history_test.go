package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateHistory(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Single history item",
			input: `
# A heading

> - Introduced in abc

Some text after the version history.
`,
			expected: `
# A heading

{{< history >}}

- Introduced in abc

{{< /history >}}

Some text after the version history.
`,
		},
		{
			name: "Multiple history items",
			input: `
## Another heading

> - Introduced in abc
> - Deprecated in xyz
> - [Removed](https://example.com) in abcdef
`,
			expected: `
## Another heading

{{< history >}}

- Introduced in abc
- Deprecated in xyz
- [Removed](https://example.com) in abcdef

{{< /history >}}
`,
		},
		{
			name: "A non-history blockquote",
			input: `
> - This is a regular list in a blockquote, without the version-related keywords
`,
			expected: `
> - This is a regular list in a blockquote, without the version-related keywords
`,
		},
		{
			name: "Multiple history sections",
			input: `
## First heading

> - Introduced in 15.5
> - Deprecated in 17.0

Some text here.

## Second heading

> - Introduced in 16.0
> - Enabled in 16.2

Some other text down here.
`,
			expected: `
## First heading

{{< history >}}

- Introduced in 15.5
- Deprecated in 17.0

{{< /history >}}

Some text here.

## Second heading

{{< history >}}

- Introduced in 16.0
- Enabled in 16.2

{{< /history >}}

Some other text down here.
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, _ := updateHistory([]byte(tc.input), "")
			require.Equal(t, tc.expected, output)
		})
	}
}
