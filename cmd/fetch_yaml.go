package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/cmd/gldocs/tasks"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch YAML data files for the Docs site.",
	Long:  `Fetch YAML data for the Docs site. See doc/post-processing.md for more information. `,
	Run:   fetchRun,
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func fetchRun(cmd *cobra.Command, args []string) {
	files := []tasks.FileSource{
		{
			URL:         "https://gitlab.com/gitlab-org/gitlab-docs/-/raw/main/content/_data/navigation.yaml",
			Destination: "data/navigation.yaml",
		},
	}

	// Download files
	tasks.FetchRemoteDataFiles(files)

	// Modify their contents
	tasks.ModifyDataFiles(files)
}
