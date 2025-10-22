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

func (c CurrencyController) GetPair(
	leftCurrencyType dtos.CurrencyType,
	rightCurrencyType dtos.CurrencyType,
) (dtos.CurrencyPair, error) {
	leftRate, err := c.repository.GetRateByType(leftCurrencyType)
	if err != nil {
		return dtos.CurrencyPair{}, err
	}

	rightRate, err := c.repository.GetRateByType(rightCurrencyType)
	if err != nil {
		return dtos.CurrencyPair{}, err
	}

	return dtos.CurrencyPair{
		Left:  leftCurrencyType,
		Right: rightCurrencyType,
		Rate:  leftRate / rightRate,
	}, nil
}
