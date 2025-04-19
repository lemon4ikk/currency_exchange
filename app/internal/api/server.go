package api

import (
	"currency_exchange/internal/handler"
	"currency_exchange/internal/service"
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	db *sql.DB
}

func NewServer(inputDB *sql.DB) *Server {
	return &Server{
		db: inputDB,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	currencyService := service.NewCurrencyService(s.db)
	currencyHandler := handler.NewCurrencyHandler(currencyService)

	exchangeService := service.NewExchangeService(s.db)
	exchangeHandler := handler.NewExchangeHandler(exchangeService)

	exchangeRateService := service.NewExchangeRateService(s.db)
	exchangeRateHandler := handler.NewExchangeRateHandler(exchangeRateService)

	mux.HandleFunc("GET /currencies", currencyHandler.AllHandler)
	mux.HandleFunc("GET /currency/{code}", currencyHandler.CodeHandler)
	mux.HandleFunc("POST /currencies", currencyHandler.NewCurrency)

	mux.HandleFunc("GET /exchangeRates", exchangeHandler.AllHandler)
	mux.HandleFunc("GET /exchangeRates/{code}", exchangeHandler.CodeHandler)
	mux.HandleFunc("POST /exchangeRates", exchangeHandler.NewExchange)
	mux.HandleFunc("PATCH /exchangeRates/{code}", exchangeHandler.UpdateHandler)

	mux.HandleFunc("GET /exchange", exchangeRateHandler.SearchHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error listening to the port: %v", err)
	}
}
