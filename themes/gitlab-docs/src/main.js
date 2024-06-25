import Vue from "vue";
import tocbot from "tocbot";
import { expandCollapse } from "./utilities/collapse";
import SurveyBanner from "./components/survey_banner.vue";
import SidebarMenu from "./components/sidebar_menu.vue";
import VersionsMenu from "./components/versions_menu.vue";
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
  expandCollapse();
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
    ignoreSelector: ".help-feedback h2, .help-feedback h3, .archives-list h3",
    collapseDepth: 3,
    // These should match our header-offset CSS variable.
    // header-offset = 4rem = 64px
    headingsOffset: 64,
    scrollSmoothOffset: -64,
  });

  // Initialize Vue components
  // Survey banner
  const bannerContainer = document.querySelector(
    '[data-vue-app="survey-banner"]',
  );
  (() =>
    new Vue({
      el: bannerContainer,
      components: {
        SurveyBanner,
      },
      render(createElement) {
        return createElement(SurveyBanner, {
          props: { text: bannerContainer.innerHTML, variant: "info" },
        });
      },
    }))();

  // Sidebar menu
  const menuContainer = document.querySelector('[data-vue-app="sidebar-menu"]');
  (() =>
    new Vue({
      el: menuContainer,
      components: {
        SidebarMenu,
      },
      render(createElement) {
        return createElement(SidebarMenu, {
          props: {
            baseUrl: document
              .querySelector('meta[name="gitlab-docs-base-url"]')
              ?.getAttribute("content"),
          },
        });
      },
    }))();

  // Versions menu
  const versionsContainer = document.querySelector(
    '[data-vue-app="versions-menu"]',
  );
  (() =>
    new Vue({
      el: versionsContainer,
      components: {
        VersionsMenu,
      },
      render(createElement) {
        return createElement(VersionsMenu);
      },
    }))();
});
