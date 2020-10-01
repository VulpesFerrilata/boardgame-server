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
		if err, ok := errors.NewStatusError(stt); ok {
			eh.HandleError(ctx, err)
			return
		}
	}

	if err, ok := err.(errors.Error); ok {
		problem := err.ToProblem(trans)
		eh.HandleError(ctx, problem)
		return
	}

	if _, ok := err.(iris.Problem); !ok {
		problem := iris.NewProblem()
		problem.Status(iris.StatusInternalServerError)
		problem.Type("about:blank")
		title, _ := trans.T("internal-error")
		problem.Title(title)
		problem.Detail(err.Error())
		eh.HandleError(ctx, problem)
		return
	}

	ctx.Problem(err)
}
