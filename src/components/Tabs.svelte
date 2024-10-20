<script lang="ts">
	import { cn } from '$lib/utils';

	import P from './P.svelte';

	export let activeTitle = '';
	export let className = '';
	export let tabs: { title: string; onClick: VoidFunction }[] = [];
</script>

<div class={cn('w-full', className)}>
	<div class="flex items-center justify-between border-secondary border-2 rounded-md">
		{#each tabs as { title, onClick }, i}
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<!-- svelte-ignore a11y-no-static-element-interactions -->
			<div
				class={cn(
					'flex items-center justify-center bg-background w-full px-3 py-2 rounded-md',
					activeTitle === title && 'bg-card rounded-lg'
				)}
				on:click={onClick}
			>
				<P className={cn(activeTitle === title ? 'text-primary' : 'text-foreground')} v="p5"
					>{title}</P
				>
			</div>
		{/each}
	</div>
</div>

<!-- isActive && 'bg-background rounded-lg' -->
