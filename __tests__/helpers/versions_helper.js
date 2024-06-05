/**
 * Creates a mock browser window object with a given path and/or href.
 *
 * @param {Object} options
 *   Can include "href" or "pathname" properties
 */
export const setWindowLocation = (options = {}) => {
  const { pathname, href } = options;
  const location = {
    ...window.location,
    ...(pathname && { pathname }),
    ...(href && { href }),
  };
  Object.defineProperty(window, "location", {
    writable: true,
    value: location,
  });
};

/**
 * Creates a mock metatag.
 *
 * @param {String} tagName
 * @param {String} tagContent
 */
export const setMetatag = (tagName, tagContent) => {
  const meta = document.createElement("meta");
  meta.setAttribute("name", tagName);
  meta.setAttribute("content", tagContent);
  document.head.appendChild(meta);
};
