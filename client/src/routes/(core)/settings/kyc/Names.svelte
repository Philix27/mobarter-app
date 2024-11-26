<script lang="ts">
	import { Button, P, TextInput } from 'components';
	import Wrapper from './Wrapper.svelte';

	let amountSelected = $state('');
	let props = $props<{}>();
</script>

<Wrapper title={'Full names'}>
	<TextInput place="John" label="First name" required bind:value={amountSelected} />
	<TextInput place="Doe" label="Last name" required bind:value={amountSelected} />
	<TextInput place="" label="Middle name" required bind:value={amountSelected} />
	<TextInput label="Date of Birth" required bind:value={amountSelected} />
	<Button size="sm" className="my-2 w-fit">Submit</Button>
</Wrapper>
