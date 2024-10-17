<script>
	// @ts-nocheck

	import Drawer from '../comps/Drawer.svelte';
	import '../app.css';
	import 'tailwindcss/tailwind.css';
	import { drawerState } from '../store/settings';
	import { http, createConfig } from '@wagmi/core';
	import { mainnet, sepolia } from '@wagmi/core/chains';

	export const config = createConfig({
		chains: [mainnet, sepolia],
		transports: {
			[mainnet.id]: http(),
			[sepolia.id]: http()
		}
	});
</script>

<div class="app">
	<!-- <WagmiProvider {config}> -->
	<main>
		{#if $drawerState}
			<Drawer />
		{/if}
		<slot />
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
