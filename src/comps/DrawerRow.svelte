<script>
	import { cn } from '../lib/utils';
	import P from '../components/P.svelte';
	import { drawerState } from '../store/settings';

	export let isActive = false;
	export let title;
	export let icon;

	function onItemClick() {
		$drawerState = false;
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
	class={cn('w-full flex items-center py-2 px-4 cursor-pointer', isActive && 'bg-accent')}
	on:click={onItemClick}
>
	<iconify-icon class="text-xl mb-1 mr-2 text-foreground" {icon} role="button" />
	<P text={title} className={'font-normal text-[14px]'} />
</div>
