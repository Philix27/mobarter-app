<script lang="ts">
	// import { useDisconnect, useAppKit, useAppKitNetwork } from '@reown/appkit/react';
	import * as controller from '@reown/appkit';
	import { onMount } from 'svelte';

	// const btn = new W3mConnectButton();
	// Ad.WagmiAdapter;
	onMount(() => {
		controller.AccountController.fetchTokenBalance().then((res) => {
			console.log('Balance', res);
		});
	});
	const appK = new controller.W3mConnectView();

	// import { useDisconnect, useAppKit, useAppKitNetwork,  } from '@reown/appkit';
	// import { networks } from '@/config';
	// const { disconnect } = useDisconnect();
	// const { open } = useAppKit();
	// const { switchNetwork } = useAppKitNetwork();
</script>

<div>
	<p>appkit-button</p>
	<appkit-button />

	<p>Action Button List</p>
	<!-- <button
		onclick={() => {
			open();
		}}>Open</button
	>
	<button
		onclick={() => {
			disconnect();
		}}>Disconnect</button
	> -->
	<!-- <button onClick={() => switchNetwork(networks[1])}>Switch</button> -->
</div>
