/**
 * @jest-environment jsdom
 */

import { shallowMount } from "@vue/test-utils";
import SurveyBanner from "../../themes/gitlab-docs/src/components/survey_banner.vue";

describe("component: Survey banner", () => {
  const propsData = { text: "Some text", variant: "info" };
  let wrapper;

  beforeEach(() => {
    wrapper = shallowMount(SurveyBanner, { propsData });
  });

  it("renders a banner", () => {
    expect(wrapper.exists(".banner")).toBe(true);
  });

  it("renders the correct banner text", () => {
    const bannerText = wrapper.find("div");
    expect(bannerText.text()).toEqual(propsData.text);
  });
});
