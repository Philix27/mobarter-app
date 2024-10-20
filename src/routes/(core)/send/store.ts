import { writable } from 'svelte/store';

export type ISendTabs = 'Phone' | 'Wallet' | 'Bank';
export const sendScreenState = writable<ISendTabs>('Wallet');
