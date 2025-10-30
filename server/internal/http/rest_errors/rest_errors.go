package resterrors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

func NewRestErr(message, err string, code int) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewUnprocessableEntityError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unprocessable entity",
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}
