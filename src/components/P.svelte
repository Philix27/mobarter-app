<script>
	// export let text;
	export let className = '';
	export let variant = 'p4';

	const p1 = 'text-3xl font-extrabold leading-tight text-muted-foreground ' + className;
	const p2 = 'text-2xl font-bold leading-tight text-muted-foreground ' + className;
	const p3 = 'text-lg font-bold leading-tight  text-muted-foreground ' + className;
	const p4 = 'text-md font-medium leading-tight text-muted-foreground ' + className;
	const p5 = 'text-sm font-regular leading-tight text-muted-foreground ' + className;
	const p6 = 'text-xs font-light leading-tight text-muted-foreground ' + className;
</script>

{#if variant == 'p1'}
	<p class={p1}><slot /></p>
{/if}

{#if variant == 'p2'}
	<p class={p2}><slot /></p>
{/if}

{#if variant == 'p3'}
	<p class={p3}><slot /></p>
{/if}

{#if variant == 'p4'}
	<p class={p4}><slot /></p>
{/if}

{#if variant == 'p5'}
	<p class={p5}><slot /></p>
{/if}

{#if variant == 'p6'}
	<p class={p6}><slot /></p>
{/if}

<style lang="postcss">
</style>
