package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/dto"
)

func NewUserResponse(userDTO *dto.UserDTO) *UserResponse {
	userResponsePb := new(UserResponse)
	userResponsePb.ID = int64(userDTO.ID)
	return userResponsePb
}

type UserResponse struct {
	*user.UserResponse
}
