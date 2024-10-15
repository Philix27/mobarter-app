import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import { cva as c } from 'class-variance-authority';

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export const cva = c;
