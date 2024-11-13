<script lang="ts">
	import { cn } from '$lib/utils';
	import 'iconify-icon';
	import { P } from 'components';

	export let isPassword: boolean = false;
	export let control: any = '';
	/** name to be used as label */
	export let name: string = '';
	/** placeholder */
	export let place: string = '';
	export let value: string | number = '';
	/** Description */
	export let desc: string = '';
	export let label: string = '';
	export let inputType: 'text' | 'number' | 'date' | 'file' | 'password' = 'text';
	export let errorMessage: string = '';
	export let className: string = '';
	export let showPassword: boolean = false;
</script>

<div class={cn('flex-[1] w-full mb-2', className)}>
	<div class="flex flex-col items-start">
		{#if label}
			<div class="pb-2">
				<P v="p5">{label}</P>
			</div>
		{/if}

		<div
			class={cn(
				`flex justify-between 
            items-center border rounded-md border-secondary 
            w-full px-2 bg-background`,
				errorMessage && 'border-red-600'
			)}
		>
			{#if isPassword}
				<iconify-icon class={cn('mx-2', errorMessage && 'text-red-600')} icon="mdi:lock" />
			{/if}
			<input
				{...control}
				placeholder={place}
				value={value}
				type={isPassword ? 'password' : !showPassword && inputType}
				class={cn(
					`flex h-10 w-full rounded-md
             border-none outline-none bg-background text-foreground active::bg-transparent
        `
				)}
				pattern={inputType === 'number' && '[0-9]*'}
				inputmode={inputType === 'number' && 'numeric'}
				{name}
			/>
			{#if isPassword}
				{#if showPassword}
					<iconify-icon
						icon="mdi:eye"
						on:click={() => {
							// Todo
							// setToggle -> false
						}}
					/>
				{:else}
					<iconify-icon
						icon="mdi:close"
						on:click={() => {
							// Todo
							// setToggle -> true
						}}
					/>
				{/if}
			{/if}
		</div>

		{#if errorMessage}
			<P className={`text-red-500 my-2`}>{errorMessage}</P>
		{/if}
		{#if desc}
			<div class="flex items-center justify-center mt-1 text-muted text-xs px-1">
				<iconify-icon icon="mdi:info" class="mr-2" />
				<P className="font-thin ">{desc}</P>
			</div>
		{/if}
	</div>
</div>
