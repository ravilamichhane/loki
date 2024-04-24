package thor

import (
	"errors"
	"syscall"
)

type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// trustedError is an error that is trusted to be returned to the client
type trustedError struct {
	Err    error
	Status int
}

func NewTrustedError(err error, status int) error {
	if err != nil {
		return &trustedError{
			err,
			status}
	}
	return &trustedError{err, status}
}

func (te *trustedError) Error() string {
	return te.Err.Error()
}

func IsTrustedError(err error) bool {
	var te *trustedError
	return errors.As(err, &te)
}

func GetTrustedError(err error) *trustedError {
	var te *trustedError
	if !errors.As(err, &te) {
		return nil
	}
	return te
}

// shutdownError is an error that is returned when the server is shutting down
type shutdownError struct {
	Message string
}

func NewShutdownError(message string) error {
	return &shutdownError{message}
}

func (se *shutdownError) Error() string {
	return se.Message
}

func IsShutdownError(err error) bool {
	var se *shutdownError
	return errors.As(err, &se)
}

func validateShutdownError(err error) bool {
	switch {
	case errors.Is(err, syscall.EPIPE):
		return false
	case errors.Is(err, syscall.ECONNRESET):
		return false
	}
	return true
}

// FieldError is an error that is returned when a field fails validation
type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

type FieldErrors []FieldError

func NewFieldError(field string, err error) FieldErrors {
	var fe FieldErrors
	if err != nil {
		fe = append(fe, FieldError{field, err.Error()})
	}

	return fe
}

func (fe FieldErrors) Error() string {

	return "Validation Error"
}

func (fe FieldErrors) Fields() map[string]string {
	m := make(map[string]string)
	for _, fld := range fe {
		m[fld.Field] = fld.Err
	}
	return m
}

func IsFieldErrors(err error) bool {
	_, ok := err.(FieldErrors)
	return ok
}

func GetFieldErrors(err error) FieldErrors {
	var fe FieldErrors
	if !errors.As(err, &fe) {
		return FieldErrors{}
	}
	return fe
}
