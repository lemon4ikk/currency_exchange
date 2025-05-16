package handler

import (
	"currency_exchange/internal/service"
	"currency_exchange/internal/validator"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CurrencyHandler struct {
	currency *service.CurrencyService
}

func NewCurrencyHandler(c *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{
		currency: c,
	}
}

func (s *CurrencyHandler) AllHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	currency, err := s.currency.AllCurrencies()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return currency, http.StatusOK, nil
}

func (s *CurrencyHandler) CodeHandler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	code := r.PathValue("code")

	fmt.Println(code, len(code))

	status, err := validator.ValidateCurrency(code)
	if err != nil {
		return nil, status, err
	}

	currency, err := s.currency.Code(code)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return currency, http.StatusOK, nil
}

func (s *CurrencyHandler) NewCurrency(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	name := values.Get("name")
	code := values.Get("code")
	sign := values.Get("sign")

	status, err := validator.ValidateCurrency(code)
	if err != nil {
		return nil, status, err
	}

	currency, err := s.currency.New(name, code, sign)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return currency, http.StatusOK, nil
}
