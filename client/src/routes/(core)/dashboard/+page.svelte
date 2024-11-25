<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import type { PageData } from './$types';
	import { cUSD } from 'celo-kit';
	import { getBalance, account, formatEtherRounded, getActiveChain, chainId } from 'lib/web3';
	import { browser } from '$app/environment';
	import { utilityCategories } from 'lib/category';

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
	let quickActions = [
		// { name: 'Airtime', link: '/payments/airtime', icon: 'fluent:call-32-regular' },
		// { name: 'Data', link: '/payments/data', icon: 'gg:data' },
		// { name: 'Cable TV', link: '/payments/cable', icon: 'ic:outline-live-tv' },
		// { name: 'Electricity', link: '/payments/electric', icon: 'lucide:utility-pole' },
		// { name: 'Sports', link: '/payments/sports', icon: 'solar:gamepad-outline' },
		// { name: 'Travel', link: '/payments/travel', icon: 'hugeicons:bitcoin-withdraw' }
		// { name: 'Savings', link: '/savings', icon: 'material-symbols-light:savings-outline' },
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
	<Nav theme={props.data.theme} showThemeToggle />
	<div class="">
		<div class="bg-primary px-3 py-2 mb-4 rounded-md">
			<div class="flex items-center justify-between">
				{#await getBalance(cUSD.address[getActiveChain($chainId)], $account.address ?? accountAddress!)}
					<p class="text-white">...</p>
				{:then data}
					<P className="text-xl font-extralight text-white"
						>{formatEtherRounded(data.value)} <span class="text-sm">{cUSD.symbol}</span></P
					>
				{:catch}
					<div></div>
				{/await}

				<div class="flex flex-col items-end justify-between">
					<a href="/balances" class="mb-4">
						<P className="text-white text-sm">Balances</P>
					</a>
					<a href="/transactions">
						<P className="text-white text-sm">Transactions</P>
					</a>
				</div>
			</div>
		</div>
		<div class=" min-h-screen h-screen overflow-y-scroll no-scrollbar scroll-smooth">
			<QuickActions items={quickActions} />
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
