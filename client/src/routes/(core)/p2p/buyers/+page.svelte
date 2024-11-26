<script>
	import { Nav } from 'components';
	import * as Tabs from '$lib/components/ui/tabs';
</script>

<div>
	<Nav title="P2P Buyers" isBack />
	<div class="">
		<div>P2P Buyers</div>
	</div>
	<Tabs.Root value="account" class="w-full ">
		<Tabs.List>
			<Tabs.Trigger value="account">Account</Tabs.Trigger>
			<Tabs.Trigger value="password">Password</Tabs.Trigger>
		</Tabs.List>
		<Tabs.Content value="account">Make changes to your account here.</Tabs.Content>
		<Tabs.Content value="password">Change your password here.</Tabs.Content>
	</Tabs.Root>
</div>
