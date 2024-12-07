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
	import Configs from './Configs.svelte';

	// Inject the Analytics functionality
	inject({ mode: dev ? 'development' : 'production' });
</script>

<!-- <svelte:head>
	<html class="overscroll-none bg-background color-scheme: dark">
		<slot />
	</html>
</svelte:head> -->
<div class="flex flex-col min-h-screen w-full items-center justify-center">
	<div class=" min-h-screen max-w-[400px] w-[400px] bg-background">
		<Configs>
			<!-- <WagmiProvider {config}> -->
			<main>
				{#if $drawerState}
					<Drawer />
				{/if}
				<slot />
				<Toaster position="top-center" />
			</main>
			<!-- </WagmiProvider> -->
		</Configs>
	</div>
</div>
