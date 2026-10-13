package transaction

import "walletAPI/internal/domain/transaction"

type TransactionRepository interface {
	Create(transaction *transaction.Transaction) error
}
