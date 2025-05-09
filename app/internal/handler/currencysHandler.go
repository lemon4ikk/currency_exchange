package handler

import (
	"currency_exchange/internal/middleware"
	"currency_exchange/internal/service"
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
	currencies, msg := s.currency.AllCurrencies()
	if msg.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, msg)
		return
	}

	middleware.WriteJSON(w, http.StatusOK, currencies)
}

func (s *CurrencyHandler) CodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	currency, msg := s.currency.Code(code)
	if msg.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, msg)
		return
	}

	middleware.WriteJSON(w, http.StatusOK, currency)
}

func (s *CurrencyHandler) NewCurrency(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	code := r.FormValue("code")
	sign := r.FormValue("sign")

	currency, msg := s.currency.New(name, code, sign)
	if msg.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, msg)
		return
	}

	middleware.WriteJSON(w, http.StatusCreated, currency)
}
