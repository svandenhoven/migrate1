# Testing the GitLab Docs site

Tests for the GitLab Docs site include tests for code and tests for links in content. For more information, see
[Documentation testing](https://docs.gitlab.com/ee/development/documentation/testing/).

Tests are run in `gitlab-docs-hugo`
[CI/CD pipeline](https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/-/pipelines), which is
configured in the project's [`.gitlab-ci.yml`](../.gitlab-ci.yml) file.

## Code tests

These code tests are included in the project:

| Test target   | Tool(s)                     | Make target          | Purpose                                    |
|:--------------|:----------------------------|:---------------------|:-------------------------------------------|
| Frontend      | Stylelint, ESLint, Prettier | `lint-frontend`      | Frontend code quality                      |
| Markdown      | markdownlint                | `lint-markdown`      | Documentation formatting and syntax checks |
| Shell scripts | ShellCheck                  | `lint-shell-scripts` | Syntax checks                              |
| YAML          | yamllint                    | `lint-yaml`          | Syntax checks                              |
| Vue           | Jest                        | `jest-tests`         | Unit tests                                 |
| Go            | Go testing, Testify         | `go-tests`           | Unit tests                                 |

### Run code tests locally

To run all tests:

```shell
make test
```

You can also run tests individually by specifying the Make target. For example, to run Go tests only:

```shell
make go-tests
```

### Install Lefthook

If you want to run the tests before pushing changes, use [Lefthook](https://github.com/evilmartians/lefthook#readme).
To install Lefthook, run:

```shell
yarn run lefthook install
```

Tests are run whenever you run `git push`.

Lefthook is configured in [`lefthook.yml`](../lefthook.yml).
