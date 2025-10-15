package api

import (
	"net/http"

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

	c.JSON(
		http.StatusOK,
		currencies,
	)
}
