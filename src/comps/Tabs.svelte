<script>
	import { cn } from '$lib/utils';

	import P from '../components/P.svelte';

	export let className = '';
	export let tabs = [{ title: '', isActive: false, onClick: () => {} }];
</script>

<div class={cn('w-full', className)}>
	<div class="flex items-center justify-between border-secondary border-2 rounded-md">
		{#each tabs as { isActive, title, onClick }, i}
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<!-- svelte-ignore a11y-no-static-element-interactions -->
			<div
				class={cn(
					'flex items-center justify-center bg-secondary w-full px-3 py-2 rounded-md',
					isActive && 'bg-background rounded-lg'
				)}
				on:click={onClick}
			>
				<P className={cn(isActive && 'text-primary')} variant="p5">{title}</P>
			</div>
		{/each}
	</div>
</div>
