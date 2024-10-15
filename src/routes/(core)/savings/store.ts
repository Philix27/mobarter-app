import { writable } from 'svelte/store';

export const receiveState = writable<'BUY' | 'WALLET'>('BUY');
