package handler

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/auth/infrastructure/go-micro/viewmodel"
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/usecase/interactor"
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
)

func NewAuthHandler(authInteractor interactor.AuthInteractor) auth.AuthHandler {
	return &authHandler{
		authInteractor: authInteractor,
	}
}

type authHandler struct {
	authInteractor interactor.AuthInteractor
}

func (ah authHandler) Authenticate(ctx context.Context, tokenRequestPb *auth.TokenRequest, claimResponsePb *auth.ClaimResponse) error {
	tokenRequestVM := viewmodel.TokenRequest{
		TokenRequest: tokenRequestPb,
	}

	claimDTO, err := ah.authInteractor.Authenticate(ctx, tokenRequestVM.ToTokenForm())
	if err != nil {
		return err
	}

	claimResponsePb = viewmodel.NewClaimResponse(claimDTO).ClaimResponse
	return nil
}
