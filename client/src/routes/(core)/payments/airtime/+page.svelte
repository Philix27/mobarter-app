<script lang="ts">
	import { BottomSheet, Button, Nav, P, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
	import { AirtimeData } from './data';
	import { cn } from 'lib/utils';
	import Row from './Row.svelte';

	type INetwork = 'MTN' | 'GLO' | 'AIRTEL';
	const getImgPath = (name: INetwork) => {
		if (name.trim().toUpperCase() === 'MTN') return '/networks/mtn2.png';
		if (name.trim().toUpperCase() === 'GLO') return '/networks/glo.png';
		if (name.trim().toUpperCase() === 'AIRTEL') return '/networks/airtel.png';
		return '/networks/mtn2.png';
	};

	let networkSelected: INetwork = 'MTN';
	$: showNetwork = false;
	$: amountSelected = 100;
	$: phoneValue = '';
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
		class="flex w-full items-center justify-between bg-secondary py-2 px-3 rounded-md"
	>
		<P className="text-sm">Select Network</P>
		<img src={getImgPath(networkSelected)} alt="" height={35} width={35} />
	</div>
	<div class="w-full grid grid-cols-4 gap-2 my-4">
		{#each AirtimeData['Nigeria'].amount as val}
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<!-- svelte-ignore a11y-no-static-element-interactions -->
			<div
				class={cn(
					`py-4 bg-secondary border-secondary border rounded-md flex items-center justify-center`,
					amountSelected === val && 'border-primary '
				)}
				on:click={() => {
					amountSelected = val;
				}}
			>
				<P className="text-[14px] font-semibold"
					>{`${AirtimeData['Nigeria'].symbol}${val.toString()}`}</P
				>
			</div>
		{/each}
	</div>
	<TextInput place="Amount" inputType="number" label="Amount" bind:value={amountSelected} />
	<TextInput
		place="Mobile number"
		inputType="number"
		label="Phone number"
		bind:value={phoneValue}
	/>
	<Button onclick={() => toast('Funds sent')}>Buy</Button>
</div>

<BottomSheet
	bind:show={showNetwork}
	onClose={() => {
		showNetwork = false;
	}}
	title="Networks"
>
	<Row
		title="MTN"
		imgSrc={getImgPath('MTN')}
		isActive={networkSelected === 'MTN'}
		onClick={() => {
			networkSelected = 'MTN';
		}}
	/>
	<Row
		title="Airtel"
		imgSrc={getImgPath('AIRTEL')}
		isActive={networkSelected === 'AIRTEL'}
		onClick={() => {
			networkSelected = 'AIRTEL';
		}}
	/>
	<Row
		title="Glo"
		imgSrc={getImgPath('GLO')}
		isActive={networkSelected === 'GLO'}
		onClick={() => {
			networkSelected = 'GLO';
		}}
	/>
</BottomSheet>
