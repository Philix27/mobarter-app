<script lang="ts">
	import { Nav } from 'components';
	import { sendScreenState, type ISendTabs } from './store';
	import ScreenBank from './ScreenBank.svelte';
	import ScreenPhone from './ScreenPhone.svelte';
	import * as Tabs from '$lib/components/ui/tabs';

	const tabs: { title: ISendTabs; onClick: VoidFunction }[] = [
		{
			title: 'Wallet',
			onClick: () => {
				$sendScreenState = 'Wallet';
			}
		},
		{
			title: 'Bank',
			onClick: () => {
				$sendScreenState = 'Bank';
			}
		},
		{
			title: 'Phone',
			onClick: () => {
				$sendScreenState = 'Phone';
			}
		}
	];
</script>

<svelte:head>
	<title>Send money</title>
	<meta name="description" content="Send money" />
</svelte:head>

<div>
	<Nav title="Send" isBack />

	<div class="">
		<Tabs.Root value="Wallet" class="w-full">
			<Tabs.Head list={['Wallet', 'Bank', 'Phone']} />

			<Tabs.Content value="Wallet">
				<ScreenBank />
			</Tabs.Content>
			<Tabs.Content value="Bank">
				<ScreenPhone />
			</Tabs.Content>
			<Tabs.Content value="Phone">
				<ScreenPhone />
			</Tabs.Content>
		</Tabs.Root>
	</div>
</div>
