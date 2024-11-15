<script>
	import { P } from 'components';
	import { utilityCategories } from 'lib/category';

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

	// const newObj = structuredClone(oldObj);
</script>

<div class="">
	{#each utilityCategories as item, index}
		<a href={item.link}>
			<div
				class="flex items-center px-2 py-[2px] w-full bg-accent border-b-[0.5px] border-background"
			>
				<div
					class="h-[40px] w-[43px] bg-background rounded-full flex items-center justify-center mr-2"
				>
					<iconify-icon icon={item.icon} class="text-xl text-primary" />
				</div>
				<div class="flex flex-col items-start p-2 w-full">
					<P className="pb-1 text-sm">{item.name}</P>
					<P className=" text-[12px] text-secondary ">{item.description}</P>
				</div>
			</div>
		</a>
	{/each}
</div>
