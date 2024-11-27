<script lang="ts">
	import { account, modal, getBalance } from 'lib/web3';
	import { Button } from 'components';
	import { browser } from '$app/environment';

	$: accountAddress = $account.address;
	// let accountAddress = $state($account.address);
	$: isMiniPay = false;

	if (browser) {
		if (window && window.ethereum) {
			// User has a injected wallet

			if (window.ethereum.isMiniPay) {
				// User is using Minipay
				isMiniPay = true;
				// Requesting account addresses
				let accounts = window.ethereum.request({
					method: 'eth_requestAccounts',
					params: []
				});

				// Injected wallets inject all available addresses,
				// to comply with API Minipay injects one address but in the form of array
				console.log(accounts[0]);
				// @ts-ignore
				// accountAddress = $account.address;
				accountAddress = accounts[0];
			}

			// User is not using MiniPay
		}
	}

	function connectWallet() {
		// const yoe = $postsX.data;
		// $dataPlans.data?.Airtime_GetDataPlans?.dataPlans?.map((val, i) => {
		// 	val?.network;
		// });
		// console.log('');
		// console.log('Afa');
		modal.open({ view: 'Connect' });
	}
</script>

<section class="bg-primary h-screen flex flex-col items-center">
	<img src="dollar.png" alt="Welcome" class="h-fit max-h-[300px] mt-[100px]" />

	{#if $account.isConnected || isMiniPay}
		<a href="/dashboard" class="my-4"><p>Dashboard</p></a>
	{:else}
		<Button variant="secondary" className="w-fit" onclick={connectWallet}>Connect</Button>
		<!-- <button on:click={}>Connect</button> -->
	{/if}
</section>
