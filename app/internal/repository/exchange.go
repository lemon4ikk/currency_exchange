package repository

import (
	"currency_exchange/internal/templates"
	"database/sql"
	"log"
	"strconv"
)

type ExchangeRepo struct {
	repo *sql.DB
}

func NewExchangeRepo(r *sql.DB) ExchangeRepo {
	return ExchangeRepo{
		repo: r,
	}
}

func (e *ExchangeRepo) Exchange(baseCode, targetCode, amount string) ([]templates.ExchangeRateAmount, error) {
	var amountExchange []templates.ExchangeRateAmount
	var baseId, targetId int
	var m templates.Msg

	baseCurrensyId := e.repo.QueryRow("SELECT ID FROM Currencies WHERE Code = ?;", baseCode)
	targetCurrensyId := e.repo.QueryRow("SELECT ID FROM Currencies WHERE Code = ?;", targetCode)

	if err := baseCurrensyId.Scan(&baseId); err != nil {
		if baseId == 0 {
			m.Message = "валюта с кодом " + baseCode + " не найдена"
			return nil, err
		}

		log.Fatalf("Scan completed with error %v", err)
	}

	if err := targetCurrensyId.Scan(&targetId); err != nil {
		if targetId == 0 {
			m.Message = "валюта с кодом " + targetCode + " не найдена"
			return nil, err
		}

		log.Fatalf("Scan completed with error %v", err)
	}

	rateSql, err := e.repo.Query("SELECT Rate FROM ExchangeRates WHERE BaseCurrencyId = ? AND TargetCurrencyId = ?;", baseId, targetId)
	if err != nil {
		return nil, err
	}
	var rate float64

	for rateSql.Next() {
		if err := rateSql.Scan(&rate); err != nil {
			log.Fatal(err)
		}
	}

	var reverseFlag bool
	var usdRateFlag bool

	if rate == 0 {
		reverseFlag = true
		usdRateFlag = true
		baseId, targetId = targetId, baseId
		reverseRateSql, err := e.repo.Query("SELECT Rate FROM ExchangeRates WHERE BaseCurrencyId = ? AND TargetCurrencyId = ?;", baseId, targetId)
		if err != nil {
			log.Fatal(err)
		}
		var revercseRate float64

		for reverseRateSql.Next() {
			if err := reverseRateSql.Scan(&revercseRate); err != nil {
				break
			}
		}

		if revercseRate == 0 {
			reverseFlag = false
		} else {
			usdRateFlag = false
		}

		baseId, targetId = targetId, baseId
		var baseUsdRate, targetUsdRate float64
		usdID := 1

		baseUsdRateSql, err := e.repo.Query("SELECT Rate FROM ExchangeRates WHERE BaseCurrencyId = ? AND TargetCurrencyId = ?;", usdID, baseId)
		if err != nil {
			log.Fatal(err)
		}

		targetUsdRateSql, err := e.repo.Query("SELECT Rate FROM ExchangeRates WHERE BaseCurrencyId = ? AND TargetCurrencyId = ?;", usdID, targetId)
		if err != nil {
			log.Fatal(err)
		}

		for baseUsdRateSql.Next() {
			if err := baseUsdRateSql.Scan(&baseUsdRate); err != nil {
				return amountExchange, err
			}
		}

		for targetUsdRateSql.Next() {
			if err := targetUsdRateSql.Scan(&targetUsdRate); err != nil {
				return amountExchange, err
			}
		}

		if baseUsdRate == 0 || targetUsdRate == 0 {
			m.Message = "обменный курс не найден"
			return amountExchange, err
		}

		if reverseFlag {
			rate = 1 / revercseRate
		}

		if usdRateFlag {
			rate = targetUsdRate / baseUsdRate
		}
	}

	a, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		log.Fatalf("ParseFloat complited with error %v", err)
	}

	convertedAmount := rate * a

	currencyInfo, err := e.repo.Query(
		`SELECT c.ID AS base_currency_id,
	    	    c.FullName AS bc_name,
	    		c.Code AS bc_code,
	    		c.Sign AS bc_sign,
	    		v.ID AS target_currency_id,
	    		v.FullName AS tg_name,
	    		v.Code AS tg_code,
	    		v.Sign AS tg_sign
	    FROM ExchangeRates e
	    JOIN Currencies c ON e.BaseCurrencyId = c.ID
	    JOIN Currencies v ON e.TargetCurrencyId = v.ID 
	    WHERE base_currency_id = ? AND target_currency_id = ?;`, baseId, targetId)
	if err != nil {
		log.Fatal(err)
	}

	if usdRateFlag {
		currencyInfo, err = e.repo.Query(
			`SELECT c.ID AS base_currency_id,
    				c.FullName AS bc_name,
				    c.Code AS bc_code,
				    c.Sign AS bc_sign,
				    v.ID AS target_currency_id,
				    v.FullName AS tg_name,
				    v.Code AS tg_code,
				    v.Sign AS tg_sign
				FROM Currencies c
				JOIN Currencies v ON v.ID = ?
				WHERE c.ID = ?;`, targetId, baseId)
		if err != nil {
			log.Fatal(err)
		}
	}

	for currencyInfo.Next() {
		var e templates.ExchangeRateAmount
		e.Amount = float32(a)
		e.ConvertedAmount = float32(convertedAmount)
		e.Rate = float32(rate)

		if err := currencyInfo.Scan(&e.BaseCurrency.ID, &e.BaseCurrency.Name, &e.BaseCurrency.Code, &e.BaseCurrency.Sign, &e.TargetCurrency.ID, &e.TargetCurrency.Name, &e.TargetCurrency.Code, &e.TargetCurrency.Sign); err != nil {
			m.Message = "ошибка"
			return amountExchange, err
		}

		if reverseFlag {
			e.BaseCurrency.ID, e.TargetCurrency.ID = e.TargetCurrency.ID, e.BaseCurrency.ID
			e.BaseCurrency.Name, e.TargetCurrency.Name = e.TargetCurrency.Name, e.BaseCurrency.Name
			e.BaseCurrency.Code, e.TargetCurrency.Code = e.TargetCurrency.Code, e.BaseCurrency.Code
			e.BaseCurrency.Sign, e.TargetCurrency.Sign = e.TargetCurrency.Sign, e.BaseCurrency.Sign
		}

		amountExchange = append(amountExchange, e)
	}

	return amountExchange, nil
}
