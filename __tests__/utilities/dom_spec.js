import { JSDOM } from "jsdom";
import { getNextUntil } from "../../themes/gitlab-docs/src/utilities/dom";

describe("getNextUntil", () => {
  let dom;
  let document;
  let container;

  beforeEach(() => {
    dom = new JSDOM(`
        <div id="container">
          <p id="start">Start</p>
          <span>Span 1</span>
          <div>Div 1</div>
          <span>Span 2</span>
          <p id="end">End</p>
          <div>Div after end</div>
        </div>
      `);
    document = dom.window.document;
    container = document.getElementById("container");
  });

  it("should return HTML between start and end elements", () => {
    const startEl = document.getElementById("start");
    const result = getNextUntil(startEl, "#end");
    expect(result).toBe(
      "<span>Span 1</span><div>Div 1</div><span>Span 2</span>",
    );
  });

  it("should return all subsequent HTML if no end selector is provided", () => {
    const startEl = document.getElementById("start");
    const result = getNextUntil(startEl);
    expect(result).toBe(
      '<span>Span 1</span><div>Div 1</div><span>Span 2</span><p id="end">End</p><div>Div after end</div>',
    );
  });

  it("should return an empty string if start element is the last child", () => {
    const lastEl = container.lastElementChild;
    const result = getNextUntil(lastEl);
    expect(result).toBe("");
  });

  it("should handle case when end selector doesn't match any element", () => {
    const startEl = document.getElementById("start");
    const result = getNextUntil(startEl, "#non-existent");
    expect(result).toBe(
      '<span>Span 1</span><div>Div 1</div><span>Span 2</span><p id="end">End</p><div>Div after end</div>',
    );
  });
});
