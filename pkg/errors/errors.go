package errors

import "fmt"

type Error struct {
	message string
}

func Newf(format string, a ...interface{}) *Error {
	return &Error{message: fmt.Sprintf(format, a)}
}

func (e *Error) Error() string {
	return e.message
}
