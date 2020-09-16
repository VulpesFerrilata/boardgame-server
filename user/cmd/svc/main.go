package main

import (
	"database/sql"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/server/grpc"

	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/container"
	"github.com/micro/go-micro/v2"
)

func main() {
	container := container.NewContainer()

	if err := container.Invoke(func(userHandler user.UserHandler, transactionMiddleware *middleware.TransactionMiddleware, translatorMiddleware *middleware.TranslatorMiddleware) error {
		// New Service
		service := micro.NewService(
			micro.Name("boardgame.user.svc"),
			micro.Version("latest"),
			micro.Server(
				grpc.NewServer(
					server.WrapHandler(translatorMiddleware.HandlerWrapper),
					server.WrapHandler(transactionMiddleware.HandlerWrapperWithTxOptions(&sql.TxOptions{})),
				),
			),
		)

		// Initialise service
		service.Init()

		// Register Handler
		if err := user.RegisterUserHandler(service.Server(), userHandler); err != nil {
			return err
		}

		// Run service
		return service.Run()
	}); err != nil {
		log.Fatal(err)
	}
}
