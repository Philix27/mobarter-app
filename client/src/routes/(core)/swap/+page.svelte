<script lang="ts">
	import { Nav } from 'components';
	import ValueSection from './ValueSection.svelte';
	import 'iconify-icon';
</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div>
	<Nav title="Swap" isBack />
	<div class="w-full relative bg-background">
		<div>
			<ValueSection
				title="You send"
				isReadOnly={false}
				value=""
				onTokenClick={() => {}}
				onChange={(e) => {}}
			/>
			<div class="my-1 flex items-center justify-center absolute bottom-[40%] left-[45%]">
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
	</div>
</div>
