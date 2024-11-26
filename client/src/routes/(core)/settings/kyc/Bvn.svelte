<script lang="ts">
	import { Button, P, TextInput } from 'components';
	import Wrapper from './Wrapper.svelte';

	let amountSelected = $state('');
	let props = $props<{ }>();
</script>

<Wrapper title={"BVN Verification"}>
	<TextInput place="Bvn" inputType="number" label="BVN No." required bind:value={amountSelected} />

	<Button size="sm" className="my-2 w-fit">Submit</Button>
</Wrapper>
