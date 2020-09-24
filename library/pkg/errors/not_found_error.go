package errors

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewNotFoundError(name string) Error {
	return &notFoundError{
		name: name,
	}
}

type notFoundError struct {
	name string
}

func (nfe notFoundError) Error() string {
	return fmt.Sprintf("%s not found", nfe.name)
}

func (nfe notFoundError) ToProblem(trans ut.Translator) iris.Problem {
	problem := iris.NewProblem()
	problem.Status(iris.StatusNotFound)
	problem.Type("about:blank")
	title, _ := trans.T("not-found-error")
	problem.Title(title)
	detail, _ := trans.T("not-found-error-detail", nfe.name)
	problem.Detail(detail)
	return problem
}

func (nfe notFoundError) ToStatus(trans ut.Translator) *status.Status {
	detail, _ := trans.T("not-found-error-detail", nfe.name)
	stt := status.New(codes.NotFound, detail)
	return stt
}
