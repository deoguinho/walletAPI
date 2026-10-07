package transaction

import "testing"

func TestTransactionValidate(t *testing.T) {

	walletID := "wallet1"
	toWalletId := "wallet2"

	tests := []struct {
		name        string
		transaction Transaction
		wantErr     bool
	}{
		{
			name: "valid deposit transaction",
			transaction: Transaction{
				Type:         DepositTransaction,
				FromWalletID: nil,
				ToWalletID:   &walletID,
				Amount:       100,
				Status:       "pending",
			},
			wantErr: false,
		},
		{
			name: "invalid deposit transaction with from wallet",
			transaction: Transaction{
				Type:         DepositTransaction,
				FromWalletID: &toWalletId,
				ToWalletID:   &walletID,
				Amount:       100,
				Status:       "pending",
			},
			wantErr: true,
		},
		{name: "valid withdrawal transaction",
			transaction: Transaction{
				Type:         WithdrawalTransaction,
				FromWalletID: &walletID,
				ToWalletID:   nil,
				Amount:       100,
				Status:       "pending",
			},
			wantErr: false,
		},
		{name: "valid transfer transaction",
			transaction: Transaction{
				Type:         TransferTransaction,
				FromWalletID: &walletID,
				ToWalletID:   &toWalletId,
				Amount:       100,
				Status:       "pending",
			},
			wantErr: false,
		},
		{name: "invalid transfer transaction with missing to wallet",
			transaction: Transaction{
				Type:         TransferTransaction,
				FromWalletID: &walletID,
				ToWalletID:   nil,
				Amount:       100,
				Status:       "pending",
			},
			wantErr: true,
		},
		{name: "invalid transfer transaction with missing from wallet",
			transaction: Transaction{
				Type:         TransferTransaction,
				FromWalletID: nil,
				ToWalletID:   &toWalletId,
				Amount:       100,
				Status:       "pending",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		err := tt.transaction.Validate()
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"%s: expected error: %v, got: %v",
					tt.name,
					tt.wantErr,
					err,
				)
			}
		})

	}
}
