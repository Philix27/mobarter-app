<script lang="ts">
	import { Nav } from 'components';
	import Transactions from './Transactions.svelte';
	import * as Tabs from '$lib/components/ui/tabs';
</script>

<div>
	<Nav title="Transactions" isBack />
	<Tabs.Root value="Wallet" class="w-full ">
		<Tabs.List>
			<Tabs.Trigger value="Wallet">Wallet</Tabs.Trigger>
			<Tabs.Trigger value="Swap">Swap</Tabs.Trigger>
			<Tabs.Trigger value="Airtime">Airtime</Tabs.Trigger>
		</Tabs.List>
		<Tabs.Content value="Wallet"><Transactions /></Tabs.Content>
		<Tabs.Content value="Swap"><Transactions /></Tabs.Content>
		<Tabs.Content value="Airtime"><Transactions /></Tabs.Content>
	</Tabs.Root>
</div>
