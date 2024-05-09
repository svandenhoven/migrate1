<script>
import { GlButton } from "@gitlab/ui";

export default {
  name: "MenuItem",
  components: {
    GlButton,
  },
  props: {
    item: {
      type: Object,
      required: true,
    },
    level: {
      type: Number,
      default: 1,
      required: true,
    },
  },
  data() {
    return {
      isOpen: this.item.isActiveTrail,
      isMounted: false,
    };
  },
  computed: {
    linkClasses() {
      return `sidebar-link gl-block gl-py-3 gl-pr-5 gl-w-full gl-text-gray-900 hover:gl-text-gray-900 hover:gl-no-underline level-${this.level}`;
    },
    activeClasses() {
      return [
        "gl-rounded",
        this.item.isActiveTrail && "sidebar-link-active-trail",
        this.item.isActive && "sidebar-link-active-item",
      ]
        .filter(Boolean)
        .join(" ");
    },
    shouldAnimate() {
      return this.isMounted && this.item.submenu;
    },
    transitionComponent() {
      return this.shouldAnimate ? "transition" : "div";
    },
  },
  mounted() {
    this.$nextTick(() => {
      this.isMounted = true;
    });

    if (this.item.isActive) {
      this.$refs.activeItem.scrollIntoView({
        behavior: "instant",
        block: "center",
      });
    }
  },
  methods: {
    toggleOpen() {
      this.isOpen = !this.isOpen;
    },
  },
};
</script>

<template>
  <div>
    <a
      v-if="!item.submenu"
      :ref="item.isActiveTrail ? 'activeItem' : null"
      :href="item.prefixedUrl"
      :class="`${linkClasses} ${activeClasses}`"
    >
      {{ item.title }}
    </a>
    <div
      v-else
      :class="`gl-flex gl-justify-between gl-items-center gl-cursor-pointer ${activeClasses}`"
    >
      <a
        :ref="item.isActive ? 'activeItem' : null"
        :href="item.prefixedUrl"
        :class="`${linkClasses} ${activeClasses}`"
      >
        {{ item.title }}
      </a>
      <gl-button
        :icon="isOpen ? 'chevron-down' : 'chevron-right'"
        :aria-label="isOpen ? 'Collapse' : 'Expand'"
        tabindex="0"
        class="submenu-toggle gl-mr-4 !gl-shadow-none !gl-p-0"
        @click="toggleOpen"
      />
    </div>
    <component :is="transitionComponent" name="submenu">
      <div v-if="item.submenu && isOpen" class="submenu">
        <menu-item
          v-for="(child, childIndex) in item.submenu"
          :key="childIndex"
          :item="child"
          :level="level + 1"
        />
      </div>
    </component>
  </div>
</template>
