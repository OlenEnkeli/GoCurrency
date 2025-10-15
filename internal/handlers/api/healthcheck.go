package api

import (
	"net/http"

	"github.com/OlenEnkeli/GoCurrency/internal/settings"
	"github.com/gin-gonic/gin"
)

type HealthcheckResponse struct {
	Status     string `json:"status"`
	AppMode    string `json:"app_mode"`
	ApiName    string `json:"api_name"`
	ApiVersion string `json:"api_version"`
}

// Healthcheck проверяет работоспособность сервиса.
// @Summary Проверка работы API
// @Description Проверяет работоспособность сервиса
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthcheckResponse
// @Router / [get]
func (h *Handler) Healthcheck(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		HealthcheckResponse{
			Status:     "ok",
			AppMode:    settings.Settings.App.Mode,
			ApiName:    "go_concurrency",
			ApiVersion: settings.Settings.API.Version,
		},
	)
}
