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
		<Tabs.List>
			<Tabs.Trigger value={TabKeys.completed}>{TabKeys.completed}</Tabs.Trigger>
			<Tabs.Trigger value={TabKeys.success}>{TabKeys.success}</Tabs.Trigger>
			<Tabs.Trigger value={TabKeys.failed}>{TabKeys.failed}</Tabs.Trigger>
			<Tabs.Trigger value={TabKeys.pending}>{TabKeys.pending}</Tabs.Trigger>
		</Tabs.List>
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
