import { writable } from 'svelte/store';

export type ITransactionsTabs = 'Airtime' | 'Wallet' | 'Swap';
export const transactionsState = writable<ITransactionsTabs>('Wallet');
