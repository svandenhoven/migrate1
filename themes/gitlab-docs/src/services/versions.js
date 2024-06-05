const DOCS_VERSIONS_ENDPOINT = "https://docs.gitlab.com/versions.json";
const ARCHIVE_VERSIONS_ENDPOINT =
  "https://archives.docs.gitlab.com/archive_versions.json";

export const SITE_VERSION = document
  .querySelector('meta[name="gitlab-docs-version"]')
  ?.getAttribute("content");

/**
 * Fetch a list of versions available on docs.gitlab.com.
 *
 * @returns Array
 */
let cachedVersions = null;

export async function getVersions() {
  if (!cachedVersions) {
    try {
      const data = await (await fetch(DOCS_VERSIONS_ENDPOINT)).json();
      cachedVersions = Object.assign(...data);
    } catch (error) {
      console.error(error);
    }
  }
  return cachedVersions || [];
}

/**
 * Fetch a list of site versions available on the Archives site.
 *
 * @returns Array
 */
export async function getArchivesVersions() {
  const versions = await fetch(ARCHIVE_VERSIONS_ENDPOINT)
    .then((response) => response.json())
    .catch((error) => console.error(error));
  return versions || [];
}
