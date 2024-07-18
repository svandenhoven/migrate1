/**
 * @jest-environment jsdom
 */

import { shallowMount } from "@vue/test-utils";
import { GlIcon } from "@gitlab/ui";
import ClipboardCopy from "../../themes/gitlab-docs/src/components/clipboard_copy.vue";
import { createContainer } from "../helpers/jest_helpers";

describe("component: CopyButton", () => {
  const propsData = { codeContent: "Sample code" };
  let wrapper;

  beforeEach(() => {
    // Add a container around the mounted component.
    // We need this to avoid tooltip errors from BootstrapVue.
    const componentData = {
      attachTo: createContainer(),
      propsData,
      stubs: {
        GlIcon: true,
      },
    };

    wrapper = shallowMount(ClipboardCopy, componentData);
  });

  it("renders a copy button", () => {
    expect(wrapper.find(".docs-codeblock-toolbar").exists()).toBe(true);
  });

  it("renders a GlIcon component", () => {
    expect(wrapper.findComponent(GlIcon).exists()).toBe(true);
  });

  it("has a screen reader text for the button", () => {
    expect(wrapper.find(".sr-only").text()).toBe("Copy to clipboard");
  });

  it("updates copiedStatus when copy is called", async () => {
    Object.assign(navigator, {
      clipboard: {
        writeText: jest.fn().mockImplementation(() => Promise.resolve()),
      },
    });

    await wrapper.vm.copy();
    expect(wrapper.vm.copiedStatus).toBe("Copied!");
  });

  it("resets copiedStatus after a delay", async () => {
    jest.useFakeTimers();
    Object.assign(navigator, {
      clipboard: {
        writeText: jest.fn().mockImplementation(() => Promise.resolve()),
      },
    });

    await wrapper.vm.copy();
    jest.advanceTimersByTime(1500);
    expect(wrapper.vm.copiedStatus).toBe("");
  });

  it("displays the correct copiedStatus in the aria-live region", async () => {
    await wrapper.setData({ copiedStatus: "Copied!" });
    expect(wrapper.find('[aria-live="polite"]').text()).toBe("Copied!");
  });
});
