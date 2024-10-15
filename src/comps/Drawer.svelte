<script>
	import { drawerState } from '../store/settings';
	import DrawerRow from './DrawerRow.svelte';

	const hideDrawer = () => {
		$drawerState = false;
	};

	const items = [
		{ title: 'Dashboard', link: '', icon: 'mdi:home' },
		{ title: 'Invoice', link: '', icon: 'mdi:settings' },
		{ title: 'Settings', link: '', icon: 'mdi:menu' },
		{ title: 'Help', link: '', icon: 'mdi:home' },
		{ title: 'Support', link: '', icon: 'mdi:home' },
		{ title: 'Wallet', link: '', icon: 'mdi:home' },
		{ title: 'Version', link: '', icon: 'mdi:home' }
	];

	const onItemClick = () => {};
</script>

<div class="w-full h-screen flex fixed top-0 left-0 bg-black/30">
	<div class={'w-[60%] max-w-[300px] bg-background rounded-r-2xl'}>
		<div class="mt-[50px] mb-[100px] flex flex-col">
			{#each items as drawerItem, i}
				<DrawerRow title={drawerItem.title} icon={drawerItem.icon} />
			{/each}
		</div>
	</div>
	<div class={'w-[40%]'} on:click={hideDrawer} />
</div>

<style>
</style>
