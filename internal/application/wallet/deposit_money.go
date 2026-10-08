package wallet

type DepositMoneyRequest struct {
	repository WalletRepository
}

func NewDepositMoney(repository WalletRepository) *DepositMoneyRequest {
	return &DepositMoneyRequest{repository: repository}
}

func (r *DepositMoneyRequest) Execute(walletID int64, amount int64) error {
	wallet, err := r.repository.GetByID(walletID)
	if err != nil {
		return err
	}

	err = wallet.Deposit(amount)
	if err != nil {
		return err
	}

	err = r.repository.Save(wallet)
	if err != nil {
		return err
	}

	return nil
}
