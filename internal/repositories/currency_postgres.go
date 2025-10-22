package repositories

import (
	"fmt"

	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
	"github.com/OlenEnkeli/GoCurrency/internal/errors"
	"github.com/OlenEnkeli/GoCurrency/internal/repositories/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CurrencyPostgres struct {
	db *sqlx.DB
}

func NewCurrencyPostgres(db *sqlx.DB) *CurrencyPostgres {
	return &CurrencyPostgres{db: db}
}

func (p *CurrencyPostgres) GetLatest() ([]dtos.Currency, error) {
	var result []models.Currency

	err := p.db.Select(
		&result,
		"SELECT id, currency_type, rate FROM current_currency;",
	)
	if err != nil {
		return nil, errors.NewInternalError(fmt.Sprintf("Can`t select from DB: %s", err))
	}

	currencies := make([]dtos.Currency, len(result))

	for index, model := range result {
		dto, err := model.ToDTO()
		if err != nil {
			return nil, err
		}

		currencies[index] = dto
	}

	return currencies, nil
}

func (p *CurrencyPostgres) GetRateByType(currencyType dtos.CurrencyType) (float32, error) {
	var result float32
	if err := p.db.Get(
		&result,
		fmt.Sprintf(
			"SELECT rate FROM current_currency WHERE currency_type='%s';",
			currencyType,
		),
	); err != nil {
		logrus.Error(err)
		return 0, errors.NewNotFoundError("Currency", "currency_type", currencyType)
	}

	return result, nil
}
