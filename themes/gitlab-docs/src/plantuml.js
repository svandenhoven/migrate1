/**
 * Convert PlantUML markup to images
 *
 * Custom implementation based on https://plantuml.com/text-encoding
 *
 * Customizations:
 * - Use GitLab's PlantUML server
 * - Return images in PNG format
 * - Adds a link around the image
 */

/**
 * Deflation compression with browser APIs.
 *
 * 1. We first create a blob containing the plant uml as the content
 * 2. We pipe the blobs content through deflate (without headers) with CompressionStream
 * 3. We convert the returned bytes to an ASCII string with `String.fromCodePoint`
 *
 * @see https://developer.mozilla.org/en-US/docs/Web/API/CompressionStream
 * @param {string} puml
 * @returns {Promise<string>}
 */
async function deflateCustom(puml) {
  const compressedStream = new Blob([puml], { type: "text/plain" })
    .stream()
    .pipeThrough(new CompressionStream("deflate-raw"));

  let deflated = "";

  for await (const arrayOfUint8 of compressedStream) {
    deflated += String.fromCodePoint(...arrayOfUint8);
  }

  return deflated;
}

const BASE_64 =
  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
const BASE_PLANT_UML =
  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_";

/**
 * Convert base64 to plant uml's encoding
 *
 * Plant UML uses a custom base64-like encoding. However, the characters map differently, for example
 * what's A in base64 is 0 in the encoding.
 *
 * We can simply replace the characters
 *
 * @see https://plantuml.com/text-encoding
 * @param string
 * @returns {string}
 */
function convertBase64ToPlantUML(string) {
  const base64 = window.btoa(string);
  let res = "";
  for (const letter of base64) {
    res += BASE_PLANT_UML[BASE_64.indexOf(letter)];
  }
  return res;
}

const plantUmlElements = document.querySelectorAll(".language-plantuml");
plantUmlElements.forEach(async (element) => {
  const parent = element.parentNode;
  const s = decodeURIComponent(encodeURIComponent(element.textContent));
  const imgElement = document.createElement("img");
  const imgData = convertBase64ToPlantUML(await deflateCustom(s));
  const imgSrc = `https://plantuml.gitlab-static.net/png/${imgData}`;
  imgElement.setAttribute("src", imgSrc);

  // Create a link element
  const linkElement = document.createElement("a");
  linkElement.setAttribute("href", imgSrc);
  linkElement.setAttribute("class", "plantuml");
  linkElement.setAttribute("target", "_blank");

  // Wrap the image element inside the link element
  linkElement.appendChild(imgElement);

  parent.replaceChild(linkElement, element);
});
