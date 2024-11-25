<script lang="ts">
	import { Nav } from 'components';
	import Bvn from './Bvn.svelte';
	import type { Component } from 'svelte';

	const stages: {
		title: string;
		Comp: Component;
	}[] = [
		{ title: 'Phone', Comp: Bvn },
		{ title: 'Email', Comp: Bvn },
		{ title: 'BVN', Comp: Bvn },
		{ title: 'NIN', Comp: Bvn },
		{ title: 'Passport', Comp: Bvn }
	];

	const ll = Bvn;
</script>

<div>
	<Nav title="Verification" isBack />
	<div class="">
		{#each stages as { title, Comp }}
			<div>{title}</div>
			<Comp />
		{/each}
	</div>
</div>
