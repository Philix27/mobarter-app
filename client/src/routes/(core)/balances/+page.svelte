<script lang="ts">
	import { Nav, P } from 'components';
	import { TokenList, CELO } from 'celo-kit';
	import { browser } from '$app/environment';
	import TokenRow from './TokenRow.svelte';
	import { chainId } from 'lib/web3';

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

	function getActiveChain() {
		if ($chainId === 42220) return 'celo';
		if ($chainId === 44787) return 'alfajores';
		if ($chainId === 62320) return 'baklava';
		return 'celo';
	}
</script>

<svelte:head>
	<title>Wallet Balance</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Balance" isBack />

	<div>
		<TokenRow symbol={CELO.symbol} tokenAddress={'0x000'} name={CELO.name} />
		{#each TokenList.filter((val) => val.symbol.toUpperCase() !== 'CELO') as item}
			<TokenRow
				symbol={item.symbol}
				tokenAddress={item.address[getActiveChain()]}
				name={item.name}
			/>
		{/each}
	</div>
</div>
