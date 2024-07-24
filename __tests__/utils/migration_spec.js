import {
  rewriteLegacyURL,
  getRedirectUrl,
} from "../../themes/gitlab-docs/src/utils/migration";

describe("rewriteLegacyURL", () => {
  it("removes ee/ from the beginning of the path and adds .html", () => {
    const input = "https://example.com/ee/foo/bar";
    const expected = "https://example.com/foo/bar.html";
    expect(rewriteLegacyURL(input)).toBe(expected);
  });

  it("does not remove ee/ if not at the beginning of the path", () => {
    const input = "https://example.com/tree/ee/foo/bar";
    const expected = "https://example.com/tree/ee/foo/bar.html";
    expect(rewriteLegacyURL(input)).toBe(expected);
  });

  it("adds .html to the end of the URL", () => {
    const input = "https://example.com/foo/bar";
    const expected = "https://example.com/foo/bar.html";
    expect(rewriteLegacyURL(input)).toBe(expected);
  });

  it("removes trailing slash and adds .html", () => {
    const input = "https://example.com/foo/bar/";
    const expected = "https://example.com/foo/bar.html";
    expect(rewriteLegacyURL(input)).toBe(expected);
  });

  it("handles URLs with query parameters", () => {
    const input = "https://example.com/ee/foo/bar?param=value";
    const expected = "https://example.com/foo/bar.html?param=value";
    expect(rewriteLegacyURL(input)).toBe(expected);
  });

  it("handles URLs with hash fragments", () => {
    const input = "https://example.com/ee/foo/bar#section";
    const expected = "https://example.com/foo/bar.html#section";
    expect(rewriteLegacyURL(input)).toBe(expected);
  });
});

describe("getRedirectUrl", () => {
  const oldDomain =
    "https://gitlab-org.gitlab.io/technical-writing-group/gitlab-docs-hugo";
  const newDomain = "https://new.docs.gitlab.com";

  test("redirects root old domain to new gitlab.com", () => {
    expect(getRedirectUrl(oldDomain)).toBe("https://new.gitlab.com");
    expect(getRedirectUrl(`${oldDomain}/`)).toBe("https://new.gitlab.com");
  });

  test("redirects old domain pages to new domain", () => {
    const path = "/some/page";
    expect(getRedirectUrl(`${oldDomain}${path}`)).toBe(`${newDomain}${path}`);
  });

  test("redirects old domain with query parameters", () => {
    const path = "/some/page?param=value";
    expect(getRedirectUrl(`${oldDomain}${path}`)).toBe(`${newDomain}${path}`);
  });

  test("redirects old domain with hash", () => {
    const path = "/some/page#section";
    expect(getRedirectUrl(`${oldDomain}${path}`)).toBe(`${newDomain}${path}`);
  });

  test("returns null for URLs not starting with old domain", () => {
    expect(getRedirectUrl("https://example.com")).toBeNull();
  });
});
