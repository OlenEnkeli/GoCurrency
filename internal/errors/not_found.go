package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	BaseError

	Entity string `json:"entity"`
	Field  string `json:"field"`
	Value  any    `json:"value"`
}

func NewNotFoundError(
	entity string,
	field string,
	value any,
) NotFoundError {
	return NotFoundError{
		BaseError: NewBaseError(fmt.Sprintf("not found %s with %s=%s", entity, field, value)),
		Entity:    entity,
		Field:     field,
		Value:     value,
	}
}

func (e NotFoundError) Print() string {
	return fmt.Sprintf("not found %s with %s=%s", e.Entity, e.Field, e.Value)
}

func (e NotFoundError) ToAPIError() APIError {
	return APIError{
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "ValidationError",
		Message:   e.Print(),
		Details: map[string]string{
			"entity": e.Entity,
			"field":  e.Field,
			"value":  fmt.Sprintf("%v", e.Value),
		},
	}
}
