package handler

import (
	"currency_exchange/internal/service"
	"currency_exchange/internal/validator"
	"io"
	"net/http"
	"net/url"
)

type ExchangeHandler struct {
	exchangeService *service.ExchangeService
}

func NewExchangeHandler(e *service.ExchangeService) *ExchangeHandler {
	return &ExchangeHandler{
		exchangeService: e,
	}
}

func (e *ExchangeHandler) SearchHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	baseCode := values.Get("from")
	targetCode := values.Get("to")
	amount := values.Get("amount")

	status, err := validator.ValidateCurrency(baseCode)
	if err != nil {
		return nil, status, err
	}

	status, err = validator.ValidateCurrency(targetCode)
	if err != nil {
		return nil, status, err
	}

	exchange, err := e.exchangeService.Exchange(baseCode, targetCode, amount)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return exchange, http.StatusOK, nil
}
