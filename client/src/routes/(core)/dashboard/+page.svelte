<script>
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import Transactions from './Transactions.svelte';

	const clickIcon = () => {
		goto('/notify');
	};
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Mobarter" icon="mdi:notifications" onIconClick={clickIcon} />
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
