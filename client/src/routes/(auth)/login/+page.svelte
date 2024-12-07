<script>
	import Button from 'components/Button.svelte';
	import TextInput from 'components/TextInput.svelte';
	import AuthWrapper from '../AuthWrapper.svelte';
</script>

<AuthWrapper
	title="Login"
	desc="Hey, welcome back"
	footerLink={{ title: 'Create account', link: '/create-account' }}
>
	<div class="justify-center w-full flex flex-col items-center gap-y-5">
		<TextInput label="Email" place="Enter email" icon="hugeicons:mail-02" />

		<div class="w-full">
			<TextInput
				label="Password"
				place="Enter password"
				isPassword
				icon="hugeicons:square-lock-02"
			/>
			<div class="w-full flex mb-5 items-center justify-end">
				<a href="/forgot-password" class="text-sm underline">Forgot password</a>
			</div>
		</div>

		<Button>Login</Button>
	</div>
</AuthWrapper>
