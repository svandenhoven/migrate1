# Documentation post-processing

During this project, content is processed after cloning to make it compatible with Hugo.
This allows us to develop on the Hugo site with actual content while still continously
publishing to the Nanoc site.

How this works:

1. When we build the site, we start by running `make clone-docs-projects`
1. `clone-docs-projects` clones source content
1. `clone-docs-projects` then invokes `scripts/content-post-process.sh`, which runs migration scripts
that modify file names, front matter, and content
1. `scripts/content-post-process.sh` then fetches and modifies YAML data from the `gitlab-docs` project

When it's time to launch the Hugo site, the content migration scripts can then be used to
update the source documentation.

YAML data files are not currently tracked in Git (see `.gitignore`) because we need
to fetch and modify them at build time. After launch, these will be part of the
project repository.

## Exclude files

Some files are not published to the website and thus should not be processed.

To exclude a file or directory, add its path to the `IgnoreFiles` slice in
[`migrate.go`](../cmd/gldocs/tasks/migrate.go).
