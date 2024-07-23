<script>
import { GlButton } from "@gitlab/ui";
import menuDataRaw from "../../../../data/navigation.yaml";
import { NAV_BREAKPOINT } from "../utils/constants";
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
      isOverlayOpen: false,
    };
  },
  created() {
    const { menu } = this.processMenuData(menuDataRaw);
    this.menuItems = Object.freeze(menu);
  },
  mounted() {
    window.addEventListener("resize", this.handleResize);
    this.initWidth = window.innerWidth;
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.handleResize);
  },
  methods: {
    toggleCollapse(action) {
      const container = document.querySelector(".template-single");

      if (action === "open" || action === "close") {
        this.isCollapsed = action === "close";
      } else {
        this.isCollapsed = !this.isCollapsed;
      }

      container.classList.toggle(
        "template-single-sidebar-collapsed",
        this.isCollapsed,
      );
    },
    toggleOverlay(action) {
      this.isOverlayOpen = action === "open";
      if (this.isOverlayOpen) {
        this.toggleCollapse("open");
      }
    },
    handleResize() {
      // Close the overlay if we've resized the window larger than the mobile breakpoint
      if (this.isOverlayOpen && window.innerWidth >= NAV_BREAKPOINT) {
        this.toggleOverlay("close");
      }
    },
    isActiveItem(item) {
      const currentPath = window.location.pathname;
      // Allow for matching extensionless URLs.
      // GitLab Pages will serve a page like /user.html at
      // both /user/ and /user.html.
      const normalizedPath = currentPath.replace(/\/$/, ".html");
      const itemUrl = `/${item.url}`;

      return [
        currentPath,
        `${currentPath}/`,
        normalizedPath,
        normalizedPath.replace(".html", ""),
      ].some((path) => path.endsWith(itemUrl));
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
  <div class="sidebar-container">
    <transition name="fade">
      <div
        v-if="isOverlayOpen"
        class="lg:gl-hidden gl-block gl-z-2 modal-backdrop"
        aria-hidden="true"
        @click="toggleOverlay('close')"
      ></div>
    </transition>
    <gl-button
      icon="hamburger"
      class="lg:gl-hidden gl-mt-6"
      aria-label="Toggle sidebar"
      @click="toggleOverlay('open')"
    />

    <aside
      class="global-nav-wrapper gl-leading-5"
      :class="{
        'sidebar-collapsed': isCollapsed,
        'sidebar-overlay-open !gl-block': isOverlayOpen,
      }"
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
        v-on="{ click: isOverlayOpen ? toggleOverlay : toggleCollapse }"
      >
        <div class="gl-h-full gl-flex gl-items-center gl-ml-3">
          <gl-button
            :icon="
              isCollapsed ? 'chevron-double-lg-right' : 'chevron-double-lg-left'
            "
            class="!gl-shadow-none !gl-p-0"
            :class="{ 'sidebar-toggle-collapsed': isCollapsed }"
          >
            <span v-if="!isCollapsed" class="gl-ml-2">Collapse sidebar</span>
          </gl-button>
        </div>
      </div>
    </aside>
  </div>
</template>
