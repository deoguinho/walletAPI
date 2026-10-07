package main

import (
	"fmt"
	"log"

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
}
