package service

import (
	"currency_exchange/internal/templates"
	"database/sql"
	"log"
	"net/http"
)

type CurrencyService struct {
	currencyDB *sql.DB
}

func NewCurrencyService(s *sql.DB) *CurrencyService {
	return &CurrencyService{
		currencyDB: s,
	}
}

func (s *CurrencyService) AllCurrencies(w http.ResponseWriter, r *http.Request) ([]templates.Currency, templates.Msg) {
	var currencies []templates.Currency
	var m templates.Msg

	rows, err := s.currencyDB.Query("SELECT * FROM Currencies")
	if err != nil {
		m.Message = "ошибка"
		log.Fatal(err)
		return nil, m
	}
	defer rows.Close()

	for rows.Next() {
		var c templates.Currency
		if err = rows.Scan(&c.ID, &c.Code, &c.FullName, &c.Sign); err != nil {
			m.Message = "ошибка"
			return nil, m
		}
		currencies = append(currencies, c)
	}

	return currencies, m
}

func (s *CurrencyService) Code(w http.ResponseWriter, r *http.Request) (templates.Currency, templates.Msg) {
	var currency templates.Currency
	var m templates.Msg
	c := r.PathValue("code")

	rows := s.currencyDB.QueryRow("SELECT * FROM Currencies WHERE Code = ?;", c)

	if err := rows.Scan(&currency.ID, &currency.Code, &currency.FullName, &currency.Sign); err != nil {
		if err == sql.ErrNoRows {
			m.Message = "валюта не найдена"
			return currency, m
		}

		m.Message = "ошибка"
		return currency, m
	}

	return currency, m
}

func (s *CurrencyService) New(w http.ResponseWriter, r *http.Request) (templates.Currency, templates.Msg) {
	var newCurrency templates.Currency
	var m templates.Msg
	name := r.FormValue("name")
	code := r.FormValue("code")
	sign := r.FormValue("sign")

	newValues := s.currencyDB.QueryRow("INSERT INTO Currencies (Code, FullName, Sign) VALUES (?, ?, ?) RETURNING ID, Code, FullName, Sign;", code, name, sign)

	if err := newValues.Scan(&newCurrency.ID, &newCurrency.Code, &newCurrency.FullName, &newCurrency.Sign); err != nil {
		if err.Error() == "UNIQUE constraint failed: Currencies.Code" {
			m.Message = "валюта с таким кодом уже существует"
			return newCurrency, m
		}

		m.Message = "ошибка"
		log.Fatal(err.Error())
		return newCurrency, m
	}

	return newCurrency, m
}
