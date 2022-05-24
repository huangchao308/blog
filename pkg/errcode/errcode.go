package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var codes = make(map[int]string)

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("code %d already exists", code))
	}
	codes[code] = msg
	return &Error{
		Code:    code,
		Msg:     msg,
		Details: make([]string, 0),
	}
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d]: %s", e.Code, e.Msg)
}

func (e *Error) GetDetails() []string {
	return e.Details
}

func (e *Error) WithDetails(details ...string) *Error {
	e.Details = append(e.Details, details...)
	return e
}

func (e *Error) StatusCode() int {
	switch e.GetCode() {
	case Success.Code:
		return http.StatusOK
	case ServerErr.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case NotFound.Code:
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenErr.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		return http.StatusUnauthorized
	case ToomanyRequests.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
