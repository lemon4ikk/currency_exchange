package handler

import (
	"currency_exchange/internal/service"
	"encoding/json"
	"net/http"
)

type exchangeHandler struct {
	exchangeRateService *service.ExchangeService
}

func NewExchangeRateHandler(e *service.ExchangeService) *exchangeHandler {
	return &exchangeHandler{
		exchangeRateService: e,
	}
}

func (e *exchangeHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
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
