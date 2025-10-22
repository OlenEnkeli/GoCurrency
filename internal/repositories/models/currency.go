package models

import (
	"time"

	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
)

type Currency struct {
	Id           int64     `db:"id"`
	CurrencyType string    `db:"currency_type"`
	Rate         float32   `db:"rate"`
	CurrencyDate time.Time `db:"currency_date"`
}

func (c *Currency) ToDTO() (dtos.Currency, error) {
	var date time.Time
	if c.CurrencyDate.IsZero() {
		date = time.Now()
	}

	dto := dtos.Currency{
		Id:           c.Id,
		CurrencyType: dtos.CurrencyType(c.CurrencyType),
		Rate:         c.Rate,
		CurrencyDate: date.Format("2006-01-02"),
	}

	return dto, dto.Validate()
}
