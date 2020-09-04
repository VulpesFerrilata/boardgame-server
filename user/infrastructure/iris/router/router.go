package router

import (
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/transaction"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/translator"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/iris/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Router interface {
	InitRoutes(app *iris.Application)
}

func NewRouter(userController controller.UserController, transactionMiddleware *transaction.TransactionMiddleware, translatorMiddleware *translator.TranslatorMiddleware) Router {
	return &router{
		userController:        userController,
		transactionMiddleware: transactionMiddleware,
		translatorMiddleware:  translatorMiddleware,
	}
}

type router struct {
	userController        controller.UserController
	transactionMiddleware *transaction.TransactionMiddleware
	translatorMiddleware  *translator.TranslatorMiddleware
}

func (r router) InitRoutes(app *iris.Application) {
	apiRoot := app.Party("/api")
	apiRoot.Use(r.transactionMiddleware.Serve, r.translatorMiddleware.Serve)
	mvcApp := mvc.New(apiRoot.Party("/user"))
	mvcApp.Handle(r.userController)
}
