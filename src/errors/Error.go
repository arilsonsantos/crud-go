package errors

import "net/http"

type ErrorDto struct {
	Message string  `json:"message"`
	Err     string  `json:"err"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ErrorDto) Error() string {
	return r.Message
}

func RestError(message string, err string, code int, causes []Cause) *ErrorDto {
	return &ErrorDto{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func BadRequestError(message string) *ErrorDto {
	return &ErrorDto{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func BadRequestCauseError(message string, causes []Cause) *ErrorDto {
	return &ErrorDto{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func InternalServerError(message string) *ErrorDto {
	return &ErrorDto{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NotFoundError(message string) *ErrorDto {
	return &ErrorDto{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func ForbiddenError(message string) *ErrorDto {
	return &ErrorDto{
		Message: message,
		Err:     "Forbiden",
		Code:    http.StatusForbidden,
	}
}
