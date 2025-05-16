package validator

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrEmptyCurrencyCode     = errors.New("currency code is empty")
	ErrInvalidCurrencyLength = errors.New("currency code must be either 3")
	ErrCurrencyNotUppercase  = errors.New("currency code must be in uppercase")
)

func ValidateCurrency(currencyCode string) (int, error) {
	if len(currencyCode) == 0 {
		return http.StatusBadRequest, ErrEmptyCurrencyCode
	}

	if len(currencyCode) != 3 && len(currencyCode) != 6 {
		return http.StatusBadRequest, ErrInvalidCurrencyLength
	}

	if currencyCode != strings.ToUpper(currencyCode) {
		return http.StatusBadRequest, ErrCurrencyNotUppercase
	}

	return 0, nil
}
