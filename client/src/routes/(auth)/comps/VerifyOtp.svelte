<script lang="ts">
	import { z } from 'zod';
	import { Button, TextInput } from 'components';

	let props = $props<{
		onSuccess?: VoidFunction;
	}>();
</script>

<div class="flex flex-col items-center w-full gap-y-4">
	<TextInput place="Enter OTP" inputType="number" />

	<Button onclick={props.onSuccess}>Verify OTP</Button>
</div>
