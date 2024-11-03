<script lang="ts">
	import { Nav } from 'components';

	import { Button, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
</script>

<svelte:head>
	<title>Sign Up</title>
	<meta name="description" content="Create account" />
</svelte:head>

<div>
	<Nav title="Login" />
	<div class="mt-4 w-full flex flex-col items-center justify-center gap-y-4">
		<TextInput place="Email" />
		<TextInput place="Password" inputType="password" />
		<TextInput place="First name" />
		<TextInput place="Last name" />

		<Button onclick={() => toast('Funds sent')}>Sign Up</Button>
	</div>
</div>
