import Vue from "vue";
import DocsBanner from "./components/survey_banner.vue";

import "../assets/css/tailwind.css";
import "../assets/css/variables.css";
import "../assets/css/typography.css";
import "../assets/css/layout.css";
import "../assets/css/main.css";

/**
 * Adds a clickable permalink to each content heading
 */
const addHeaderPermalinks = () => {
  document
    .querySelectorAll(".docs-content h2, h3, h4, h5, h6")
    .forEach((heading) => {
      const { id } = heading;
      if (!id) return;

      const anchor = document.createElement("a");
      anchor.href = `#${id}`;
      anchor.title = "Permalink";
      anchor.classList.add("anchor");

      heading.appendChild(anchor);
    });
};

document.addEventListener("DOMContentLoaded", () => {
  addHeaderPermalinks();

  /**
   * Initialize Vue SFCs
   */
  const surveyBanner = document.querySelector('[data-vue-app="survey-banner"]');
  (() =>
    new Vue({
      el: surveyBanner,
      components: {
        DocsBanner,
      },
      render(createElement) {
        return createElement(DocsBanner, {
          props: { text: surveyBanner.innerHTML, variant: "info" },
        });
      },
    }))();
});
