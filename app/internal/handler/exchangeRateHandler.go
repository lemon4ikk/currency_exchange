package handler

import (
	"currency_exchange/internal/middleware"
	"currency_exchange/internal/service"
	"net/http"
)

type ExchangeRateHandler struct {
	exchange *service.ExchangeRateService
}

func NewExchangeRateHandler(c *service.ExchangeRateService) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		exchange: c,
	}
}

func (c *ExchangeRateHandler) AllHandler(w http.ResponseWriter, r *http.Request) {
	e, msg := c.exchange.AllExchange(w, r)
	if msg.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, msg)
		return
	}

	middleware.WriteJSON(w, http.StatusOK, e)
}

func (c *ExchangeRateHandler) CodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	e, msg := c.exchange.CodeExchange(code)
	var s int
	if msg.Message != "" {
		if msg.Message == "ошибка" {
			s = http.StatusInternalServerError
		} else {
			s = http.StatusNotFound
		}

		middleware.WriteJSON(w, s, msg)
		return
	}

	middleware.WriteJSON(w, http.StatusOK, e)
}

func (c *ExchangeRateHandler) NewExchange(w http.ResponseWriter, r *http.Request) {
	baseCurrencyCode := r.FormValue("baseCurrencyCode")
	targetCurrencyCode := r.FormValue("targetCurrencyCode")
	rate := r.FormValue("rate")

	e, msg := c.exchange.NewExchange(baseCurrencyCode, targetCurrencyCode, rate)
	if msg.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, msg)
	}

	middleware.WriteJSON(w, http.StatusCreated, e)
}

func (c *ExchangeRateHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	rate := r.FormValue("rate")

	e, msg := c.exchange.UpdateExchange(code, rate)
	if msg.Message != "" {
		middleware.WriteJSON(w, http.StatusInternalServerError, msg)
	}

	middleware.WriteJSON(w, http.StatusOK, e)
}
