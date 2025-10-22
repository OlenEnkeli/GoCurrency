package errors

import (
	"fmt"
	"net/http"
)

type ETLError struct {
	BaseError

	Provider    string `json:"provider"`
	StatusCode  int    `json:"status_code"`
	OriginalURL string `json:"original_url"`
}

func NewETLError(
	provider string,
	statusCode int,
	originalURL string,
) ETLError {
	return ETLError{
		BaseError:   NewBaseError(fmt.Sprintf("call to API %s failed: %s response (%s)", provider, statusCode, originalURL)),
		Provider:    provider,
		StatusCode:  statusCode,
		OriginalURL: originalURL,
	}
}

func (e ETLError) Print() string {
	return fmt.Sprintf("call to API %s failed: %s response (%s)", e.Provider, e.OriginalURL, e.StatusCode)
}

func (e ETLError) ToAPIError() APIError {
	return APIError{
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "ValidationError",
		Message:   e.Print(),
		Details: map[string]any{
			"provider":     e.Provider,
			"status_code":  e.StatusCode,
			"original_url": e.OriginalURL,
		},
	}
}
