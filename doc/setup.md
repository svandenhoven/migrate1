# Set up local development and preview

Set up your workstation to preview content changes on the Gitab Docs website.

If you intend to do development work on the site, see
[GitLab docs site development](development.md) for additional setup
guidelines.

## Prerequisites

Prerequisites:

- [Git](https://git-scm.com)
- Make. For example, [GNU Make](https://www.gnu.org/software/make/).

If these prerequisites aren't installed:

- macOS users can install them by using [Homebrew](https://brew.sh), which is also useful for installing other packages.
- Linux users can to install them by using their distribution's package manager.

## Managing system dependencies

This project supports both `asdf` and `mise` for installing required language runtimes and other system dependencies.

For information on installing:

- `asdf`, see <https://asdf-vm.com/guide/getting-started.html>.
- `mise`, see <https://mise.jdx.dev/getting-started.html>.

If you don't install either `asdf` or `mise`, you must install required language runtimes and other system dependencies yourself.
See [`.tool-versions`](../.tool-versions) for the required dependencies.

If you use `asdf` or `mise`, these dependencies will be installed automatically when you follow
the steps below.

## Set up your workstation

To set up your workstation for GitLab Docs:

1. Verify that you have `asdf` or `mise` installed, or that you've installed system dependencies manually.
1. Clone this project (`gitlab-docs-hugo`).

   Until we launch the Hugo site, you should clone the project to a separate
   location from your existing project checkouts, as it will modify docs markdown files on
   build.

   For example, if your existing projects are in `~/dev`, you could clone
   `gitlab-docs-hugo` to `~/dev/docs-new/gitlab-docs-hugo`.

1. Ensure you are in the `gitlab-docs-hugo` folder. Then run:

   ```shell
   make setup
   ```

1. The website build pulls content from project directories that are
   checked out in the same location as `gitlab-docs-hugo`. You can use your
   existing cloned repositories, or you can clone them all at once by running this command
   from within the `gitlab-docs-hugo` folder:

   ```shell
   make clone-docs-projects
   ```

   In either case, your resulting file structure should look like this:

   ```markdown
   .
   ├── charts-gitlab/
   ├── gitlab/
   ├── gitlab-docs-hugo/
   ├── gitlab-runner/
   ├── gitlab-operator/
   └── omnibus-gitlab/
   ```

1. Compile the site, and view a local preview:

   ```shell
   make view
   ```

You can now browse the site at [`http://localhost:1313`](http://localhost:1313). Changes you make to markdown content or
website source files should be visible immediately.

### Troubleshooting

If you encounter problems with your local site, run `make setup` first to make sure
dependencies are all installed and up-to-date.

For further assistance, GitLab teammembers can reach out in the `#docs-tooling` Slack channel.
Community contributors can [file an issue](https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/-/issues)
in the GitLab Docs project.
