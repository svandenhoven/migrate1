package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateIndexLinks(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Basic link",
			input:          "This is an [index page](index.md) link",
			expectedOutput: "This is an [index page](_index.md) link",
		},
		{
			name:           "Link with a longer path",
			input:          "This is an [index page](../../foo/index.md) link with a longer path",
			expectedOutput: "This is an [index page](../../foo/_index.md) link with a longer path",
		},
		{
			name:           "Link with anchors",
			input:          "This one has an [anchor](../index.md#bar)",
			expectedOutput: "This one has an [anchor](../_index.md#bar)",
		},
		{
			name:           "Link with anchors and a query string",
			input:          "This one has an [anchor](../index.md?removal_milestone=17.0#gitlab-170)",
			expectedOutput: "This one has an [anchor](../_index.md?removal_milestone=17.0#gitlab-170)",
		},
		{
			name:           "Absolute link",
			input:          "This link is [absolute](https://example.com/index.md)",
			expectedOutput: "This link is [absolute](https://example.com/index.md)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := updateIndexLinks([]byte(tc.input), "")
			require.Nil(t, err)
			require.Equal(t, tc.expectedOutput, output)
		})
	}
}
