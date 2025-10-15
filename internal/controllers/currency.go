package controllers

import (
	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
	"github.com/OlenEnkeli/GoCurrency/internal/repositories"
)

type CurrencyController struct {
	repository *repositories.Repository
}

func NewCurrencyController(repository *repositories.Repository) *CurrencyController {
	return &CurrencyController{
		repository: repository,
	}
}

func (c CurrencyController) GetLatest() ([]dtos.Currency, error) {
	return c.repository.GetLatest()
}
