<script lang="ts">
	import { goto } from '$app/navigation';
	import { Nav, P } from 'components';
	import QuickActions from './QuickActions.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div>
	<Nav theme={data.theme} showThemeToggle />
	<div class="">
		<div class=" min-h-screen h-screen overflow-y-scroll no-scrollbar scroll-smooth">
			<QuickActions />
		</div>
	</div>
</div>
