package wallet

import (
	"errors"
	"time"
)

const startBalance = 0

var ErrWalletNotFound = errors.New("wallet not found")

type Wallet struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewWallet(userID int64, balance int64) *Wallet {
	return &Wallet{
		UserID:  userID,
		Balance: startBalance,
	}
}

// Deposito
func (w *Wallet) Deposit(amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	w.Balance += amount
	return nil
}

// Saque
func (w *Wallet) Withdraw(amount int64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	if amount > w.Balance {
		return errors.New("insufficient balance")
	}

	w.Balance -= amount
	return nil

}
