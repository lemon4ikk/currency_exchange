package handler

import (
	"currency_exchange/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type ExchangeHandler struct {
	exchange *service.ExchangeService
}

func NewExchangeHandler(c *service.ExchangeService) *ExchangeHandler {
	return &ExchangeHandler{
		exchange: c,
	}
}

func (c *ExchangeHandler) AllHandler(w http.ResponseWriter, r *http.Request) {
	e, msg := c.exchange.AllExchange(w, r)
	if msg.Message != "" {
		w.Header().Set("Contetnt-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (c *ExchangeHandler) CodeHandler(w http.ResponseWriter, r *http.Request) {
	e, msg := c.exchange.CodeExchange(w, r)
	if msg.Message != "" {
		w.Header().Set("Contetnt-Type", "application/json")
		if msg.Message == "ошибка" {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (c *ExchangeHandler) NewExchange(w http.ResponseWriter, r *http.Request) {
	e, msg := c.exchange.NewExchange(w, r)
	if msg.Message != "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (c *ExchangeHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	e, err := c.exchange.UpdateExchange(w, r)
	if err != nil {
		log.Fatalf("method NewExchange complited with error: %v", err)
	}

	json.NewEncoder(w).Encode(e)
	json.NewEncoder(w).Encode(http.StatusOK)
}
