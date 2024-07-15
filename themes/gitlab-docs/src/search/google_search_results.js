import Vue from "vue";
import GoogleResults from "../components/search/google_results.vue";
import { activateKeyboardShortcut } from "./utils";

document.addEventListener("DOMContentLoaded", () => {
  activateKeyboardShortcut();

  (() =>
    new Vue({
      el: ".js-search-results",
      components: {
        GoogleResults,
      },
      render(createElement) {
        return createElement(GoogleResults);
      },
    }))();
});
