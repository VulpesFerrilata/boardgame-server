package errors

import "strings"

func NewValidationError(errs ...string) error {
	return &ValidationError{
		errs: errs,
	}
}

type ValidationError struct {
	errs []string
}

func (ve ValidationError) Error() string {
	return strings.Join(ve.errs, "\n")
}

func (ve *ValidationError) Append(err string) {
	ve.errs = append(ve.errs, err)
}

func (ve ValidationError) HasErrors() bool {
	return len(ve.errs) > 0
}
