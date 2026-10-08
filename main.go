package main

import (
	"fmt"
	"log"
	"net/http"

	"walletAPI/internal/application/wallet"
	"walletAPI/internal/handler"
	"walletAPI/internal/infra/database"
)

func main() {
	databaseURL := "postgres://wallet:wallet@localhost:5432/wallet"

	pool, err := database.NewPostgresPool(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	fmt.Println("Conectado ao PostgreSQL!")
	repository := database.NewWalletRepository(pool)

	//Wallet Handler
	walletHandler := handler.NewWalletHandler(
		wallet.NewCreateWallet(repository),
		wallet.NewDepositMoney(repository),
		wallet.NewWithdrawMoney(repository),
		wallet.NewTransferMoney(repository, database.NewTransactionManager(pool)),
	)
	http.HandleFunc("POST /wallets", walletHandler.CreateWallet)
	http.HandleFunc("POST /wallets/{id}/deposit", walletHandler.DepositMoney)
	http.HandleFunc("POST /wallets/{id}/withdraw", walletHandler.WithdrawMoney)
	http.HandleFunc("POST /wallets/transfer", walletHandler.TransferMoney)

	log.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
