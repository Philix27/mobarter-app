<script>
	import P from '../../../components/P.svelte';
	import 'iconify-icon';
	const transactions = [
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		},
		{
			amount: 90,
			date: '23-10-09'
		}
	];
	const longList = [...transactions, ...transactions, ...transactions];
	// const oldObj = { a: 4, b: 9 };

	// const newObj = structuredClone(oldObj);
</script>

<div class="">
	<!-- <h1 class="mt-4 mb-2 font-bold">Transactions</h1> -->
	<div class="flex items-center justify-between mt-5 mb-3 mx-2">
		<P className="text-sm">Transactions</P>
		<a href="/transactions">
			<iconify-icon icon="material-symbols:history" class="text-xl text-primary" />
		</a>
	</div>
	<div class="overflow-y-scroll h-full max-h-[800px] no-scrollbar">
		{#each longList as item, index}
			<a href={`/transactions/${index}`}>
				<div class="flex items-center p-2 w-full bg-card mb-[2px] rounded-md">
					<div class="size-[40px] bg-background rounded-full flex items-center justify-center">
						<iconify-icon icon="mdi:dollar" class="text-xl" />
					</div>
					<div class="flex items-center justify-between p-2 w-full">
						<P>{item.amount.toString()}</P>
						<P>{item.date}</P>
					</div>
				</div>
			</a>
		{/each}
	</div>
</div>
