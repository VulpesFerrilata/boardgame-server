package service

import (
	"github.com/VulpesFerrilata/grpc/protoc/auth"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
)

func NewAuthService(translatorMiddleware *middleware.TranslatorMiddleware) auth.AuthService {
	service := micro.NewService(
		micro.Name("boardgame.auth.svc.client"),
		micro.Version("latest"),
		micro.Client(
			grpc.NewClient(
				client.WrapCall(
					translatorMiddleware.CallWrapper,
				),
			),
		),
	)

	return auth.NewAuthService("boardgame.auth.svc", service.Client())
}
