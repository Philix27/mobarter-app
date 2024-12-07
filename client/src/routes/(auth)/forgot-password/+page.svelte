<script lang="ts">
	import { goto } from '$app/navigation';
	import AuthWrapper from '../AuthWrapper.svelte';
	import SendOtp from '../comps/SendOtp.svelte';
	import VerifyOtp from '../comps/VerifyOtp.svelte';
	import EnterPassword from './EnterPassword.svelte';

	let activeScreen: 'sendOtp' | 'verifyOtp' | 'enterPassword' = 'sendOtp';
</script>

<AuthWrapper
	title="Forgot Password"
	desc="Oops, you forgot your password?"
	footerLink={{ title: 'Have you remembered your password?', link: '/login' }}
>
	<div class="flex items-center justify-center w-full flex-col">
		<div class="flex flex-col items-center w-full">
			{#if activeScreen === 'sendOtp'}
				<SendOtp onSuccess={() => (activeScreen = 'verifyOtp')} />
			{:else if activeScreen === 'verifyOtp'}
				<VerifyOtp onSuccess={() => (activeScreen = 'enterPassword')} />
			{:else if activeScreen === 'enterPassword'}
				<EnterPassword onSuccess={() => goto('/login')} />
			{/if}
		</div>
	</div>
</AuthWrapper>
