<script>
import { GlButton } from "@gitlab/ui";
import menuDataRaw from "../../../../data/navigation.yaml";
import MenuItem from "./sidebar_menu_item.vue";

export default {
  components: {
    GlButton,
    MenuItem,
  },
  props: {
    baseUrl: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      menuItems: [],
      isCollapsed: false,
    };
  },
  created() {
    const { menu } = this.processMenuData(menuDataRaw);
    this.menuItems = Object.freeze(menu);
  },
  methods: {
    toggleCollapse() {
      this.isCollapsed = !this.isCollapsed;

      // Set a class on the page wrapper.
      // This allows us to modify the grid dimensions
      const container = document.querySelector(".template-single");
      if (container) {
        container.classList.toggle(
          "template-single-sidebar-collapsed",
          this.isCollapsed,
        );
      }
    },
    isActiveItem(item) {
      const currentPath = window.location.pathname;
      const relativeCurrentPath = currentPath.replace(this.baseUrl, "");

      return (
        relativeCurrentPath.endsWith(`/${item.url}`) ||
        relativeCurrentPath.endsWith(`/${item.url}/`)
      );
    },
    processMenuData(currMenu) {
      let isActiveTrail = false;

      // Add attributes to link items
      const menu = currMenu.map((item) => {
        const { url, submenu } = item;

        // Add the appropriate base URL to internal links
        const prefixedUrl = url.toLowerCase().startsWith("http")
          ? url
          : `${this.baseUrl}${url}`;

        let itemActiveTrail = false;
        let isActive = false;

        if (this.isActiveItem(item)) {
          itemActiveTrail = true;
          isActive = true;
        }

        // Run this recursively for submenus
        let processedSubmenu = submenu;
        if (Array.isArray(submenu)) {
          const { menu: submenuItems, isActiveTrail: submenuTrail } =
            this.processMenuData(submenu);
          processedSubmenu = submenuItems;
          itemActiveTrail = itemActiveTrail || submenuTrail;
        }

        isActiveTrail = isActiveTrail || itemActiveTrail;

        return {
          ...item,
          prefixedUrl,
          isActive,
          submenu: processedSubmenu,
          isActiveTrail: itemActiveTrail,
        };
      });

      return { menu, isActiveTrail };
    },
  },
};
</script>

<template>
  <aside
    class="global-nav-wrapper gl-leading-5"
    :class="{ 'sidebar-collapsed': isCollapsed }"
  >
    <div class="global-nav gl-fixed gl-overflow-y-scroll gl-px-2">
      <nav class="gl-my-3" aria-label="Main">
        <menu-item
          v-for="(item, index) in menuItems"
          :key="index"
          :item="item"
          :level="1"
        />
      </nav>
    </div>

    <div
      class="sidebar-toggle gl-bottom-0 gl-fixed gl-cursor-pointer hover:gl-bg-gray-50"
      @click="toggleCollapse"
    >
      <div class="gl-h-full gl-flex gl-items-center gl-ml-3">
        <gl-button
          :icon="
            isCollapsed ? 'chevron-double-lg-right' : 'chevron-double-lg-left'
          "
          class="!gl-shadow-none !gl-p-0"
          :class="isCollapsed ? 'sidebar-toggle-collapsed' : 'gl-mr-2'"
        >
          <template v-if="!isCollapsed">Collapse sidebar</template>
        </gl-button>
      </div>
    </div>
  </aside>
</template>
