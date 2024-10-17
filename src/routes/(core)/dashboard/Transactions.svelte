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
	<h1 class="mt-4 mb-2 font-bold">Transactions</h1>

	<div class="overflow-y-scroll h-full max-h-[800px] no-scrollbar">
		{#each longList as item}
			<div class="flex items-center p-2 w-full bg-card mb-1 rounded-md">
				<div class="size-[40px] bg-background rounded-full flex items-center justify-center">
					<iconify-icon icon="mdi:dollar" class="text-xl" />
				</div>
				<div class="flex items-center justify-between p-2 w-full">
					<P>{item.amount.toString()}</P>
					<P>{item.date}</P>
				</div>
			</div>
		{/each}
	</div>
</div>
