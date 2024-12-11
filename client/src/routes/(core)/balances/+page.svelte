<script lang="ts">
	import { Nav, P } from 'components';
	import { globalStore } from 'lib/store';
	import * as Tabs from '$lib/components/ui/tabs';
	import Celo from './Celo.svelte';

	const TabKeys = {
		celo: 'Celo',
		eth: 'Ethereum',
		abitrium: 'Abitrium',
		optimism: 'Optimisim',
		starknet: 'Starknet'
	};

	function getList() {
		const list = [TabKeys.celo, TabKeys.eth];
		if ($globalStore.env !== 'MINIPAY') {
			list.push(TabKeys.optimism, TabKeys.starknet);
		}

		return list;
	}
</script>

<svelte:head>
	<title>Wallet Balance</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Balance" isBack />

	<Tabs.Root value={TabKeys.celo} class="w-full">
		<Tabs.List>
			<Tabs.Head list={getList()} />
		</Tabs.List>
		<Tabs.Content value={TabKeys.celo}>
			<Celo />
		</Tabs.Content>
		<Tabs.Content value={TabKeys.eth}>
			<div>eth</div>
		</Tabs.Content>

		{#if $globalStore.env !== 'MINIPAY'}
			<Tabs.Content value={TabKeys.optimism}>
				<div>optimism</div>
			</Tabs.Content>
		{/if}
		{#if $globalStore.env !== 'MINIPAY'}
			<Tabs.Content value={TabKeys.starknet}>
				<div>starknet</div>
			</Tabs.Content>
		{/if}
	</Tabs.Root>
</div>
