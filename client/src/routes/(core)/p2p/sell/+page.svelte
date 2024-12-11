<script lant="ts">
	import { Nav } from 'components';
	import * as Tabs from '$lib/components/ui/tabs';
	import Content from '../comps/Content.svelte';
</script>

<div>
	<Nav title="Sell your token" isBack />

	<Tabs.Root value="CUSD" class="w-full ">
		<Tabs.Head list={['CUSD', 'USDT', 'CELO']} />

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
