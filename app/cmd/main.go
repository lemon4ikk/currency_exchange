package main

import (
	"currency_exchange/internal/api"
	"currency_exchange/internal/repository"
	"log"
)

const (
	pathDb = "internal/repository/database.db"
)

func main() {
	db, err := repository.InitDB(pathDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	server := api.NewServer(db.Db)
	server.Run()
}
