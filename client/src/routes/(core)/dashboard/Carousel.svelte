<script lang="ts">
	// import * as Card from '$lib/components/ui/card/index.js';
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import type { CarouselAPI } from '$lib/components/ui/carousel/context.js';
	import Autoplay from 'embla-carousel-autoplay';

	let api: CarouselAPI;
	let current = 0;
	let count = 0;

	$: if (api) {
		count = api.scrollSnapList().length;
		current = api.selectedScrollSnap() + 1;

		api.on('select', () => {
			current = api.selectedScrollSnap() + 1;
		});
	}
</script>

<div>
	<Carousel.Root
		bind:api
		class="w-full"
		plugins={[
			Autoplay({
				delay: 10_000
			})
		]}
		opts={{
			align: 'center',
			loop: true
		}}
	>
		<Carousel.Content>
			{#each ['banner/merry.png', 'banner/merryO.png', 'banner/merryW.png', 'banner/coming.png'] as url, i (i)}
				<Carousel.Item>
					<div class="bg-primary w-full h-[70px] rounded-md">
						<img src={url} alt="img" class="w-full h-full object-cover rounded-md" />
					</div>
				</Carousel.Item>
			{/each}
		</Carousel.Content>
	</Carousel.Root>
</div>
