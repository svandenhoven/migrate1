# Set up local development and preview

Set up your workstation to develop for the GitLab Docs site and preview content changes.

## Prerequisites

Prerequisites:

- [Git](https://git-scm.com)
- Make. For example, [GNU Make](https://www.gnu.org/software/make/).

If these prerequisites aren't installed:

- macOS users can install them by using [Homebrew](https://brew.sh), which is also useful for installing other packages.
- Linux users can to install them by using their distribution's package manager.

## Set up your workstation

To set up your workstation for GitLab Docs:

1. Complete the [core `asdf` installation instructions](https://asdf-vm.com/guide/getting-started.html).
1. To make `asdf` available for later steps, close the terminal window you used to install `asdf` and open a new terminal window.
1. Clone this project (`gitlab-docs-hugo`).
1. In the checkout of `gitlab-docs-hugo`, run:

   ```shell
   make setup
   ```
