package rest_err

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"err"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiden(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbiden",
		Code:    http.StatusForbidden,
	}
}
