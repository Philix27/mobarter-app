<script>
	import P from '../../../components/P.svelte';
	import 'iconify-icon';
	const transactions = [
		{
			amount: 90,
			date: '23-10-09'
		}
	];

	const oldObj = { a: 4, b: 9 };

	const newObj = structuredClone(oldObj);
</script>

<div class="text-column">
	<h1>Transactions</h1>
	<iconify-icon icon="mdi:home" />
	{#each transactions as item}
		<div>
			<P>{item.amount}</P>
		</div>
		>
	{/each}

	<p>
		This is a <a href="https://kit.svelte.dev">SvelteKit</a> app. You can make your own by typing the
		following into your command line and following the prompts:
	</p>

	<pre>npm create svelte@latest</pre>

	<p>
		The page you're looking at is purely static HTML, with no client-side interactivity needed.
		Because of that, we don't need to load any JavaScript. Try viewing the page's source, or opening
		the devtools network panel and reloading.
	</p>

	<p>
		The <a href="/sverdle">Sverdle</a> page illustrates SvelteKit's data loading and form handling. Try
		using it with JavaScript disabled!
	</p>
</div>
