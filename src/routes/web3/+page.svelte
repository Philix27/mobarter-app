<script>
	// import Counter from './Counter.svelte';
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';
	import { account, chainId, modal } from 'lib/web3';
	import Counter from '../Counter.svelte';

	// import { connected } from 'svelte-wagmi';

	function connectWallet() {
		modal.open({ view: 'Connect' });
	}
</script>

<svelte:head>
	<title>Mobarter</title>
	<meta name="description" content="On & Off ramping" />
</svelte:head>

<section>
	<h1>
		<span class="welcome">
			<picture>
				<source srcset={welcome} type="image/webp" />
				<img src={welcome_fallback} alt="Welcome" />
			</picture>
		</span>

		to your new<br />SvelteKit appex
	</h1>

	<Counter />

	<a href="/dashboard" class="my-4"><p>Dashboard</p></a>

	{#if $account.isConnected}
		<p>Connected</p>
	{:else}
		<button on:click={connectWallet} class="p-2 bg-orange-500 rounded-lg">Not Connected</button>
		<!-- <button on:click={}>Connect</button> -->
	{/if}
</section>

<div class="w-full">
	{#if $account.address}
		<div class="card space-x-20 py-4">
			{#if $account.chain}
				<div class="grid grid-cols-2 text-center">
					<div>Chain ID: <span class="font-bold"> {$chainId}</span></div>

					<div>Chain Name: <span class="font-bold"> {$account.chain?.name}</span></div>
					<div>
						Chain Native Currency: <span class="font-bold">
							{$account.chain?.nativeCurrency.decimals}</span
						>
					</div>
					<div>
						Chain Native Decimals: <span class="font-bold">
							{$account.chain?.nativeCurrency.name}</span
						>
					</div>
				</div>
			{/if}
		</div>
	{/if}
</div>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	}
</style>
