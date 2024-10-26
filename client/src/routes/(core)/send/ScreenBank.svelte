<script>
	import { Button, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
</script>

<div class="mt-4 w-full flex flex-col items-center justify-center gap-y-4">
	<TextInput place="Select Bank Account" />
	<TextInput place="Amount" inputType="number" />
	<Button onclick={() => toast('Funds sent')}>Hey man</Button>
</div>
