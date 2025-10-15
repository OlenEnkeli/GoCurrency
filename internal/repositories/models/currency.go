package models

import (
	"time"

	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
)

type Currency struct {
	Id            int64     `db:"id"`
	CurrencyType  string    `db:"currency_type"`
	CurrencyValue float32   `db:"currency_value"`
	CurrencyDate  string    `db:"currency_date"`
	CreatedAt     time.Time `db:"created_at"`
}

func (c *Currency) ToDTO() (dtos.Currency, error) {
	dto := dtos.Currency{
		Id:            c.Id,
		CurrencyType:  dtos.CurrencyType(c.CurrencyType),
		CurrencyValue: c.CurrencyValue,
		CurrencyDate:  c.CurrencyDate,
	}

	return dto, dto.Validate()
}
