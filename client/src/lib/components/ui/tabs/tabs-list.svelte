<script lang="ts">
	import { Tabs as TabsPrimitive } from 'bits-ui';
	import { cn } from '$lib/utils.js';

	type $$Props = TabsPrimitive.ListProps;

	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<TabsPrimitive.List
	class={cn(
		' text-muted-foreground flex items-center justify-center w-full border-b border-secondary',
		className
	)}
	{...$$restProps}
>
	<slot />
</TabsPrimitive.List>
