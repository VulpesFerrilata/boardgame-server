package router

import (
	"database/sql"

	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/iris/controller"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/iris/error_handler"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Router interface {
	InitRoutes(app *iris.Application)
}

func NewRouter(authController controller.AuthController, transactionMiddleware *middleware.TransactionMiddleware, translatorMiddleware *middleware.TranslatorMiddleware) Router {
	return &router{
		authController:        authController,
		transactionMiddleware: transactionMiddleware,
		translatorMiddleware:  translatorMiddleware,
	}
}

type router struct {
	authController        controller.AuthController
	transactionMiddleware *middleware.TransactionMiddleware
	translatorMiddleware  *middleware.TranslatorMiddleware
}

func (r router) InitRoutes(app *iris.Application) {
	apiRoot := app.Party("/api")
	apiRoot.Use(
		r.transactionMiddleware.ServeWithTxOptions(&sql.TxOptions{}),
		r.translatorMiddleware.Serve,
	)
	mvcApp := mvc.New(apiRoot.Party("/auth"))
	mvcApp.ErrorHandler = error_handler.ErrorHandlerWrapper(
		error_handler.NewDefaultErrorHandler(r.translatorMiddleware),
	)
	mvcApp.Handle(r.authController)
}
