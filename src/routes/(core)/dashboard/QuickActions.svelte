<script>
	import 'iconify-icon';
	import { P } from 'components';

	let quickActions = [
		{ id: '1', name: 'Send', link: '/send', icon: 'mdi:send' },
		{ id: '2', name: 'Swap', link: '/swap', icon: 'tdesign:swap' },
		{ id: '3', name: 'Receive', link: '/receive', icon: 'hugeicons:money-receive-circle' },
		{ id: '4', name: 'Withdraw', link: '/withdraw', icon: 'hugeicons:bitcoin-withdraw' },
		{ id: '5', name: 'Savings', link: '/savings', icon: 'material-symbols-light:savings-outline' },
		{ id: '6', name: 'Airtime', link: '/airtime', icon: 'mdi:phone' },
		{ id: '7', name: 'Utility', link: '/utility', icon: 'lucide:utility-pole' }
	];
</script>

<div class="grid grid-cols-4 gap-x-2 gap-y-2 w-full mb-5">
	{#each quickActions as quick, i}
		<a
			href={quick.link}
			class="p-4 flex flex-col items-center justify-center bg-accent rounded-md cursor-pointer"
		>
			<iconify-icon class="text-xl text-primary mb-2" icon={quick.icon} />
			<!-- <p>{quick.name}</p> -->
			<P v="p6">{quick.name}</P>
		</a>
	{/each}
</div>
