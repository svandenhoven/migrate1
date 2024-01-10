/**
 * Site-wide functions.
 */

/**
 * Adds a clickable permalink to each content heading
 */
const addHeaderPermalinks = () => {
  document
    .querySelectorAll(".docs-content h2, h3, h4, h5, h6")
    .forEach((heading) => {
      const { id } = heading;
      if (!id) return;

      const anchor = document.createElement("a");
      anchor.href = `#${id}`;
      anchor.title = "Permalink";
      anchor.classList.add("anchor");

      heading.appendChild(anchor);
    });
};

document.addEventListener("DOMContentLoaded", () => {
  addHeaderPermalinks();
});
