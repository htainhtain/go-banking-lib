package app_err

import "net/http"

type AppErr struct {
	Code    int
	Message string
}

func (a *AppErr) Error() string {
	return a.Message
}

type ErrMessage struct {
	Message string `json:"message"`
}

func (a *AppErr) ErrorResponse() *ErrMessage {
	return &ErrMessage{
		a.Message,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidateError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewUnauthorizedError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}
