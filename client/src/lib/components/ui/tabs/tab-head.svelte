<script lang="ts">
	import { List, Trigger } from '.';
	let props = $props<{ list: string[] }>();
</script>

<List>
	{#each props.list as item}
		<Trigger value={item}>{item}</Trigger>
	{/each}
</List>
