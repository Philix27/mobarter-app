<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import Categories from './Categories.svelte';
	import { account, getBalance } from 'lib/web3';
	import type { PageData } from './$types';
	import { cUSD } from 'celo-kit';
	import { onMount } from 'svelte';
	import Transactions from '../transactions/Transactions.svelte';

	export let data: PageData;

	const clickIcon = () => {
		goto('/notify');
	};

	let balance: number = 0;

	onMount(() => {
		getBalance(cUSD.address.alfajores as `0x${string}`, $account.address as `0x${string}`)
			.then((val) => {
				balance = parseInt(val.value.toString());
			})
			.catch((e) => {
				balance = 0.0;
			});
	});
</script>

<svelte:head>
	<title>Dashboard</title>
</svelte:head>

<div>
	<Nav theme={data.theme} showThemeToggle />
	<div class="">
		<div class="flex items-center justify-between mb-3 mx-4">
			<P className="text-3xl font-extralight">${balance}</P>

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
