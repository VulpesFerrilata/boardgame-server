package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	catan "catan/proto/catan"
)

type Catan struct{}

func (e *Catan) Handle(ctx context.Context, msg *catan.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *catan.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
