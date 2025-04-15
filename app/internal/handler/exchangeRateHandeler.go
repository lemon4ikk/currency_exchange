package handler

import (
	"currency_exchange/internal/service"
	"encoding/json"
	"net/http"
)

type exchangeRateHandler struct {
	exchangeRateService *service.ExchangeRateService
}

func NewExchangeRateHandler(e *service.ExchangeRateService) *exchangeRateHandler {
	return &exchangeRateHandler{
		exchangeRateService: e,
	}
}

func (e *exchangeRateHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	res, m := e.exchangeRateService.Exchange(w, r)

	if m.Message != "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(m)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}
}
