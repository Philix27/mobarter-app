import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import Icons from 'unplugin-icons/vite';
import { nodePolyfills } from 'vite-plugin-node-polyfills';


export default defineConfig({
	plugins: [
		sveltekit(),
		Icons({
			compiler: 'svelte'
		}),
		nodePolyfills({
			exclude: ['fs'],
			globals: {
				Buffer: true,
				global: true,
				process: true
			},
			protocolImports: true
		})
	],
	optimizeDeps: {
		include: ['dayjs/plugin/relativeTime.js', 'dayjs',]
	},
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
// vite.config.js
