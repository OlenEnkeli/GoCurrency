package errors

import (
	"fmt"
	"net/http"
)

type InternalError struct {
	BaseError

	Message string `json:"message"`
}

func NewInternalError(
	message string,
) InternalError {
	return InternalError{
		BaseError: NewBaseError(message),
		Message:   message,
	}
}

func (e InternalError) Print() string {
	return fmt.Sprintf("InternalServerError unhandled: %s", e.Message)
}

func (e InternalError) ToAPIError() APIError {
	return APIError{
		HTTPCode:  http.StatusUnprocessableEntity,
		ErrorCode: "InternalError",
		Message:   e.Message,
	}
}
