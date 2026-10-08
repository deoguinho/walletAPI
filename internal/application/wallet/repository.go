package wallet

import "walletAPI/internal/domain/wallet"

type WalletRepository interface {
	Create(wallet *wallet.Wallet) error
	GetByID(id int64) (*wallet.Wallet, error)
	Save(wallet *wallet.Wallet) error
}
