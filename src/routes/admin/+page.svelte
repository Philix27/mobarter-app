<script lang="ts">
	import { Nav } from 'components';
	import 'iconify-icon';
	import { adminScreenState, type IAdminScreen } from './store';
	import Card from './Card.svelte';
	import CreditBuy from './CreditBuy.svelte';
	import CreditSell from './CreditSell.svelte';
	import Email from './Email.svelte';
	import Kyc from './Kyc.svelte';
	import Logs from './Logs.svelte';

	const sections: { title: string; icon: string; link: IAdminScreen }[] = [
		{ title: 'Buyers', icon: 'icons8:buy', link: 'CreditBuy' },
		{ title: 'Sellers', icon: 'mdi:dollar', link: 'CreditSell' },
		{ title: 'Support', icon: 'mdi:support', link: 'Support' },
		{ title: 'Emails', icon: 'mdi:email-outline', link: 'Email' },
		{ title: 'Kyc', icon: 'mdi:lock', link: 'Kyc' },
		{ title: 'Newsletter', icon: 'noto:newspaper', link: 'NewsLetter' },
		{ title: 'Log', icon: 'octicon:log-16', link: 'Logs' }
	];

	function clickTab(screen: IAdminScreen) {
		$adminScreenState = screen;
	}
</script>

<div>
	<Nav title="Admin" isBack />
	<div class="grid md:grid-cols-7 grid-cols-4 w-full md:gap-4 gap-2">
		{#each sections as item}
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<div
				role="button"
				on:click={() => clickTab(item.link)}
				class="bg-card rounded-md md:p-2 px-1 py-2 flex flex-col items-center justify-center h-full hover:border-primary hover:border"
			>
				<iconify-icon class="text-foreground text-[20px] mb-2" icon={item.icon} />
				<p class="text-xs">{item.title}</p>
			</div>
		{/each}
	</div>
	<div>
		{#if $adminScreenState === 'Card'}
			<Card />
		{:else if $adminScreenState === 'CreditBuy'}
			<CreditBuy />
		{:else if $adminScreenState === 'CreditSell'}
			<CreditSell />
		{:else if $adminScreenState === 'Email'}
			<Email />
		{:else if $adminScreenState === 'Kyc'}
			<Kyc />
		{:else if $adminScreenState === 'Logs'}
			<Logs />
		{/if}
	</div>
</div>
