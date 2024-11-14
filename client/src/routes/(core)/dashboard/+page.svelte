<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import Transactions from './Transactions.svelte';
	import { account, getBalance } from 'lib/web3';
	import type { PageData } from './$types';
	import { cUSD } from 'celo-kit';
	import { onMount } from 'svelte';

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
	<Nav title="Mobarter" theme={data.theme} onIconClick={clickIcon} isHome={true} />

	<div class="">
		<div class="flex items-center justify-between mb-3 mx-4">
			<P className="text-3xl font-extralight">${balance}</P>

			<a href="/balances">
				<P className="text-primary text-sm">Balances</P>
			</a>
		</div>
		<QuickActions />
		<Transactions />
	</div>
</div>
