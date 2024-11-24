<script lang="ts">
	import { Nav, P } from 'components';
	import { TokenList, CELO } from 'celo-kit';
	import { browser } from '$app/environment';
	import TokenRow from './TokenRow.svelte';

	$: accountAddress = '';
	$: isMiniPay = false;

	if (browser) {
		if (window && window.ethereum) {
			if (window.ethereum.isMiniPay) {
				isMiniPay = true;

				let accounts = window.ethereum.request({
					method: 'eth_requestAccounts',
					params: []
				});

				console.log(accounts[0]);
				// @ts-ignore
				accountAddress = accounts[0];
			}
		}
	}
</script>

<svelte:head>
	<title>Wallet Balance</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Balance" isBack />

	<div>
		<TokenRow symbol={CELO.symbol} tokenAddress={''} name={CELO.name} />
		{#each TokenList.filter((val) => val.symbol.toUpperCase() !== 'CELO') as item}
			<TokenRow symbol={item.symbol} tokenAddress={item.address.celo} name={item.name} />
		{/each}
	</div>
</div>
