<script>
	// import Counter from './Counter.svelte';
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';
	import { account, chainId, modal } from 'lib/web3';
	import { Button } from 'components';

	// import { connected } from 'svelte-wagmi';

	function connectWallet() {
		modal.open({ view: 'Connect' });
	}
</script>

<section class="bg-primary h-screen flex flex-col items-center ">
	<img src="dollar.png" alt="Welcome" class="h-fit max-h-[300px] mt-[100px]" />
	<a href="/dashboard" class="my-4"><p>Dashboard</p></a>

	{#if $account.isConnected}
		<p>Already Connected</p>
	{:else}
		<Button variant="secondary" className="w-fit" onclick={connectWallet}>Connect</Button>
		<!-- <button on:click={}>Connect</button> -->
	{/if}
</section>

<style>
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
