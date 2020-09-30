package errors

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
)

type Error interface {
	error
	ToProblem(trans ut.Translator) iris.Problem
}
