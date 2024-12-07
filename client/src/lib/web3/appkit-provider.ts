import { cookieToInitialState, type Config } from '@wagmi/core';
import { createAppKit } from '@reown/appkit/react';
import { cookieStorage, createStorage } from 'wagmi';
import { WagmiAdapter } from '@reown/appkit-adapter-wagmi';
import { mainnet, arbitrum } from '@reown/appkit/networks';
import type { AppKitNetwork } from '@reown/appkit/networks';

// Get projectId from https://cloud.reown.com
export const projectId = process.env.NEXT_PUBLIC_PROJECT_ID || 'b56e18d47c72ab683b10814fe9495694'; // this is a public projectId only to use on localhost

if (!projectId) {
	throw new Error('Project ID is not defined');
}

export const networks = [mainnet, arbitrum] as [AppKitNetwork, ...AppKitNetwork[]];

//Set up the Wagmi Adapter (Config)
export const wagmiAdapter = new WagmiAdapter({
	storage: createStorage({
		storage: cookieStorage
	}),
	ssr: true,
	projectId,
	networks
});

export const config = wagmiAdapter.wagmiConfig;

// Set up metadata
const metadata = {
	name: 'next-reown-appkit',
	description: 'next-reown-appkit',
	url: 'https://github.com/0xonerb/next-reown-appkit-ssr', // origin must match your domain & subdomain
	icons: ['https://avatars.githubusercontent.com/u/179229932']
};

// Create the modal
export const modal = createAppKit({
	adapters: [wagmiAdapter],
	projectId,
	networks,
	metadata,
	themeMode: 'light',
	features: {
		analytics: true // Optional - defaults to your Cloud configuration
	}
});

const initialState = cookieToInitialState(wagmiAdapter.wagmiConfig as Config);
// const initialState = cookieToInitialState(wagmiAdapter.wagmiConfig as Config, cookies);
