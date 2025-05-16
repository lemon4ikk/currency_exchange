package repository

import (
	"currency_exchange/internal/templates"
	"database/sql"
	"errors"
)

type ExchangeRateRepo struct {
	repo *sql.DB
}

func NewExchangeRateRepo(r *sql.DB) ExchangeRateRepo {
	return ExchangeRateRepo{
		repo: r,
	}
}

var ErrExchangeRateNotFound error = errors.New("Exchange rate not found")
var ErrExchangeRateDoesNotExist error = errors.New("Currency from the currency pair does not exist in the database")

func (e *ExchangeRateRepo) AllExchange() ([]templates.ExchangeRate, error) {
	var exchange []templates.ExchangeRate

	currencyInfo, err := e.repo.Query(
		`SELECT e.ID AS id,
				bc.ID AS bc_id,
			  	bc.FullName AS bc_name,
			  	bc.Code AS bc_code,
			  	bc.Sign AS bc_sign,
				tc.ID AS tc_id,
			  	tc.FullName AS tc_name,
			  	tc.Code AS tc_code,
			  	tc.Sign AS tc_sign,
			  	e.Rate AS rate
		FROM ExchangeRates e
		JOIN Currencies bc ON e.BaseCurrencyId = bc.ID
		JOIN Currencies tc ON e.TargetCurrencyId = tc.ID;`)
	if err != nil {
		return nil, err
	}

	for currencyInfo.Next() {
		var e templates.ExchangeRate
		if err = currencyInfo.Scan(&e.ID, &e.BaseCurrency.ID, &e.BaseCurrency.Name, &e.BaseCurrency.Code, &e.BaseCurrency.Sign, &e.TargetCurrency.ID, &e.TargetCurrency.Name, &e.TargetCurrency.Code, &e.TargetCurrency.Sign, &e.Rate); err != nil {
			return nil, err
		}

		exchange = append(exchange, e)
	}

	return exchange, nil
}

func (er *ExchangeRateRepo) CodeExchange(code string) ([]templates.ExchangeRate, error) {
	var exchange []templates.ExchangeRate
	base := string(code[:3])
	target := string(code[3:])

	currencyInfo := er.repo.QueryRow(
		`SELECT e.ID AS id,
		 	    cb.ID AS bc_id,
	    	    cb.FullName AS bc_name,
	    		cb.Code AS bc_code,
	    		cb.Sign AS bc_sign,
	    		ct.ID AS tc_id,
	    		ct.FullName AS tg_name,
	    		ct.Code AS tg_code,
	    		ct.Sign AS tg_sign,
	    		e.Rate AS rate
	    FROM ExchangeRates e
	    JOIN Currencies cb ON e.BaseCurrencyId = cb.ID
	    JOIN Currencies ct ON e.TargetCurrencyId = ct.ID 
	    WHERE bc_code = ? AND tg_code = ?;`, base, target)

	var e templates.ExchangeRate
	if err := currencyInfo.Scan(&e.ID, &e.BaseCurrency.ID, &e.BaseCurrency.Name, &e.BaseCurrency.Code, &e.BaseCurrency.Sign, &e.TargetCurrency.ID, &e.TargetCurrency.Name, &e.TargetCurrency.Code, &e.TargetCurrency.Sign, &e.Rate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrExchangeRateNotFound
		}

		return nil, err
	}

	exchange = append(exchange, e)

	return exchange, nil
}

