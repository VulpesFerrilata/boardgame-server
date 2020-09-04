package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"catan/handler"
	"catan/subscriber"

	catan "catan/proto/catan"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("boardgame.catan.service.catan"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	catan.RegisterCatanHandler(service.Server(), new(handler.Catan))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("boardgame.catan.service.catan", service.Server(), new(subscriber.Catan))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
