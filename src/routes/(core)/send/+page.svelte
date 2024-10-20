<script lang="ts">
	import { Nav, Tabs } from 'components';
	import { sendScreenState, type ISendTabs } from './store';
	import ScreenBank from './ScreenBank.svelte';
	import ScreenWallet from './ScreenWallet.svelte';
	import ScreenPhone from './ScreenPhone.svelte';
	import { getTitle } from 'components/TabTitle';

	const tabs: { title: ISendTabs; onClick: VoidFunction }[] = [
		{
			title: 'Wallet',
			onClick: () => {
				$sendScreenState = 'Wallet';
			}
		},
		{
			title: 'Bank',
			onClick: () => {
				$sendScreenState = 'Bank';
			}
		},
		{
			title: 'Phone',
			onClick: () => {
				$sendScreenState = 'Phone';
			}
		}
	];
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Send" isBack />
	<div class="">
		<Tabs {tabs} activeTitle={getTitle(tabs, $sendScreenState)} />
		{#if $sendScreenState === 'Bank'}
			<ScreenBank />
		{:else if $sendScreenState === 'Phone'}
			<ScreenPhone />
		{:else if $sendScreenState === 'Wallet'}
			<ScreenWallet />
		{/if}
	</div>
</div>

<!-- activeTitle={$sendScreenState === 'BANK'
	? 'Bank'
	: $sendScreenState === 'PHONE'
		? 'Phone'
		: 'Wallet'} -->

<!-- activeTitle={tabs
				.filter((val) => val.title.toString() === $sendScreenState)[0]
				.title.toString()} -->
