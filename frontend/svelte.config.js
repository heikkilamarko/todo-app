import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', '.svg'],
	kit: {
		adapter: adapter({
			fallback: 'app.html',
			precompress: true
		}),
		prerender: {
			entries: []
		}
	}
};

export default config;
