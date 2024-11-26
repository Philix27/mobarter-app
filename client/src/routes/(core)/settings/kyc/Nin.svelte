<script lang="ts">
	import { Button, P, TextInput } from 'components';
	import Wrapper from './Wrapper.svelte';

	let ninNum = $state('');
	let props = $props<{}>();
</script>

<Wrapper title={'Nin Verification'}>
	<TextInput
		place="Nin no."
		inputType="number"
		label="National Identity Number"
		required
		bind:value={ninNum}
	/>

	<Button size="sm" className="my-2 w-fit">Submit</Button>
</Wrapper>
