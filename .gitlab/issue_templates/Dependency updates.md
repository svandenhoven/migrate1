## Dependency updates

Remove sections below as needed. Generally we want to batch related updates together and create separate MRs
for different parts of the stack. For example, it makes sense to update Node.js modules and Yarn together. However,
if you need to update Node.js modules and Go, you should create MRs.

### Node modules

- [ ] Run `npx npm-check-updates -u -x '*vue*, eslint'` to bump versions for everything except Vue and eslint.
- [ ] Run `yarn` to update `yarn.lock`.
- [ ] Run `make lint-frontend` and `make jest-tests` and verify frontend tests still pass.
- [ ] Run `make view` and browse around the site a bit. Check for new console errors or other problems.

### Yarn

[Release History](https://github.com/yarnpkg/berry/releases)

- [ ] Run `yarn set version stable` to update the `packageManager` field in `package.json`.

### Go version

[Release History](https://go.dev/doc/devel/release)

- [ ] Update `go.mod`
- [ ] Update `.tool-versions`
- [ ] If this is a major version, update `GO_VERSION_PREVIOUS` and `GO_VERSION` in `.gitlab-ci.yml`
- [ ] If this is a major version, update the versions in `.go-matrix-job` in `.gitlab-ci.yml`
- [ ] Check that linting and tests still pass: `make lint-go && make go-tests`

### Go modules

- [ ] Run `go get -t -u ./...` to update all dependencies, including test dependencies
- [ ] Run `go mod tidy` to remove now unneeded dependencies from `go.sum`
- [ ] Check that linting and tests still pass: `make lint-go && make go-tests`

### Linux distros

Update the relevant variable in `.gitlab-ci.yml` (e.g, `ALPINE_VERSION`). Verify pipelines all still pass.

- [ ] Alpine ([releases](https://alpinelinux.org/releases/))
- [ ] Debian ([releases](https://www.debian.org/releases/))

### Other tools

Update `.tool-versions` and the relevant variable in `.gitlab-ci.yml` (e.g, `HUGO_VERSION`).

- [ ] Hugo ([releases](https://github.com/gohugoio/hugo/releases))
- [ ] node.js ([releases](https://nodejs.org/en/download/package-manager)). We want to use the LTS version of node.
- [ ] Shellcheck ([releases](https://github.com/koalaman/shellcheck/releases))
- [ ] yq ([releases](https://github.com/mikefarah/yq/releases))
- [ ] golangci-lint ([releases](https://github.com/golangci/golangci-lint/releases))

Update `.gitlab-ci.yml` variables only:

- [ ] Lychee
- [ ] asdf ([releases](https://github.com/asdf-vm/asdf/releases))
