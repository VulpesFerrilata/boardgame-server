package handler

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/user"
	"github.com/VulpesFerrilata/boardgame-server/user/infrastructure/go-micro/viewmodel"
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
	userRequestVM := viewmodel.UserRequest{
		UserRequest: userRequestPb,
	}
	userDTO, err := uh.userInteractor.GetUserById(ctx, userRequestVM.ToUserForm())
	if err != nil {
		return err
	}
	userResponsePb = viewmodel.NewUserResponse(userDTO).UserResponse
	return nil
}

func (uh userHandler) GetUserByCredential(ctx context.Context, credentialRequestPb *user.CredentialRequest, userResponsePb *user.UserResponse) error {
	credentialRequestVM := viewmodel.CredentialRequest{
		CredentialRequest: credentialRequestPb,
	}
	userDTO, err := uh.userInteractor.GetUserByCredential(ctx, credentialRequestVM.ToLoginForm())
	if err != nil {
		return err
	}
	userResponsePb = viewmodel.NewUserResponse(userDTO).UserResponse
	return nil
}
