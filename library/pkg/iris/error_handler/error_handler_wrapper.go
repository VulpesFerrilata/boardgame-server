package error_handler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero/di"
)

func ErrorHandlerWrapper(errorHandlers ...ErrorHandler) di.ErrorHandler {
	return di.ErrorHandlerFunc(func(ctx iris.Context, err error) {
		for _, errorHandler := range errorHandlers {
			problem := errorHandler.HandleError(ctx, err)
			if problem.Validate() {
				ctx.Problem(problem)
				return
			}
		}
	})
}
