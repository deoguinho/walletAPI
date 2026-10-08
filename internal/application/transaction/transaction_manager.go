package transaction

import (
	"walletAPI/internal/application/wallet"
)

type DBTX interface{}

type TransactionManager interface {
	WithTransaction(fn func(
		repository wallet.WalletRepository,
		transactionRepository TransactionRepository) error) error
}
