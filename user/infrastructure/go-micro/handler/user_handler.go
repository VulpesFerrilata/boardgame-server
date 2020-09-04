package handler

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/go-micro/converter"
	"github.com/VulpesFerrilata/boardgame-server/user/internal/usecase/interactor"
)

func NewUserHandler(userInteractor interactor.UserInteractor) user.UserHandler {
	return userHandler{
		userInteractor: userInteractor,
	}
}

type userHandler struct {
	userInteractor interactor.UserInteractor
}

func (uh userHandler) GetUserById(ctx context.Context, userRequestPb *user.UserRequest, userResponsePb *user.UserResponse) error {
	userForm := converter.ConvertUserRequestPbToUserForm(userRequestPb)
	userDTO, err := uh.userInteractor.GetUserById(ctx, userForm)
	if err != nil {
		return err
	}
	userResponsePb = converter.ConvertUserDtoToUserResponsePb(userDTO)
	return nil
}

func (uh userHandler) GetUserByCredential(ctx context.Context, credentialRequestPb *user.CredentialRequest, userResponsePb *user.UserResponse) error {
	loginForm := converter.ConvertCredentialRequestPbToLoginForm(credentialRequestPb)
	userDTO, err := uh.userInteractor.GetUserByCredential(ctx, loginForm)
	if err != nil {
		return err
	}
	userResponsePb = converter.ConvertUserDtoToUserResponsePb(userDTO)
	return nil
}
