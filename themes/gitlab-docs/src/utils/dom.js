/**
 * Get HTML between two elements.
 *
 * @param {Element} el
 * @param {String} selector
 * @returns {String} HTML between the two given elements
 */
export const getNextUntil = (el, selector) => {
  const siblings = [];
  let next = el.nextElementSibling;

  while (next) {
    if (selector && next.matches(selector)) break;
    siblings.push(next.outerHTML);
    next = next.nextElementSibling;
  }

  return siblings.join("");
};
