<script lang="ts">
	import { cn } from '$lib/utils.js';

	let className: string | undefined | null = undefined;
	export { className as class };
</script>

<div
	class={cn(
		` 
		bg-card w-full rounded-md p-2 mb-1 shadow-sm
		`,
		className
	)}
>
	<slot />
</div>
