package cmd

import (
	"log"

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
	if len(args) < 1 {
		log.Fatal("Usage: go run cmd/gldocs/main.go migrate [migration_name] <directory_path>")
	}

	dirPath := args[len(args)-1]
	migrationName := "all"
	if len(args) > 1 {
		migrationName = args[0]
	}

	if err := tasks.Migrate(dirPath, migrationName); err != nil {
    log.Fatal(err)
  }
}
