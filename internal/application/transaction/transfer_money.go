package transaction

import (
	"errors"
	"strconv"
	"walletAPI/internal/application/wallet"
	"walletAPI/internal/domain/transaction"
)

type TransferMoney struct {
	repository            wallet.WalletRepository
	transactionManager    TransactionManager
	transactionRepository TransactionRepository
}

func NewTransferMoney(

	repository wallet.WalletRepository,
	transactionManager TransactionManager,
	transactionRepository TransactionRepository) *TransferMoney {
	return &TransferMoney{
		repository:            repository,
		transactionManager:    transactionManager,
		transactionRepository: transactionRepository,
	}
}

func (tm *TransferMoney) Execute(fromWalletID int64, toWalletID int64, amount int64) error {
	if fromWalletID == toWalletID {
		return errors.New("cannot transfer to the same wallet")
	}

	return tm.transactionManager.WithTransaction(func(repository wallet.WalletRepository, transactionRepository TransactionRepository) error {

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

		fromWalletIDString := strconv.FormatInt(fromWalletID, 10)
		toWalletIDString := strconv.FormatInt(toWalletID, 10)
		tx := &transaction.Transaction{
			FromWalletID: &fromWalletIDString,
			ToWalletID:   &toWalletIDString,
			Amount:       amount,
			Type:         transaction.TransferTransaction,
			Status:       transaction.CompletedStatus,
		}

		err = transactionRepository.Create(tx)
		if err != nil {
			return err
		}
		return nil
	})

}
