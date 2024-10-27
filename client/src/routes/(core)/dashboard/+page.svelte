<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import Transactions from './Transactions.svelte';

	const clickIcon = () => {
		goto('/notify');
	};

	import type { PageData } from './$types';

	export let data: PageData;

	// @ts-ignore
	// let { data } = $props();
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav
		title={data.theme}
		icon="mdi:notifications"
		theme={data.theme}
		onIconClick={clickIcon}
		isHome={true}
	/>
	<!-- <Nav title="Mobarter" icon="mdi:notifications" onIconClick={clickIcon} isHome={true} /> -->
	<div class="">
		<div class="flex items-center justify-between mb-3 mx-4">
			<P className="text-3xl font-extralight">$0.00</P>
			<a href="/balances">
				<P className="text-primary text-sm">Balances</P>
			</a>
		</div>
		<QuickActions />
		<Transactions />
	</div>
</div>
