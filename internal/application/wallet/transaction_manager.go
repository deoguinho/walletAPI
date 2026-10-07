package wallet

type TransactionManager interface {
	withTransaction(func() error) error
}
