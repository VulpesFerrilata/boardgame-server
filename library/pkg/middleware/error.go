package middleware

import (
	"context"

	"github.com/micro/go-micro/v2/server"
	"google.golang.org/grpc/status"
)

type Error interface {
	ToStatus() status.Status
}

type ErrorMiddleware struct {
}

func (em ErrorMiddleware) HandlerWrapper(f server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := f(ctx, req, rsp)
		if err, ok := err.(Error); ok {
			stt := err.ToStatus()
			return stt.Err()
		}
		return err
	}
}
