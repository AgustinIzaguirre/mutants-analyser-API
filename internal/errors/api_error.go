package errors
import (
	"net/http"
)

type ApiError interface {
	GetStatus() int
	GetError() string
}

type apiError struct {
	Status int `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func new(status int, errorMessage string) ApiError{
	return &apiError{Status: status, ErrorMessage: errorMessage}
}

func (apiError *apiError) GetStatus() int {
	return apiError.Status
}

func (apiError *apiError) GetError() string {
	return apiError.ErrorMessage
}

func NewBadRequestError(errorMessage string) ApiError {
	return new(http.StatusBadRequest, errorMessage)
}

func NewForbiddenError(errorMessage string) ApiError {
	return new(http.StatusForbidden, errorMessage)
}

func NewInternalServerError(errorMessage string) ApiError {
	return new(http.StatusInternalServerError, errorMessage)
}