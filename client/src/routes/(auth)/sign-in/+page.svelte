<script lang="ts">
	import { Nav } from 'components';

	import { Button, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
</script>

<svelte:head>
	<title>Login</title>
	<meta name="description" content="Signin to Mobarter" />
</svelte:head>

<div>
	<Nav title="Login" />
	<div class="mt-4 w-full flex flex-col items-center justify-center gap-y-4">
		<TextInput place="Email" />
		<TextInput place="Password" inputType="password" />

		<Button onclick={() => toast('Funds sent')}>Login</Button>
	</div>
</div>
