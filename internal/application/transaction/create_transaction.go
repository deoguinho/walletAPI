package transaction

import (
	"walletAPI/internal/domain/transaction"
)

type createTransaction struct {
	repository TransactionRepository
}

func newCreateTransaction(repository TransactionRepository) *createTransaction {
	return &createTransaction{repository: repository}

}
func (r *createTransaction) Execute(fromWalletID int64, toWalletID int64, amount int64, Type string, status string) (*transaction.Transaction, error) {
	w := transaction.NewTransaction(
		fromWalletID, toWalletID, amount, Type, status,
	)

	err := r.repository.Create(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
