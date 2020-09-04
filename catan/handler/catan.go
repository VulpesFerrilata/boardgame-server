package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	catan "catan/proto/catan"
)

type Catan struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Catan) Call(ctx context.Context, req *catan.Request, rsp *catan.Response) error {
	log.Info("Received Catan.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Catan) Stream(ctx context.Context, req *catan.StreamingRequest, stream catan.Catan_StreamStream) error {
	log.Infof("Received Catan.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&catan.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Catan) PingPong(ctx context.Context, stream catan.Catan_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&catan.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
