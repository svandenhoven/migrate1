/**
 * @jest-environment jsdom
 */
import { mount } from "@vue/test-utils";
import flushPromises from "flush-promises";
import {
  mockResults,
  mockNoResults,
  mockHistoryCookie,
} from "../__mocks__/search_results_mock";
import SearchForm from "../../themes/gitlab-docs/src/components/search/google_search_form.vue";
import {
  trackPageHistory,
  RECENT_HISTORY_ITEMS,
} from "../../themes/gitlab-docs/src/search/history";
import {
  getCookie,
  setCookie,
} from "../../themes/gitlab-docs/src/utils/cookies";
import { fetchResults } from "../../themes/gitlab-docs/src/services/google_search_api";
import { createContainer, setMetatag } from "../helpers/jest_helpers";

jest.mock("../../themes/gitlab-docs/src/services/google_search_api", () => ({
  fetchResults: jest.fn(),
  MAX_RESULTS_PER_PAGE: 10,
}));

describe("themes/gitlab-docs/src/components/search/google_search_form.vue", () => {
  let wrapper;

  // Use fake timers to mock debounce behavior.
  beforeAll(() => {
    jest.useFakeTimers();
  });
  afterAll(() => {
    jest.runAllTimers();
  });

  beforeEach(() => {
    jest.clearAllMocks();

    setMetatag("gitlab-docs-base-url", "http://localhost/");

    // Add a container around the mounted component.
    // We need this to avoid tooltip errors from BootstrapVue.
    const componentData = {
      attachTo: createContainer(),
      propsData: {
        numResults: 10,
      },
    };
    wrapper = mount(SearchForm, componentData);
  });

  afterEach(() => {
    document.querySelector('meta[name="gitlab-docs-base-url"]').remove();
    wrapper.destroy();
  });

  it("runs a search when the user types in a query", async () => {
    fetchResults.mockResolvedValueOnce(mockResults);

    const input = wrapper.find("input");
    input.setValue("how does jest work");
    await input.trigger("keyup");
    jest.advanceTimersByTime(500); // debounce

    expect(fetchResults).toHaveBeenCalledTimes(1);
  });

  it('displays "No results found" message when there are no search results', async () => {
    fetchResults.mockResolvedValueOnce(mockNoResults);

    const input = wrapper.find("input");
    input.setValue("non-existent query");
    await input.trigger("keyup");
    jest.advanceTimersByTime(500);
    await flushPromises();

    expect(fetchResults).toHaveBeenCalledTimes(1);
    expect(wrapper.find('[data-testid="no-results"]').exists()).toBe(true);
  });

  it('displays "See all results" link when there are more results than shown', async () => {
    fetchResults.mockResolvedValueOnce(mockResults);

    const input = wrapper.find("input");
    input.setValue("test");
    await input.trigger("keyup");
    jest.advanceTimersByTime(500);
    await flushPromises();

    expect(wrapper.vm.hasMoreResults).toBe(true);

    const moreResultsLink = wrapper.find('[data-testid="more-results"]');
    expect(moreResultsLink.exists()).toBe(true);
    expect(moreResultsLink.attributes("href")).toBe(
      "http://localhost/search.html/?q=test",
    );
  });
});

describe("themes/gitlab-docs/src/search/history", () => {
  afterEach(() => {
    // Delete the cookie after each test
    document.cookie =
      "pageHistory=; expires=Mon, 12 June 2023 00:00:00 UTC; path=/;";
  });

  it("should set a cookie with the current page URL and title", () => {
    // Set up the DOM
    document.title = "Test Page | GitLab";
    const location = {
      ...window.location,
      pathname: "/test-page",
    };
    Object.defineProperty(window, "location", {
      writable: true,
      value: location,
    });

    trackPageHistory();

    // Check that the cookie was set correctly
    const cookieValue = getCookie("pageHistory");
    expect(cookieValue).not.toBeNull();
    const pageHistory = JSON.parse(cookieValue);
    expect(pageHistory).toHaveLength(1);
    expect(pageHistory[0].path).toBe("/test-page");
    expect(pageHistory[0].title).toBe("Test Page");
  });

  it("should limit the number of items in the history to RECENT_HISTORY_ITEMS", () => {
    document.title = "Test Page | GitLab";

    // Set a cookie with RECENT_HISTORY_ITEMS pages in it, then track this page
    setCookie("pageHistory", JSON.stringify(mockHistoryCookie), 365);
    trackPageHistory();

    // Check that the cookie still contains RECENT_HISTORY_ITEMS
    const cookieValue = getCookie("pageHistory");
    expect(cookieValue).not.toBeNull();
    const pageHistory = JSON.parse(cookieValue);
    expect(pageHistory).toHaveLength(RECENT_HISTORY_ITEMS);
  });

  it("should not add duplicate history items", () => {
    document.title = "Test Page | GitLab";

    // Set a cookie with the current page URL
    const initialPageHistory = [{ path: "/test-page", title: "Test Page" }];
    setCookie("pageHistory", JSON.stringify(initialPageHistory), 365);

    trackPageHistory();

    // Check that the cookie was updated correctly, with one instance of Test Page
    const cookieValue = getCookie("pageHistory");
    expect(cookieValue).not.toBeNull();
    const pageHistory = JSON.parse(cookieValue);
    expect(pageHistory).toHaveLength(1);
    expect(pageHistory[0].path).toBe("/test-page");
    expect(pageHistory[0].title).toBe("Test Page");
  });
});
