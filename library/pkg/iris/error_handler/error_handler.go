package error_handler

import (
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero/di"
	"google.golang.org/grpc/status"
)

func NewDefaultErrorHandler(translatorMiddleware *middleware.TranslatorMiddleware) di.ErrorHandler {
	return &errorHandler{
		translatorMiddleware: translatorMiddleware,
	}
}

type errorHandler struct {
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (eh errorHandler) HandleError(ctx iris.Context, err error) {
	if err == nil {
		return
	}
	trans := eh.translatorMiddleware.Get(ctx.Request().Context())

	if stt, ok := status.FromError(err); ok {
		if serverErr, ok := errors.NewStatusError(stt); ok {
			err = serverErr
		}
	}

	if serverErr, ok := err.(errors.Error); ok {
		problem := serverErr.ToProblem(trans)
		err = problem
	}

	if _, ok := err.(iris.Problem); !ok {
		problem := iris.NewProblem()
		problem.Status(iris.StatusInternalServerError)
		problem.Type("about:blank")
		title, _ := trans.T("internal-error")
		problem.Title(title)
		problem.Detail(err.Error())
		err = problem
	}

	ctx.Problem(err)

}
