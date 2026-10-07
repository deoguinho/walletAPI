package wallet

import (
	"errors"
	"time"
)

type Wallet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
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
