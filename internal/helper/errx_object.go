package helper

import "fmt"

type ErrorxObj struct {
	StatusCode int
	Message    string
}

func construct(code int, message string) Errorx {
	return &ErrorxObj{
		StatusCode: code,
		Message:    message,
	}
}

func (e *ErrorxObj) Error() string {
	return fmt.Sprintf("ERR: %s", e.Message)
}

func (e *ErrorxObj) GetStatusCode() int {
	return e.StatusCode
}
