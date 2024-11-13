<script>
	import { account, chainId, modal } from 'lib/web3';
	import { Button } from 'components';

	// import { connected } from 'svelte-wagmi';
	import { browser } from '$app/environment';

	$: accountAddress = '';

	if (browser) {
		if (window && window.ethereum) {
			// User has a injected wallet

			if (window.ethereum.isMiniPay) {
				// User is using Minipay

				// Requesting account addresses
				let accounts = window.ethereum.request({
					method: 'eth_requestAccounts',
					params: []
				});

				// Injected wallets inject all available addresses,
				// to comply with API Minipay injects one address but in the form of array
				console.log(accounts[0]);
				// @ts-ignore
				accountAddress = $account.address;
			}

			// User is not using MiniPay
		}
	}

	function connectWallet() {
		modal.open({ view: 'Connect' });
	}

	// The code must run in a browser environment and not in node environment

	// User does not have a injected wallet
</script>

<section class="bg-primary h-screen flex flex-col items-center">
	<img src="dollar.png" alt="Welcome" class="h-fit max-h-[300px] mt-[100px]" />
	<a href="/dashboard" class="my-4"><p>Dashboard</p></a>
	<p>Address: {accountAddress}</p>

	{#if $account.isConnected}
		<p>Already Connected</p>
	{:else}
		<Button variant="secondary" className="w-fit" onclick={connectWallet}>Connect</Button>
		<!-- <button on:click={}>Connect</button> -->
	{/if}
</section>
