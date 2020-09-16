package error_handler

import (
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/kataras/iris/v12"
)

type ErrorHandler interface {
	HandleError(ctx iris.Context, err error) iris.Problem
}

func NewDefaultErrorHandler(translatorMiddleware *middleware.TranslatorMiddleware) ErrorHandler {
	return &defaultErrorHandler{
		translatorMiddleware: translatorMiddleware,
	}
}

type defaultErrorHandler struct {
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (deh defaultErrorHandler) HandleError(ctx iris.Context, err error) iris.Problem {
	problem := iris.NewProblem()
	trans := deh.translatorMiddleware.Get(ctx.Request().Context())
	switch err.(type) {
	case *errors.ValidationError:
		title, _ := trans.T("validation-error")
		problem.Status(iris.StatusUnprocessableEntity)
		problem.Type("about:blank")
		problem.Title(title)
		problem.Detail(err.Error())
	default:
		title, _ := trans.T("internal-error")
		problem.Status(iris.StatusInternalServerError)
		problem.Type("about:blank")
		problem.Title(title)
		problem.Detail(err.Error())
	}
	return problem
}
