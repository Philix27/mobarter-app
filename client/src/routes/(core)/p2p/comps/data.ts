import type { IBankAccount } from 'types';
export type CryptoCurrencies = 'CELO' | 'CUSD' | 'USDT';
export type FiatCurrencies = 'NGN' | 'KES' | 'CED';
export type TradeType = 'BUY' | 'SELL';

export const ListOfAgentsAds: {
	name: string;
	limitUpper: number;
	limitLower: number;
	totalSale: number;
	completedSale: number;
	successRate: number;
	fiatCurrency: FiatCurrencies;
	fiatRate: number;
	cryptoCurrency: CryptoCurrencies;
	cryptoRate: number;
	averageResponseTime: number;
	instructions: string;
	tradeType: TradeType;
	bankAccounts: IBankAccount[];
}[] = [
	{
		name: 'Skimmy Janson',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'CELO',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'SELL',
		bankAccounts: []
	},
	{
		name: 'Solomon King',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'CUSD',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'BUY',
		bankAccounts: []
	},
	{
		name: 'Marie Curie',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'CELO',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'SELL',
		bankAccounts: []
	},
	{
		name: 'Lobert Oxford',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'USDT',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'BUY',
		bankAccounts: []
	},
	{
		name: 'Cane Marble',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'CUSD',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'SELL',
		bankAccounts: []
	},
	{
		name: 'Cane Marble',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'CUSD',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'BUY',
		bankAccounts: []
	},
	{
		name: 'Cane Marble',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'USDT',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'BUY',
		bankAccounts: []
	},
	{
		name: 'Cane Marble',
		limitUpper: 2000,
		limitLower: 5000,
		totalSale: 0,
		completedSale: 0,
		successRate: 94.7,
		fiatCurrency: 'NGN',
		fiatRate: 1200,
		cryptoCurrency: 'CUSD',
		cryptoRate: 0,
		averageResponseTime: 0,
		instructions: '',
		tradeType: 'BUY',
		bankAccounts: []
	}
];
