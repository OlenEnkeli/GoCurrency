package api

import (
	_ "github.com/OlenEnkeli/GoCurrency/docs"
	routers "github.com/OlenEnkeli/GoCurrency/internal/api/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Currency API
// @version 1.0
// @description Демо-проект на Go с поддержкой PostgreSQL и MongoDB

func NewRouter() *gin.Engine {
	handler := &routers.Handler{}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", handler.Healthcheck)

	return router
}
