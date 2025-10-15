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
	Id            int64        `db:"id"             json:"id"`
	CurrencyType  CurrencyType `db:"currency_type"  json:"currency_type"`
	CurrencyValue float32      `db:"currency_value" json:"currency_value"`
	CurrencyDate  string       `db:"currency_date"  json:"currency_date,omitempty"`
}

func (c Currency) Validate() error {
	switch c.CurrencyType {
	case USD, RUB, EUR, JPY:
		return nil
	default:
		return errors.NewValidationError(
			"currency_type",
			"Must be on of USD, RUB, EUR, JPY",
			c.CurrencyType,
		)
	}
}
