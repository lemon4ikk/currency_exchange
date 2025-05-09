package service

import (
	"currency_exchange/internal/repository"
	"currency_exchange/internal/templates"
)

type ExchangeService struct {
	repo repository.ExchangeRepo
}

func NewExchangeService(r repository.ExchangeRepo) ExchangeService {
	return ExchangeService{
		repo: r,
	}
}

func (e *ExchangeService) Exchange(baseCode, targetCode, amount string) ([]templates.ExchangeRateAmount, templates.Msg) {
	return e.repo.Exchange(baseCode, targetCode, amount)
}
