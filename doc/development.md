# GitLab docs site development

## Theme development

The `gitlab-docs` theme contains the site's CSS and JS code.

### JavaScript and CSS

JavaScript and CSS code is bundled with Vite. To work on frontend code locally, run Vite in watch mode:

```shell
yarn run dev
```

To live reload changes, run Vite in watch mode, and in another terminal, run Hugo in development mode:

```shell
make view
```

### CSS usage

This project uses vanilla CSS, and we aim to keep custom styling minimal in order to reduce maintenance needs and
maintain consistency with GitLab UI. Try to use
[utility classes](https://docs.gitlab.com/ee/development/fe_guide/style/scss.html#utility-classes) where possible.

## Add a new product

NOTE:
We encourage you to create an issue and connect with the Technical Writing team before you add a new product to the
documentation site, as there may be planning information that the team can help with, including integrating any new
content into the site's global navigation menu.

To add an additional set of product documentation to <https://docs.gitlab.com> from a separate GitLab repository (beyond
any product documentation already added to the site):

1. Edit [`data/products.yaml`](../data/products.yaml) and complete the following steps:

   1. Add an entry to `products` similar to the other listed entries. For example:

      ```yaml
      gitlab:
        repo: 'https://gitlab.com/gitlab-org/gitlab.git'
        content_dir: 'doc'
        default_branch: 'master'
        clone_dir: '../gitlab'
        docs_dir: 'doc'
      ```

      Where:

      - `repo`: The repository for the new product. Use the "Clone with HTTPS" URL for the project.
      - `content_dir`: The relative path where the docs reside within the repo.
      - `default_branch`: Default Git branch for the project.
      - `clone_dir`: Target destination for the project repo. Usually this will just be the project name, located in the
        parent directory. Be sure to avoid duplicating `clone_dir` between products.
      - `docs_dir`: Directory name where docs content is located.

   1. Edit [`hugo.yaml`] and add the new content source under `mounts`:

      ```yaml
      module:
      mounts:
      # <existing sources>
      # Add a new entry like this at the end:
      - source: ../my-new-product/docs
        target: content
      ```

      Where:

      - `source`: The relative path to the new content source. This will be the `clone_dir` you used in `products.yaml`,
        plus the directory name for where the documentation files are located (`docs_dir` in the previous step).
      - `target`: This should always be `content`.

## Exclude a directory

To exclude a directory so the contents aren't published to the docs site, add the path
as an `excludeFiles` setting in the Hugo [config file](../config/_default/hugo.yaml).
For more information, see the Hugo
[module config documentation](https://gohugo.io/hugo-modules/configuration/#module-configuration-mounts).
