<script>
import { GlDisclosureDropdown, GlDisclosureDropdownGroup } from "@gitlab/ui";
import { compareVersions } from "compare-versions";
import { getVersions, getArchivesVersions } from "../services/versions";
import { isProduction, relativeCurrentPath } from "../helpers/environment";
import { HUGO_LAUNCH_VERSION } from "../helpers/constants";

export default {
  components: {
    GlDisclosureDropdown,
    GlDisclosureDropdownGroup,
  },
  data() {
    return {
      versions: {},
      archiveVersions: {},
      menuItems: [],
    };
  },
  computed: {
    generatedMenuItems() {
      return [
        // Group 1
        {
          items: [
            {
              text: `${this.versions.next} (not yet released)`,
              href: this.versionedPagePath(this.activeVersion),
              extraAttrs: {
                "data-testid": "next-version",
              },
            },
          ],
        },
        // Group 2
        {
          items: [
            {
              text: `${this.versions.current} (recently released)`,
              href: this.versionedPagePath(this.versions.current),
              extraAttrs: {
                "data-testid": "current-version",
              },
            },
            ...this.versions.last_minor.map((version) => ({
              text: version,
              href: this.versionedPagePath(version),
            })),
          ],
        },
        // Group 3
        {
          items: [
            ...this.versions.last_major.map((version) => ({
              text: version,
              href: this.versionedPagePath(version),
            })),
          ],
        },
      ];
    },
    activeVersion() {
      return document.querySelector('meta[name="gitlab-docs-version"]')
        ?.content;
    },
    legacyPath() {
      return document.querySelector('meta[name="gitlab-docs-legacy-path"]')
        ?.content;
    },
  },
  async created() {
    // If we're on the production domain, populate the menu with links to older versions of this page.
    // Archived and self-hosted instances only show the link to the Archives page.
    if (isProduction()) {
      try {
        this.versions = await getVersions();
        this.archiveVersions = await getArchivesVersions();
        this.menuItems = this.generatedMenuItems;
      } catch (err) {
        console.error(`Failed to fetch versions.json: ${err}`);
      }
    }

    this.menuItems.push({
      items: [{ text: "Archives", href: "https://docs.gitlab.com/archives" }],
    });
  },
  methods: {
    // Returns the path to the current page for a given version
    versionedPagePath(versionNumber) {
      let path = relativeCurrentPath();

      // Temporary handling for pre-Hugo paths
      // We can drop this when the menu no longer includes Nanoc versions of pages
      if (
        this.legacyPath &&
        compareVersions(versionNumber, HUGO_LAUNCH_VERSION) < 0
      ) {
        path = this.legacyPath;
      }

      // If we're viewing a released version,
      // do not include its version prefix when creating the other links.
      if (this.activeVersion !== this.versions.next) {
        path = `/${path
          .split("/")
          .filter((n) => n)
          .slice(1) // Drop the active version prefix
          .join("/")}/`;
      }

      if (versionNumber && versionNumber !== this.versions.next) {
        path = `/${versionNumber}${path}`;
      }

      if (this.archiveVersions.includes(versionNumber)) {
        path = `https://archives.docs.gitlab.com${path}`;
      } else {
        path = `https://docs.gitlab.com${path}`;
      }

      return path;
    },
  },
};
</script>

<template>
  <gl-disclosure-dropdown
    :toggle-text="`v${activeVersion}`"
    class="md:gl-mr-4 md:gl-ml-3"
    data-testid="versions-menu"
  >
    <gl-disclosure-dropdown-group
      v-for="(group, index) in menuItems"
      :key="index"
      :group="group"
    />
  </gl-disclosure-dropdown>
</template>
