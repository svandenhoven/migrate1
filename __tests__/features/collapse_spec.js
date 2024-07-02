import { JSDOM } from "jsdom";
import { expandCollapse } from "../../themes/gitlab-docs/src/features/collapse";

describe("expandCollapse", () => {
  let dom;
  let container;

  beforeEach(() => {
    dom = new JSDOM(`
      <div>
        <button data-toggle="collapse" data-target="#target1" aria-expanded="false">Toggle 1</button>
        <div id="target1" class="collapse" aria-hidden="true">Target 1</div>
        <button data-toggle="collapse" data-target="#target2" aria-expanded="false">Toggle 2</button>
        <div id="target2" class="collapse" aria-hidden="true">Target 2</div>
      </div>
    `);
    global.document = dom.window.document;
    container = dom.window.document.querySelector("div");
  });

  afterEach(() => {
    dom.window.close();
  });

  test("expands and collapses target elements when triggers are clicked", () => {
    expandCollapse();

    const trigger1 = container.querySelector('[data-target="#target1"]');
    const trigger2 = container.querySelector('[data-target="#target2"]');
    const target1 = container.querySelector("#target1");
    const target2 = container.querySelector("#target2");

    // Initial state
    expect(target1.classList.contains("show")).toBe(false);
    expect(target1.getAttribute("aria-hidden")).toBe("true");
    expect(trigger1.getAttribute("aria-expanded")).toBe("false");
    expect(target2.classList.contains("show")).toBe(false);
    expect(target2.getAttribute("aria-hidden")).toBe("true");
    expect(trigger2.getAttribute("aria-expanded")).toBe("false");

    // Click trigger1
    trigger1.click();
    expect(target1.classList.contains("show")).toBe(true);
    expect(target1.getAttribute("aria-hidden")).toBe("false");
    expect(trigger1.getAttribute("aria-expanded")).toBe("true");
    expect(target2.classList.contains("show")).toBe(false);
    expect(target2.getAttribute("aria-hidden")).toBe("true");
    expect(trigger2.getAttribute("aria-expanded")).toBe("false");

    // Click trigger2
    trigger2.click();
    expect(target1.classList.contains("show")).toBe(true);
    expect(target1.getAttribute("aria-hidden")).toBe("false");
    expect(trigger1.getAttribute("aria-expanded")).toBe("true");
    expect(target2.classList.contains("show")).toBe(true);
    expect(target2.getAttribute("aria-hidden")).toBe("false");
    expect(trigger2.getAttribute("aria-expanded")).toBe("true");

    // Click trigger1 again
    trigger1.click();
    expect(target1.classList.contains("show")).toBe(false);
    expect(target1.getAttribute("aria-hidden")).toBe("true");
    expect(trigger1.getAttribute("aria-expanded")).toBe("false");
    expect(target2.classList.contains("show")).toBe(true);
    expect(target2.getAttribute("aria-hidden")).toBe("false");
    expect(trigger2.getAttribute("aria-expanded")).toBe("true");
  });
});
