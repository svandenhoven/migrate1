package tasks

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateTabs(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Single tab",
			input: `
Some text before the tabs.

::Tabs

:::TabTitle Tab One

This is content in tab one.

::EndTabs

Some text after the tabs.
`,
			expected: `
Some text before the tabs.

{{< tabs >}}

{{< tab title="Tab One" >}}

This is content in tab one.

{{< /tabs >}}

Some text after the tabs.
`,
		},
		{
			name: "Multiple tabs",
			input: `
::Tabs

:::TabTitle Tab One

This is content in tab one.

:::TabTitle Tab Two

This is content in tab two.

:::TabTitle Tab Three

This is content in tab three.

::EndTabs
`,
			expected: `
{{< tabs >}}

{{< tab title="Tab One" >}}

This is content in tab one.

{{< tab title="Tab Two" >}}

This is content in tab two.

{{< tab title="Tab Three" >}}

This is content in tab three.

{{< /tabs >}}
`,
		},
		{
			name: "Multiple tab sets",
			input: `
::Tabs

:::TabTitle Tab One

Some content in the first tab.

:::TabTitle Tab Two

Content in tab two.

::EndTabs

Some other text here.

::Tabs

:::TabTitle Tab Set Two, Tab One

Some content in tab set two.

:::TabTitle Tab Set Two, Tab Two

Content in tab set two, tab two.

::EndTabs
`,
			expected: `
{{< tabs >}}

{{< tab title="Tab One" >}}

Some content in the first tab.

{{< tab title="Tab Two" >}}

Content in tab two.

{{< /tabs >}}

Some other text here.

{{< tabs >}}

{{< tab title="Tab Set Two, Tab One" >}}

Some content in tab set two.

{{< tab title="Tab Set Two, Tab Two" >}}

Content in tab set two, tab two.

{{< /tabs >}}
`,
		},
		{
			name:     "Tabs in a codeblock",
			input:    "```markdown\n::Tabs\n\n:::TabTitle Tab One\n\nThis is content in tab one.\n\n::EndTabs\n```",
			expected: "```markdown\n::Tabs\n\n:::TabTitle Tab One\n\nThis is content in tab one.\n\n::EndTabs\n```",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, _ := updateTabs([]byte(tc.input), "")
			require.Equal(t, tc.expected, output)
		})
	}
}

func TestUpdateTabsWithMalformedInput(t *testing.T) {
	malformedTestCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Malformed ::Tabs (missing indentation)",
			input: `
Some text before.
::Tabs
:::TabTitle Tab One
Content
::EndTabs
`,
			expected: `
Some text before.
{{< tabs >}}

{{< tab title="Tab One" >}}

Content
{{< /tabs >}}
`,
		},
		{
			name: "Malformed ::EndTabs (missing indentation)",
			input: `
Some text before.
    ::Tabs
    :::TabTitle Tab One
    Content
::EndTabs
`,
			expected: `
Some text before.
    {{< tabs >}}

    {{< tab title="Tab One" >}}

    Content
{{< /tabs >}}
`,
		},
		{
			name: "Malformed :::TabTitle (missing title)",
			input: `
::Tabs
:::TabTitle
Content
::EndTabs
`,
			expected: `
{{< tabs >}}

{{< tab title="Content" >}}

{{< /tabs >}}
`,
		},
		{
			name:     "Completely wrong ::Tabs",
			input:    "::TabsWrong\n",
			expected: "::TabsWrong\n",
		},
		{
			name:     "Completely wrong :::TabTitle",
			input:    ":::TabTitleWrong\n",
			expected: ":::TabTitleWrong\n",
		},
	}

	for _, tc := range malformedTestCases {
		t.Run(tc.name, func(t *testing.T) {
			// Capture stdout to check for error messages
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			output, err := updateTabs([]byte(tc.input), "test.md")

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read captured output
			var buf bytes.Buffer
			_, copyErr := io.Copy(&buf, r)
			require.NoError(t, copyErr, "Error copying captured output")
			capturedOutput := buf.String()

			require.NoError(t, err)
			require.Equal(t, tc.expected, output)

			// Check if an error message was printed for malformed input
			if strings.Contains(tc.name, "wrong") {
				require.Contains(t, capturedOutput, "Malformed tab structure in test.md")
			}
		})
	}
}
