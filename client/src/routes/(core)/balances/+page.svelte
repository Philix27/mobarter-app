<script lang="ts">
	import { Nav, P } from 'components';
	import { TokenList } from 'celo-kit';
	import { getBalance, account } from 'lib/web3';
</script>

<svelte:head>
	<title>Wallet Balance</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Balance" isBack />

	<div>
		{#each TokenList as item}
			<div
				class="flex items-center px-3 py-[2px] w-full bg-accent border-b-[0.5px] border-background"
			>
				<div class="h-[30px] w-[30px] rounded-full flex items-center justify-center mr-2">
					<img src={`/tokens/${item.symbol}.svg`} alt="cUSD" />
				</div>
				<div class="flex items-center justify-between w-full">
					<div class="flex flex-col items-start p-2">
						<P className="text-sm">{item.name}</P>

						<P className="text-secondary text-xs mt-1 font-thin">Celo Network</P>
					</div>
					<div class="flex">
						<!-- {#await $getBalance("0x3432", "0x3432")}
							<p>... </p>
						{:then value}
						<p>Done</p>
						{/await} -->
						<P className="text-sm text-primary font-extralight">0.12 {item.symbol}</P>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
