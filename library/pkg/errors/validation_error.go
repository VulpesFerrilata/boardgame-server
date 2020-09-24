package errors

import (
	"bytes"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidationError interface {
	Error
	WithFieldError(field string, err string)
	HasErrors() bool
}

func NewValidationError() ValidationError {
	return &validationError{}
}

type validationError struct {
	fieldErrs map[string][]string
}

func (ve validationError) Error() string {
	buff := bytes.NewBufferString("")

	for _, errs := range ve.fieldErrs {
		for _, err := range errs {
			buff.WriteString(err)
			buff.WriteString("\n")
		}
	}

	return strings.TrimSpace(buff.String())
}

func (ve *validationError) WithFieldError(field string, err string) {
	if ve.fieldErrs == nil {
		ve.fieldErrs = make(map[string][]string)
	}
	errs, ok := ve.fieldErrs[field]
	if !ok {
		errs = make([]string, 0)
	}
	errs = append(errs, err)
	ve.fieldErrs[field] = errs
}

func (ve validationError) HasErrors() bool {
	for _, errs := range ve.fieldErrs {
		if len(errs) > 0 {
			return true
		}
	}
	return false
}

func (ve validationError) ToProblem(trans ut.Translator) iris.Problem {
	problem := iris.NewProblem()
	problem.Status(iris.StatusUnprocessableEntity)
	problem.Type("about:blank")
	title, _ := trans.T("validation-error")
	problem.Title(title)
	detail, _ := trans.T("validation-error-detail")
	problem.Detail(detail)
	problem.Key("errors", ve.fieldErrs)
	return problem
}

func (ve validationError) ToStatus(trans ut.Translator) (*status.Status, error) {
	detail, _ := trans.T("validation-error-detail")
	stt := status.New(codes.InvalidArgument, detail)
	badRequest := &errdetails.BadRequest{
		FieldViolations: make([]*errdetails.BadRequest_FieldViolation, 0),
	}
	for field, errs := range ve.fieldErrs {
		for _, err := range errs {
			fieldViolation := &errdetails.BadRequest_FieldViolation{
				Field:       field,
				Description: err,
			}
			badRequest.FieldViolations = append(badRequest.FieldViolations, fieldViolation)
		}
	}
	return stt.WithDetails(badRequest)
}
