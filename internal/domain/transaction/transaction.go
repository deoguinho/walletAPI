package transaction

import (
	"fmt"
	"time"
)

type TransactionType string
type TransactionStatus string

const (
	DepositTransaction    TransactionType = "deposit"
	WithdrawalTransaction TransactionType = "withdrawal"
	TransferTransaction   TransactionType = "transfer"
)

const (
	PendingStatus   TransactionStatus = "pending"
	CompletedStatus TransactionStatus = "completed"
	FailedStatus    TransactionStatus = "failed"
)

type Transaction struct {
	ID           string            `json:"id"`
	FromWalletID *string           `json:"from_wallet_id"`
	ToWalletID   *string           `json:"to_wallet_id"`
	Amount       int64             `json:"amount"`
	Type         TransactionType   `json:"type"`
	Status       TransactionStatus `json:"status"`
	CreatedAt    time.Time         `json:"created_at"`
}

func (t *Transaction) Validate() error {
	if t.FromWalletID == t.ToWalletID {
		return fmt.Errorf("invalid transaction: from and to wallet IDs cannot be the same")
	}

	if t.Amount <= 0 {
		return fmt.Errorf("invalid transaction amount: must be greater than zero")
	}

	types := map[TransactionType]bool{
		DepositTransaction:    true,
		WithdrawalTransaction: true,
		TransferTransaction:   true,
	}

	if !types[t.Type] {
		return fmt.Errorf("invalid transaction type: %s", t.Type)
	}

	statuses := map[TransactionStatus]bool{
		PendingStatus:   true,
		CompletedStatus: true,
		FailedStatus:    true,
	}

	if !statuses[t.Status] {
		return fmt.Errorf("invalid transaction status: %s", t.Status)
	}

	switch t.Type {
	case DepositTransaction:
		if t.ToWalletID == nil || t.FromWalletID != nil {
			return fmt.Errorf("invalid deposit transaction: missing to wallet ID")
		}
	case WithdrawalTransaction:
		if t.FromWalletID == nil || t.ToWalletID != nil {
			return fmt.Errorf("invalid withdrawal transaction: missing from wallet ID")
		}
	case TransferTransaction:
		if t.FromWalletID == nil || t.ToWalletID == nil {
			return fmt.Errorf("invalid transfer transaction: missing from or to wallet ID")
		}
	}

	return nil
}
