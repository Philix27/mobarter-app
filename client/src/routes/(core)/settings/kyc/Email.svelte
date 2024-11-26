<script lang="ts">
	import { Button, P, TextInput } from 'components';
	import Wrapper from './Wrapper.svelte';

	let email = $state('');
	let otpCode = $state('');
	let props = $props<{}>();
</script>

<Wrapper title={'Email Verification'}>
	<TextInput place="example@gmail.com" label="Email address" required bind:value={email} />
	<TextInput label="OTP" place="Enter OTP code" inputType="number" required bind:value={otpCode} />

	<Button size="sm" className="my-2 w-fit">Submit</Button>
</Wrapper>
