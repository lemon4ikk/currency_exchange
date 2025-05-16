package service

import (
	"currency_exchange/internal/repository"
	"currency_exchange/internal/templates"
	"net/http"
)

type ExchangeRateService struct {
	repo repository.ExchangeRateRepo
}

func NewExchangeRateService(r repository.ExchangeRateRepo) *ExchangeRateService {
	return &ExchangeRateService{
		repo: r,
	}
}

func (e *ExchangeRateService) AllExchange(w http.ResponseWriter, r *http.Request) ([]templates.ExchangeRate, error) {
	return e.repo.AllExchange()
}

func (e *ExchangeRateService) CodeExchange(code string) ([]templates.ExchangeRate, error) {
	return e.repo.CodeExchange(code)
}

func (e *ExchangeRateService) NewExchange(baseCurrencyCode, targetCurrencyCode, rate string) ([]templates.ExchangeRate, error) {
	return e.repo.NewExchange(baseCurrencyCode, targetCurrencyCode, rate)
}

func (e *ExchangeRateService) UpdateExchange(code, rate string) ([]templates.ExchangeRate, error) {
	return e.repo.UpdateExchange(code, rate)
}
