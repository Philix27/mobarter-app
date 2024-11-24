<script lang="ts">
	import { Nav, P, Tabs, getTitle } from 'components';
	import { receiveState, type IReceiveTabs } from './store';
	import { QRCodeImage } from 'svelte-qrcode-image';
	import { browser } from '$app/environment';
	import { account } from 'lib/web3';

	const tabs: { title: IReceiveTabs; onClick: VoidFunction }[] = [
		{
			title: 'Wallet',
			onClick: () => {
				$receiveState = 'Wallet';
			}
		},
		{
			title: 'Buy',
			onClick: () => {
				$receiveState = 'Buy';
			}
		}
	];

	let isMiniPay = $state(false);
	let accountAddress = $state($account.address);
	if (browser) {
		if (window && window.ethereum) {
			// User has a injected wallet

			if (window.ethereum.isMiniPay) {
				// User is using Minipay
				isMiniPay = true;
				// Requesting account addresses
				let accounts = window.ethereum.request({
					method: 'eth_requestAccounts',
					params: []
				});

				// Injected wallets inject all available addresses,
				// to comply with API Minipay injects one address but in the form of array
				console.log(accounts[0]);
				// @ts-ignore
				// accountAddress = $account.address;
				accountAddress = accounts[0];
			}

			// User is not using MiniPay
		}
	}
</script>

<Nav title="Receive" isBack />
<div class="">
	<!-- <Tabs {tabs} activeTitle={getTitle(tabs, $receiveState)} /> -->
	<div class="w-full flex items-center justify-center">
		<div class="size-[300px] bg-white p-2 rounded-lg my-5">
			<QRCodeImage text="0x20F50b8832f87104853df3FdDA47Dd464f885a49" width={200} />
		</div>
	</div>

	<div class="flex flex-col items-center justify-between">
		<P>Wallet Address</P>
		<div class="w-[70%] p-2 break-words my-4 text-wrap bg-card rounded-md">
			<P
				v="p5"
				className={` 
			text-sm text-center`}
			>
				{accountAddress?.toUpperCase()}
			</P>
		</div>
	</div>
</div>
