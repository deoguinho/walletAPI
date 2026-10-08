package wallet

import "errors"

type TransferMoney struct {
	repository         WalletRepository
	transactionManager TransactionManager
}

func NewTransferMoney(
	repository WalletRepository,
	transactionManager TransactionManager) *TransferMoney {
	return &TransferMoney{
		repository:         repository,
		transactionManager: transactionManager,
	}
}

func (tm *TransferMoney) Execute(fromWalletID int64, toWalletID int64, amount int64) error {
	if fromWalletID == toWalletID {
		return errors.New("cannot transfer to the same wallet")
	}

	return tm.transactionManager.WithTransaction(func(repository WalletRepository) error {

		fromWallet, err := repository.GetByID(fromWalletID)
		if err != nil {
			return err
		}

		toWallet, err := repository.GetByID(toWalletID)
		if err != nil {
			return err
		}

		err = fromWallet.Withdraw(amount)
		if err != nil {
			return err
		}

		err = toWallet.Deposit(amount)
		if err != nil {
			return err
		}

		err = repository.Save(fromWallet)
		if err != nil {
			return err
		}

		err = repository.Save(toWallet)
		if err != nil {
			return err
		}

		return nil
	})
}
