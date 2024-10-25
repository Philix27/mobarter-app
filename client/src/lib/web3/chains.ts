import { defineChain, fallback, http, type Transport } from 'viem';
import type { ConfiguredChainId } from './client';
import {
	arbitrum,
	aurora,
	avalanche,
	base,
	bsc,
	gnosis,
	mainnet,
	optimism,
	polygon,
	zora,
	goerli,
	ronin,
	saigon,
	celo,
	celoAlfajores,
	sepolia
} from 'viem/chains';
import { zkSync } from 'viem/zksync';

// #region State variables

//#endregion

//#region ChainList
const fhenixConfig = defineChain({
	id: 8008135,
	name: 'Fhenix',
	network: 'Fhenix',
	nativeCurrency: { name: 'tFHE', symbol: 'tFHE', decimals: 18 },
	rpcUrls: {
		public: {
			http: ['https://api.helium.fhenix.zone']
		},
		default: {
			http: ['https://api.helium.fhenix.zone']
		}
	},
	blockExplorers: {
		default: {
			name: 'Fhenix',
			url: ' https://explorer.helium.fhenix.zone'
		}
	}
});

export const chains = [
	arbitrum,
	aurora,
	avalanche,
	base,
	bsc,
	gnosis,
	mainnet,
	optimism,
	polygon,
	zkSync,
	zora,
	goerli,
	ronin,
	saigon,
	fhenixConfig,
	celo,
	celoAlfajores,
	sepolia
] as const;
//#endregion
export const transports = chains.reduce(
	(acc, { id }) => {
		// const url = rpcUrls[id];
		const url = 'rpcUrls[id]';
		acc[id] = url ? fallback([http(url), http()]) : http();
		return acc;
	},
	{} as Record<ConfiguredChainId, Transport>
);
