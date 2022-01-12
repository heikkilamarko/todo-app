import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  extensions: [".svelte", ".svg"],
  kit: {
    target: "#svelte",
    files: {
      template: "src/index.html",
    },
    adapter: adapter({
      fallback: "index.html",
    }),
  },
};

export default config;
