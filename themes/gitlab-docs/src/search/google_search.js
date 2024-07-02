import Vue from "vue";
import GoogleSearchForm from "../components/search/google_search_form.vue";
import { activateKeyboardShortcut, scrollToQuery } from "./utils";

document.addEventListener("DOMContentLoaded", () => {
  scrollToQuery();
  activateKeyboardShortcut();

  (() =>
    new Vue({
      el: ".js-google-search-form",
      components: { GoogleSearchForm },
      render(createElement) {
        return createElement(GoogleSearchForm, {
          props: {
            numResults: 7,
          },
        });
      },
    }))();
});
