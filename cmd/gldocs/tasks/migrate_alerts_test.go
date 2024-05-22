package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateAlerts(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Single alert",
			input: `
Some text before the alert.

NOTE:
This is a note.

Some text after the alert.
`,
			expected: `
Some text before the alert.

{{< alert type="note" >}}

This is a note.

{{< /alert >}}

Some text after the alert.
`,
		},
		{
			name: "Multiple alerts",
			input: `
NOTE:
This is a note.

Some text between alerts.

WARNING:
This is a warning.

FLAG:
This is a flag.

`,
			expected: `
{{< alert type="note" >}}

This is a note.

{{< /alert >}}

Some text between alerts.

{{< alert type="warning" >}}

This is a warning.

{{< /alert >}}

{{< alert type="flag" >}}

This is a flag.

{{< /alert >}}

`,
		},
		{
			name: "Valid disclaimer",
			input: `
DISCLAIMER:
This page contains information related to upcoming products, features, and functionality. It is important to note that the information presented is for informational purposes only. Please do not rely on this information for purchasing or planning purposes. As with all projects, the items mentioned on this page are subject to change or delay. The development, release, and timing of any products, features, or functionality remain at the sole discretion of GitLab Inc.

`,
			expected: `
{{< alert type="disclaimer" />}}

`,
		},
		{
			name: "Invalid disclaimer",
			input: `
DISCLAIMER:
This is some text that should be in a different type of alert box.
`,
			expected: `
DISCLAIMER:
This is some text that should be in a different type of alert box.
`,
		},
		{
			name: "Availability details",
			input: `
DETAILS:
**Tier:** Free, Premium, Ultimate
**Offering:** GitLab.com, Self-managed, GitLab Dedicated
**Status:** Experiment

`,
			expected: `
{{< details >}}

- Tier: Free, Premium, Ultimate
- Offering: GitLab.com, Self-managed, GitLab Dedicated
- Status: Experiment

{{< /details >}}

`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, _ := updateAlerts([]byte(tc.input), "")
			require.Equal(t, tc.expected, output)
		})
	}
}

func TestCheckParts(t *testing.T) {
	testCases := []struct {
		name        string
		input       []string
		expectedErr bool
	}{
		{
			name:        "Valid parts",
			input:       []string{"NOTE:\nThis is a note.\n\n", "  ", "This is a note.\n"},
			expectedErr: false,
		},
		{
			name:        "Missing indentation and content",
			input:       []string{"NOTE:\n\n"},
			expectedErr: true,
		},
		{
			name:        "Missing content",
			input:       []string{"NOTE:\nThis is a note.\n\n", "  "},
			expectedErr: true,
		},
		{
			name:        "Empty parts",
			input:       []string{},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkParts(tc.input)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
