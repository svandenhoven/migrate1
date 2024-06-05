/**
 * @jest-environment jsdom
 */

import { mount } from "@vue/test-utils";
import flushPromises from "flush-promises";
import VersionsMenu from "../../themes/gitlab-docs/src/components/versions_menu.vue";
import * as versions from "../../themes/gitlab-docs/src/services/versions";
import * as environment from "../../themes/gitlab-docs/src/helpers/environment";
import { mockVersions, mockArchiveVersions } from "../__mocks__/versions_mock";
import { setWindowPath, setVersionMetatag } from "../helpers/versions_helper";

describe("component: VersionsMenu", () => {
  const productionUrl = "https://docs.gitlab.com";
  const archivesUrl = "https://archives.docs.gitlab.com";

  beforeEach(() => {
    jest.clearAllMocks();
    jest.spyOn(environment, "isProduction").mockReturnValue(true);

    // Mock data from external JSON files
    jest.spyOn(versions, "getVersions").mockReturnValue(mockVersions);
    jest
      .spyOn(versions, "getArchivesVersions")
      .mockReturnValue(mockArchiveVersions);
  });
  afterEach(() => {
    document.querySelector('meta[name="gitlab-docs-version"]').remove();
  });

  it("Fetches versions.json and displays current version", async () => {
    const wrapper = mount(VersionsMenu);
    setVersionMetatag(mockVersions.next);
    await flushPromises();

    expect(versions.getVersions).toHaveBeenCalledTimes(1);
    expect(versions.getArchivesVersions).toHaveBeenCalledTimes(1);

    const nextVersion = wrapper.find('[data-testid="next-version"]').text();
    expect(nextVersion).toContain(mockVersions.next);
  });

  it("Generates correct menu links from the homepage", async () => {
    setWindowPath("/");
    setVersionMetatag(mockVersions.next);

    const wrapper = mount(VersionsMenu);
    await wrapper.setData({
      versions: mockVersions,
    });

    expect(wrapper.vm.versionedPagePath("")).toBe(`${productionUrl}/`);
    expect(wrapper.vm.versionedPagePath(mockVersions.current)).toBe(
      `${productionUrl}/${mockVersions.current}/`,
    );

    mockVersions.last_minor.forEach((v) => {
      expect(wrapper.vm.versionedPagePath(v)).toBe(`${productionUrl}/${v}/`);
    });

    expect(wrapper.vm.versionedPagePath(mockVersions.last_major[0])).toBe(
      `${productionUrl}/16.11/`,
    );
    expect(wrapper.vm.versionedPagePath(mockVersions.last_major[1])).toBe(
      `${archivesUrl}/15.11/`,
    );
  });

  it("Generates correct menu links from an interior page", async () => {
    setWindowPath("/ee/user/project/issue_board");
    setVersionMetatag(mockVersions.next);

    const wrapper = mount(VersionsMenu);
    await wrapper.setData({ versions: mockVersions });

    expect(wrapper.vm.versionedPagePath("")).toBe(
      `${productionUrl}/ee/user/project/issue_board`,
    );
    expect(wrapper.vm.versionedPagePath(mockVersions.current)).toBe(
      `${productionUrl}/${mockVersions.current}/ee/user/project/issue_board`,
    );

    mockVersions.last_minor.forEach((v) => {
      expect(wrapper.vm.versionedPagePath(v)).toBe(
        `${productionUrl}/${v}/ee/user/project/issue_board`,
      );
    });

    expect(wrapper.vm.versionedPagePath(mockVersions.last_major[0])).toBe(
      `${productionUrl}/16.11/ee/user/project/issue_board`,
    );
    expect(wrapper.vm.versionedPagePath(mockVersions.last_major[1])).toBe(
      `${archivesUrl}/15.11/ee/user/project/issue_board`,
    );
  });

  it("Generates correct menu links from an older version on docs.gitlab.com", async () => {
    setWindowPath("/15.2/runner");
    setVersionMetatag("15.2");

    const wrapper = mount(VersionsMenu);
    await wrapper.setData({
      versions: mockVersions,
      archiveVersions: mockArchiveVersions,
    });

    expect(wrapper.vm.versionedPagePath("")).toBe(`${productionUrl}/runner`);
    expect(wrapper.vm.versionedPagePath(mockVersions.current)).toBe(
      `${productionUrl}/${mockVersions.current}/runner`,
    );

    mockVersions.last_minor.forEach((v) => {
      expect(wrapper.vm.versionedPagePath(v)).toBe(
        `${productionUrl}/${v}/runner`,
      );
    });
    expect(wrapper.vm.versionedPagePath(mockVersions.last_major[0])).toBe(
      `${productionUrl}/16.11/runner`,
    );
    expect(wrapper.vm.versionedPagePath(mockVersions.last_major[1])).toBe(
      `${archivesUrl}/15.11/runner`,
    );
  });

  it("Shows simplified menu on non-production sites", async () => {
    jest.spyOn(environment, "isProduction").mockReturnValue(false);
    setVersionMetatag("14.10");
    const wrapper = mount(VersionsMenu);
    expect(
      wrapper.find('[data-testid="versions-menu"] a:nth-child(2)').exists(),
    ).toBe(false);
  });

  it("Uses legacyPath for versions older than HUGO_LAUNCH_VERSION", async () => {
    setWindowPath("/user/project/issue_board");
    setVersionMetatag(mockVersions.next);

    // Set a legacy path meta tag
    const legacyPath = "/ee/user/project/issue_board.html";
    const metaTag = document.createElement("meta");
    metaTag.setAttribute("name", "gitlab-docs-legacy-path");
    metaTag.setAttribute("content", legacyPath);
    document.head.appendChild(metaTag);

    const wrapper = mount(VersionsMenu);
    await wrapper.setData({ versions: mockVersions });

    // Older than HUGO_LAUNCH_VERSION
    const olderVersion = "16.10";
    expect(wrapper.vm.versionedPagePath(olderVersion)).toBe(
      `${productionUrl}/${olderVersion}${legacyPath}`,
    );
    // Equal to HUGO_LAUNCH_VERSION
    const equalVersion = "18.0";
    expect(wrapper.vm.versionedPagePath(equalVersion)).toBe(
      `${productionUrl}/${equalVersion}/user/project/issue_board`,
    );
    // Newer than HUGO_LAUNCH_VERSION
    const newerVersion = "18.1";
    expect(wrapper.vm.versionedPagePath(newerVersion)).toBe(
      `${productionUrl}/${newerVersion}/user/project/issue_board`,
    );

    // Clean up the meta tag
    document.head.removeChild(metaTag);
  });
});
