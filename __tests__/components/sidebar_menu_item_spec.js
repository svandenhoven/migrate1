/**
 * @jest-environment jsdom
 */

import { mount } from "@vue/test-utils";
import SidebarMenuItem from "../../themes/gitlab-docs/src/components/sidebar_menu_item.vue";

describe("component: SidebarMenuItem", () => {
  let wrapper;
  let mockItem;

  beforeEach(() => {
    mockItem = {
      title: "Test Item",
      url: "test-item",
      submenu: [
        {
          title: "Test Subitem",
          url: "test-subitem",
        },
      ],
    };

    wrapper = mount(SidebarMenuItem, {
      propsData: {
        item: mockItem,
        level: 1,
      },
    });

    // Mock the scrollIntoView method
    Element.prototype.scrollIntoView = jest.fn();
  });

  it("toggles a submenu when the button is clicked", async () => {
    const submenuToggle = wrapper.find(".submenu-toggle");
    expect(wrapper.find(".submenu").exists()).toBe(false);

    await submenuToggle.trigger("click");
    expect(wrapper.find(".submenu").exists()).toBe(true);

    await submenuToggle.trigger("click");
    expect(wrapper.find(".submenu").exists()).toBe(false);
  });
});
