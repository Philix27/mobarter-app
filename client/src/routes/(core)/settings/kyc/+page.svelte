<script lang="ts">
	import { Nav } from 'components';
	import Bvn from './Bvn.svelte';
	import Email from './Email.svelte';
	import * as Accordion from '$lib/components/ui/accordion';
	import Nin from './Nin.svelte';
	import Phone from './Phone.svelte';
	import Names from './Names.svelte';
</script>

<div>
	<Nav title="Verification" isBack />
	<div class="overflow-y-scroll max-h-screen scroll-smooth no-scrollbar">
		<Accordion.Root>
			<Names />
			<Email />
			<Phone />
			<Bvn />
			<Nin />
		</Accordion.Root>
	</div>
</div>

<style>
	/* Hide scrollbar for Chrome, Safari and Opera */
	.no-scrollbar::-webkit-scrollbar {
		display: none;
	}

	/* Hide scrollbar for IE, Edge and Firefox */
	.no-scrollbar {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}
</style>
