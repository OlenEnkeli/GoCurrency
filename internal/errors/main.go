package errors

type Error interface {
	error
	Print() string
	ToAPIError() APIError
}

type BaseError struct {
	error

	message string
}

func NewBaseError(message string) BaseError {
	return BaseError{message: message}
}

func (b BaseError) Error() string {
	return b.message
}
