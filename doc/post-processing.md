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

## Data files on CI

Because we refresh YAML data files as part of the content clone process (`make clone-docs-projects`),
these are not available for frontend jobs running on CI during the migration project.

If you need to reference a YAML data file from a frontend file:

- See the `create-navigation-stub` Yarn script as an example for creating a stub data
file. This approach is good for tests that provide their own mock dat, or lint checks.
- The `prebuild:fetch_data` CI job is used to fetch actual data prior to running the
frontend build.

These can be removed when the Hugo site launches. At that time we will be committing these
data files into the Git repository for this project and no longer needing to fetch them
from a separate source.
