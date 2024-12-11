<script lang="ts">
	import { Nav } from 'components';
	import * as Tabs from '$lib/components/ui/tabs';
	import Airtime from './Airtime.svelte';
	import DataPlan from './DataPlan.svelte';

	const TabKeys = {
		tab1: 'Airtime',
		tab2: 'Data'
	};
</script>

<div>
	<Nav title="Airtime" isBack />

	<Tabs.Root value={TabKeys.tab1} class="w-full">
		<Tabs.Head list={[TabKeys.tab1, TabKeys.tab2]} />
		<Tabs.Content value={TabKeys.tab1}>
			<Airtime />
		</Tabs.Content>
		<Tabs.Content value={TabKeys.tab2}>
			<DataPlan />
		</Tabs.Content>
	</Tabs.Root>
</div>
