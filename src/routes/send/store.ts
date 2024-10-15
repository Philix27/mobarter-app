import { writable } from 'svelte/store';

export const sendScreenState = writable<'PHONE' | 'WALLET' | 'BANK'>('WALLET');
