import Vue from "vue";
import TabbedContent from "./components/tabbed_content.vue";
import "../assets/css/landing.css";

// Tab render for landing page.
const tabsContainer = document.querySelector('[data-vue-app="docs-tabs"]');
(() =>
  new Vue({
    el: tabsContainer,
    components: {
      TabbedContent,
    },
    render(createElement) {
      let { tabTitles, tabContents } = tabsContainer.dataset;
      tabTitles = tabTitles?.split(",");
      tabContents = tabContents?.split("</ul>,");

      return createElement(TabbedContent, {
        props: {
          tabTitles,
          tabContents,
          responsive: true,
        },
      });
    },
  }))();
