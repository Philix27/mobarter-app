<script lang="ts">
	import { Button, P, TextInput } from 'components';
	import Wrapper from './Wrapper.svelte';

	let phone = $state('');
	let otpCode = $state('');
	let props = $props<{}>();
</script>

<Wrapper title={'Phone Verification'}>
	<TextInput place="08101234567" label="Phone number" required bind:value={phone} />
	<TextInput label="OTP" place="Enter OTP code" inputType="number" required bind:value={otpCode} />

	<Button size="sm" className="my-2 w-fit">Submit</Button>
</Wrapper>
