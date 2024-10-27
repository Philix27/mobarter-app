<script lang="ts">
	import { getTitle, Nav, P, Tabs } from 'components';
	import Transactions from '../dashboard/Transactions.svelte';
	import type { ITransactionsTabs } from './store';
	import { transactionsState } from './store';

	const tabs: { title: ITransactionsTabs; onClick: VoidFunction }[] = [
		{
			title: 'Wallet',
			onClick: () => {
				$transactionsState = 'Wallet';
			}
		},
		{
			title: 'Swap',
			onClick: () => {
				$transactionsState = 'Swap';
			}
		},
		{
			title: 'Airtime',
			onClick: () => {
				$transactionsState = 'Airtime';
			}
		}
	];
</script>

<div>
	<Nav title="Transactions" isBack />
	<Tabs {tabs} activeTitle={getTitle(tabs, $transactionsState)} />
	<Transactions />
</div>
