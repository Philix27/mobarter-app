<script lang="ts">
	import { Nav, P } from 'components';
	import { QRCodeImage } from 'svelte-qrcode-image';
	import { browser } from '$app/environment';
	import { account } from 'lib/web3';
	import { globalStore } from 'lib/store';
</script>

<Nav title="Receive" isBack />
<div class="">
	<div class="w-full flex items-center justify-center">
		<div class="size-[300px] bg-white p-2 rounded-lg my-5">
			{#if $globalStore.walletAddress}
				<QRCodeImage text={$globalStore.walletAddress || 'hello'} width={200} />
			{/if}
		</div>
	</div>

	<div class="flex flex-col items-center justify-between">
		<P>Wallet Address</P>
		<div class="w-[70%] p-2 break-words my-4 text-wrap bg-card rounded-md">
			{#if $globalStore.walletAddress}
				<P
					v="p5"
					className={` 
			text-sm text-center`}
				>
					{$globalStore.walletAddress!.toUpperCase()}
				</P>
			{/if}
		</div>
	</div>
</div>
