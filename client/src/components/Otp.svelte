<script lang="ts">
	// import * as InputOTP from '$lib/components/ui/input-otp/index.js';
	import * as Bits from 'bits-ui';
	import { InputOTP, InputOTPSlot, InputOTPGroup } from 'lib/components/ui/input-otp';
	import { slice } from 'viem';
</script>

<!-- <Root maxlength={6} pattern={Bits.REGEXP_ONLY_DIGITS}></Root> -->

<InputOTP maxlength={6} pattern={Bits.REGEXP_ONLY_DIGITS}>
	{#snippet children({ cells })}
		<InputOTPGroup>
			{#each cells.slice(0, 2) as cell}
				<InputOTPSlot {cell} />
			{/each}
		</InputOTPGroup>
	{/snippet}
</InputOTP>
