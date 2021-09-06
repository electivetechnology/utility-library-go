package connect

import "fmt"

type ApiError struct {
	Message string
	Status  int
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Api returned %d response with message: %s", e.Status, e.Message)
}

func NewApiError(message string) ApiError {
	return ApiError{Message: message}
}

func NewInternalError(message string) ApiError {
	return ApiError{Message: message, Status: 500}
}
