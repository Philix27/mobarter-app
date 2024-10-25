<script lang="ts">
	//   import { TextP } from "@/comps";
	// import { Token } from "@/lib";
	// import React, { memo } from "react";
	// import { Balance, TokenSelector } from "../../client/app/(others)/_comps";

	import { P } from 'components';

	export let title;
	export let isReadOnly = false;
	export let balance: string = '';
	export let tokenAddress: string = '';
	export let value: string;
	export let onTokenClick: VoidFunction;
	export let onChange: (val: string) => void;
</script>

<div class="bg-card w-full mb-1 flex flex-col items-center justify-between rounded-lg px-3 py-4">
	<div class="flex items-center justify-between w-full">
		<P>{title}</P>
		<!-- <TextP class="text-muted text-[10px]">
          <Balance tokenAddress={props.tokenAddress as `0x${string}`} />
        </TextP> -->
		<P className="text-muted text-[10px]">..bal</P>
	</div>

	<div class="w-full flex items-center justify-between">
		<input
			type="number"
			placeholder="0.00"
			class={`outline-none w-full
                border-none bg-transparent font-light
                px-2 py-2 text-2xl tracking-wide mr-2`}
			pattern={'[0-9]*'}
			inputmode="numeric"
			readOnly={isReadOnly}
			{value}
		/>
		<!-- <TokenSelector onClick={props.onTokenClick} token={props.token} /> -->
	</div>

	<div class="flex justify-between items-center w-full">
		<P className="text-primary text-[10px]">~ $9000 Equi</P>
		<P className="text-primary text-[10px] font-bold">MAX</P>
	</div>
</div>
<!-- 
onChange={(e) => {
	onChange(e.target.value);
}} -->
