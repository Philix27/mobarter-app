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

<div class="pb-[20px]">
	{#each longList as item, index}
		<a href={`/transactions/${index}`}>
			<div class="flex items-center p-2 w-full bg-accent mb-[2px] rounded-md">
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
