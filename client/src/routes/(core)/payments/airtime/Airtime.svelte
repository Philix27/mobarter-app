<script lang="ts">
	import { BottomSheet, Button, P, TextInput, BottomRow } from 'components';
	import { AirtimeData } from 'lib/data.js';
	import { cn } from 'lib/utils';
	import { z } from 'zod';

	const formSchema = z.object({
		amountSelected: z.number({ message: 'Must be a number' }),
		phoneValue: z.number({ message: 'Not more than 11 letters' }).max(11)
	});

	const handleSubmit = () => {
		const result = formSchema.safeParse({ amountSelected, phoneValue });
		if (!result.success) {
			showConfirm = true;
			// zErrors = result.error.format().amountSelected?._errors;
			// let uop = parseFormError('', cerrors);
			// errors = result.error.errors.map((error) => {
			// 	return {
			// 		field: error.path[0],
			// 		message: error.message
			// 	};
			// });
			return;
		} else {
			//! Continue with form submission...
		}
	};

	type INetwork = 'MTN' | 'GLO' | 'AIRTEL';
	type IConfirm = 'Not Confirmed' | 'Confirm';

	const getImgPath = (name: INetwork) => {
		if (name.trim().toUpperCase() === 'MTN') return '/networks/mtn2.png';
		if (name.trim().toUpperCase() === 'GLO') return '/networks/glo.png';
		if (name.trim().toUpperCase() === 'AIRTEL') return '/networks/airtel.png';
		return '/networks/mtn2.png';
	};

	let networkSelected: INetwork = $state('MTN');
	let showNetwork = $state(false);
	let amountSelected = $state(100);
	let phoneValue = $state('');

	let showConfirm = $state(false);
</script>

<div class="w-full flex flex-col items-center justify-center gap-y-4">
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={() => {
			showNetwork = true;
		}}
		class="flex w-full items-center justify-between bg-card py-2 px-3 rounded-md"
	>
		<P className="text-sm">Select Network</P>
		<img src={getImgPath(networkSelected)} alt="" height={35} width={35} />
	</div>
	<div class="w-full grid grid-cols-4 gap-2 my-4">
		<!-- svelte-ignore legacy_code -->
		{#each AirtimeData['Nigeria'].amount as val}
			<!-- svelte-ignore legacy_code -->
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<div
				class={cn(
					`py-4 bg-card border-secondary 
					border rounded-md flex items-center justify-center`,
					amountSelected === val && 'border-primary '
				)}
				onclick={() => {
					amountSelected = val;
				}}
			>
				<P className="text-[14px] font-semibold"
					>{`${AirtimeData['Nigeria'].symbol}${val.toString()}`}</P
				>
			</div>
		{/each}
	</div>

	<TextInput
		place="Amount"
		inputType="number"
		label="Amount"
		required
		bind:value={amountSelected}
	/>
	<!-- errorMessage={errors['amountSelected']} -->
	<TextInput
		place="Mobile number"
		inputType="number"
		label="Phone number"
		required
		bind:value={phoneValue}
	/>
	<!-- onclick={() => toast('Funds sent')} -->
	<Button onclick={handleSubmit} btype="submit">Buy</Button>
</div>

<BottomSheet
	bind:show={showConfirm}
	onClose={() => {
		showConfirm = false;
	}}
	title="Confirm"
>
	<div class="w-full flex flex-col items-center justify-center gap-y-3">
		<p>You are about this thousand naira $7.98</p>

		<Button
			onclick={() => {
				showConfirm = false;
			}}
			btype="submit">Buy</Button
		>
	</div>
</BottomSheet>
<BottomSheet
	bind:show={showNetwork}
	onClose={() => {
		showNetwork = false;
		showConfirm = false;
	}}
	title="Networks"
>
	<BottomRow
		title="MTN"
		imgSrc={getImgPath('MTN')}
		isActive={networkSelected === 'MTN'}
		onClick={() => {
			networkSelected = 'MTN';
		}}
	/>
	<BottomRow
		title="Airtel"
		imgSrc={getImgPath('AIRTEL')}
		isActive={networkSelected === 'AIRTEL'}
		onClick={() => {
			networkSelected = 'AIRTEL';
		}}
	/>
	<BottomRow
		title="Glo"
		imgSrc={getImgPath('GLO')}
		isActive={networkSelected === 'GLO'}
		onClick={() => {
			networkSelected = 'GLO';
		}}
	/>
</BottomSheet>
