<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import Categories from './Categories.svelte';
	import type { PageData } from './$types';
	import { cUSD } from 'celo-kit';
	import Transactions from '../transactions/Transactions.svelte';
	import { getBalance, account, formatEtherRounded, getActiveChain, chainId } from 'lib/web3';
	import { browser } from '$app/environment';

	let props = $props<{ data: PageData }>();
	const clickIcon = () => {
		goto('/notify');
	};

	let isMiniPay = $state(false);
	let accountAddress = $state($account.address);
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

	let balance: number = 0;
</script>

<svelte:head>
	<title>Dashboard</title>
</svelte:head>

<div>
	<Nav theme={props.data.theme} showThemeToggle />
	<div class="">
		<div class="flex items-center justify-between mb-3 mx-4">
			{#await getBalance(cUSD.address[getActiveChain($chainId)], $account.address ?? accountAddress!)}
				<p>...</p>
			{:then data}
				<P className="text-3xl font-extralight">${formatEtherRounded(data.value)} {cUSD.symbol}</P>
			{/await}

			<a href="/balances">
				<P className="text-primary text-sm">Balances</P>
			</a>
		</div>
		<div class=" min-h-screen h-screen overflow-y-scroll no-scrollbar scroll-smooth">
			<QuickActions />
			<div class="">
				<div class="flex items-center justify-between mt-5 mb-3 mx-2">
					<P className="text-sm">Utilities</P>
					<div></div>
				</div>
				<Categories />
				<!-- <div class="max-h-[300px] overflow-y-scroll no-scrollbar">
					<Categories />
				</div> -->
			</div>

			<div>
				<div class="flex items-center justify-between mt-5 mb-3 mx-2">
					<P className="text-sm">Transactions</P>
					<a href="/transactions">
						<iconify-icon icon="material-symbols:history" class="text-xl text-primary" />
					</a>
				</div>
				<div class="max-h-[300px] overflow-y-scroll no-scrollbar">
					<Transactions />
				</div>
			</div>
		</div>
	</div>
</div>
