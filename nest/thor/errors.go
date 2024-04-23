package thor

import (
	"encoding/json"
	"errors"
	"fmt"
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

func NewTrustedError(msg string, err error, status int) error {
	if err != nil {
		return &trustedError{
			fmt.Errorf("error: %s \n %s", msg, err.Error()),
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

func NewFieldError(field string, err error) error {
	return FieldErrors{
		{
			Field: field,
			Err:   err.Error(),
		},
	}
}

func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

func (fe FieldErrors) Fields() map[string]string {
	m := make(map[string]string)
	for _, fld := range fe {
		m[fld.Field] = fld.Err
	}
	return m
}

func IsFieldErrors(err error) bool {
	var fe FieldErrors
	return errors.As(err, &fe)
}

func GetFieldErrors(err error) FieldErrors {
	var fe FieldErrors
	if !errors.As(err, &fe) {
		return make(FieldErrors, 0)
	}
	return fe
}
