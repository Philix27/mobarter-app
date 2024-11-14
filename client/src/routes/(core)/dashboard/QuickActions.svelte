<script lang="ts">
	import { P } from 'components';

	let quickActions = [
		{ id: '6', name: 'Airtime', link: '/payments/airtime', icon: 'fluent:call-32-regular' },
		{ id: '7', name: 'Data', link: '/payments/data', icon: 'gg:data' },
		{ id: '8', name: 'Cable TV', link: '/payments/cable', icon: 'ic:outline-live-tv' },
		{ id: '9', name: 'Electricity', link: '/payments/electric', icon: 'lucide:utility-pole' },
		{ id: '10', name: 'Sports', link: '/payments/sports', icon: 'solar:gamepad-outline' },
		{ id: '4', name: 'Travel', link: '/payments/travel', icon: 'hugeicons:bitcoin-withdraw' },
		// { id: '5', name: 'Savings', link: '/savings', icon: 'material-symbols-light:savings-outline' },
		// { id: '4', name: 'Withdraw', link: '/withdraw', icon: 'hugeicons:bitcoin-withdraw' },
		// { id: '1', name: 'Send', link: '/send', icon: 'majesticons:send-line' },
		// { id: '2', name: 'Swap', link: '/swap', icon: 'tdesign:swap' },
		// { id: '3', name: 'Receive', link: '/receive', icon: 'hugeicons:money-receive-circle' }
	];
</script>

<div class="grid grid-cols-4 gap-x-2 gap-y-2 w-full mb-5">
	{#each quickActions as quick, i}
		<a
			href={quick.link}
			class="p-4 flex flex-col items-center justify-center bg-accent rounded-md cursor-pointer"
		>
			<iconify-icon class="text-xl text-primary mb-2" icon={quick.icon} />
			<P v="p6">{quick.name}</P>
		</a>
	{/each}
</div>
