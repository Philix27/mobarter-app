<script lang="ts">
	import { Nav } from 'components';

	import { Button, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
</script>

<svelte:head>
	<title>Reset Password</title>
	<meta name="description" content="Reset Password" />
</svelte:head>

<div>
	<Nav title="Reset Password" />
	<div class="mt-4 w-full flex flex-col items-center justify-center gap-y-4">
		<TextInput place="Email" />

		<Button onclick={() => toast('Funds sent')}>Login</Button>
	</div>
</div>
