package wallet

import "testing"

func TestWalletWithdraw(t *testing.T) {
	tests := []struct {
		name    string
		wallet  Wallet
		amount  int64
		wantErr bool
	}{
		{name: "insufficient balance",
			wallet: Wallet{
				Balance: 1000,
			},
			amount:  1500,
			wantErr: true,
		}, {name: "sufficient balance",
			wallet: Wallet{
				Balance: 2000,
			},
			amount:  1500,
			wantErr: false,
		},
		{name: "negative withdrawal amount",
			wallet: Wallet{
				Balance: 2000,
			},
			amount:  -500,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		err := tt.wallet.Withdraw(tt.amount)
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("Wallet.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Wallet balance after withdrawal: %d", tt.wallet.Balance)
		})
	}
}

func TestWalletDeposit(t *testing.T) {
	w := Wallet{
		Balance: 1000,
	}

	err := w.Deposit(500)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if w.Balance != 1500 {
		t.Fatalf("expected balance 1500, got %d", w.Balance)
	}
}

func TestWalletDepositZeroAmount(t *testing.T) {
	w := Wallet{
		Balance: 1000,
	}

	err := w.Deposit(0)
	if err == nil {
		t.Fatal("expected error for zero deposit amount, got nil")
	}
}

func TestDepositSufficientBalance(t *testing.T) {
	w := Wallet{
		Balance: 1000,
	}
	const depositAmount int64 = 500

	err := w.Deposit(depositAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if w.Balance != 1500 {
		t.Fatalf("expected balance 1500, got %d", w.Balance)
	}

}

func TestDepositNegativeAmount(t *testing.T) {
	w := Wallet{
		Balance: 1000,
	}
	const depositAmount int64 = -500
	const expectedErrorMessage = "amount must be greater than zero"

	err := w.Deposit(depositAmount)
	if err == nil {
		t.Fatal("expected error for negative deposit amount, got nil")
	}
}

func TestDepositZeroAmount(t *testing.T) {
	w := Wallet{
		Balance: 1000,
	}
	const depositAmount int64 = 0
	const expectedErrorMessage = "amount must be greater than zero"

	err := w.Deposit(depositAmount)
	if err == nil {
		t.Fatal("expected error for zero deposit amount, got nil")
	}
}
