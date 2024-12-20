package transactions

type Currency int

const (
	Celo_CUSD Currency = iota + 1
	Celo_CELO
	Eth_USDT
	Optimism_USDT
	Base_USDT
	Arbitrium_USDT
)

func VerifyTransfer(walletAddress string, transactionHash string, currency Currency) {

}
