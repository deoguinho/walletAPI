package database

import (
	"context"
	"walletAPI/internal/domain/transaction"
)

type TransactionRepository struct {
	db DBTX
}

func NewTransactionRepository(db DBTX) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create implements [transaction.TransactionRepository]
func (r *TransactionRepository) Create(transaction *transaction.Transaction) error {
	row := r.db.QueryRow(context.Background(),
		"INSERT INTO transactions (from_wallet_id, to_wallet_id, amount, type, status) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at", transaction.FromWalletID, transaction.ToWalletID, transaction.Amount, transaction.Type, transaction.Status)
	if err := row.Scan(&transaction.ID, &transaction.CreatedAt); err != nil {
		return err
	}
	return nil
}
