<script lang="ts">
	import { P } from 'components';
	import { cn } from 'lib/utils';
	import * as Accordion from '$lib/components/ui/accordion';

	let props = $props<{ title: string }>();
</script>

<Accordion.Root>
	<Accordion.Item value={props.title}>
		<Accordion.Trigger>
			<P className="no-underline">{props.title}</P>
		</Accordion.Trigger>
		<Accordion.Content>
			<!-- svelte-ignore slot_element_deprecated -->
			<div class={cn('bg-card rounded-lg p-2 mb-2 flex flex-col items-center justify-center')}>
				<slot />
			</div>
		</Accordion.Content>
	</Accordion.Item>
</Accordion.Root>
