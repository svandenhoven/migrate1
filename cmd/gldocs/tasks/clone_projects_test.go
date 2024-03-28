package tasks

import (
	"fmt"
	"testing"
)

func TestShouldSkipClone(t *testing.T) {
	testCases := []struct {
		name               string
		CI_PIPELINE_SOURCE string
		branch             string
		product            Product
		expectedReturn     bool
	}{
		{
			name:               "Remote branch is the same as product default",
			CI_PIPELINE_SOURCE: "trigger",
			branch:             "main",
			product:            Product{DefaultBranch: "main"},
			expectedReturn:     true,
		},
		{
			name:               "Pipeline triggered via API (multi-project pipeline)",
			CI_PIPELINE_SOURCE: "push",
			branch:             "main",
			product:            Product{DefaultBranch: "main"},
			expectedReturn:     false,
		},
		{
			name:               "Remote branch and default branch are different",
			CI_PIPELINE_SOURCE: "trigger",
			branch:             "my-new-feature",
			product:            Product{DefaultBranch: "main"},
			expectedReturn:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("CI_PIPELINE_SOURCE", tc.CI_PIPELINE_SOURCE)

			actualReturn := tc.product.ShouldSkipClone(tc.branch)

			if actualReturn != tc.expectedReturn {
				t.Errorf("Unexpected result. Got (%t), want (%t)",
					actualReturn, tc.expectedReturn)
			}
		})
	}
}

func TestProductCloneInfo(t *testing.T) {
	testCases := []struct {
		name            string
		productName     string
		product         Product
		envVarSuffix    string
		mergeRequestIID string
		branchName      string
		expectedBranch  string
		expectedRef     string
	}{
		{
			name:            "Hourly pipeline, or MR to gitlab-docs",
			productName:     "",
			envVarSuffix:    "",
			product:         Product{DefaultBranch: "main"},
			mergeRequestIID: "",
			branchName:      "main",
			expectedBranch:  "main",
			expectedRef:     "heads/main",
		},
		{
			name:            "Stable branch pipeline",
			productName:     "",
			envVarSuffix:    "",
			product:         Product{DefaultBranch: "main"},
			mergeRequestIID: "",
			branchName:      "16.6",
			expectedBranch:  "16.6",
			expectedRef:     "heads/16.6",
		},
		{
			name:            "Test with merge request",
			productName:     "gitlab",
			envVarSuffix:    "EE",
			product:         Product{DefaultBranch: "master"},
			mergeRequestIID: "123",
			branchName:      "feature-branch",
			expectedBranch:  "feature-branch",
			expectedRef:     "merge-requests/123/head",
		},
		{
			name:            "Test without merge request",
			productName:     "operator",
			envVarSuffix:    "OPERATOR",
			product:         Product{DefaultBranch: "main"},
			mergeRequestIID: "",
			branchName:      "",
			expectedBranch:  "main",
			expectedRef:     "heads/main",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(fmt.Sprintf("MERGE_REQUEST_IID_%s", tc.envVarSuffix), tc.mergeRequestIID)
			t.Setenv(fmt.Sprintf("BRANCH_%s", tc.envVarSuffix), tc.branchName)

			// Call the function and check the results
			actualBranch, actualRef := tc.product.CloneInfo(tc.productName)

			if actualBranch != tc.expectedBranch || actualRef != tc.expectedRef {
				t.Errorf("Unexpected result. Got (%s, %s), want (%s, %s)",
					actualBranch, actualRef, tc.expectedBranch, tc.expectedRef)
			}
		})
	}
}
