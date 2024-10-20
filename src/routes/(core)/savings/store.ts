import { writable } from 'svelte/store';

export type ISavingsStore = 'Deposit' | 'Withdraw';
export const savingsStore = writable<ISavingsStore>('Deposit');
