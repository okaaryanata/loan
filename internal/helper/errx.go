package helper

import (
	"fmt"
)

type Errorx interface {
	Error() string
	GetStatusCode() int
}

func NewErrorx(statusCode int, message string) Errorx {
	return construct(statusCode, message)
}

func NewErrorxFromErr(err error) Errorx {
	return construct(500, err.Error())
}

func NewErrxFromErri(statusCode int, err error) Errorx {
	return construct(statusCode, err.Error())
}

func NewErrorxf(format string, args ...interface{}) Errorx {
	return construct(500, fmt.Sprintf(format, args...))
}

func NewErrorxif(statusCode int, format string, args ...interface{}) Errorx {
	return construct(statusCode, fmt.Sprintf(format, args...))
}
