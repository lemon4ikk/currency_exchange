package repository

import "currency_exchange/internal/templates"

type Currency interface {
	AllCurrencies() ([]templates.Currency, templates.Msg)
	Code() (templates.Currency, templates.Msg)
	New() (templates.Currency, templates.Msg)
}