func (e *ExchangeRateRepo) NewExchange(baseCurrencyCode, targetCurrencyCode, rate string) ([]templates.ExchangeRate, error) {
	var b, t int
	var m templates.Msg

	baseCurrensyId := e.repo.QueryRow("SELECT ID FROM Currencies WHERE Code = ?;", baseCurrencyCode)
	targetCurrensyId := e.repo.QueryRow("SELECT ID FROM Currencies WHERE Code = ?;", targetCurrencyCode)

	if err := baseCurrensyId.Scan(&b); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrExchangeRateDoesNotExist
		}
	}

	if err := targetCurrensyId.Scan(&t); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrExchangeRateDoesNotExist
		}
	}

	e.repo.QueryRow("INSERT INTO ExchangeRates (BaseCurrencyId, TargetCurrencyId, Rate) VALUES (?, ?, ?);", b, t, rate)

	var newExchange []templates.ExchangeRate

	currencyInfo, err := e.repo.Query(
		`SELECT e.ID AS id,
				cb.ID AS bc_id,
	   			cb.FullName AS bc_name,
	   			cb.Code AS bc_code,
	   			cb.Sign AS bc_sign,
	   			ct.ID AS tc_id,
	    		ct.FullName AS tg_name,
	    		ct.Code AS tg_code,
	    		ct.Sign AS tg_sign,
	   			e.Rate AS rate
		FROM ExchangeRates e
		JOIN Currencies cb ON e.BaseCurrencyId = cb.ID
		JOIN Currencies ct ON e.TargetCurrencyId = ct.ID 
		WHERE bc_id = ? AND tc_id = ?;`, b, t)
	if err != nil {
		m.Message = "ошибка"
		return nil, err
	}

	for currencyInfo.Next() {
		var e templates.ExchangeRate
		if err = currencyInfo.Scan(&e.ID, &e.BaseCurrency.ID, &e.BaseCurrency.Name, &e.BaseCurrency.Code, &e.BaseCurrency.Sign, &e.TargetCurrency.ID, &e.TargetCurrency.Name, &e.TargetCurrency.Code, &e.TargetCurrency.Sign, &e.Rate); err != nil {
			m.Message = "ошибка"
			return nil, err
		}

		newExchange = append(newExchange, e)
	}

	return newExchange, nil
}

func (e *ExchangeRateRepo) UpdateExchange(c, rate string) ([]templates.ExchangeRate, error) {
	var baseId, targetId int
	var m templates.Msg

	baseCurrencyCode := string(c[0]) + string(c[1]) + string(c[2])
	targetCurrencyCode := string(c[3]) + string(c[4]) + string(c[5])

	baseCurrensyId := e.repo.QueryRow("SELECT ID FROM Currencies WHERE Code = ?;", baseCurrencyCode)
	targetCurrensyId := e.repo.QueryRow("SELECT ID FROM Currencies WHERE Code = ?;", targetCurrencyCode)

	if err := baseCurrensyId.Scan(&baseId); err != nil {
		if err == sql.ErrNoRows {
			m.Message = "валюта с кодом " + baseCurrencyCode + " из валютной пары не существует в БД"
			return nil, err
		}
	}

	if err := targetCurrensyId.Scan(&targetId); err != nil {
		if err == sql.ErrNoRows {
			m.Message = "валюта с кодом " + targetCurrencyCode + " из валютной пары не сузествует в БД"
			return nil, err
		}
	}

	e.repo.Exec("UPDATE ExchangeRates SET Rate = ? WHERE BaseCurrencyId = ? AND TargetCurrencyId = ?;", rate, baseId, targetId)

	var newExchange []templates.ExchangeRate

	currencyInfo, err := e.repo.Query(
		`SELECT e.ID AS id,
				cb.ID AS bc_id,
			    cb.FullName AS bc_name,
			    cb.Code AS bc_code,
			    cb.Sign AS bc_sign,
			    ct.ID AS tc_id,
	    		ct.FullName AS tg_name,
	    		ct.Code AS tg_code,
	    		ct.Sign AS tg_sign,
			    e.Rate AS rate
		FROM ExchangeRates e
		JOIN Currencies cb ON e.BaseCurrencyId = cb.ID
		JOIN Currencies ct ON e.TargetCurrencyId = ct.ID 
		WHERE bc_id = ? AND tc_id = ?;`, baseId, targetId)
	if err != nil {
		m.Message = "ошибка"
		return nil, err
	}

	for currencyInfo.Next() {
		var e templates.ExchangeRate
		if err = currencyInfo.Scan(&e.ID, &e.BaseCurrency.ID, &e.BaseCurrency.Name, &e.BaseCurrency.Code, &e.BaseCurrency.Sign, &e.TargetCurrency.ID, &e.TargetCurrency.Name, &e.TargetCurrency.Code, &e.TargetCurrency.Sign, &e.Rate); err != nil {
			m.Message = "ошибка"
			return nil, err
		}

		newExchange = append(newExchange, e)
	}

	return newExchange, nil
}
