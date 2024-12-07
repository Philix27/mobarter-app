<script lang="ts">
	import { goto } from '$app/navigation';
	import AuthWrapper from '../AuthWrapper.svelte';
	import EnterPassword from '../comps/EnterPassword.svelte';
	import SendOtp from '../comps/SendOtp.svelte';
	import VerifyOtp from '../comps/VerifyOtp.svelte';
	let activeScreen: 'sendOtp' | 'verifyOtp' | 'enterPassword' = 'sendOtp';
</script>

<AuthWrapper
	title="Create Account"
	desc="Let's get started?"
	footerLink={{ title: 'Already have an account', link: '/login' }}
>
	<div class="flex items-center justify-center w-full flex-col">
		{#if activeScreen === 'sendOtp'}
			<SendOtp onSuccess={() => (activeScreen = 'verifyOtp')} />
		{:else if activeScreen === 'verifyOtp'}
			<VerifyOtp onSuccess={() => (activeScreen = 'enterPassword')} />
		{:else if activeScreen === 'enterPassword'}
			<EnterPassword onSuccess={() => goto('/dashboard')} />
		{/if}
	</div>
</AuthWrapper>
