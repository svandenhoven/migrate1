import Vue from "vue";
import tocbot from "tocbot";
import DocsBanner from "./components/survey_banner.vue";
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
   * Set up the table of contents
   * @see https://tscanlin.github.io/tocbot/
   */
  tocbot.init({
    tocSelector: ".js-toc",
    positionFixedSelector: ".js-toc",
    contentSelector: ".docs-content",
    headingSelector: "h2, h3, h4",
    ignoreSelector: ".help-feedback h2, .help-feedback h3",
    collapseDepth: 3,
    // These should match our header-offset CSS variable.
    // header-offset = 4rem = 64px
    headingsOffset: 64,
    scrollSmoothOffset: -64,
  });

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
