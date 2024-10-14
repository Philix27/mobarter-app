<script>
	import P from '../components/P.svelte';
	import 'iconify-icon';

	export let title;
	// export let icon = '';
	// export let isBack = false;
	// export let onIconClick = () => {};
</script>



<div
	class={`
        h-[50px]  px-5 flex items-center 
        justify-between 
        fixed top-0 w-full 
        z-10 bg-background
      `}
>
	<div class="flex items-center gap-x-2">
		<iconify-icon class="text-xl mb-2 text-foreground mb-2" icon="mdi:menu" />
		<iconify-icon class="text-xl mb-2 text-foreground mb-2" icon="mdi:left" />
		<!-- {isBack ? (
          <LuChevronLeft
            size={22}
            onClick={() => {
              router.back();
            }}
          />
        ) : (
          <IoMenu
            size={23}
            onClick={() => {
              store.update({ drawerIsOpen: true });
            }}
          />
        )} -->
	</div>

	<P
		variant="p4"
		className="text-card-foreground my-0 font-extrabold tracking-wide font-sans"
		text={title}
	/>
	<div class="flex items-center gap-x-3">
		<iconify-icon class="text-xl mb-2 text-foreground mb-2" icon="mdi:left" />
		<!-- {props.icon ? <Icon onClick={props.onIconClick} size={22} /> : <div />} -->
	</div>
</div>
