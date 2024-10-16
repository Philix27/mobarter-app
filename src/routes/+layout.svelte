<script>
	// @ts-nocheck

	import Drawer from '../comps/Drawer.svelte';
	import '../app.css';
	import 'tailwindcss/tailwind.css';
	import { drawerState } from '../store/settings';
	// import { WagmiProvider } from '@byteatatime/wagmi-svelte';
	// import { WagmiProvider, createWagmiConfig, http } from '@byteatatime/wagmi-svelte';
	// import { mainnet, sepolia } from '@byteatatime/wagmi-svelte/chains';

	// import { WagmiProvider } from 'wagmi-svelte';

	const config = createWagmiConfig({
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
