package database

import (
	"context"
	"errors"
	"walletAPI/internal/domain/wallet"

	"github.com/jackc/pgx/v5"
)

type WalletRepository struct {
	db DBTX
}

func NewWalletRepository(db DBTX) *WalletRepository {
	return &WalletRepository{db: db}
}

// Create implements [wallet.WalletRepository].
func (r *WalletRepository) Create(wallet *wallet.Wallet) error {
	row := r.db.QueryRow(context.Background(), "INSERT INTO wallets (user_id, balance) VALUES ($1, $2) RETURNING id, created_at", wallet.UserID, wallet.Balance)
	if err := row.Scan(&wallet.ID, &wallet.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (r *WalletRepository) GetByID(id int64) (*wallet.Wallet, error) {
	row := r.db.QueryRow(context.Background(), "SELECT id, user_id, balance, created_at FROM wallets WHERE id = $1", id)
	var w wallet.Wallet
	if err := row.Scan(&w.ID, &w.UserID, &w.Balance, &w.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, wallet.ErrWalletNotFound
		}
		return nil, err
	}

	return &w, nil
}

func (r *WalletRepository) Save(w *wallet.Wallet) error {
	tag, err := r.db.Exec(context.Background(), "UPDATE wallets SET balance = $2 WHERE id = $1", w.ID, w.Balance)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return wallet.ErrWalletNotFound
	}
	return nil
}
