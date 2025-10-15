package errors

import (
	"fmt"
	"net/http"
)

type ValidationError struct {
	BaseError

	Field   string `json:"field"`
	Message string `json:"message"`
	Value   any    `json:"values"`
}

func NewValidationError(
	field string,
	message string,
	value any,
) ValidationError {
	return ValidationError{
		BaseError: NewBaseError(fmt.Sprintf("validation error on field %s: %s", field, message)),
		Field:     field,
		Message:   message,
		Value:     value,
	}
}

func (e ValidationError) Print() string {
	return fmt.Sprintf("ValidationError on field %s=%s: %s", e.Field, e.Value, e.Message)
}

func (e ValidationError) ToAPIError() APIError {
	return APIError{
		HTTPCode:  http.StatusUnprocessableEntity,
		ErrorCode: "ValidationError",
		Message:   fmt.Sprintf("Cannot validate field %s: %s", e.Field, e.Message),
		Details: map[string]string{
			e.Field: e.Message,
		},
	}
}
