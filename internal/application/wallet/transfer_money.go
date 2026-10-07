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

func (tm *TransferMoney) Execute(fromWalletID, toWalletID int64, amount int64) error {
	if fromWalletID == toWalletID {
		return errors.New("cannot transfer to the same wallet")
	}

	return tm.transactionManager.withTransaction(func() error {

		// Retrieve the source wallet
		fromWallet, err := tm.repository.GetByID(fromWalletID)
		if err != nil {
			return err
		}

		// Retrieve the destination wallet
		toWallet, err := tm.repository.GetByID(toWalletID)
		if err != nil {
			return err
		}

		// Remove the amount from the source wallet
		err = fromWallet.Withdraw(amount)
		if err != nil {
			return err
		}

		// Add the amount to the destination wallet
		err = toWallet.Deposit(amount)
		if err != nil {
			// If deposit fails, rollback the withdrawal
			fromWallet.Deposit(amount)
			return err
		}

		//save the updated wallets
		err = tm.repository.Save(fromWallet)
		if err != nil {
			return err
		}
		err = tm.repository.Save(toWallet)
		if err != nil {
			return err
		}

		return nil
	})
}
