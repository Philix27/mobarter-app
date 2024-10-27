<script>
	import { Nav, P } from 'components';
	import { page } from '$app/stores';
</script>

<div>
	<Nav title={`Transactions - ${$page.params.id}`} isBack />
	<div class="w-full relative bg-background">
		<P>Transaction History</P>
		<P>Transaction hash</P>
		<P>Type</P>
		<P>From</P>
		<P>To</P>
		<P>Currency</P>
		<P>{$page.params.id}</P>
	</div>
</div>
