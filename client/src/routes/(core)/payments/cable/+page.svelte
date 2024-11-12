<script>
	import { Button, Nav, TextInput } from 'components';
	import { toast } from 'svelte-sonner';
</script>

<svelte:head>
	<title>Receive Funds</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="w-full flex flex-col items-center justify-center gap-y-4">
	<Nav title="Cable" isBack />
	<TextInput place="Select Airtime" label="Select Airtime" />
	<TextInput place="Mobile number" inputType="number" label="Phone number" />
	<TextInput place="Amount" inputType="number" label="Amount" />
	<Button onclick={() => toast('Funds sent')}>Buy</Button>
</div>
