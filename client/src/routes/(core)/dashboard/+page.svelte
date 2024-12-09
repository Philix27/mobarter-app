<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import type { PageData } from './$types';
	import { cUSD } from 'celo-kit';
	import { getBalance, account, formatEtherRounded, getActiveChain, chainId } from 'lib/web3';
	import { browser } from '$app/environment';
	import { utilityCategories } from 'lib/category';
	import Carousel from './Carousel.svelte';

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

	let quickActions = [
		{ name: 'Buy', link: '/p2p/sellers', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Sell', link: '/p2p/buyers', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Withdraw', link: '/withdraw', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Send', link: '/send', icon: 'majesticons:send-line' },
		{ name: 'Swap', link: '/swap', icon: 'tdesign:swap' },
		{ name: 'Receive', link: '/receive', icon: 'hugeicons:money-receive-circle' }
	];
</script>

<svelte:head>
	<title>Dashboard</title>
</svelte:head>

<div>
	<Nav theme={props.data.theme} showThemeToggle gotoUrl="/login" />
	<div class="">
		<div class="bg-primary px-3 py-2 mb-4 rounded-md">
			<div class="flex items-center justify-between">
				{#await getBalance(cUSD.address[getActiveChain($chainId)], $account.address ?? accountAddress!)}
					<P className="text-xl font-extralight text-white"
						>0.. <span class="text-sm">{cUSD.symbol}</span></P
					>
				{:then data}
					<P className="text-xl font-extralight text-white"
						>{formatEtherRounded(data.value)} <span class="text-sm">{cUSD.symbol}</span></P
					>
				{:catch}
					<P className="text-xl font-extralight text-white"
						>0.. <span class="text-sm">{cUSD.symbol}</span></P
					>
				{/await}

				<div class="flex flex-col items-end justify-between">
					<a href="/balances" class="mb-8">
						<P className="text-white text-sm">Balances</P>
					</a>
					<a href="/transactions">
						<P className="text-white text-sm">Transactions</P>
					</a>
				</div>
			</div>
		</div>
		<div class="min-h-screen h-screen overflow-y-scroll no-scrollbar scroll-smooth">
			<QuickActions items={quickActions} />
			<Carousel />
			<div class="">
				<div class="flex items-center justify-between mt-5 mb-3 mx-2">
					<P className="text-sm">Utilities</P>
					<div></div>
				</div>
				<QuickActions items={utilityCategories} />
			</div>
		</div>
	</div>
</div>
