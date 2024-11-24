<script>
	import { P } from 'components';
	import { utilityCategories } from 'lib/category';

	// const newObj = structuredClone(oldObj);
</script>

<div class="grid grid-cols-4 gap-x-2 gap-y-2 w-full mb-5">
	{#each utilityCategories as item, index}
		<a
			href={item.link}
			class={`
			p-4 flex flex-col 
			items-center justify-center 
			shadow-md bg-accent
			rounded-sm cursor-pointer`}
		>
			<iconify-icon icon={item.icon} class="text-xl text-primary mb-2" />
			<P v="p6" className="pb-1 text-xs">{item.name}</P>
		</a>
	{/each}
</div>
