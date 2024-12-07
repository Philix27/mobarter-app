<script lang="ts">
	import { Button, TextInput } from 'components';

	let props = $props<{ onSuccess?: VoidFunction }>();
</script>

<div class="flex flex-col items-center w-full gap-y-4">
	<TextInput place="Enter Email" icon="hugeicons:mail-02" />
	<Button onclick={props.onSuccess}>Send OTP</Button>
</div>
