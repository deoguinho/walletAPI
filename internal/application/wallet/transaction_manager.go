package wallet

type DBTX interface{}

type TransactionManager interface {
	WithTransaction(fn func(repository WalletRepository) error) error
}
