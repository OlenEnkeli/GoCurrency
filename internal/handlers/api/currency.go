package api

import (
	"net/http"

	"github.com/OlenEnkeli/GoCurrency/internal/dtos"
	"github.com/gin-gonic/gin"
)

// GetLatestCurrencies Получение списка актуальных валют
// @Summary Получение текущего курса валют
// @Description Получение списка актуальных валют (курса к доллару)
// @Tags currencies
// @Accept json
// @Produce json
// @Success 200 {array} dtos.Currency
// @Router /currencies/latest/ [get]
func (h *Handler) GetLatestCurrencies(c *gin.Context) {
	currencies, err := h.controller.Currency.GetLatest()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, currencies)
}

// GetPair Получение курса валютной пары
// @Summary Получение курса валютной пары
// @Description Получение отношения валюты left к валюте right
// @Tags currencies
// @Accept json
// @Produce json
// @Success 200 {object} dtos.CurrencyPair
// @Router /currencies/pair/ [get]
// @Param left  query string true "Left currency type"
// @Param right query string true "Right currency type"
func (h *Handler) GetPair(c *gin.Context) {
	left := dtos.CurrencyType(c.Query("left"))
	if err := left.Validate(); err != nil {
		c.Error(err)
		return
	}

	right := dtos.CurrencyType(c.Query("right"))
	if err := right.Validate(); err != nil {
		c.Error(err)
		return
	}

	result, err := h.controller.Currency.GetPair(left, right)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
