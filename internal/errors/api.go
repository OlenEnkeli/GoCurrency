package errors

import "github.com/gin-gonic/gin"

type APIError struct {
	HTTPCode  int    `json:"http_code"`
	ErrorCode string `json:"error"`
	Message   string `json:"message"`
	Details   any    `json:"details,omitempty"`
}

func (e APIError) JSON() gin.H {
	return gin.H{
		"error":   e.ErrorCode,
		"message": e.Message,
		"details": e.Details,
	}
}
