<script lang="ts">
	import { BottomSheet, Button, Nav, P, TextInput } from 'components';
	import { toast } from 'svelte-sonner';

	type INetwork = 'MTN' | 'GLO' | 'AIRTEL';
	const getImgPath = (name: INetwork) => {
		if (name.trim().toUpperCase() === 'MTN') return '/networks/mtn2.png';
		if (name.trim().toUpperCase() === 'GLO') return '/networks/glo.png';
		if (name.trim().toUpperCase() === 'AIRTEL') return '/networks/airtel.png';
		return '/networks/mtn2.png';
	};

	let networkSelected: INetwork = 'MTN';
	$: showNetwork = false;
</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="w-full flex flex-col items-center justify-center gap-y-4">
	<Nav title="Airtime" isBack />
	<!-- svelte-ignore missing-declaration -->
	<!-- svelte-ignore a11y-interactive-supports-focus -->
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div
		on:click={() => {
			showNetwork = true;
		}}
		class="flex w-full items-center justify-between mb-4"
	>
		<P className="text-primary">Select Network</P>
		<img src={getImgPath(networkSelected)} alt="" height={35} width={35} />
	</div>
	<TextInput place="Select Airtime" label="Select Airtime" />
	<TextInput place="Mobile number" inputType="number" label="Phone number" />
	<TextInput place="Amount" inputType="number" label="Amount" />
	<Button onclick={() => toast('Funds sent')}>Buy</Button>
</div>

<BottomSheet
	bind:show={showNetwork}
	onClose={() => {
		showNetwork = false;
	}}
	title="Networks"
>
	<div></div>
	<P>Hello</P>
</BottomSheet>
