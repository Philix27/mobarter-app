<script lang="ts">
	import { P } from 'components';

	let quickActions = [
		// { name: 'Airtime', link: '/payments/airtime', icon: 'fluent:call-32-regular' },
		// { name: 'Data', link: '/payments/data', icon: 'gg:data' },
		// { name: 'Cable TV', link: '/payments/cable', icon: 'ic:outline-live-tv' },
		// { name: 'Electricity', link: '/payments/electric', icon: 'lucide:utility-pole' },
		// { name: 'Sports', link: '/payments/sports', icon: 'solar:gamepad-outline' },
		// { name: 'Travel', link: '/payments/travel', icon: 'hugeicons:bitcoin-withdraw' }
		// { name: 'Savings', link: '/savings', icon: 'material-symbols-light:savings-outline' },
		{ name: 'Buy', link: '/p2p/sellers', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Sell', link: '/p2p/buyers', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Withdraw', link: '/withdraw', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Send', link: '/send', icon: 'majesticons:send-line' },
		{ name: 'Swap', link: '/swap', icon: 'tdesign:swap' },
		{ name: 'Receive', link: '/receive', icon: 'hugeicons:money-receive-circle' }
	];
</script>

<div class="grid grid-cols-4 gap-x-2 gap-y-2 w-full mb-5">
	{#each quickActions as quick, i}
		<a
			href={quick.link}
			class={`
			p-4 flex flex-col 
			items-center justify-center 
			bg-accent shadow-md 
			rounded-sm cursor-pointer`}
		>
			<iconify-icon class="text-xl text-primary mb-2" icon={quick.icon} />
			<P v="p6">{quick.name}</P>
		</a>
	{/each}
</div>
