<script lang="ts">
	import P from './P.svelte';
	import 'iconify-icon';
	import { goto } from '$app/navigation';
	import { drawerState } from '../store/settings';
	import { enhance } from '$app/forms';
	import { page } from '$app/stores';
	import type { SubmitFunction } from '@sveltejs/kit';

	const props = $props<{
		theme?: string;
		title?: string;
		icon?: string;
		isHome?: boolean;
		showThemeToggle?: boolean;
		isBack?: boolean;
		onIconClick?: () => void;
	}>();
	/**
	 * @param {string} route
	 * @param {boolean} replaceState
	 */
	// function routeToPage(route, replaceState) {
	// 	goto(`/${route}`, { replaceState });
	// }

	function goBack(defaultRoute = '/dashboard') {
		// function goBack() {
		if (props.isBack) {
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

	const submitUpdateTheme: SubmitFunction = ({ action }) => {
		const theme = action.searchParams.get('theme');

		if (theme) {
			// if (theme === 'light') {
			// 	document.documentElement.classList.replace('light', theme);
			// } else {
			// 	document.documentElement.classList.replace('dark', theme);
			// }

			document.documentElement.setAttribute('class', `color-scheme: ${theme}`);
			// document.documentElement.getElementsByTagName("html")
			// .classList.replace('dark', theme);
			// document.documentElement.classList.replace('color-scheme: light', `color-scheme: ${theme}`);

			document.documentElement.setAttribute('data-theme', theme);
		}
	};
</script>

<div
	class={`
        h-[50px] flex items-center 
        justify-between 
        fixed top-0 w-full  left-0
        z-10 bg-background px-5
      `}
>
	{#if props.isBack}
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

	{#if props.title}
		<P
			v="p4"
			className="text-card-foreground my-0 font-extrabold tracking-wide font-sans ml-4 text-sm"
			>{props.title}</P
		>
	{/if}

	<div class="flex items-center">
		{#if props.showThemeToggle}
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<!-- svelte-ignore a11y-interactive-supports-focus -->
			<form action="" method="POST" use:enhance={submitUpdateTheme}>
				{#if props.theme === 'light'}
					<button
						formaction="/?/setTheme&theme=dark&redirectTo={$page.url.pathname}"
						class="outline-none border-none p-0 mt-1 mr-4"
					>
						<iconify-icon icon="tabler:moon" class="text-xl text-foreground" />
					</button>
				{:else}
					<button
						class="outline-none border-none p-0 mt-1 mr-4"
						formaction="/?/setTheme&theme=light&redirectTo={$page.url.pathname}"
					>
						<iconify-icon icon="lets-icons:sun-light" class="text-xl text-foreground" />
					</button>
				{/if}
			</form>
		{/if}
		{#if props.isHome}
			<iconify-icon
				class="text-xl text-foreground mr-2"
				icon="system-uicons:bell"
				on:click={props.onIconClick}
			/>
		{/if}
		{#if props.icon}
			<iconify-icon
				class="text-xl text-foreground mr-2"
				icon={props.icon}
				on:click={props.onIconClick}
			/>
		{/if}
	</div>
</div>
