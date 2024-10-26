import { createThirdwebClient } from 'thirdweb';
import * as Wallet from 'thirdweb/wallets';
import * as Chains from 'thirdweb/chains';

const client = createThirdwebClient({
	clientId: '....'
});

const wallets = [
	Wallet.inAppWallet({
		auth: {
			options: ['email', 'passkey', 'phone', 'google']
		}
	}),
	Wallet.createWallet('io.metamask'),
	Wallet.createWallet('com.coinbase.wallet'),
	Wallet.createWallet('me.rainbow'),
	Wallet.createWallet('io.rabby'),
	Wallet.createWallet('io.zerion.wallet')
];

Wallet.inAppWallet({
	auth: {
		options: ['discord', 'apple']
	}
});

const wallet = Wallet.walletConnect();

export const connect = await wallet.connect({
	client,
	projectId: '',
	chain: Chains.celo,
	optionalChains: [Chains.celo, Chains.celoAlfajoresTestnet],
	appMetadata: {
		name: 'thirdweb powered dApp',
		url: 'https://thirdweb.com',
		description: 'thirdweb powered dApp',
		logoUrl: 'https://thirdweb.com/favicon.ico'
	}
});

// function Example() {
//   return (
//     <ConnectButton
//       client={client}
//       wallets={wallets}
//       theme={darkTheme({
//         colors: {
//           borderColor: "#222020",
//           modalBg: "#121212",
//           accentText: "#f8300d",
//           separatorLine: "#2c2b2b",
//           tertiaryBg: "#0a0a0a",
//           primaryText: "#a0a0ab",
//         },
//       })}
//       connectModal={{
//         size: "compact",
//         showThirdwebBranding: false,
//       }}
//     />
//   );
// }
