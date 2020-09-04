package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/container"
	"github.com/micro/go-micro/v2"
)

func main() {
	container := container.NewContainer()

	if err := container.Invoke(func(userHandler user.UserHandler) error {
		// New Service
		service := micro.NewService(
			micro.Name("boardgame.user.svc"),
			micro.Version("latest"),
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
