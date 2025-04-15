package main

import (
	"currency_exchange/internal/app"
	"currency_exchange/internal/repository"
	"log"
)

const (
	pathDb = "database.db"
)

func main() {
	db, err := repository.InitDB(pathDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	server := app.NewServer(db.Db)
	server.Run()
}
