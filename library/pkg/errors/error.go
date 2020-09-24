package errors

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error interface {
	error
	ToProblem(trans ut.Translator) iris.Problem
}

func NewErrorFromStatus(stt *status.Status) (Error, bool) {
	switch stt.Code() {
	case codes.InvalidArgument:
		for _, detail := range stt.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				validationErr := NewValidationError()
				for _, fieldViolation := range t.GetFieldViolations() {
					validationErr.WithFieldError(fieldViolation.Field, fieldViolation.Description)
				}
				if validationErr.HasErrors() {
					return validationErr, true
				}
			}
		}
	case codes.NotFound:
		return NewNotFoundError(stt.Message()), true
	}
	return nil, false
}
