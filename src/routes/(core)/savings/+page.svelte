<script>
	import Tabs from '../../../comps/Tabs.svelte';
	import Nav from '../../../comps/Nav.svelte';
	import { receiveState } from './store';

	const tabs = [
		{
			title: 'Wallet',
			isActive: $receiveState === 'WALLET',
			onClick: () => {
				$receiveState = 'WALLET';
			}
		},
		{
			title: 'Buy',
			isActive: $receiveState === 'BUY',
			onClick: () => {
				$receiveState = 'BUY';
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
		<Tabs {tabs} />
		<div>Send Money</div>
	</div>
</div>

<style lang="postcss">
</style>
