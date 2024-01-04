# Documentation post-processing

To publish the source documentation content files and maintain compatibility with `/help`, post-processing tasks occur using the
[`content-post-process.sh`](../scripts/content-post-process.sh) script:

- Rename `index.md` files to `_index.md`. For more information, see [Index pages: _index.md](https://gohugo.io/content-management/organization/#index-pages-_indexmd).

## Temporary migration tasks

Scripts in the `scripts/migration` directory run as part of a Hugo build and change content in the source documentation Markdown files. These scripts:

- Allow building for testing and preview a Hugo-based site using documentation intended for a Nanoc-based site.
- Will eventually be used to migrate documentation permanently to a Hugo-based site build.
