package wallet

import "walletAPI/internal/domain/wallet"

type CreateWallet struct {
	repository WalletRepository
}

func NewCreateWallet(repository WalletRepository) *CreateWallet {
	return &CreateWallet{repository: repository}
}

func (r *CreateWallet) Execute(userID int64) (*wallet.Wallet, error) {
	w := wallet.NewWallet(userID, 0)
	err := r.repository.Create(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
