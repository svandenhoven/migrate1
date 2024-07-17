/* global GOOGLE_SEARCH_KEY */

import { searchResultQueryParam, cleanTitle } from "../search/utils";
import { docsBaseURL } from "../utils/environment";

export const GPS_ENDPOINT =
  "https://www.googleapis.com/customsearch/v1/siterestrict?";
export const GPS_ID = "97494f9fe316a426d";
export const MAX_RESULTS_PER_PAGE = 10;
export const MAX_TOTAL_RESULTS = 100;

/**
 * Rewrite link paths for search results
 *
 * These links need to be relative, not absolute,
 * which is how they're returned from Google.
 *
 * @param {string} link
 *   The result URL from Google, which will look something like
 *   https://docs.gitlab.com/ee/update/deprecations.html
 *
 * @return {string}
 *   A relative link for use within the current Docs instance
 */
const rewriteResultLink = (link, query) => {
  const queryParams = searchResultQueryParam(query, link);
  const resultLink = link.replace("https://docs.gitlab.com/", docsBaseURL());
  return `${resultLink}${queryParams}`;
};

export const fetchResults = async (query, filters, pageNumber, numResults) => {
  if (!query || query.length < 2 || typeof query !== "string") {
    return [];
  }

  // Construct the query string for additional filters if needed.
  const filterQuery = filters.length
    ? `+more:pagemap:metatags-gitlab-docs-section:${filters.join(",")}`
    : "";

  const response = await fetch(
    GPS_ENDPOINT +
      new URLSearchParams({
        key: GOOGLE_SEARCH_KEY,
        cx: GPS_ID,
        q: `${query.replaceAll(" ", "+")}${filterQuery}`,
        num: numResults,
        start: (pageNumber - 1) * numResults + 1,
      }),
  );
  const results = await response.json();

  // Clean up the title and link on each result.
  if (results.items) {
    results.items = results.items.map((item) => ({
      ...item,
      formattedTitle: cleanTitle(item.htmlTitle),
      relativeLink: rewriteResultLink(item.link, query),
      breadcrumbs:
        (item.pagemap.metatags &&
          item.pagemap.metatags[0] &&
          item.pagemap.metatags[0]["gitlab-docs-breadcrumbs"]) ||
        "",
    }));
  }

  return results;
};
