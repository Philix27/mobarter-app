<script>
	import P from '../components/P.svelte';
	import 'iconify-icon';
	import { goto } from '$app/navigation';
	import { drawerState } from '../store/settings';

	export let title;
	export let isBack = false;
	export let icon = '';
	export let onIconClick = () => {};

	/**
	 * @param {string} route
	 * @param {boolean} replaceState
	 */
	function routeToPage(route, replaceState) {
		goto(`/${route}`, { replaceState });
	}

	function goBack(defaultRoute = '/dashboard') {
		// function goBack() {
		if (isBack) {
			const ref = document.referrer;
			// console.log('Cliecked backe', { ref });
			// goto(ref.length > 0 ? ref : defaultRoute);
			goto('/dashboard');
		}
	}

	function onShowDrawer() {
		$drawerState = true;
	}
	function onCloseDrawer() {
		$drawerState = false;
	}
</script>

<div
	class={`
        h-[50px] flex items-center 
        justify-between 
        fixed top-0 w-full  left-0
        z-10 bg-background px-5
      `}
>
	{#if isBack}
		<!-- svelte-ignore a11y-click-events-have-key-events -->
		<!-- svelte-ignore a11y-interactive-supports-focus -->
		<iconify-icon
			class="text-foreground text-[20px]"
			icon="mingcute:left-line"
			on:click={goBack}
			role="button"
		/>
	{:else}
		<!-- svelte-ignore a11y-click-events-have-key-events -->
		<!-- svelte-ignore a11y-interactive-supports-focus -->
		<iconify-icon
			class="text-foreground text-[20px]"
			icon="mdi:menu"
			on:click={onShowDrawer}
			role="button"
		/>
	{/if}

	<P variant="p4" className="text-card-foreground my-0 font-extrabold tracking-wide font-sans"
		>{title}
	</P>
	<div class="flex items-center">
		{#if icon}
			<iconify-icon class="text-xl text-foreground mr-2" {icon} on:click={onIconClick} />
		{/if}
	</div>
</div>
