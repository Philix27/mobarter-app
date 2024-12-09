<script lang="ts">
	import { ListOfAgentsAds, type CryptoCurrencies, type TradeType } from './data';
	import * as Ads from 'components/ads-card';

	let props = $props<{ token: CryptoCurrencies; tradeType: TradeType }>();

	const isValidToken = (val: CryptoCurrencies) =>
		val.toString().toLowerCase() === props.token.toString().toLowerCase();

	const isValidTradeType = (val: TradeType) =>
		val.toString().toLowerCase() === props.tradeType.toString().toLowerCase();

	const filterOptions = ['Country', 'Bank Transfer'];

	let ads = $derived(
		ListOfAgentsAds.filter((val) => {
			return isValidToken(val.cryptoCurrency);
		})
		// return val.cryptoCurrency.toString().toLowerCase() === props.token.toString().toLowerCase();
	);
</script>

<div class="h-screen overflow-y-scroll no-scrollbar pb-[200px]">
	<!-- <div class="h-[50px]">
		Hello
		{ads.length}
	</div> -->
	{#each ads as item}
		<Ads.Card>
			<Ads.Row title="Name" value={item.name} />
			<Ads.Row title="Limit" value={`${item.limitLower} - ${item.limitUpper}`} />
			<Ads.Row title="Cryptocurrency" value={item.cryptoCurrency} />
			<Ads.Row title="Limit" value={`${item.limitLower} - ${item.limitUpper}`} />
			<Ads.Row title="Fiat" value={item.fiatCurrency} />
			<Ads.Row title="Success Rate" value={item.successRate.toString()} />
		</Ads.Card>
	{/each}
</div>
