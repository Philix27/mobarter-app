<script>
	import { drawerState } from '../store/settings';
	import DrawerRow from './DrawerRow.svelte';

	const hideDrawer = () => {
		$drawerState = false;
	};

	const items = [
		{ title: 'Dashboard', link: '/dashboard', icon: 'mdi:home' },
		{ title: 'Settings', link: '/settings', icon: 'mdi:settings' },
		{ title: 'Help', link: '/help', icon: 'mdi:help' },
		{ title: 'Support', link: '/support', icon: 'mdi:support' },
		{ title: 'Admin', link: '/admin', icon: 'mdi:support' }
	];

	const onItemClick = () => {};
</script>

<div class="w-full h-screen flex fixed top-0 left-0 bg-black/30 z-20">
	<div class={'w-[55%] max-w-[300px] bg-background rounded-r-2xl'}>
		<div class="mt-[50px] mb-[100px] flex flex-col">
			{#each items as { link, title, icon }, i}
				<DrawerRow {title} {icon} {link} />
			{/each}
		</div>
	</div>
	<div class={'w-[40%]'} on:click={hideDrawer} />
</div>
