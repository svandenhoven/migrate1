# Set up local development and preview

Set up your workstation to develop for the GitLab Docs site and preview content changes.

## Prerequisites

Prerequisites:

- [Git](https://git-scm.com)
- Make. For example, [GNU Make](https://www.gnu.org/software/make/).

If these prerequisites aren't installed:

- macOS users can install them by using [Homebrew](https://brew.sh), which is also useful for installing other packages.
- Linux users can to install them by using their distribution's package manager.

## Optional: Install `asdf` or `mise`

This project supports both `asdf` and `mise` for installing required language runtimes and other system dependencies.

For information on installing:

- `asdf`, see <https://asdf-vm.com/guide/getting-started.html>.
- `mise`, see <https://mise.jdx.dev/getting-started.html>.

If you don't install either `asdf` or `mise`, you must install required language runtimes and other system dependencies
yourself.

## Set up your workstation

To set up your workstation for GitLab Docs:

1. Install either `asdf` or `mise`.
1. Clone this project (`gitlab-docs-hugo`).
1. In the checkout of `gitlab-docs-hugo`, run:

   ```shell
   make setup
   ```

1. The website build will pull content from project directories that are
   checked out in the same location as `gitlab-docs-hugo`. You can use your
   existing cloned repositories, or you can clone them all at once by running this command:

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
