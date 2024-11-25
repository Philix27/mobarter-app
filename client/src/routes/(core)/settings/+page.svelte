<script lang="ts">
	import { Nav, P } from 'components';
	import 'iconify-icon';

	const sections: {
		title: string;
		subtitle?: string;
		link: string;
		icon: string;
		onClick?: VoidFunction;
	}[] = [
		{
			title: 'Profile',
			link: '/settings/profile',
			icon: 'mdi:account'
		},
		{
			title: 'Verification',
			subtitle: 'KYC verification',
			link: '/settings/kyc',
			icon: 'mdi:account'
		},
		{
			title: 'Theme',
			link: '/settings/profile',
			icon: 'mdi:color'
		},
		{
			title: 'Bank accounts',
			subtitle: 'Manage all your bank accounts',
			link: '/settings/bank',
			icon: 'mdi:account'
		},

		{
			title: 'Version',
			subtitle: 'App version no: 0.0.1',
			link: '/settings/profile',
			icon: 'mdi:account'
		},
		{
			title: 'Logout',
			subtitle: 'Disconnect from Mobarter',
			link: '#',
			icon: 'mdi:logout'
		}
	];
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<Nav title="Settings" isBack icon="mdi:dollar" />
<div class="">
	{#each sections as item}
		<a href={item.link}>
			<div class="flex items-center p-2 w-full bg-card mb-[2px] rounded-md cursor-pointer">
				<div
					class="h-[40px] w-[50px] bg-background rounded-full flex items-center justify-center mr-2"
				>
					<iconify-icon icon={item.icon} class="text-xl" />
				</div>
				<div class="flex flex-col items-start p-2 w-full">
					<P>{item.title.toString()}</P>
					{#if item.subtitle}
						<P className="text-muted text-xs mt-1 font-thin">{item.subtitle}</P>
					{/if}
				</div>
			</div>
		</a>
	{/each}
</div>
