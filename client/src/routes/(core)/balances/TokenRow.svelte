<script lang="ts">
	import { browser } from '$app/environment';
	import { getBalance, account } from 'lib/web3';
	import { P } from 'components';
	import { TokenList } from 'celo-kit';

	let props = $props<{
		name: string;
		symbol: string;
		tokenAddress: string;
	}>();

	let isMiniPay = $state(false);
	let accountAddress = $state($account.address);
	if (browser) {
		if (window && window.ethereum) {
			// User has a injected wallet

			if (window.ethereum.isMiniPay) {
				// User is using Minipay
				isMiniPay = true;
				// Requesting account addresses
				let accounts = window.ethereum.request({
					method: 'eth_requestAccounts',
					params: []
				});

				// Injected wallets inject all available addresses,
				// to comply with API Minipay injects one address but in the form of array
				console.log(accounts[0]);
				// @ts-ignore
				// accountAddress = $account.address;
				accountAddress = accounts[0];
			}

			// User is not using MiniPay
		}
	}
</script>

<div class="flex items-center px-3 py-[2px] w-full bg-accent border-b-[0.5px] border-background">
	<div class="h-[30px] w-[30px] rounded-full flex items-center justify-center mr-2">
		<img src={`/tokens/${props.symbol}.svg`} alt="cUSD" />
	</div>
	<div class="flex items-center justify-between w-full">
		<div class="flex flex-col items-start p-2">
			<P className="text-sm">{props.name}</P>

			<P className="text-secondary text-xs mt-1 font-thin">Celo Network</P>
		</div>
		<div class="flex">
			<!-- {#await getBalance('0x765DE816845861e75A25fCA122bb6898B8B1282a', $account.address!)} -->
			{#await getBalance(props.tokenAddress, $account.address ?? accountAddress!)}
				<p>...</p>
			{:then data}
				<P className="text-sm text-primary">
					{data.value}
					{props.symbol}</P
				>
			{/await}
		</div>
	</div>
</div>
