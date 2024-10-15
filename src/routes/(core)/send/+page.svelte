<script>
	import Tabs from '../../../comps/Tabs.svelte';
	import Nav from '../../../comps/Nav.svelte';
	import { sendScreenState } from './store';
	import ScreenBank from './ScreenBank.svelte';
	import ScreenWallet from './ScreenWallet.svelte';
	import ScreenPhone from './ScreenPhone.svelte';

	const tabs = [
		{
			title: 'Wallet',
			isActive: $sendScreenState === 'WALLET',
			onClick: () => {
				$sendScreenState = 'WALLET';
			}
		},
		{
			title: 'Bank Account',
			isActive: $sendScreenState === 'BANK',
			onClick: () => {
				$sendScreenState = 'BANK';
			}
		},
		{
			title: 'Phone',
			isActive: $sendScreenState === 'PHONE',
			onClick: () => {
				$sendScreenState = 'PHONE';
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
		<Tabs {tabs} />

		{#if $sendScreenState === 'BANK'}
			<ScreenBank />
		{:else if $sendScreenState === 'PHONE'}
			<ScreenPhone />
		{:else if $sendScreenState === 'WALLET'}
			<ScreenWallet />
		{/if}
	</div>
</div>

<style lang="postcss">
</style>
