<script lang="ts">
	import { cn } from '$lib/utils';
	import { P } from 'components';
	import { cva, type VariantProps } from 'class-variance-authority';

	export let onclick: VoidFunction = () => {};
	// $: onclick();

	export let className: string = '';
	export let isLoading: boolean = false;
	export let variant: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link' =
		'default';
	export let size: 'default' | 'sm' | 'lg' | 'icon' = 'default';
	export let btype: 'submit' | 'button' = 'submit';

	const buttonVariants = cva(
		`inline-flex items-center
  justify-center rounded-md
  text-sm font-medium transition-colors
  focus-visible:outline-none
  focus-visible:ring-2 focus-visible:ring-ring
  focus-visible:ring-offset-2 disabled:opacity-50
  disabled:pointer-events-none ring-offset-background w-[70%]`,
		{
			variants: {
				variant: {
					default: 'bg-primary text-primary-foreground hover:bg-primary/60 hover:border-secondary',
					destructive: 'bg-destructive text-destructive-foreground hover:bg-destructive/90',
					outline:
						'border border-input border-primary hover:bg-accent hover:text-accent-foreground',
					secondary: 'bg-secondary text-secondary-foreground hover:bg-secondary/80',
					ghost: 'hover:bg-accent hover:text-accent-foreground',
					link: 'underline-offset-4 hover:underline text-primary'
				},
				size: {
					default: 'h-10 py-2 px-4',
					sm: 'h-9 px-3 rounded-md',
					lg: 'h-11 px-8 rounded-md',
					icon: 'h-10 w-10'
				}
			},
			defaultVariants: {
				variant: 'default',
				size: 'default'
			}
		}
	);
	//
</script>

<button
	type={btype}
	disabled={isLoading}
	class={cn(buttonVariants({ variant, size, className }))}
	on:click={onclick}
>
	{#if isLoading}
		<!-- <Spinner color="#fff" size={40} />  -->
		<P>...</P>
	{:else}
		<slot />
	{/if}
</button>

<!-- import { Slot } from '@radix-ui/react-slot';
import { cva, type VariantProps } from 'class-variance-authority';
import { cn } from '@/lib';
import { Spinner } from '../spinner'; -->

<!-- export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  asChild?: boolean;
  isLoading?: boolean;
} -->

<!-- const AppButton = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, disabled, variant, size, isLoading, children, asChild = false, ...props }, ref) => {
    const Comp = asChild ? Slot : 'button';
    return (
      <Comp disabled={isLoading} className={cn(buttonVariants({ variant, size, className }))} ref={ref} {...props}>
        {isLoading ? <Spinner color="#fff" size={40} /> : children}
      </Comp>
    );
  }
);
AppButton.displayName = 'Button'; -->

<!-- export { AppButton, buttonVariants }; -->
