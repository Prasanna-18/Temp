package utilerrors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int
	AMessage string
	AError   string
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.AError
}

func NewApiError(statuscode int, message string) ApiError {
	return &apiError{
		AStatus: statuscode,
		AMessage: message,
	}
}


func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus: http.StatusInternalServerError,
		AMessage: message,
	}
}

func NewNotfoundApiError(message string) ApiError {
	return &apiError{
		AStatus: http.StatusNotFound,
		AMessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		AStatus: http.StatusBadRequest,
		AMessage: message,
	}
}