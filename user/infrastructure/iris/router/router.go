package router

import (
	"database/sql"

	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/iris/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Router interface {
	InitRoutes(app *iris.Application)
}

func NewRouter(userController controller.UserController, transactionMiddleware *middleware.TransactionMiddleware, translatorMiddleware *middleware.TranslatorMiddleware) Router {
	return &router{
		userController:        userController,
		transactionMiddleware: transactionMiddleware,
		translatorMiddleware:  translatorMiddleware,
	}
}

type router struct {
	userController        controller.UserController
	transactionMiddleware *middleware.TransactionMiddleware
	translatorMiddleware  *middleware.TranslatorMiddleware
	errorMiddleware       *middleware.ErrorMiddleware
}

func (r router) InitRoutes(app *iris.Application) {
	apiRoot := app.Party("/api")
	apiRoot.Use(
		r.transactionMiddleware.ServeWithTxOptions(&sql.TxOptions{}),
		r.translatorMiddleware.Serve,
	)
	mvcApp := mvc.New(apiRoot.Party("/user"))
	mvcApp.HandleError(r.errorMiddleware.ErrorHandler)
	mvcApp.Handle(r.userController)
}
