package api

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/OlenEnkeli/GoCurrency/internal/errors"
	"github.com/OlenEnkeli/GoCurrency/internal/settings"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func handleError(err error, c *gin.Context) {
	switch e := err.(type) {
	case errors.Error:
		apiError := e.ToAPIError()
		c.JSON(apiError.HTTPCode, apiError.JSON())
	case error:
		apiError := errors.NewInternalError(e.Error()).ToAPIError()
		c.JSON(apiError.HTTPCode, apiError.JSON())
	default:
		apiError := errors.NewInternalError(fmt.Sprintf("Unknown error: %v", e)).ToAPIError()
		c.AbortWithStatusJSON(apiError.HTTPCode, apiError.JSON())
	}
}

func handlePanic(recovered any, c *gin.Context) {
	if settings.Settings.App.Mode != "prod" {
		debug.PrintStack()
	}

	var err error
	switch e := recovered.(type) {
	case error:
		err = e
	case string:
		err = fmt.Errorf("panic: %s", e)
	default:
		err = fmt.Errorf("panic: %v", e)
	}

	apiError := errors.NewInternalError(err.Error()).ToAPIError()
	c.AbortWithStatusJSON(apiError.HTTPCode, apiError.JSON())
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				handlePanic(recovered, c)
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			lastError := c.Errors[len(c.Errors)-1]
			handleError(lastError.Err, c)
			c.Abort()
		}
	}
}
