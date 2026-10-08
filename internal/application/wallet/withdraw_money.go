package wallet

type WithdrawMoney struct {
	repository WalletRepository
}

func NewWithdrawMoney(repository WalletRepository) *WithdrawMoney {
	return &WithdrawMoney{repository: repository}
}

func (r *WithdrawMoney) Execute(walletID int64, amount int64) error {
	wallet, err := r.repository.GetByID(walletID)
	if err != nil {
		return err
	}

	err = wallet.Withdraw(amount)
	if err != nil {
		return err
	}

	err = r.repository.Save(wallet)
	if err != nil {
		return err
	}
	return nil
}
