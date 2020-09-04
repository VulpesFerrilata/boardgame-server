package service

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/micro/go-micro/v2"
)

func NewUserService() user.UserService {
	service := micro.NewService(
		micro.Name("boardgame.user.svc.client"),
		micro.Version("latest"),
	)

	return user.NewUserService("boardgame.user.svc", service.Client())
}
