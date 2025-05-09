package main

import (
	"currency_exchange/internal/api"
	"currency_exchange/internal/handler"
	"currency_exchange/internal/repository"
	"currency_exchange/internal/service"
	"log"
)

const (
	pathDb = "../internal/repository/database.db"
)

func main() {
	db, err := repository.InitDB(pathDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	currencyRepo := repository.NewCurrencyRepo(db.Db)
	currencyServ := service.NewCurrencyService(currencyRepo)
	currencyHandler := handler.NewCurrencyHandler(&currencyServ)

	exchangeRateRepo := repository.NewExchangeRateRepo(db.Db)
	exchangeRateServ := service.NewExchangeRateService(exchangeRateRepo)
	exchangeRateHandler := handler.NewExchangeRateHandler(exchangeRateServ)

	exchangeRepo := repository.NewExchangeRepo(db.Db)
	exchangeServ := service.NewExchangeService(exchangeRepo)
	exchangeHandler := handler.NewExchangeHandler(&exchangeServ)

	server := api.NewServer(currencyHandler, exchangeRateHandler, exchangeHandler)
	server.Run()
}
