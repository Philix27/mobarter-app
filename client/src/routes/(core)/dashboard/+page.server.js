// import type { PageLoadEvent } from './$types';
// import type { PageServerLoad } from './$types';

// @ts-ignore
export function load({ cookies }) {
	const theme = cookies.get('colorTheme');

	return {
		theme: theme
	};
}

// export const load = (async ({ cookies }) => {
// 	const tastyCookie = cookies.get('tastyCookie');

// 	return {
// 		tastyCookie
// 	};
// }) satisfies PageServerLoad;
