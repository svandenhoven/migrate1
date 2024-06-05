/**
 * Utilities for determining site environment.
 */

export function isProduction() {
  const prodHosts = ["docs.gitlab.com", "localhost"];
  return prodHosts.includes(window.location.hostname);
}
