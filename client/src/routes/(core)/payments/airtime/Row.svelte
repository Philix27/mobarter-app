<script lang="ts">
	import { P } from 'components';
	import { cn } from 'lib/utils';

	// const props = $props<{
	// 	isPassword?: boolean;
	// 	control?: any;
	// 	name?: string;
	// 	place?: string;
	// 	content?: string;
	// 	desc?: string;
	// 	label?: string;
	// 	type?: 'text' | 'number' | 'date' | 'file' | 'password';
	// 	errorMessage?: string;
	// 	className?: string;
	// 	showPassword?: string;
	// }>();

	// let { answer } = $props();

	export let title = '';
	export let imgSrc = '';
	export let isActive = false;
	export let onClick = () => {};
</script>

<div
	class="w-full flex items-center justify-between border-b border-secondary py-2"
	on:click={onClick}
>
	<div class="flex items-center justify-center">
		<div
			class="w-[20px] h-[20px] rounded-full flex items-center justify-center border-[3px] border-secondary mr-2"
		>
			<div
				class={cn(
					'w-[13px] h-[13px] rounded-full border border-secondary',
					isActive && 'bg-primary border-primary'
				)}
			></div>
		</div>
		<P className="text-xs">{title}</P>
	</div>
	<img src={imgSrc} alt="network" height={32} width={32} />
</div>
