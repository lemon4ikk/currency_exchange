package handler

import (
	"currency_exchange/internal/repository"
	"currency_exchange/internal/service"
	"currency_exchange/internal/validator"
	"io"
	"net/http"
	"net/url"
)

type ExchangeRateHandler struct {
	exchange *service.ExchangeRateService
}

func NewExchangeRateHandler(c *service.ExchangeRateService) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		exchange: c,
	}
}

func (c *ExchangeRateHandler) AllHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	exchangeRates, err := c.exchange.AllExchange(w, r)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return exchangeRates, http.StatusOK, nil
}

func (c *ExchangeRateHandler) CodeHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	code := r.PathValue("code")

	status, err := validator.ValidateCurrency(code)
	if err != nil {
		return nil, status, err
	}

	exchangeRate, err := c.exchange.CodeExchange(code)
	if err != nil {
		s := http.StatusInternalServerError

		if err == repository.ErrExchangeRateNotFound {
			s = http.StatusNotFound
		}

		return nil, s, err
	}

	return exchangeRate, http.StatusOK, nil
}

func (c *ExchangeRateHandler) NewExchange(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	baseCurrencyCode := values.Get("baseCurrencyCode")
	targetCurrencyCode := values.Get("targetCurrencyCode")
	rate := values.Get("rate")

	status, err := validator.ValidateCurrency(baseCurrencyCode)
	if err != nil {
		return nil, status, err
	}

	status, err = validator.ValidateCurrency(targetCurrencyCode)
	if err != nil {
		return nil, status, err
	}

	exchangeRate, err := c.exchange.NewExchange(baseCurrencyCode, targetCurrencyCode, rate)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return exchangeRate, http.StatusOK, nil
}

func (c *ExchangeRateHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	code := values.Get("code")
	rate := values.Get("rate")

	status, err := validator.ValidateCurrency(code)
	if err != nil {
		return nil, status, err
	}

	exchangeRate, err := c.exchange.UpdateExchange(code, rate)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return exchangeRate, http.StatusOK, nil
}
