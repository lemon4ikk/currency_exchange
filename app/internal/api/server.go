package api

import (
	"currency_exchange/internal/handler"
	"currency_exchange/internal/middleware"
	"log"
	"net/http"
)

type Server struct {
	CurrencyHandler     *handler.CurrencyHandler
	ExchangeRateHandler *handler.ExchangeRateHandler
	ExchangeHandler     *handler.ExchangeHandler
}

func NewServer(currencyHandler *handler.CurrencyHandler, exchangeRateHandler *handler.ExchangeRateHandler, exchangeHandler *handler.ExchangeHandler) *Server {
	return &Server{
		CurrencyHandler:     currencyHandler,
		ExchangeRateHandler: exchangeRateHandler,
		ExchangeHandler:     exchangeHandler,
	}
}

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /currencies", middleware.WriteJSON(s.CurrencyHandler.AllHandler))
	mux.HandleFunc("GET /currency/{code}", middleware.WriteJSON(s.CurrencyHandler.CodeHandler))
	mux.HandleFunc("POST /currencies", middleware.WriteJSON(s.CurrencyHandler.NewCurrency))

	mux.HandleFunc("GET /exchangeRates", middleware.WriteJSON(s.ExchangeRateHandler.AllHandler))
	mux.HandleFunc("GET /exchangeRates/{code}", middleware.WriteJSON(s.ExchangeRateHandler.CodeHandler))
	mux.HandleFunc("POST /exchangeRates", middleware.WriteJSON(s.ExchangeRateHandler.NewExchange))
	mux.HandleFunc("PATCH /exchangeRates/{code}", middleware.WriteJSON(s.ExchangeRateHandler.UpdateHandler))

	mux.HandleFunc("GET /exchange", middleware.WriteJSON(s.ExchangeHandler.SearchHandler))
}

func (s *Server) Run() {
	mux := http.NewServeMux()
	s.registerRoutes(mux)

	handler := middleware.Logging(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error listening to the port: %v", err)
	}
}
