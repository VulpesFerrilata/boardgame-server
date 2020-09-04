package service

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
	"github.com/micro/go-micro/v2"
)

func NewAuthService() auth.AuthService {
	service := micro.NewService(
		micro.Name("boardgame.auth.svc.client"),
		micro.Version("latest"),
	)

	return auth.NewAuthService("boardgame.auth.svc", service.Client())
}
