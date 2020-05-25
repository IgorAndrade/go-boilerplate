package apiErrors

import (
	"errors"
	"fmt"
	"io"
)

const (
	NoType        = ErrorType("None")
	BadRequest    = ErrorType("BadRequest")
	NotFound      = ErrorType("NotFound")
	Timeout       = ErrorType("Timeout")
	InternalError = ErrorType("InternalError")
	HTTPError     = ErrorType("HTTP Error")
)

type ErrorType string

// New creates a new customError
func (t ErrorType) New(msg string) error {
	return customError{errorType: t, originalError: errors.New(msg), contextInfo: make(map[string]string)}
}

func (t ErrorType) NewError(err error) error {
	return customError{errorType: t, originalError: err, contextInfo: make(map[string]string)}
}

// New creates a new customError with formatted message
func (t ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)

	return customError{errorType: t, originalError: err, contextInfo: make(map[string]string)}
}

// Wrap creates a new wrapped error
func (t ErrorType) Wrap(err error, msg string) error {
	return t.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (t ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return AddErrorContext(t.NewError(err), "details", fmt.Sprintf(msg, args...))
}

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   map[string]string
}

func (e customError) Error() string {
	return fmt.Sprintf("%v: %s \n %v", e.errorType, e.originalError.Error(), e.contextInfo)
}

// New creates a no type error
func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg), contextInfo: make(map[string]string)}
}

func (e customError) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(e.Error()))
	return int64(n), err
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return customError{errorType: NoType,
		originalError: errors.New(fmt.Sprintf(msg, args...)),
		contextInfo:   make(map[string]string)}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	if customErr, ok := err.(customError); ok {
		customErr.contextInfo[field] = message
		return customError{errorType: customErr.errorType,
			originalError: customErr.originalError,
			contextInfo:   customErr.contextInfo}
	}
	context := map[string]string{field: message}
	return customError{errorType: NoType, originalError: err, contextInfo: context}
}

func GetErrorContext(err error) map[string]string {
	if customErr, ok := err.(customError); ok {
		return customErr.contextInfo
	}
	return make(map[string]string)
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}
	return NoType
}

func Is(e error, t ErrorType) bool {
	eType := GetType(e)
	return string(eType) == string(t)
}
