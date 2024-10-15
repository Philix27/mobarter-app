<script>
	import Nav from '../../comps/Nav.svelte';
	import QuickActions from './QuickActions.svelte';
	import Transactions from './Transactions.svelte';
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Mobarter" />
	<div class="mt-[50px]">
		<QuickActions />
		<Transactions />
	</div>
</div>

<style lang="postcss">
</style>
