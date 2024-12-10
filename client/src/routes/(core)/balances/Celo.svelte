<script lang="ts">
	import TokenRow from './TokenRow.svelte';
	import { TokenList, CELO } from 'celo-kit';
	import { browser } from '$app/environment';
	import { account, chainId } from 'lib/web3';

	let accountAddress = $state($account.address);
	if (browser) {
		if (window && window.ethereum) {
			let accounts = window.ethereum.request({
				method: 'eth_requestAccounts',
				params: []
			});

			console.log(accounts[0]);
			// @ts-ignore
			accountAddress = accounts[0];
		}
	}

	function getActiveChain() {
		if ($chainId === 42220) return 'celo';
		if ($chainId === 44787) return 'alfajores';
		if ($chainId === 62320) return 'baklava';
		return 'celo';
	}
</script>

<div>
	<TokenRow symbol={CELO.symbol} tokenAddress={'0x000'} name={CELO.name} />
	{#each TokenList.filter((val) => val.symbol.toUpperCase() !== 'CELO') as item}
		<TokenRow symbol={item.symbol} tokenAddress={item.address[getActiveChain()]} name={item.name} />
	{/each}
</div>
