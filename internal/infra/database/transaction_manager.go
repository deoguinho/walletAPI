package database

import (
	"context"

	"walletAPI/internal/application/wallet"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionManager struct {
	db *pgxpool.Pool
}

func NewTransactionManager(db *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) WithTransaction(
	fn func(repository wallet.WalletRepository) error,
) error {
	tx, err := tm.db.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	repository := NewWalletRepository(tx)

	if err := fn(repository); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
