<script>
	import P from './P.svelte';
	import 'iconify-icon';

	export let title = '';
	export let show = false;
	export let onClose = () => {};
</script>

{#if show}
	<div
		class="fixed top-0 left-0 bg-black/30 w-full bottom-0 z-30 h-screen flex flex-col justify-between"
	>
		<div class="bg-teal-600" />
		<motion.div
			initial={{ y: 2000 }}
			animate={{ y: 0 }}
			transition={{ ease: 'easeInOut', duration: 0.4 }}
			class="bg-background w-full p-4 rounded-t-[20px] min-h-[calc(30vh)] border-t border-secondary"
		>
			<div
				class={`
      w-full 
      rounded-md
      flex items-center justify-between
      px-2 pb-2`}
			>
				{#if title}
					<P>
						{title}
					</P>
				{:else}
					<div />
				{/if}

				<div
					class="rounded-[10px] h-[20px] w-[20px] p-[4px] bg-card flex items-center justify-center"
					on:click={onClose}
					role="button"
					tabindex={1}
				>
					<iconify-icon class="text-[15px]" icon="mdi:close" />
				</div>
			</div>
			<hr class="border-secondary" />
			<div class="px-2 py-5">
				<slot />
			</div>
		</motion.div>
	</div>
{/if}
