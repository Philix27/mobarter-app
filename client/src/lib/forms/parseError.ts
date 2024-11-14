import type { ZodFormattedError } from 'zod';

export const parseFormError = <T>(name: string, error: ZodFormattedError<T>) => {
	// @ts-ignore
	return error[name]?._errors[0] ?? '';
};
