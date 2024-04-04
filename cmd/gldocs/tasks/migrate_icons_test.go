package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateIcons(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Simple icon",
			input:          "This is a **{pencil}** icon.",
			expectedOutput: "This is a {{< icon name=\"pencil\" >}} icon.",
		},
		{
			name:           "Icon with extra attributes",
			input:          "This is a **{severity-critical, 18, gl-fill-red-800}** icon.",
			expectedOutput: "This is a {{< icon name=\"severity-critical\" >}} icon.",
		},
		{
			name:           "Multiple icons",
			input:          "We have **{pencil}** and **{severity-critical, 18, gl-fill-red-800}** icons.",
			expectedOutput: "We have {{< icon name=\"pencil\" >}} and {{< icon name=\"severity-critical\" >}} icons.",
		},
		{
			name:           "No icons",
			input:          "This text doesn't have any icons.",
			expectedOutput: "This text doesn't have any icons.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := updateIcons([]byte(tc.input), "")
			require.Nil(t, err)
			require.Equal(t, tc.expectedOutput, output)
		})
	}
}
