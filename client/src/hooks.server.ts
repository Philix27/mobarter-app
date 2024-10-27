import type { Handle } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	let theme: string | null = null;
	const newTheme = event.url.searchParams.get('theme');
	const cookieTheme = event.cookies.get('colorTheme');

	if (newTheme) {
		theme = newTheme;
	} else {
		if (cookieTheme) {
			theme = cookieTheme;
		} else {
			theme = 'dark';
		}
	}

	if (theme) {
		return await resolve(event, {
			transformPageChunk: (input) => {
				// return input.html.replace('data-theme=""', `data-theme="${theme}"`);
				return input.html.replace('class="color-scheme:"', `class="color-scheme: ${theme}"`);
			}
		});
	}
	return await resolve(event);
}) satisfies Handle;
