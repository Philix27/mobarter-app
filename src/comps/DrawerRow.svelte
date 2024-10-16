<script>
	import { cn } from '../lib/utils';
	import P from '../components/P.svelte';
	import { drawerState } from '../store/settings';
	import { goto } from '$app/navigation';

	export let isActive = false;
	export let title;
	export let icon;
	export let link = '/dashboard';

	function onItemClick() {
		$drawerState = false;
		goto(link);
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
	class={cn('w-full flex items-center py-2 px-4 cursor-pointer', isActive && 'bg-accent')}
	on:click={onItemClick}
>
	<iconify-icon class="text-xl mb-1 mr-2 text-foreground" {icon} role="button" />
	<P className={'font-normal text-[15px]'}>{title}</P>
</div>
