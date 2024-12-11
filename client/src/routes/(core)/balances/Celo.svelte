<script lang="ts">
	import TokenRow from './TokenRow.svelte';
	import { TokenList, CELO } from 'celo-kit';
	import { chainId } from 'lib/web3';

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
