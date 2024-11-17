<script lang="ts">
	import { account, chainId, modal, getBalance } from 'lib/web3';
	import { Button } from 'components';
	import { browser } from '$app/environment';
	import { queryStore, gql, getContextClient } from '@urql/svelte';
	import type { GetCountryQuery, GetCountryQueryStore } from '../generated/graphql';

	const posts = queryStore<GetCountryQuery>({
		client: getContextClient(),
		query: gql`
			query GetCountry {
				CountryList {
					name
				}
			}
		`
	});

	$: accountAddress = '';
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
		const yoe = $posts.data;
		modal.open({ view: 'Connect' });
	}
</script>

<section class="bg-primary h-screen flex flex-col items-center">
	<img src="dollar.png" alt="Welcome" class="h-fit max-h-[300px] mt-[100px]" />
	<a href="/dashboard" class="my-4"><p>Dashboard</p></a>
	{#if !$posts.data}
		<p>Loading....</p>
	{:else if $posts.data}
		<!-- <p>{$posts.data}</p> -->
		<p>{$posts.data.CountryList!.name}</p>
	{:else}
		<p>Nothing...</p>
	{/if}
	<p>Address: {accountAddress}</p>
	<p>Address Wag: {$account.address}</p>

	{#if $account.isConnected}
		<p>Already Connected</p>
	{:else}
		<Button variant="secondary" className="w-fit" onclick={connectWallet}>Connect</Button>
		<!-- <button on:click={}>Connect</button> -->
	{/if}
</section>
