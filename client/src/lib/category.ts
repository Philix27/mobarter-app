export const utilityCategories: {
	id: number;
	name: string;
	code: string;
	description: string;
	link: string;
	icon: string;
	country_code: 'NG';
}[] = [
	{
		id: 1,
		name: 'Airtime',
		code: 'AIRTIME',
		description: 'Airtime',
		country_code: 'NG',
		link: '/payments/airtime',
		icon: 'fluent:call-32-regular'
	},
	{
		id: 2,
		name: 'Mobile Data Service',
		code: 'MOBILEDATA',
		description: 'Mobile Data Service',
		country_code: 'NG',
		link: '/payments/data',
		icon: 'gg:data'
	},
	{
		id: 3,
		name: 'Cable Bill Payment',
		code: 'CABLEBILLS',
		description: 'Cable Bill Payment',
		country_code: 'NG',
		link: '/payments/cable',
		icon: 'ic:outline-live-tv'
	},
	{
		id: 4,
		name: 'Internet Service',
		code: 'INTSERVICE',
		description: 'Internet Service',
		country_code: 'NG',
		link: '/payments/internet',
		icon: 'mdi:dollar'
	},
	{
		id: 5,
		name: 'Utility Bills',
		code: 'UTILITYBILLS',
		description: 'Utility Bills',
		country_code: 'NG',
		link: '/payments/utility',
		icon: 'lucide:utility-pole'
	},
	{
		id: 6,
		name: 'Tax Payment',
		code: 'TAX',
		description: 'Tax Payment',
		country_code: 'NG',
		link: '/payments/tax',
		icon: 'mdi:dollar'
	},
	{
		id: 7,
		name: 'Donations',
		code: 'DONATIONS',
		description: 'Donations',
		country_code: 'NG',
		link: '/payments/donations',
		icon: 'mdi:dollar'
	},
	{
		id: 8,
		name: 'Transport and Logistics',
		code: 'TRANSLOG',
		description: 'Transport and Logistics',
		country_code: 'NG',
		link: '/payments/transport',
		icon: 'hugeicons:bitcoin-withdraw'
	},
	{
		id: 9,
		name: 'Dealer Payments',
		code: 'DEALPAY',
		description: 'Dealer Payments',
		country_code: 'NG',
		link: '/payments/dealer',
		icon: 'mdi:dollar'
	},
	{
		id: 17,
		name: 'Religious Institutions',
		code: 'RELINST',
		description: 'Religious Institutions',
		country_code: 'NG',
		link: '/payments/religion',
		icon: 'mdi:dollar'
	},
	{
		id: 18,
		name: 'Schools & Professional Bodies',
		code: 'SCHPB',
		description: 'Schools & Professional Bodies',
		country_code: 'NG',
		link: '/payments/school',
		icon: 'mdi:dollar'
	}
];