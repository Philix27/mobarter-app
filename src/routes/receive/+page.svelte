<script>
	import P from '../../components/P.svelte';
	import Nav from '../../comps/Nav.svelte';
</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<Nav title="Receive" isBack />
<div class="mt-[50px]">
	<div class="flex items-center justify-between">
		<P text="Wallet address" />
		<div class="max-w-[50%] p-2 break-words my-4 text-wrap">
			<P
				text="0x20F50b8832f87104853df3FdDA47Dd464f885a49"
				variant="p5"
				className={`truncate text-ellipsis 
			text-sm max-w-[90%] text-center`}
			/>
		</div>
	</div>
</div>
