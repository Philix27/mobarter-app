<script lang="ts">
	type IRow = {
		title: string;
		subtitle?: string;
		icon?: string;
		imgPath?: string;
		imgComp?: Slottable;
		color?: string;
		trailingText?: string;
		trailingComp?: Slottable;
		onClick?: VoidFunction;
		hideArrow?: boolean;
		isLast?: boolean;
	};
</script>
