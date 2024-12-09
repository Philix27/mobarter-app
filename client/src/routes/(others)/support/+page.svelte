<script>
	import { Nav } from 'components';
	import { SupportChat } from './data';
	import TextInput from 'components/TextInput.svelte';
</script>

<svelte:head>
	<title>Support</title>
	<meta name="description" content="Customer's care" />
</svelte:head>

<div class="relative px-5 h-screen">
	<Nav title="Support" isBack />
	<div class=" overflow-y-scroll h-[calc(100vh-50px)] no-scrollbar pb-[100px] pt-[55px]">
		{#each SupportChat as item}
			<div class="w-full mb-2">
				{#if item.from === 'me'}
					<div class="flex items-center justify-between w-full">
						<div></div>
						<p class="bg-card px-2 rounded-md py-1 max-w-[80%] text-wrap break-words w-fit">
							{item.msg}
						</p>
					</div>
				{:else}
					<div class="flex items-center justify-between w-full">
						<p class="bg-secondary px-2 rounded-md py-1 max-w-[80%] w-fit text-wrap break-words">
							{item.msg}
						</p>
						<div></div>
					</div>
				{/if}
			</div>
		{/each}
	</div>
	<div class="absolute bottom-0 left-0 w-full px-6 py-2 bg-background ">
		<TextInput place="Enter message" />
	</div>
</div>
