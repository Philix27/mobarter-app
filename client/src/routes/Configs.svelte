<script lang="ts">
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	import { Client, cacheExchange, fetchExchange, setContextClient } from '@urql/svelte';

	const urqlClient = new Client({
		url: 'http://localhost:8080/graphql',
		exchanges: [cacheExchange, fetchExchange]
	});
	setContextClient(urqlClient);

	const queryClient = new QueryClient();

	// Inject the Analytics functionality
</script>

<QueryClientProvider client={queryClient}>
	<slot />
</QueryClientProvider>
