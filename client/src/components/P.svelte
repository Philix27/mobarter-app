<script>
	// export let text;
	export let className = '';
	export let v = 'p4';

	const p1 = 'text-3xl font-extrabold leading-tight text-muted-foreground ' + className;
	const p2 = 'text-2xl font-bold leading-tight text-muted-foreground ' + className;
	const p3 = 'text-lg font-bold leading-tight  text-muted-foreground ' + className;
	const p4 = 'text-md font-medium leading-tight text-muted-foreground ' + className;
	const p5 = 'text-sm font-regular leading-tight text-muted-foreground ' + className;
	const p6 = 'text-xs font-light leading-tight text-muted-foreground ' + className;
</script>

{#if v == 'p1'}
	<p class={p1}><slot /></p>
{/if}

{#if v == 'p2'}
	<p class={p2}><slot /></p>
{/if}

{#if v == 'p3'}
	<p class={p3}><slot /></p>
{/if}

{#if v == 'p4'}
	<p class={p4}><slot /></p>
{/if}

{#if v == 'p5'}
	<p class={p5}><slot /></p>
{/if}

{#if v == 'p6'}
	<p class={p6}><slot /></p>
{/if}

<style lang="postcss">
</style>
