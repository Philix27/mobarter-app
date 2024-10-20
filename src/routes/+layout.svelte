<script lang="ts">
	// @ts-nocheck

	import { Drawer } from 'components';
	import '../app.css';
	import 'tailwindcss/tailwind.css';
	import { drawerState } from '../store/settings';
	import { Toaster } from 'svelte-sonner';
	// Import the Analytics package, and the SvelteKit dev variable.
	import { dev } from '$app/environment';
	import { inject } from '@vercel/analytics';

	// Inject the Analytics functionality
	inject({ mode: dev ? 'development' : 'production' });
</script>

<div class="app">
	<!-- <WagmiProvider {config}> -->
	<main>
		{#if $drawerState}
			<Drawer />
		{/if}
		<slot />
		<Toaster position="top-center" />
	</main>
	<!-- </WagmiProvider> -->
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 1rem;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}
</style>
