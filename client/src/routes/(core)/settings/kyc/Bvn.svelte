<script>
	import { Button, TextInput } from 'components';

	let amountSelected = $state();
</script>

<div>
	<TextInput place="Bvn" inputType="number" label="BVN No." required bind:value={amountSelected} />

	<Button>Submit</Button>
</div>
