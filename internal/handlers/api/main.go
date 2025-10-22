package api

import (
	_ "github.com/OlenEnkeli/GoCurrency/docs"
	"github.com/OlenEnkeli/GoCurrency/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Currency API
// @version 1.0
// @description Демо-проект на Go с поддержкой PostgreSQL и MongoDB

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) *Handler {
	return &Handler{controller: controller}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(ErrorMiddleware())
	router.Use(CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", h.Healthcheck)

	currencies := router.Group("/currencies/")
	{
		currencies.GET("/latest/", h.GetLatestCurrencies)
		currencies.GET("/pair/", h.GetPair)
	}

	return router
}
