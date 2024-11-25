<script lang="ts">
	import { Button, Nav } from 'components';
	import ValueSection from './ValueSection.svelte';
	import 'iconify-icon';

	// import { Mento } from '@mento-protocol/mento-sdk';
	// import { account } from 'lib/web3';
	import { JsonRpcProvider } from 'ethers';

	const getPairs = async () => {

		// const provider = new BrowserProvider('https://alfajores-forno.celo-testnet.org');
		// const provider = $account.connector?.getProvider();
		// const provider = new JsonRpcProvider('https://alfajores-forno.celo-testnet.org');
		// const mento = await Mento.create(provider);

		// const pairs = await mento.getTradeablePairs();
		// console.log('pairs', pairs);
	};

</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Swap" isBack />
	<!-- <Button onclick={() => {getPairs()}}>Try get pair</Button> -->
	<div class="w-full relative bg-background flex flex-col items-center justify-center">
		<div class="w-full">
			<ValueSection
				title="You send"
				isReadOnly={false}
				value=""
				onTokenClick={() => {}}
				onChange={(e) => {}}
			/>
			<div class="my-1 flex items-center justify-center absolute bottom-[50%] left-[45%]">
				<div class="bg-card px-3 py-2 rounded-lg border-4 border-background">
					<iconify-icon class="text-xl text-primary" icon="tdesign:swap" />
				</div>
			</div>
			<ValueSection
				title="You receive"
				isReadOnly={true}
				value=""
				onTokenClick={() => {}}
				onChange={(e) => {}}
			/>
		</div>
		<Button className="mt-4">Swap</Button>
	</div>
</div>
