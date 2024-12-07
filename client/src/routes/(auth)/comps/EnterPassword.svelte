<script lang="ts">
	import { Button, TextInput } from 'components';
	let props = $props<{ onSuccess?: VoidFunction }>();
</script>

<div class="flex flex-col items-center w-full gap-y-4">
	<TextInput label="Password" isPassword icon="hugeicons:square-lock-02" />
	<TextInput label="Confirm password" isPassword icon="hugeicons:square-lock-02" />

	<Button onclick={props.onSuccess}>Reset password</Button>
</div>
