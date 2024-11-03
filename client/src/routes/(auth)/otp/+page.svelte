<script lang="ts">
	import { Nav } from 'components';

	import { Button, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
</script>

<svelte:head>
	<title>OTP</title>
	<meta name="description" content="Verify OTP" />
</svelte:head>

<div>
	<Nav title="Verify OTP" />
	<div class="mt-4 w-full flex flex-col items-center justify-center gap-y-4">
		<TextInput place="number" inputType="number" />

		<Button onclick={() => toast('Funds sent')}>Sign Up</Button>
	</div>
</div>
