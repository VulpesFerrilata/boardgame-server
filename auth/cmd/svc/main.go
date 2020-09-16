package main

import (
	"database/sql"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/server/grpc"

	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/container"
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware"
	"github.com/micro/go-micro/v2"
)

func main() {
	container := container.NewContainer()

	if err := container.Invoke(func(authHandler auth.AuthHandler, transactionMiddleware *middleware.TransactionMiddleware, translatorMiddleware *middleware.TranslatorMiddleware) error {
		// New Service
		service := micro.NewService(
			micro.Name("boardgame.auth.svc"),
			micro.Version("latest"),
			micro.Server(
				grpc.NewServer(
					server.WrapHandler(transactionMiddleware.HandlerWrapperWithTxOptions(&sql.TxOptions{})),
					server.WrapHandler(translatorMiddleware.HandlerWrapper),
				),
			),
		)

		// Initialise service
		service.Init()

		// Register Handler
		if err := auth.RegisterAuthHandler(service.Server(), authHandler); err != nil {
			return err
		}

		// Run service
		return service.Run()
	}); err != nil {
		log.Fatal(err)
	}
}
