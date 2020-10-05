package viewmodel

import (
	"github.com/VulpesFerrilata/boardgame-server/auth/internal/usecase/form"
	"github.com/VulpesFerrilata/boardgame-server/grpc/protoc/auth"
)

func NewTokenRequest(tokenRequestPb *auth.TokenRequest) *TokenRequest {
	return &TokenRequest{
		tokenRequestPb: tokenRequestPb,
	}
}

type TokenRequest struct {
	tokenRequestPb *auth.TokenRequest
}

func (tr TokenRequest) ToTokenForm() *form.TokenForm {
	tokenForm := new(form.TokenForm)
	tokenForm.Token = tr.tokenRequestPb.GetToken()
	return tokenForm
}
