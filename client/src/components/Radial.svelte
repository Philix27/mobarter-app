<script>
	import { cn } from '$lib/utils';
	import 'iconify-icon';

	export let className = '';
	export let isChecked = true;
</script>

{#if isChecked}
	<iconify-icon class={cn('text-primary', className)} icon="mdi:close" />
{:else}
	<iconify-icon class={cn(className)} icon="mdi:close" />
{/if}
