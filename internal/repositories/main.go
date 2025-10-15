package repositories

import (
	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
	"github.com/jmoiron/sqlx"
)

type Currency interface {
	GetLatest() ([]dtos.Currency, error)
}

type Repository struct {
	Currency
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Currency: NewCurrencyPostgres(db),
	}
}
