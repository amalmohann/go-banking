package errs

import "net/http"

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NotFoundError(message string) *AppError {
	return &AppError{Status: http.StatusNotFound, Message: message}
}

func InternalServerError(message string) *AppError {
	return &AppError{Status: http.StatusInternalServerError, Message: message}
}
