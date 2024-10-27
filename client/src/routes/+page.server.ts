// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
// export const prerender = false;

import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	setTheme: async ({ url, cookies }) => {
		const theme = url.searchParams.get('theme');
		const redirectTo = url.searchParams.get('redirectTo');
		console.log(theme);

		if (theme) {
			cookies.set('colorTheme', theme, {
				path: '/',
				maxAge: 60 * 60 * 24 * 365
			});
		}

		throw redirect(303, redirectTo ?? '/');
	}
};
