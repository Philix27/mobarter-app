import { persisted } from 'svelte-persisted-store';

type IData = {
	isMiniPay: boolean;
	device?: 'MOBILE' | 'DESKTOP' | 'TABLET';
	env?: 'MINIPAY' | 'BROWSER' | 'MOBILE_APP' | 'UNDEFINED';
	walletAddress?: string;
};

export const globalStore = persisted<IData>('studio', {
	isMiniPay: false
});
