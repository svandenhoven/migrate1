/**
 * Helpers for migrating to Hugo
 * Functions here can be dropped after the site is launched.
 */

/**
 * Temporary handler for legacy URLs
 *
 * When we launch, many of our URLs change, but Google's
 * results will still be pointing at the legacy paths
 * until Google crawls the site again.
 *
 * We will need to rewrite links for this for at least
 * a few days after we launch, until Google has re-indexed
 * the site with the new URLs.
 */
export const rewriteLegacyURL = (url) => {
  const parsedUrl = new URL(url);

  // Check if the path starts with '/ee/' and remove it if it does
  if (parsedUrl.pathname.startsWith("/ee/")) {
    parsedUrl.pathname = parsedUrl.pathname.substring(3);
  }

  // Remove trailing slash if present
  parsedUrl.pathname = parsedUrl.pathname.replace(/\/$/, "");

  // Add .html if not already present
  if (!parsedUrl.pathname.endsWith(".html")) {
    parsedUrl.pathname += ".html";
  }

  return parsedUrl.toString();
};

/**
 * Redirect to the new test domain from the old one.
 *
 * This sort of redirect is not supported in GitLab Pages,
 * even with the domain-level redirects feature flag enabled.
 *
 * See https://gitlab.com/gitlab-org/gitlab/-/issues/14243#note_1905175772
 */
export const getRedirectUrl = (currentUrl) => {
  const oldDomain =
    "https://gitlab-org.gitlab.io/technical-writing-group/gitlab-docs-hugo";
  const newDomain = "https://new.docs.gitlab.com";

  if (currentUrl.startsWith(oldDomain)) {
    if (currentUrl === `${oldDomain}/` || currentUrl === oldDomain) {
      return "https://new.docs.gitlab.com";
    }
    return currentUrl.replace(oldDomain, newDomain);
  }
  return null; // No redirect needed
};

export const primaryDomainRedirect = () => {
  const redirectUrl = getRedirectUrl(window.location.href);
  if (redirectUrl) {
    window.location.replace(redirectUrl);
  }
};
