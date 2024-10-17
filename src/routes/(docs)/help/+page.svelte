<script lang="ts">
	import P from '../../../components/P.svelte';
	import Nav from '../../../comps/Nav.svelte';

	const pages: { name: string; link: string; icon: string; subtitle?: string }[] = [
		{
			name: 'About Us',
			link: '/about',
			icon: 'mdi:group',
			subtitle: 'About the company'
		},
		{
			name: 'Terms of service',
			link: '/tos',
			icon: 'mdi:lock',
			subtitle: 'Terms and Agreements'
		},
		{
			name: 'Privacy policy',
			link: '/pp',
			icon: 'mdi:safe',
			subtitle: 'How we use your data'
		}
	];
</script>

<svelte:head>
	<title>Help</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="">
	<Nav title="Help" />

	{#each pages as page}
		<a href={page.link}>
			<div class="w-full p-2 mb-1 bg-card flex rounded-md">
				<div class="size-[40px] rounded-full flex items-center justify-center bg-background mr-4">
					<iconify-icon class="text-md text-foreground" icon={page.icon} />
				</div>
				<div>
					<P className="font-light">{page.name}</P>
					{#if page.subtitle}
						<P className="text-sm mt-1 text-muted font-thin">{page.subtitle}</P>
					{/if}
				</div>
			</div>
		</a>
	{/each}
</div>
