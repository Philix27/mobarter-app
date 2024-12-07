<script lang="ts">
	import type { Component } from 'svelte';

	let props = $props<{
		title: string;
		desc: string;
		children: Component;
		footerComps?: Component;
		footerLink: {
			title: string;
			link: string;
		};
	}>();

</script>

<div class="min-h-screen flex flex-col px-6 py-5 mt-[20px]">
	<img src="/logo.jpg" class="size-[50px] mb-10" />
	<div class="w-full flex justify-center flex-col mb-4">
		<p class="text-3xl font-bold">{props.title}</p>
		<p class="text-lg text-muted my-4">
			{props.desc}
		</p>
	</div>
	<div>
		{@render props.children()}
	</div>

	<div class="mt-20 flex items-center justify-center w-full flex-col">
		<a href={props.footerLink.link} class="text-primary text-sm">{props.footerLink.title}</a>
	</div>
</div>
