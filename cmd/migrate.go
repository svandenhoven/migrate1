package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/cmd/gldocs/tasks"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate docs content from Nanoc to Hugo.",
	Long:  `Migrate docs content from Nanoc to Hugo. See doc/post-processing.md for more information.`,
	Run:   migrateRun,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrateRun(cmd *cobra.Command, args []string) {
	tasks.MigrateFrontmatter(args)
}
