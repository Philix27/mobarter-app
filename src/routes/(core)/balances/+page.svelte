<script lang="ts">
	import { Nav, P } from 'components';
	import { TokenList } from 'celo-kit';

	// const tokenBalances: {
	// 	title: string;
	// } = [];
</script>

<svelte:head>
	<title>Wallet Balance</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Balance" isBack />

	<div>
		{#each TokenList as token}
			<div class='p-2 border-b-[0.1px] border-secondary flex items-center justify-between'>
				<P>{token.name}</P>
				<P className="text-sm text-primary font-extralight">0.12 {token.symbol}</P>
			</div>
		{/each}
	</div>
</div>

<style lang="postcss">
</style>
