<script lang="ts">
	import { Nav } from 'components';
	import * as Tabs from '$lib/components/ui/tabs';
	import Content from '../comps/Content.svelte';

	const TabKeys = {
		completed: 'All',
		success: 'Success',
		pending: 'Pending',
		failed: 'Failed'
	};
</script>

<div>
	<Nav title="Orders" isBack />

	<Tabs.Root value={TabKeys.completed} class="w-full">
		<Tabs.Head list={[TabKeys.completed, TabKeys.success, TabKeys.failed, TabKeys.pending]} />
		<Tabs.Content value={TabKeys.completed}>
			<div>Completed</div>
		</Tabs.Content>
		<Tabs.Content value={TabKeys.success}>
			<div>Success</div>
		</Tabs.Content>
		<Tabs.Content value={TabKeys.pending}>
			<div>Pending</div>
		</Tabs.Content>
		<Tabs.Content value={TabKeys.failed}>
			<div>Failed</div>
		</Tabs.Content>
	</Tabs.Root>
</div>
