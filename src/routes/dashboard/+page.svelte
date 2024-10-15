<script>
	import { goto } from '$app/navigation';
	import Nav from '../../comps/Nav.svelte';
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
	<div class="mt-[40px]">
		<QuickActions />
		<Transactions />
	</div>
</div>
