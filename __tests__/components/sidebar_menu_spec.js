/**
 * @jest-environment jsdom
 */

/* eslint-disable global-require */

import { mount } from "@vue/test-utils";
import SidebarMenu from "../../themes/gitlab-docs/src/components/sidebar_menu.vue";

jest.mock("../../data/navigation.yaml", () => {
  return require("../__mocks__/navigation_mock").default;
});

describe("component: SidebarMenu", () => {
  let wrapper;
  let mockData;

  const baseUrl = "https://docs.gitlab.com/";

  beforeEach(() => {
    wrapper = mount(SidebarMenu, {
      propsData: {
        baseUrl,
      },
    });
    mockData = require("../__mocks__/navigation_mock").default;

    // Mock the scrollIntoView method
    Element.prototype.scrollIntoView = jest.fn();
  });

  it("renders menu items based on the provided data", () => {
    const menuItems = wrapper.findAll(".sidebar-link");
    expect(menuItems.length).toBe(mockData.length);
  });

  it("toggles the sidebar collapsed state when the toggle button is clicked", async () => {
    const toggleButton = wrapper.find(".sidebar-toggle");
    expect(wrapper.find(".global-nav-wrapper").classes()).not.toContain(
      "sidebar-collapsed",
    );

    await toggleButton.trigger("click");
    expect(wrapper.find(".global-nav-wrapper").classes()).toContain(
      "sidebar-collapsed",
    );

    await toggleButton.trigger("click");
    expect(wrapper.find(".global-nav-wrapper").classes()).not.toContain(
      "sidebar-collapsed",
    );
  });

  it("applies the active trail class to the correct menu items based on the current URL", async () => {
    // Mock the window.location.pathname
    const currentPath = "/topics/git/clone/";
    Object.defineProperty(window, "location", {
      value: {
        pathname: currentPath,
      },
      writable: true,
    });

    // Remount the component to trigger the created hook with the mocked pathname
    await wrapper.vm.$nextTick();
    wrapper = mount(SidebarMenu, {
      propsData: {
        baseUrl,
      },
    });

    // Find the menu items with the active trail class
    const activeTrailItems = wrapper.findAll("a.sidebar-link-active-trail");

    // Assert that the correct menu items have the active trail class
    expect(activeTrailItems.length).toBe(3);
    expect(activeTrailItems.at(0).text()).toContain("Use GitLab");
    expect(activeTrailItems.at(1).text()).toContain("Learn Git");
    expect(activeTrailItems.at(2).text()).toContain("Clone a Git repository");

    // Assert that the active item has the correct class
    const activeItem = wrapper.find(".sidebar-link-active-item");
    expect(activeItem.exists()).toBe(true);
    expect(activeItem.text()).toContain("Clone a Git repository");
  });

  it("uses the correct URL prefix for relative links", () => {
    const relativeUrlItem = wrapper.vm.menuItems.find(
      (item) => item.url === "user/",
    );
    expect(relativeUrlItem.prefixedUrl).toBe(`${baseUrl}user/`);
  });

  it("uses the correct URL prefix for external links", () => {
    const externalUrlItem = wrapper.vm.menuItems.find(
      (item) => item.url === "https://design.gitlab.com",
    );
    expect(externalUrlItem.prefixedUrl).toBe("https://design.gitlab.com");
  });

  it("uses the correct URL prefix for relative links when the site is in a subdirectory", () => {
    const subdirBaseUrl = "https://docs.gitlab.com/17.0";

    wrapper = mount(SidebarMenu, {
      propsData: {
        baseUrl: subdirBaseUrl,
      },
    });

    const relativeUrlItem = wrapper.vm.menuItems.find(
      (item) => item.url === "user/",
    );
    expect(relativeUrlItem.prefixedUrl).toBe(`${subdirBaseUrl}user/`);
  });
});
