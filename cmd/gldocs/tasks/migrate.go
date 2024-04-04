package tasks

import (
	"fmt"
)

type MigrationFunc func(files []string)

// Register each migration here
var migrations = map[string]MigrationFunc{
	"frontmatter": MigrateFrontmatter,
	"icons":       MigrateIcons,
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
