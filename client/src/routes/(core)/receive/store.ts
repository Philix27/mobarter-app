import { writable } from 'svelte/store';

export type IReceiveTabs = "Buy" | "Wallet"
export const receiveState = writable<IReceiveTabs>('Buy');
