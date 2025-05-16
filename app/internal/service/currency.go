package service

import (
	"currency_exchange/internal/repository"
	"currency_exchange/internal/templates"
)

type CurrencyService struct {
	repo repository.CurrencyRepo
}

func NewCurrencyService(r repository.CurrencyRepo) CurrencyService {
	return CurrencyService{
		repo: r,
	}
}

func (s *CurrencyService) AllCurrencies() ([]templates.Currency, error) {
	return s.repo.All()
}

func (s *CurrencyService) Code(code string) (templates.Currency, error) {
	return s.repo.Code(code)
}

func (s *CurrencyService) New(name, code, sign string) (templates.Currency, error) {
	return s.repo.New(name, code, sign)
}
