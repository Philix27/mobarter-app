<script lang="ts">
	import { PinInput as InputOTPPrimitive } from 'bits-ui';
	import type { ComponentProps } from 'svelte';
	import { cn } from '$lib/utils.js';

	let {
		ref = $bindable(null),
		cell,
		class: className,
		...restProps
	}: ComponentProps<typeof InputOTPPrimitive.Cell> = $props();
</script>

<InputOTPPrimitive.Cell
	{cell}
	bind:ref
	class={cn(
		'border-secondary relative flex px-6 py-3 items-center justify-center border-y border-r text-md font-semibold transition-all first:rounded-l-md first:border-l last:rounded-r-md',
		cell.isActive && 'ring-ring ring-offset-background z-10 ring-2 text-foreground',
		className
	)}
	{...restProps}
>
	{cell.char}
	{#if cell.hasFakeCaret}
		<div class="pointer-events-none absolute inset-0 flex items-center justify-center">
			<div class="animate-caret-blink bg-foreground h-4 w-px duration-1000"></div>
		</div>
	{/if}
</InputOTPPrimitive.Cell>
