<script>
	import { Button, TextInput } from 'components';
	import { sendTransaction } from 'lib/web3';

	let amount = $state(100);
	let destinationWallet = $state('');

	const handleSubmit = () => {
		sendTransaction(destinationWallet, amount.toString());
		console.log("Send transactions clicked")
		// const result = formSchema.safeParse({ amountSelected, phoneValue });
		// if (!result.success) {
		// 	showConfirm = true;
		// 	// zErrors = result.error.format().amountSelected?._errors;
		// 	// let uop = parseFormError('', cerrors);
		// 	// errors = result.error.errors.map((error) => {
		// 	// 	return {
		// 	// 		field: error.path[0],
		// 	// 		message: error.message
		// 	// 	};
		// 	// });
		// 	return;
		// } else {
		// 	//! Continue with form submission...
		// }
	};
</script>

<div class="mt-4 w-full flex flex-col items-center justify-center gap-y-4">
	<TextInput
		place="Enter wallet address"
		desc="Ensure it is a verified wallet"
		bind:value={destinationWallet}
		required
	/>
	<TextInput place="Amount" inputType="number" bind:value={amount} required />
	<Button onclick={handleSubmit} btype="submit">Hey man</Button>
</div>
