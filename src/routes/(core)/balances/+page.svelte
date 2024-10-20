<script lang="ts">
	import { Nav, P } from 'components';
	import { TokenList } from 'celo-kit';
</script>

<svelte:head>
	<title>Wallet Balance</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Balance" isBack />

	<div>
		{#each TokenList as item}
			<div class="flex items-center px-3 py-1 w-full bg-card mb-[2px] rounded-md">
				<div class="h-[40px] w-[50px] rounded-full flex items-center justify-center mr-2">
					<img src={`/tokens/${item.symbol}.svg`} alt="cUSD" />
				</div>
				<div class="flex items-center justify-between w-full">
					<div class="flex flex-col items-start p-2">
						<P>{item.name}</P>

						<P className="text-muted text-xs mt-1 font-thin">Celo Network</P>
					</div>
					<div class="flex">
						<P className="text-sm text-primary font-extralight">0.12 {item.symbol}</P>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
