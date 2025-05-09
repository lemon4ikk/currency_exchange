package handler

import (
	"currency_exchange/internal/middleware"
	"currency_exchange/internal/service"
	"net/http"
)

type ExchangeHandler struct {
	exchangeService *service.ExchangeService
}

func NewExchangeHandler(e *service.ExchangeService) *ExchangeHandler {
	return &ExchangeHandler{
		exchangeService: e,
	}
}

func (e *ExchangeHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	baseCode := r.FormValue("from")
	targetCode := r.FormValue("to")
	amount := r.FormValue("amount")

	res, m := e.exchangeService.Exchange(baseCode, targetCode, amount)

	if m.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, m)
		return
	}

	middleware.WriteJSON(w, http.StatusOK, res)
}
