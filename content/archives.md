---
title: GitLab Docs Archives
feedback: false
toc: true
---

This page has links to online versions of the GitLab documentation. To browse a version that is not available online,
see the [offline archives](#offline-archives).

They’re available for download so that you can browse through them locally, or [self-host](/administration/docs_self_host)
them.

Version-specific search was introduced in 15.6 for the [online archives](https://archives.docs.gitlab.com/). Offline
search also works in self-hosted environments as of 16.6.

## Default Version

The [default version](https://docs.gitlab.com/) of this website is built from the documentation directories on the default
branches of:

- [GitLab](https://gitlab.com/gitlab-org/gitlab/-/tree/master/doc)
- [Omnibus GitLab](https://gitlab.com/gitlab-org/omnibus-gitlab/-/tree/master/doc)
- [GitLab Runner](https://gitlab.com/gitlab-org/gitlab-runner/-/tree/main/docs)
- [GitLab Charts](https://gitlab.com/gitlab-org/charts/gitlab/-/tree/master/doc)
- [GitLab Operator](https://gitlab.com/gitlab-org/cloud-native/gitlab-operator/doc)

All of these are brought together by the [GitLab Docs](https://gitlab.com/gitlab-org/gitlab-docs/) project, which
regularly deploys this content to [docs.gitlab.com](https://docs.gitlab.com/).

## Previously released versions

[Supported versions](https://about.gitlab.com/support/statement-of-support/#version-support) of GitLab Docs are available
online on the [GitLab Docs Archives website](https://archives.docs.gitlab.com/).

## Offline archives

The following archives are available and can be browsed offline. You'll need to have
[Docker](https://docs.docker.com/get-docker/) installed to access them.

<!-- Call Hugo template here to grab versions and docker commands -->
{{< archiveList >}}
