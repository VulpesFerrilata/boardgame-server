package router

import (
	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/iris/controller"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/transaction"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/translator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Router interface {
	InitRoutes(app *iris.Application)
}

func NewRouter(authController controller.AuthController, transactionMiddleware *transaction.TransactionMiddleware, translatorMiddleware *translator.TranslatorMiddleware) Router {
	return &router{
		authController:        authController,
		transactionMiddleware: transactionMiddleware,
		translatorMiddleware:  translatorMiddleware,
	}
}

type router struct {
	authController        controller.AuthController
	transactionMiddleware *transaction.TransactionMiddleware
	translatorMiddleware  *translator.TranslatorMiddleware
}

func (r router) InitRoutes(app *iris.Application) {
	apiRoot := app.Party("/api")
	apiRoot.Use(r.transactionMiddleware.Serve, r.translatorMiddleware.Serve)
	mvcApp := mvc.New(apiRoot.Party("/auth"))
	mvcApp.Handle(r.authController)
}
