<script lang="ts">
	import { BottomSheet, Button, Nav, P, TextInput, BottomRow } from 'components';
	import { AirtimeData } from 'lib/data.js';
	import { cn } from 'lib/utils';
	import { z } from 'zod';
	import type { GetDataPlansInput, GetDataPlansQuery } from 'generated/graphql';
	import { GetDataPlansDocument, NetworkProviders } from 'generated/graphql';
	import { getContextClient, queryStore } from '@urql/svelte';

	let dataPlans = $derived(
		queryStore<GetDataPlansQuery, GetDataPlansInput>({
			client: getContextClient(),
			query: GetDataPlansDocument,
			variables: {
				network: NetworkProviders.Mtn,
				category: 'None'
			},
			pause: false
		})
	);

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

	let { data, form } = $props();

	let networkSelected: INetwork = $state('MTN');
	let dataPlanSelected: string = $state('');
	let showNetwork = $state(false);
	let amountSelected = $state(100);
	let phoneValue = $state('');

	let showConfirm = $state(false);
	let showDataPlans = $state(false);
</script>

<svelte:head>
	<title>Data Plans</title>
</svelte:head>

<div class="w-full flex flex-col items-center justify-center gap-y-4">
	<Nav title="Data Plans" isBack />
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={() => {
			showNetwork = true;
		}}
		class="flex w-full items-center justify-between bg-secondary py-2 px-3 rounded-md"
	>
		<P className="text-sm">Select Network</P>
		<img src={getImgPath(networkSelected)} alt="" height={35} width={35} />
	</div>
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onclick={() => {
			showDataPlans = true;
		}}
		class="flex w-full items-center justify-between bg-secondary py-4 px-3 rounded-md"
	>
		<P className="text-sm">Select Plan</P>
		<P className="text-sm font-semibold ">{dataPlanSelected}</P>
	</div>

	<TextInput
		place="Amount"
		inputType="number"
		label="Amount"
		isReadOnly
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
<BottomSheet
	bind:show={showDataPlans}
	onClose={() => {
		showDataPlans = false;
	}}
	title="Data Plans"
>
	{#if $dataPlans.fetching}
		<p>Loading dataPlans....</p>
	{:else}
		<div class="flex flex-col overflow-y-scroll scroll-smooth gap-4">
			{#each $dataPlans.data!.Airtime_GetDataPlans!.dataPlans!.filter((val) => val!.network === networkSelected) as item}
				<BottomRow
					title={item!.plan!}
					imgSrc={getImgPath(networkSelected)}
					isActive={dataPlanSelected === item?.plan}
					onClick={() => {
						dataPlanSelected = item!.plan!;
						amountSelected = parseInt(item!.amount!);
					}}
				/>
			{/each}
		</div>
	{/if}
</BottomSheet>
