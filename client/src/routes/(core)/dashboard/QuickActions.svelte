<script lang="ts">
	import { P } from 'components';

	let props = $props<{
		items: {
			name: string;
			link: string;
			icon: string;
		}[];
	}>();
</script>

<div class="grid grid-cols-4 gap-x-2 gap-y-2 w-full mb-5">
	{#each props.items as quick, i}
		<a
			href={quick.link}
			class={`
			px-4 pb-2 pt-4 flex flex-col 
			items-center justify-center 
			bg-accent shadow-md 
			rounded-sm cursor-pointer`}
		>
			<iconify-icon class="text-xl text-primary mb-4" icon={quick.icon} />
			<P v="p6">{quick.name}</P>
		</a>
	{/each}
</div>
