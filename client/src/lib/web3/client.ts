//#region imports
import { defaultWagmiConfig, createWeb3Modal } from '@web3modal/wagmi';
import {
	getAccount,
	getChainId,
	getBalance as _getBalance,
	watchChainId,
	createConfig,
	injected,
	watchAccount as _watchAccount,
	reconnect,
	sendTransaction as _sendTransaction,
	disconnect as _disconnect,
	switchChain as _switchChain,
	type CreateConnectorFn
} from '@wagmi/core';
import { readable, writable, type Writable } from 'svelte/store';
import { celo } from 'viem/chains';
import { walletConnect } from '@wagmi/connectors';
import { browser } from '$app/environment';
import { chains, transports } from './chains';
import { ethers } from 'ethers';
// import { Web3AuthConnectorInstance } from './web3Connector';
//#endregion

const metadata = {
	name: 'Mobarter',
	description: 'Manage your crypto assets',
	url: 'https://mobarter.vercel.app/',
	icons: ['https://avatars.githubusercontent.com/u/37784886']
};

const ssr = !browser;
export const projectId = 'env.VITE_PROJECT_ID';
// export const projectId = import.meta.env.VITE_PROJECT_ID;

export const wagmiConfig = defaultWagmiConfig({
	chains,
	projectId,
	metadata,
	enableCoinbase: false,
	enableInjected: true
});

export type ConfiguredChain = (typeof chains)[number];
export type ConfiguredChainId = ConfiguredChain['id'];

// Initialize chain-specific transports based on configured RPC URLs

const connectors: CreateConnectorFn[] = ssr
	? []
	: [walletConnect({ projectId, metadata, showQrModal: false }), injected()];
// : [Web3AuthConnectorInstance([celoAlfajores, celo, sepolia]), injected()];

// const wgConfig = createConfig({
// 	chains,
// 	transports,
// 	ssr: true,
// 	connectors
// });

reconnect(wagmiConfig);

export const modal = createWeb3Modal({
	wagmiConfig,
	projectId,
	themeMode: 'dark', // light/dark mode
	themeVariables: {
		//--w3m-font-family
		'--w3m-accent': '#999BA1', // Button colour surface-500
		'--w3m-color-mix': '#071A25', // Modal colour mix primary-300
		'--w3m-color-mix-strength': 50, // Strength of colour
		'--w3m-font-size-master': '8px', // Font size
		'--w3m-border-radius-master': '999px' // border rounding
		// --w3m-z-index
	},
	defaultChain: celo,
	featuredWalletIds: [],
	enableAnalytics: false
});

export const chainId = readable(getChainId(wagmiConfig), (set) =>
	watchChainId(wagmiConfig, { onChange: set })
);

export function getActiveChain(num: number) {
	if (num === 42220) return 'celo';
	if (num === 44787) return 'alfajores';
	if (num === 62320) return 'baklava';
	return 'celo';
}

export const account = readable(getAccount(wagmiConfig), (set) =>
	_watchAccount(wagmiConfig, { onChange: set })
);

export const provider = readable<unknown | undefined>(undefined, (set) =>
	_watchAccount(wagmiConfig, {
		onChange: async (account) => {
			if (!account.connector) return set(undefined);
			set(await account.connector?.getProvider());
		}
	})
);

export const getBalance = (
	tokenAddress: `0x${string}` | string,
	userAddress: `0x${string}` | string
) => {
	// return readable(
	// 	_getBalance(wagmiConfig, {
	// 		address: userAddress,
	// 		token: tokenAddress
	// 	})
	// );
	if (tokenAddress === '0x000') {
		return _getBalance(wagmiConfig, {
			address: userAddress as `0x${string}`
		});
	}
	return _getBalance(wagmiConfig, {
		address: userAddress as `0x${string}`,
		token: tokenAddress as `0x${string}`
	});
};

export const supported_chains = writable<string[]>([]);

export function switchChain(chainId: ConfiguredChainId) {
	return _switchChain(wagmiConfig, { chainId });
}

export function disconnect() {
	return _disconnect(wagmiConfig);
}

export function getChain(chainId: number) {
	return chains.find(({ id }) => id === chainId);
}
export function sendTransaction(to: string, value: string) {
	return _sendTransaction(wagmiConfig, {
		to: to as `0x${string}`,
		value: ethers.parseEther(value)
	});
}
