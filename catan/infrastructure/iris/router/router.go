package router

import (
	"database/sql"

	"github.com/VulpesFerrilata/catan/infrastructure/iris/controller"

	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
)

type Router interface {
	InitRoutes(app *iris.Application)
}

func NewRouter(catanController controller.CatanController,
	transactionMiddleware *middleware.TransactionMiddleware,
	translatorMiddleware *middleware.TranslatorMiddleware,
	errorMiddleware *middleware.ErrorMiddleware) Router {
	return &router{
		catanController:       catanController,
		transactionMiddleware: transactionMiddleware,
		translatorMiddleware:  translatorMiddleware,
		errorMiddleware:       errorMiddleware,
	}
}

type router struct {
	catanController       controller.CatanController
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
	wsAPI := apiRoot.Party("/ws")
	mvcApp := mvc.New(wsAPI)
	mvcApp.HandleWebsocket(r.catanController)
	wsServer := websocket.New(websocket.DefaultGorillaUpgrader, mvcApp)
	wsAPI.Get("/catan", websocket.Handler(wsServer))
}
