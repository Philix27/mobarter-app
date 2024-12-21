<script lang="ts">
	import { account, modal } from 'lib/web3';
	import { Button } from 'components';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { globalStore } from 'lib/store';
	export const prerender = 'auto';
	$: accountAddress = $account.address;
	// let accountAddress = $state($account.address);
	$: isMiniPay = false;
	let deferredPrompt: Event;
	let deferredPrompt2: Event | any;

	if (browser) {
		window.addEventListener('beforeinstallprompt', (e) => {
			deferredPrompt2 = e;
		});

		if (window && window.ethereum) {
			// User has a injected wallet

			if (window.ethereum.isMiniPay) {
				// User is using Minipay
				isMiniPay = true;

				$globalStore.env = 'MINIPAY';

				// Requesting account addresses
				let accounts = window.ethereum.request({
					method: 'eth_requestAccounts',
					params: []
				});

				goto('/dashboard');

				// Injected wallets inject all available addresses,
				// to comply with API Minipay injects one address but in the form of array
				// @ts-ignore
				// accountAddress = $account.address;
				accountAddress = accounts[0];
				$globalStore.walletAddress = accounts[0];
			}

			// User is not using MiniPay
		}

		if ('serviceWorker' in navigator && 'PushManager' in window) {
			window.addEventListener('beforeinstallprompt', (e) => {
				e.preventDefault();

				deferredPrompt = e;
			});
		}
	}

	function connectWallet() {
		modal.open({ view: 'Connect' });
	}
</script>

<section class="h-screen flex flex-col items-center px-5">
	<img src="dollar.png" alt="Welcome" class="h-fit max-h-[300px] mt-[100px]" />
	<div class="w-full flex items-center justify-center flex-col">
		<p class="text-2xl font-bold">Seamless Crypto Trading</p>
		<p class="text-base text-muted text-center my-4">
			An easy, secure, and reliable way to trade crypto for local currency.
		</p>
	</div>
	<div class="mt-5 px-5 flex flex-col gap-y-5 items-center justify-between w-[80%] gap-x-4">
		<Button
			className="w-full"
			onclick={() => {
				goto('/login');
			}}>Login</Button
		>
		<Button
			variant="secondary"
			className="w-full"
			onclick={() => {
				goto('/create-account');
			}}>Create Account</Button
		>
	</div>
</section>
