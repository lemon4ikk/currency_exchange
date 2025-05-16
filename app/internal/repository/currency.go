package repository

import (
	"currency_exchange/internal/templates"
	"database/sql"
	"errors"
)

type CurrencyRepo struct {
	repo *sql.DB
}

var ErrCurrencyNotFound error = errors.New("Currency not found")

func NewCurrencyRepo(r *sql.DB) CurrencyRepo {
	return CurrencyRepo{
		repo: r,
	}
}

func (s *CurrencyRepo) All() ([]templates.Currency, error) {
	var currencies []templates.Currency
	rows, err := s.repo.Query("SELECT * FROM Currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c templates.Currency
		if err = rows.Scan(&c.ID, &c.Code, &c.FullName, &c.Sign); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrCurrencyNotFound
			}

			return nil, err
		}
		currencies = append(currencies, c)
	}

	return currencies, nil
}

func (s *CurrencyRepo) Code(code string) (templates.Currency, error) {
	var currency templates.Currency

	rows := s.repo.QueryRow("SELECT * FROM Currencies WHERE Code = ?;", code)

	if err := rows.Scan(&currency.ID, &currency.Code, &currency.FullName, &currency.Sign); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return currency, ErrCurrencyNotFound
		}

		return currency, err
	}

	return currency, nil
}

func (s *CurrencyRepo) New(name, code, sign string) (templates.Currency, error) {
	var newCurrency templates.Currency

	newValues := s.repo.QueryRow("INSERT INTO Currencies (Code, FullName, Sign) VALUES (?, ?, ?) RETURNING ID, Code, FullName, Sign;", code, name, sign)

	if err := newValues.Scan(&newCurrency.ID, &newCurrency.Code, &newCurrency.FullName, &newCurrency.Sign); err != nil {
		if err.Error() == "UNIQUE constraint failed: Currencies.Code" {
			return newCurrency, err
		}

		return newCurrency, err
	}

	return newCurrency, nil
}
