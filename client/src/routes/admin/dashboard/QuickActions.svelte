<script lang="ts">
	import { P } from 'components';

	let quickActions = [
	
		{ name: 'Users', link: '/admin/users', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'KYC', link: '/admin/kyc', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Email', link: '/admin/email', icon: 'hugeicons:bitcoin-withdraw' },
		{ name: 'Rates', link: '/admin/rates', icon: 'majesticons:send-line' },
		{ name: 'Transactions', link: '/admin/transactions', icon: 'tdesign:swap' },
		{ name: 'Blocked', link: '/admin/users-blocked', icon: 'hugeicons:money-receive-circle' }
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
