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
    // Create a mock container element
    const container = document.createElement("div");
    container.classList.add("template-single");
    document.body.appendChild(container);

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
    const currentPath = "/topics/git/clone.html";
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

  it("correctly identifies active item with trailing slash", async () => {
    Object.defineProperty(window, "location", {
      value: { pathname: "/topics/git/" },
      writable: true,
    });

    await wrapper.vm.$nextTick();
    wrapper = mount(SidebarMenu, { propsData: { baseUrl } });

    const activeItem = wrapper.find(".sidebar-link-active-item");
    expect(activeItem.exists()).toBe(true);
    expect(activeItem.text()).toContain("Learn Git");
  });

  it("correctly identifies active item without a trailing slash", async () => {
    Object.defineProperty(window, "location", {
      value: { pathname: "/topics/git" },
      writable: true,
    });

    await wrapper.vm.$nextTick();
    wrapper = mount(SidebarMenu, { propsData: { baseUrl } });

    const activeItem = wrapper.find(".sidebar-link-active-item");
    expect(activeItem.exists()).toBe(true);
    expect(activeItem.text()).toContain("Learn Git");
  });

  it("uses the correct URL prefix for relative links", () => {
    const relativeUrlItem = wrapper.vm.menuItems.find(
      (item) => item.url === "user.html",
    );
    expect(relativeUrlItem.prefixedUrl).toBe(`${baseUrl}user.html`);
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
      (item) => item.url === "user.html",
    );
    expect(relativeUrlItem.prefixedUrl).toBe(`${subdirBaseUrl}user.html`);
  });

  it("toggles the overlay state when the hamburger button is clicked", async () => {
    const hamburgerButton = wrapper.find('[data-testid="hamburger-icon"]');
    expect(wrapper.vm.isOverlayOpen).toBe(false);

    await hamburgerButton.trigger("click");
    expect(wrapper.vm.isOverlayOpen).toBe(true);
    expect(wrapper.find(".global-nav-wrapper").classes()).not.toContain(
      "sidebar-collapsed",
    );

    await wrapper.find(".modal-backdrop").trigger("click");
    expect(wrapper.vm.isOverlayOpen).toBe(false);
  });

  it("closes the overlay when the window is resized larger than the mobile breakpoint", async () => {
    // Set the initial window width below the mobile breakpoint
    window.innerWidth = 768;

    // Open the overlay
    wrapper.vm.toggleOverlay("open");
    expect(wrapper.vm.isOverlayOpen).toBe(true);

    // Resize the window larger than the mobile breakpoint
    window.innerWidth = 1024;
    window.dispatchEvent(new Event("resize"));
    await wrapper.vm.$nextTick();

    expect(wrapper.vm.isOverlayOpen).toBe(false);
  });
});
