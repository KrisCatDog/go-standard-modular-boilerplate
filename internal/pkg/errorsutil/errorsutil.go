package errorsutil

import (
	"fmt"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
)

// InternalError represents an internal error that has wrapped some necessary fields.
type InternalError struct {
	err  error
	msg  string
	code api.CommonErrorCode
}

// Wrapf accept and returns the instance of InternalError.
func Wrapf(err error, msg string, code api.CommonErrorCode) error {
	return &InternalError{
		err:  err,
		msg:  msg,
		code: code,
	}
}

// Error returns actual combined and formatted error messages from InternalError instance.
func (e InternalError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.err)
	}

	return e.msg
}

// Code returns an internal code from InternalError instance.
func (e *InternalError) Code() api.CommonErrorCode {
	return e.code
}
