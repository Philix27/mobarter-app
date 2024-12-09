<script lant="ts">
	import { Nav } from 'components';
	import * as Tabs from '$lib/components/ui/tabs';
	import Content from '../comps/Content.svelte';
</script>

<div>
	<Nav title="Sell your token" isBack />

	<Tabs.Root value="account" class="w-full ">
		<Tabs.List>
			<Tabs.Trigger value="CUSD">CUSD</Tabs.Trigger>
			<Tabs.Trigger value="USDT">USDT</Tabs.Trigger>
			<Tabs.Trigger value="CELO">CELO</Tabs.Trigger>
		</Tabs.List>
		<Tabs.Content value="CUSD">
			<Content token="CUSD" tradeType="SELL" />
		</Tabs.Content>
		<Tabs.Content value="USDT">
			<Content token="USDT" tradeType="SELL" />
		</Tabs.Content>
		<Tabs.Content value="CELO">
			<Content token="CELO" tradeType="SELL" />
		</Tabs.Content>
	</Tabs.Root>
</div>
