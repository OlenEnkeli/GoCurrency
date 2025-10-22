package dtos

import (
	"github.com/OlenEnkeli/GoCurrency/internal/errors"
)

type CurrencyType string

const (
	USD CurrencyType = "USD"
	RUB CurrencyType = "RUB"
	EUR CurrencyType = "EUR"
	JPY CurrencyType = "JPY"
)

type Currency struct {
	Id           int64        `json:"id"`
	CurrencyType CurrencyType `json:"type"`
	Rate         float32      `json:"rate"`
	CurrencyDate string       `json:"date"`
}

type CurrencyPair struct {
	Left  CurrencyType `json:"left"`
	Right CurrencyType `json:"right"`
	Rate  float32      `json:"rate"`
}

func (c CurrencyType) Validate() error {
	switch c {
	case USD, RUB, EUR, JPY:
		return nil
	default:
		return errors.NewValidationError(
			"currency_type",
			"Must be on of USD, RUB, EUR, JPY",
			c,
		)
	}
}

func (c Currency) Validate() error {
	return c.CurrencyType.Validate()
}
