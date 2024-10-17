<script>
	import { Nav, P, Tabs } from 'components';
	import { receiveState } from './store';
	import { QRCodeImage } from 'svelte-qrcode-image';

	const tabs = [
		{
			title: 'Wallet',
			isActive: $receiveState === 'WALLET',
			onClick: () => {
				$receiveState = 'WALLET';
			}
		},
		{
			title: 'Buy',
			isActive: $receiveState === 'BUY',
			onClick: () => {
				$receiveState = 'BUY';
			}
		}
	];
</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<Nav title="Receive" isBack />
<div class="">
	<Tabs {tabs} />
	<div class="w-full flex items-center justify-center">
		<div class="size-[300px] bg-white p-2 rounded-lg my-5">
			<QRCodeImage text="0x20F50b8832f87104853df3FdDA47Dd464f885a49" width={200} />
		</div>
	</div>

	<div class="flex items-center justify-between">
		<P>Wallet Address</P>
		<div class="max-w-[50%] p-2 break-words my-4 text-wrap">
			<P
				variant="p5"
				className={`truncate text-ellipsis 
			text-sm max-w-[90%] text-center`}
			>
				0x20F50b8832f87104853df3FdDA47Dd464f885a49
			</P>
		</div>
	</div>
</div>
