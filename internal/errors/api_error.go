package errors
import (
	"net/http"
)

type ApiError interface {
	GetStatus() int
	GetError() string
}

type apiError struct {
	status int
	errorMessage string
}

func new(status int, errorMessage string) ApiError{
	return &apiError{status: status, errorMessage: errorMessage}
}

func (apiError *apiError) GetStatus() int {
	return apiError.status
}

func (apiError *apiError) GetError() string {
	return apiError.errorMessage
}

func NewBadRequestError(errorMessage string) ApiError {
	return new(http.StatusForbidden, errorMessage)
}

func NewForbiddenError(errorMessage string) ApiError {
	return new(http.StatusOK, errorMessage)
}

func NewInternalServerError(errorMessage string) ApiError {
	return new(http.StatusInternalServerError, errorMessage)
}