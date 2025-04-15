package handler

import (
	"currency_exchange/internal/service"
	"encoding/json"
	"net/http"
)

type CurrencyHandler struct {
	currency *service.CurrencyService
}

func NewCurrencyHandler(c *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{
		currency: c,
	}
}

func (s *CurrencyHandler) AllHandler(w http.ResponseWriter, r *http.Request) {
	currencies, msg := s.currency.AllCurrencies(w, r)
	if msg.Message != "" {
		w.Header().Set("Contetnt-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencies)
}

func (s *CurrencyHandler) CodeHandler(w http.ResponseWriter, r *http.Request) {
	currency, msg := s.currency.Code(w, r)
	if msg.Message != "" {
		w.Header().Set("Contetnt-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currency)
}

func (s *CurrencyHandler) NewCurrency(w http.ResponseWriter, r *http.Request) {
	currency, msg := s.currency.New(w, r)
	if msg.Message != "" {
		w.Header().Set("Contetnt-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)
}
