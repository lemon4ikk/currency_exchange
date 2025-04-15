package templates

type Exchange struct {
	ID               int     `json:"id"`
	BaseCurrencyId   int     `json:"baseCurrencyId"`
	TargetCurrencyId int     `json:"targetCurrencyId"`
	Rate             float32 `json:"rate"`
}

type ExchangeRate struct {
	ID           int `json:"id"`
	BaseCurrency struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Sign string `json:"sign"`
	} `json:"baseCurrency"`
	TargetCurrency struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Sign string `json:"sign"`
	} `json:"targetCurrency"`
	Rate float32 `json:"rate"`
}

type ExchangeRateAmount struct {
	BaseCurrency struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Sign string `json:"sign"`
	} `json:"baseCurrency"`
	TargetCurrency struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
		Sign string `json:"sign"`
	} `json:"targetCurrency"`
	Rate            float32 `json:"rate"`
	Amount          float32 `json:"amount"`
	ConvertedAmount float32 `json:"convertedAmount"`
}
