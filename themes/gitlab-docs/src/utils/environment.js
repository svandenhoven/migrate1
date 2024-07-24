/**
 * Utilities for determining site environment.
 */

export const docsBaseURL = () =>
  document
    .querySelector('meta[name="gitlab-docs-base-url"]')
    ?.getAttribute("content");

/**
 * Check if we're on a production domain
 *
 * We can drop gitlab-org.gitlab.io once we move back to
 * the gitlab-docs project.
 *
 * @returns Boolean
 */
export function isProduction() {
  const prodHosts = ["docs.gitlab.com", "localhost", "new.docs.gitlab.com"];
  return prodHosts.includes(window.location.hostname);
}

/**
 * Get the URL path for the active page,
 * relative to the base URL
 *
 * @returns String
 */
export function relativeCurrentPath() {
  return `/${window.location.href.replace(docsBaseURL(), "")}`;
}
