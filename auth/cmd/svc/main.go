package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/container"
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
	"github.com/micro/go-micro/v2"
)

func main() {
	container := container.NewContainer()

	if err := container.Invoke(func(authHandler auth.AuthHandler) error {
		// New Service
		service := micro.NewService(
			micro.Name("boardgame.auth.svc"),
			micro.Version("latest"),
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
