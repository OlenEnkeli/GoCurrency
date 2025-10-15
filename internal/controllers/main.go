package controllers

import (
	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
	"github.com/OlenEnkeli/GoCurrency/internal/repositories"
)

type Currency interface {
	GetLatest() ([]dtos.Currency, error)
}

type Controller struct {
	Currency
}

func NewController(repository *repositories.Repository) *Controller {
	return &Controller{
		Currency: NewCurrencyController(repository),
	}
}
