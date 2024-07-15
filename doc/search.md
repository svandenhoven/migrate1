# Search

GitLab Docs supports two search backends, Google Programmable Search and Lunr.js. Environment variables are passed at
build time to set the search backend.

- The primary production site, docs.gitlab.com, runs Google Programmable Search.
- Lunr.js. Archives, review apps, and self-hosted sites run Lunr.js.

## Google Programmable Search

The Technical Writing team introduced Google Programmable Search in May 2023 as part of the 16.0 release (see [epic](https://gitlab.com/groups/gitlab-org/-/epics/9420)).

### Implementation details

Search forms are built with [`search-box-by-type`](https://gitlab-org.gitlab.io/gitlab-ui/?path=/docs/base-search-box-by-type--docs)
GitLab UI components.
The forms make [API requests](../themes/gitlab-docs/src/services/google_search_api.js) to the
[Google Custom Search Site Restricted JSON API](https://developers.google.com/custom-search/v1/site_restricted_api).

Search results appear beneath the form as you type on the homepage and interior pages.
They can also choose "View all results" to run their query from the [advanced search page](https://docs.gitlab.com/search).

The advanced search page includes up to 100 results.
These are rendered alongside custom filters that use a [section metatag](../themes/gitlab-docs/layouts/partials/head.html)
to filter pages by navigation section.

#### Vue components

- [Search forms on the homepage and interior content pages](../themes/gitlab-docs/src/components/search/google_search_form.vue)
- [Search form and results on the advanced search page](../themes/gitlab-docs/src/components/search/google_results.vue)
- [Search filters on the advanced search page](../themes/gitlab-docs/src/components/search/search_filters.vue)

#### Configuration

Google Programmable Search has two areas of admininstration on Google:

1. The [Programmable Search Engine control panel](https://programmablesearchengine.google.com/controlpanel/overview?cx=97494f9fe316a426d)
is used to configure settings for search behavior.
1. Billing and credentials (API keys) are managed from the [Google Cloud Console](https://console.cloud.google.com).
Changes or access requests for the console can be requested from the GitLab Infrastructure team by
[creating an issue](https://gitlab.com/gitlab-com/gl-infra/reliability/-/issues) in the Reliability team project.
Access to this is not required to do development work on search.

Configuration from the Programmable Search Engine (PSE) control panel is tracked in [`config/search/google`](../config/search/google/)
as XML files.
These files must be manually uploaded and downloaded from the PSE control panel to change actual search configuration.

You can read more about what each file is used for on the [Programmable Search Engine documentation site](https://developers.google.com/custom-search/docs/basics).

##### Modify synonyms

Sometimes we may want to provide synoyms if users are searching for terms that differ from what's actually written.

To add a synonym:

1. Review Google's [best practices documentation](https://developers.google.com/custom-search/docs/queries#expanding-search-queries-with-synonyms).
1. Modify [`config/search/google/synonyms.xml`](/config/search/google/synonyms.xml) as needed. For formatting guidelines,
see [Creating Synonyms](https://developers.google.com/custom-search/docs/queries#creating-synonyms)
on the PSE documentation site.
1. Create an MR with your change, and add `configuration` and `docsearch::google` labels.
1. After approved and before merging the MR, a [CODEOWNER](../CODEOWNERS) for the `config/search/google` directory
must visit the PSE dashboard,and under the [Query Enhancement card](https://programmablesearchengine.google.com/controlpanel/searchfeatures?cx=97494f9fe316a426d#query-enhancement-card),
expand "Upload/Download Synonyms in XML files" and select **Upload** to upload the new XML file.

Keep in mind that this only affects Google-powered search on docs.gitlab.com, not general Google search.
It's likely a better idea to consider changing content rather than search configuration.

#### Authentication

The production site and review apps fetch the Google Programmable Search API key from the `GOOGLE_SEARCH_KEY` variable
in CI environments.
This key is set as a CI variable. GitLab Docs maintainers can change that key by navigating to Settings > CI/CD  > Variables
in the [GitLab Docs Hugo](https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo) project.

As this API key is used in frontend code, it is a public key, but storing it as a variable allows for rotating the key
without making code changes.

#### Analytics

- [Statistics & Logs](https://programmablesearchengine.google.com/cse/statistics/stats?cx=97494f9fe316a426d)

## Lunr.js Search

[Lunr.js](https://lunrjs.com/) is available as an alternative search backend for archived and self-hosted GitLab Docs installations.
Lunr search can also be used in offline or air-gapped environments.
Lunr search requires an additional build step to create a search index.

Documentation review apps and versions of the site running on <https://archives.docs.gitlab.com> use Lunr search by default.

## Development

### Local environment

You can build your local Nanoc site to use a specific search backend
by setting the `SEARCH_BACKEND` environment variable at compile time.

- Use Google search: `SEARCH_BACKEND="google" make view`. Default if `SEARCH_BACKEND` is not set.
- Use Lunr.js search: `SEARCH_BACKEND="lunr" make view`.

#### Local build with Google Programmable Search

Running Google search on your local site requires an API key. To set this locally:

1. Copy the key from the project [CI/CD settings](https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/-/settings/ci_cd).
1. Add the key as a local environment variable in `~/.zshrc` (or whatever your shell settings file is) like this:
`export GOOGLE_SEARCH_KEY="abc123"`
1. Run  `make view` to build and preview locally.

Alternatively, you can pass the API key for a build like this:

```shell
GOOGLE_SEARCH_KEY="abc123" SEARCH_BACKEND="google" make view
```

Developers can also use [HackyStack](https://about.gitlab.com/handbook/infrastructure-standards/realms/sandbox/)
on Google to spin up a sandbox GCP project.
This allows for creating and testing a separate Programmable Search Engine, with separate API keys, and separate billing.

#### Local build with Lunr.js search

To create a local build with Lunr.js search:

1. `make setup`
1. `SEARCH_BACKEND="lunr" make view`

### Review apps

To build a review app with Google Programmable Search as the search backend,
include the string `gps` in the name of your Git branch.
