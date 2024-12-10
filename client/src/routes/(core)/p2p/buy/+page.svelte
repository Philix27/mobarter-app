<script lang="ts">
	import { Nav } from 'components';
	import * as Tabs from '$lib/components/ui/tabs';
	import Content from '../comps/Content.svelte';
</script>

<div>
	<Nav title="Buy coin" isBack />

	<Tabs.Root value="CUSD" class="w-full" >
		<Tabs.List>
			<Tabs.Trigger value="CUSD">CUSD</Tabs.Trigger>
			<Tabs.Trigger value="USDT">USDT</Tabs.Trigger>
			<Tabs.Trigger value="CELO">CELO</Tabs.Trigger>
		</Tabs.List>
		<Tabs.Content value="CUSD">
			<Content token="CUSD" tradeType="BUY" />
		</Tabs.Content>
		<Tabs.Content value="USDT">
			<Content token="USDT" tradeType="BUY" />
		</Tabs.Content>
		<Tabs.Content value="CELO">
			<Content token="CELO" tradeType="BUY" />
		</Tabs.Content>
	</Tabs.Root>
</div>
