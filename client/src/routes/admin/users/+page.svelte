<script lang="ts">
	import { Nav, P } from 'components';
	import { usersList } from './data';
	import Modal from './Modal.svelte';
	import Kyc from './Kyc.svelte';
	import Transactions from './Transactions.svelte';
	import Contact from './Contact.svelte';

	let props = $state<{
		showModal: 'KYC' | 'Transactions' | 'INFO' | 'CLOSE';
	}>({
		showModal: 'CLOSE'
	});
</script>

<div>
	<Nav showThemeToggle />
	<div class="">
		<div class=" min-h-screen h-screen overflow-y-scroll no-scrollbar scroll-smooth">
			{#each usersList as item}
				<div class="bg-card p-2 border-b border hover:bg-accent flex items-center justify-between">
					<div
						class="w-[25%]"
						onclick={() => {
							props.showModal = 'INFO';
						}}
					>
						<P className="text-sm">{item.firstName} {item.lastName}</P>
					</div>
					<div class="w-[25%] flex gap-x-2">
						<iconify-icon class="text-xl text-primary mb-2" icon={'hugeicons:bitcoin-withdraw'} />
						<iconify-icon class="text-xl text-primary mb-2" icon={'hugeicons:bitcoin-withdraw'} />
						<iconify-icon class="text-xl text-primary mb-2" icon={'hugeicons:bitcoin-withdraw'} />
						<iconify-icon class="text-xl text-primary mb-2" icon={'hugeicons:bitcoin-withdraw'} />
					</div>
					<div
						class="w-[25%] grid grid-cols-2"
						onclick={() => {
							props.showModal = 'KYC';
						}}
					>
						<P className="text-[10px] mb-1 text-muted">NIN</P>
						<P className="text-[10px] mb-1 text-muted">BVN</P>
						<P className="text-[10px] mb-1 text-muted">PASSPORT</P>
						<P className="text-[10px] mb-1 text-muted">ADDRESS</P>
					</div>
					<div class="w-[25%] flex items-end justify-items-end">
						<div>
							<P className="text-xs mb-2 text-muted">{item.email}</P>
							<P className="text-xs text-muted">{item.phone}</P>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
<Modal
	title="KYC"
	show={props.showModal === 'KYC'}
	onClose={() => {
		props.showModal = 'CLOSE';
	}}
>
	<Kyc /></Modal
>
<Modal
	title="Transactions"
	show={props.showModal === 'Transactions'}
	onClose={() => {
		props.showModal = 'CLOSE';
	}}><Transactions /></Modal
>
<Modal
	title="More Info"
	show={props.showModal === 'INFO'}
	onClose={() => {
		props.showModal = 'CLOSE';
	}}><Contact /></Modal
>
