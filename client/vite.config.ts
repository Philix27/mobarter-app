import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import Icons from 'unplugin-icons/vite';
import { nodePolyfills } from 'vite-plugin-node-polyfills';
// import codegen from 'vite-plugin-graphql-codegen';

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
		include: [
			'dayjs/plugin/relativeTime.js',
			'dayjs',
			'@apollo/client/core',
			'@apollo/client/cache',
			'@apollo/client/link/ws',
			'@apollo/client/link/context',
			'@apollo/client/link/error',
			'@apollo/client/utilities'
		],
		exclude: ['@apollo/client', 'svelte-apollo', 'subscriptions-transport-ws', '@urql/svelte']
	},
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
// vite.config.js
