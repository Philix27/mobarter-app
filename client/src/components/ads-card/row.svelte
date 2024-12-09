<script lang="ts">
	import { cn } from '$lib/utils.js';

	let props = $props<{ className?: string; title: string; value: string }>();
</script>

<div class={cn('flex items-center justify-between w-full border-b-[0.5px] py-1', props.className)}>
	<p class="text-sm">{props.title}</p>
	<p class="text-sm">{props.value}</p>
</div>
