---
title: Shortcodes
feedback: false
---

[Shortcodes](https://gohugo.io/content-management/shortcodes/)
are snippets of template code that we can include in our Markdown
content to display non-standard elements on a page, such as alert
boxes or tabs. Shortcodes are the Hugo equivalent of
[Filters](https://nanoc.app/doc/filters/) in Nanoc.

This page is intended to be a development reference during
the Hugo project, as well as a guide for changes we'll need to
make to the [Documentation Style Guide](https://docs.gitlab.com/ee/development/documentation/styleguide)
when we launch the new site.

## Alert boxes

[Nanoc version: Alert boxes](https://docs.gitlab.com/ee/development/documentation/styleguide/#alert-boxes)

### Note

{{< alert type="note" >}}

Here's some note text.

You can add any standard markdown inside any of
these alerts. For example, here's a list:

- One
- Two
- Five hundred

{{< /alert >}}

```text
{{</* alert type="note" */>}}

Here's some note text.

You can add any standard markdown inside any of
these alerts. For example, here's a list:

- One
- Two
- Five hundred

{{</* /alert */>}}
```

### Warning

{{< alert type="warning" >}}

This is a warning!

{{< /alert >}}

```text
{{</* alert type="warning" */>}}

This is a warning!

{{</* /alert */>}}
```

### Flag

{{< alert type="flag" >}}

This is a feature flag.

{{< /alert >}}

```text
{{</* alert type="flag" */>}}

This is a feature flag.

{{</* /alert */>}}
```

### Disclaimer

You do not need to add any text for the disclaimer alert type, it is populated from a template.

{{< alert type="disclaimer" />}}

```text
{{</* alert type="disclaimer" /*/>}}
```

## Availability details

{{< details >}}

- Tier: Premium, Ultimate with [GitLab Duo Pro](https://docs.gitlab.com/ee/subscriptions/subscription-add-ons.html)
- Offering: GitLab.com, Self-managed, GitLab Dedicated
- Status: Experiment
- Hotdogs: Yes
- Some other thing: Absolutely not

{{< /details >}}

```text
{{</* details */>}}

- Tier: Premium, Ultimate with [GitLab Duo Pro](https://docs.gitlab.com/ee/subscriptions/subscription-add-ons.html)
- Offering: GitLab.com, Self-managed, GitLab Dedicated
- Status: Experiment
- Hotdogs: Yes
- Some other thing: Absolutely not

{{</* /details */>}}
```

[Nanoc version: Product availability](https://docs.gitlab.com/ee/development/documentation/styleguide/availability_details.html)

## Version history

{{< history >}}

- Introduced in [GitLab 17.0](https://about.gitlab.com/releases/2024/05/16/gitlab-17-0-released/)
- Deprecated in GitLab 17.1

{{< /history >}}

```text
{{</* history */>}}

- Introduced in [GitLab 17.0](https://about.gitlab.com/releases/2024/05/16/gitlab-17-0-released/)
- Deprecated in GitLab 17.1

{{</* /history */>}}
```

[Nanoc version: Version history](https://docs.gitlab.com/ee/development/documentation/versions.html)

### Availability details with version history

{{< details >}}

- Tier: some tiers
- Offering: some offerings
- Status: Experiment
- Hotdogs: Yes

{{< /details >}}

{{< history >}}

- Introduced in xyz
- Deprecated in abc

{{< /history >}}

```text
{{</* details */>}}

- Tier: some tiers
- Offering: some offerings
- Status: Experiment
- Hotdogs: Yes

{{</* /details */>}}

{{</* history */>}}

- Introduced in xyz
- Deprecated in abc

{{</* /history */>}}
```

## Icons

[Nanoc version: Icons](https://docs.gitlab.com/ee/development/documentation/styleguide/#gitlab-svg-icons)

This is an icon: {{< icon name="tanuki" >}}

```text
This is an icon: {{</* icon name="tanuki" */>}}
```
