package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/cmd/gldocs/tasks"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone source projects for the Docs site.",
	Long:  `Clone source projects for the Docs site. See https://docs.gitlab.com/ee/development/documentation/site_architecture/ for more information. `,
	Run:   cloneRun,
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func cloneRun(cmd *cobra.Command, args []string) {
	tasks.CloneProjects()
}
