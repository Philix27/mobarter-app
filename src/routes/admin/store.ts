import { writable } from 'svelte/store';

export type IAdminScreen =
	| 'Card'
	| 'CreditBuy'
	| 'CreditSell'
	| 'Support'
	| 'Kyc'
	| 'Logs'
	| 'Email'
	| 'NewsLetter'
	| 'DepositBuy'
	| 'DepositSell';

export const adminScreenState = writable<IAdminScreen>('Support');
