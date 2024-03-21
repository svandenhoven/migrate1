# GitLab docs site maintenance

Some of the issues that the GitLab technical writing team handles to maintain
`https://docs.gitlab.com` include:

- The deployment process.
- Temporary event or survey banners.

## Deployment process

We use [GitLab Pages](https://docs.gitlab.com/ee/user/project/pages/) to build and host this website.

The site is built and deployed automatically in GitLab CI/CD jobs.
See [`.gitlab-ci.yml`](../.gitlab-ci.yml)
for the current configuration. The project has [scheduled pipelines](https://docs.gitlab.com/ee/ci/pipelines/schedules.html)
that build and deploy the site every hour.

## Survey banner

In case there's a survey that needs to reach a big audience, the docs site has
the ability to host a banner for that purpose. When it is enabled, it's shown
at the top of every interior page of the docs site.

To publish a survey, edit [`banner.yaml`](/data/banner.yaml) and:

1. Set `show_banner` to `true`.
1. Under `description`, add what information you want to appear in the banner.
   Markdown is supported.

To unpublish a survey, edit [`banner.yaml`](/content/_data/banner.yaml) and
set `show_banner` to `false`.
