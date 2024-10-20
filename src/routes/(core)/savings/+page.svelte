<script lang="ts">
	import { Nav, Tabs } from 'components';
	import { savingsStore, type ISavingsStore } from './store';
	import { getTitle } from 'components/TabTitle';

	const tabs: { title: ISavingsStore; onClick: VoidFunction }[] = [
		{
			title: "Withdraw",
			onClick: () => {
				$savingsStore = 'Withdraw';
			}
		},
		{
			title: 'Deposit',
			onClick: () => {
				$savingsStore = 'Deposit';
			}
		}
	];
</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Savings" isBack />
	<div class="">
		<Tabs {tabs} activeTitle={getTitle(tabs, $savingsStore)} />
		<div>Send Money</div>
	</div>
</div>

<style lang="postcss">
</style>
