<script>
import { GlIcon, GlTooltipDirective as GlTooltip } from "@gitlab/ui";

export default {
  directives: {
    GlTooltip,
  },
  components: {
    GlIcon,
  },
  props: {
    codeContent: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      copiedStatus: "",
    };
  },
  methods: {
    async copy() {
      // Hide the hover tooltip
      // https://bootstrap-vue.org/docs/directives/tooltip#hiding-and-showing-tooltips-via-root-events
      this.$root.$emit("bv::hide::tooltip");

      try {
        this.copiedStatus = "Copied!";
        await window.navigator.clipboard.writeText(this.codeContent);
        setTimeout(() => {
          // Hide the "Copied!" tooltip after 1.5s
          this.copiedStatus = "";
          this.$root.$emit("bv::hide::tooltip");
        }, 1500);
      } catch (err) {
        console.error("Failed to copy: ", err);
      }
    },
  },
};
</script>

<template>
  <div
    class="docs-codeblock-toolbar gl-flex gl-justify-end gl-bg-gray-100 gl-py-1 gl-pr-3"
  >
    <span v-gl-tooltip.click :title="copiedStatus" class="gl-cursor-pointer">
      <gl-icon
        v-gl-tooltip.focus.hover
        name="copy-to-clipboard"
        aria-label="Copy to clipboard"
        class="clipboard-icon hover:gl-fill-purple-600 focus:gl-fill-purple-600 gl-transition-colors"
        tabindex="0"
        title="Copy"
        @click="copy()"
        @keydown.enter="copy()"
      />
      <span class="sr-only">Copy to clipboard</span>
      <span aria-live="polite" class="sr-only">
        {{ copiedStatus }}
      </span>
    </span>
  </div>
</template>
