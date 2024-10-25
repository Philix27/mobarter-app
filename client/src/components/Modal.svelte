<script>
</script>

<div
	class={`bg-black/50 w-full 
      fixed left-0 z-10
      top-0 mt-[50px]
      h-screen
      flex flex-col items-center
      justify-center`}
>
	<div class="bg-background w-[90%] md:w-[50%] lg:w-[75%] rounded-lg px-4 py-4">
		<slot />
	</div>
</div>
