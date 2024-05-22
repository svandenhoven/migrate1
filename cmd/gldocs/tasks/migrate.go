package tasks

import (
	"fmt"
)

// These pages are not published with the site but exist in docs directories.
// No need to process or warn about these.
var IgnorePaths = []string{
	"../gitlab-operator/doc/adr",
	"../gitlab/drawers",
}

type MigrationFunc func(files []string)

// Register each migration here
var migrations = map[string]MigrationFunc{
	"frontmatter": MigrateFrontmatter,
	"icons":       MigrateIcons,
	"alerts":      MigrateAlerts,
	"indexLinks":  MigrateIndexLinks,
}

func Migrate(dirPath string, migrationName string) error {
	files, err := FindMarkdownFiles(dirPath)
	if err != nil {
		return fmt.Errorf("Error finding markdown files: %v", err)
	}
	uniqueFiles := uniqueFilesOnly(files)

	if migrationName == "all" {
		for _, migrationFunc := range migrations {
			migrationFunc(uniqueFiles)
		}
	} else if migrationFunc, ok := migrations[migrationName]; ok {
		migrationFunc(uniqueFiles)
	} else {
		return fmt.Errorf("Unknown migration: %s\n", migrationName)
	}

	return nil
}
