<script>
	import P from '../components/P.svelte';
	import 'iconify-icon';
	import { goto } from '$app/navigation';

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
</script>

<div
	class={`
        h-[50px] flex items-center 
        justify-between 
        fixed top-0 w-full 
        z-10 bg-background
      `}
>
	<div class="flex items-center">
		{#if isBack}
			<iconify-icon class="text-xl text-foreground" icon="mingcute:left-line" on:click={goBack} />
		{:else}
			<iconify-icon class="text-xl mb-2 text-foreground" icon="mdi:menu" />
		{/if}
	</div>

	<P
		variant="p4"
		className="text-card-foreground my-0 font-extrabold tracking-wide font-sans"
		text={title}
	/>
	<div class="flex items-center gap-x-3">
		{#if icon}
			<button type="button" on:click={onIconClick}>
				<iconify-icon class="text-xl mb-2 text-foreground" {icon} />
			</button>
		{/if}
	</div>
</div>
