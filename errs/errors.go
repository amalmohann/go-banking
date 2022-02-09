package errs

import "net/http"

type AppError struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
}

func (e AppError) ToResponse() *AppError {
	return &AppError{Message: e.Message}
}

func NotFoundError(message string) *AppError {
	return &AppError{Status: http.StatusNotFound, Message: message}
}

func InternalServerError(message string) *AppError {
	return &AppError{Status: http.StatusInternalServerError, Message: message}
}

func ValidationError(message string) *AppError {
	return &AppError{Status: http.StatusUnprocessableEntity, Message: message}
}
