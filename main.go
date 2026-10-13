package main

import (
	"fmt"
	"log"
	"net/http"

	_ "walletAPI/docs"
	"walletAPI/internal/application/transaction"
	"walletAPI/internal/application/wallet"
	"walletAPI/internal/config"
	handler "walletAPI/internal/handler/wallet"
	"walletAPI/internal/infra/database"

	_ "walletAPI/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Wallet API
// @version 1.0
// @description API de carteira digital desenvolvida em Go.
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.Load()

	pool, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	fmt.Println("Conectado ao PostgreSQL!")
	repository := database.NewWalletRepository(pool)
	transactionRepository := database.NewTransactionRepository(pool)

	//[Wallet Handler]
	walletHandler := handler.NewWalletHandler(
		wallet.NewCreateWallet(repository),
		wallet.NewDepositMoney(repository),
		wallet.NewWithdrawMoney(repository),
		transaction.NewTransferMoney(repository, database.NewTransactionManager(pool), transactionRepository),
	)

	//[Wallet Route]
	http.HandleFunc("POST /wallets", walletHandler.CreateWallet)
	http.HandleFunc("POST /wallets/{id}/deposit", walletHandler.DepositMoney)
	http.HandleFunc("POST /wallets/{id}/withdraw", walletHandler.WithdrawMoney)
	http.HandleFunc("POST /wallets/transfer", walletHandler.TransferMoney)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	log.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
